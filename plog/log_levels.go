package plog

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type Params map[string]interface{}

func Fatal(msg string, err error, params Params) {
	parsedError := getAppError(msg, err)
	params = union(params, parsedError.Params)

	if levelFatal >= plogLevel {
		params["timestamp"] = time.Now()
		params["error"] = parsedError
		params["src"] = "o"
		logr.WithFields(logrus.Fields(params)).Fatalln(msg)
	}
}

func ErrorAndSendToSentry(msg string, err error, params Params) {
	if err == nil {
		err = errors.New("Nil error passed from origin")
	}
	parsedError := getAppError(msg, err)
	params = union(params, parsedError.Params)

	errorParams(msg, parsedError, params)
}

func Error(msg string, err error, params Params) {
	if err == nil {
		err = errors.New("Nil error passed from origin")
	}
	parsedError := getAppError(msg, err)
	params = union(params, parsedError.Params)

	errorParams(msg, parsedError, params)
}

func errorParams(msg string, err error, params Params) {
	params["timestamp"] = time.Now()
	params["error"] = err
	params["src"] = "o"
	if levelError >= plogLevel {
		logr.WithFields(logrus.Fields(params)).Errorln(msg)
	}
}

func Warn(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	if levelWarn >= plogLevel {
		params["source"] = "o"
		params["timestamp"] = time.Now()
		logr.WithFields(logrus.Fields(params)).Warnln(msg)
	}
}

func Info(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	params["timestamp"] = time.Now()
	if levelInfo >= plogLevel {
		logr.WithFields(logrus.Fields(params)).Infoln(msg)
	}
}

func BroadcastInfo(tag string, msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	params["timestamp"] = time.Now()
	params["level"] = "info"
	if levelInfo >= plogLevel {
		logr.WithFields(logrus.Fields(params)).Infoln(msg)
	}
}

// Deprecated : Use Info instead
func InfoD(tag string, args ...interface{}) {
	if levelInfo >= plogLevel {
		logr.WithFields(logrus.Fields{
			"src":       "o",
			"args":      fmt.Sprintf("%+v", args),
			"timestamp": time.Now(),
		}).Infoln(tag)
	}
}

func Debug(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	if levelDebug >= plogLevel {
		_, fn, line, _ := runtime.Caller(1)
		params["source"] = "o"
		params["timestamp"] = time.Now()
		params["file"] = fn
		params["line"] = line
		logr.WithFields(logrus.Fields(params)).Debugln(msg)
	}
}

func Trace(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	if levelTrace >= plogLevel {
		_, fn, line, _ := runtime.Caller(1)
		params["level"] = "trace"
		params["source"] = "o"
		params["file"] = fn
		params["line"] = line
		params["timestamp"] = time.Now()
		logr.WithFields(logrus.Fields(params)).Traceln(msg)
	}
}
