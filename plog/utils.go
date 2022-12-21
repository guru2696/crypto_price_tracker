package plog

import (
	"errors"
	"net/http"

	"crypto_price_tracker/plerrors"
	"crypto_price_tracker/plog/log_tags"
)

type message struct {
	Tag    log_tags.Tag `json:"Tag"`
	Params interface{}  `json:"params,omitempty"`
}

var MP = MessageWithParam
var M = Message

// deprecated
func MessageWithParam(tag *log_tags.Tag, param interface{}) message {
	return message{
		Tag:    *tag,
		Params: param,
	}
}

// deprecated
func Message(param interface{}) message {
	return message{
		Tag:    *log_tags.MESSAGE,
		Params: param,
	}
}

func getAppError(msg string, err error) (parsedError *plerrors.AppError) {
	if err == nil {
		err = errors.New("Nil error passed from origin")
	}
	switch err.(type) {
	case *plerrors.AppError:
		parsedError = err.(*plerrors.AppError)
		if parsedError == nil {
			parsedError = plerrors.NewAppError("", "system_error", msg, http.StatusInternalServerError, "Nil error passed from origin", nil)
		}
	case *plerrors.ServiceError:
		parsedError = plerrors.NewAppError("", "service_error", msg, http.StatusBadRequest, err.Error(), nil)
	default:
		parsedError = plerrors.NewAppError("", "system_error", msg, http.StatusInternalServerError, err.Error(), nil)
	}
	return parsedError
}

func union(primaryMap Params, secondaryMap Params) Params {
	if primaryMap == nil {
		primaryMap = Params{}
	}
	if secondaryMap == nil {
		return primaryMap
	}
	for key, value := range secondaryMap {
		primaryMap[key] = value
	}
	return primaryMap
}
