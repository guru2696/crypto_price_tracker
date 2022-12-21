package plerrors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

type StatusCode int

const (
	BadInput StatusCode = iota
	InternalServiceError
	Unauthorized
	NoResultFound
	MultiStatus
	StatusConflict
	StatusNotFound
	ExpectationFailed
)

var statusMap = map[StatusCode]int{
	BadInput:             http.StatusBadRequest,
	InternalServiceError: http.StatusInternalServerError,
	Unauthorized:         http.StatusUnauthorized,
	NoResultFound:        http.StatusBadRequest,
	MultiStatus:          http.StatusMultiStatus,
	StatusConflict:       http.StatusConflict,
	StatusNotFound:       http.StatusNotFound,
	ExpectationFailed:    http.StatusExpectationFailed,
}

var statusCodeMap = map[int]StatusCode{
	http.StatusBadRequest:          BadInput,
	http.StatusInternalServerError: InternalServiceError,
	http.StatusUnauthorized:        Unauthorized,
	http.StatusMultiStatus:         MultiStatus,
	http.StatusConflict:            StatusConflict,
	http.StatusNotFound:            StatusNotFound,
	http.StatusExpectationFailed:   ExpectationFailed,
}

func HttpStatusToError(httpStatus int) StatusCode {
	if val, ok := statusCodeMap[httpStatus]; ok {
		return val
	}
	return InternalServiceError
}

func (s StatusCode) ToHttpStatus() int {
	return statusMap[s]
}

type ServiceError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success bool      `json:"success"`
	Error   ErrorBody `json:"error"`
}

type DataResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

type MetaData map[string]interface{}

func (md *MetaData) String() string {
	bty, _ := json.Marshal(*md)
	return string(bty)
}

func (md *MetaData) Set(key string, val interface{}) {
	if strings.Contains(key, ".") || strings.Contains(key, "[") {
		md.SetNested(key, val)
		return
	}

	if *md == nil {
		*md = MetaData{}
	}
	(*md)[key] = val
}

func (md *MetaData) SetNested(key string, val interface{}) error {
	keys := strings.Split(strings.Replace(key, "[", ".[", -1), ".")
	jByte, err := json.Marshal(md)
	if err != nil {
		return err
	}
	valByte, err := json.Marshal(val)
	if err != nil {
		return err
	}
	rByte, err := jsonparser.Set(jByte, valByte, keys...)
	if err != nil {
		return err
	}
	return json.Unmarshal(rByte, &md)
}

type ErrorBody struct {
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	MetaData MetaData `json:"meta_data"`
}

func (s ServiceError) Error() string {
	return "Service Error : " + s.Code + " : " + s.Message
}

func (eb ErrorBody) Error() string {
	return "Error : " + eb.Code + " : " + eb.Message + ":" + eb.MetaData.String()
}

// --------------------------------------------------------------------------------//
type stack []uintptr

type causeErr struct {
	msg   string
	trace *stack
}

func (ce *causeErr) Error() string {
	return ce.msg
}

func (ce *causeErr) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*ce.trace))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*ce.trace)[i])
	}
	return f
}

type IncorrectUUIDError struct {
	ServiceError
	UUIDFields []string `json:"fields"`
}

func ErrIncorrectUUID(fields ...string) IncorrectUUIDError {
	return IncorrectUUIDError{
		ServiceError: ServiceError{"GE_0001", "Invalid UUID passed"},
		UUIDFields:   fields,
	}
}

type RequestParamMissingError struct {
	ServiceError
	UUIDFields []string `json:"fields"`
}

type DoFailError struct {
	Meta interface{}
	Err  error
}

func (s DoFailError) Error() string {
	return fmt.Sprintf("Failing for: %s, meta:%+v", s.Err.Error(), s.Meta)
}

func NewDoFailError(err error, meta ...interface{}) DoFailError {
	return DoFailError{Err: err, Meta: meta}
}

type NoFailError struct {
	Err  error
	Meta interface{}
}

func (s NoFailError) Error() string {
	return fmt.Sprintf("Failing for: %s, meta:%+v", s.Err.Error(), s.Meta)
}

func NewNoFailError(err error, meta ...interface{}) NoFailError {
	return NoFailError{Err: err, Meta: meta}
}

//--------------------------------------------------------------------------------//

var ErrMalformedJson = ServiceError{
	Code:    "GE_0001",
	Message: "Malformed json",
}

var ErrMalformedJsonPayload = ServiceError{
	Code:    "GE_0001",
	Message: "One or more parameter in payload is invalid. For example, a parameter is string instead integer.",
}

type AppError struct {
	Code          string                 `json:"code"`
	Message       string                 `json:"message"`               // Message to be display to the end user without debugging information
	DetailedError string                 `json:"detailed_error"`        // Internal error string to help the developer
	RequestId     string                 `json:"request_id,omitempty"`  // The RequestId that's also set in the header
	StatusCode    StatusCode             `json:"status_code,omitempty"` // The http status code
	Where         string                 `json:"-"`                     // The function where it happened in the form of Struct.Func
	IsOAuth       bool                   `json:"is_oauth,omitempty"`    // Whether the error is OAuth specific
	Params        map[string]interface{} `json:"params,omitempty"`
	cerr          *causeErr
	CallerService string `json:"caller_service,omitempty"`
}

func NewAppError(where string, code string, message string, status StatusCode, details string, params map[string]interface{}) *AppError {
	ap := &AppError{}
	ap.Code = code
	ap.Params = params
	if message == "" {
		ap.Message = code
	} else {
		ap.Message = message
	}
	ap.Where = where
	ap.DetailedError = details
	ap.StatusCode = status
	ap.IsOAuth = false
	ap.cerr = &causeErr{trace: callers(), msg: details}
	return ap
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}

func (er *AppError) AppendParams(params map[string]interface{}) {
	if er.Params == nil {
		er.Params = make(map[string]interface{})
	}
	if params == nil {
		return
	}
	for key, value := range params {
		er.Params[key] = value
	}
}

func IsInternalServiceError(pErr *AppError) bool {
	if pErr == nil {
		return false
	}

	if pErr.StatusCode == InternalServiceError {
		return true
	}

	return false
}

func IsNoResultFoundError(pErr *AppError) bool {
	if pErr == nil {
		return false
	}

	if pErr.StatusCode == NoResultFound {
		return true
	}

	return false
}
