package plog

import (
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	tag    string
	level  levelType
	logger *logrus.Logger
}

func NewLogger(name string, levelString string) *CustomLogger {
	level, ok := levels[levelString]
	if !ok {
		panic("Invalid level string: " + levelString)
	}
	return &CustomLogger{name, level, logr}
}

func (c *CustomLogger) Fatal(msg string, err error, params Params) {
	parsedError := getAppError(msg, err)
	params = union(params, parsedError.Params)

	if levelFatal >= c.level {
		params["source"] = c.tag
		params["error"] = parsedError
		params["timestamp"] = time.Now()
		c.logger.WithFields(logrus.Fields(params)).Fatalln(msg)
	}
}

func (c *CustomLogger) Error(msg string, err error, params Params) {
	parsedError := getAppError(msg, err)
	params = union(params, parsedError.Params)

	params["timestamp"] = time.Now()
	params["error"] = parsedError
	params["src"] = c.tag
	if levelError >= c.level {
		c.logger.WithFields(logrus.Fields(params)).Errorln(msg)
	}
}

func (c *CustomLogger) Warn(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	params["source"] = c.tag
	params["timestamp"] = time.Now()
	if levelWarn >= c.level {
		c.logger.WithFields(logrus.Fields(params)).Warnln(msg)
	}
}

func (c *CustomLogger) Info(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	params["source"] = c.tag
	params["timestamp"] = time.Now()
	if levelInfo >= c.level {
		c.logger.WithFields(logrus.Fields(params)).Infoln(msg)
	}
}

func (c *CustomLogger) Debug(msg string, params Params) {
	if params == nil {
		params = Params{}
	}
	if levelDebug >= c.level {
		_, fn, line, _ := runtime.Caller(1)
		params["source"] = c.tag
		params["timestamp"] = time.Now()
		params["file"] = fn
		params["line"] = line
		c.logger.WithFields(logrus.Fields(params)).Debugln(msg)
	}
}

func (c *CustomLogger) Trace(msg string, params Params) {
	if params == nil {
		params = Params{}
	}

	if levelTrace >= c.level {
		_, fn, line, _ := runtime.Caller(1)
		params["level"] = "trace"
		params["source"] = c.tag
		params["file"] = fn
		params["line"] = line
		params["timestamp"] = time.Now()
		c.logger.WithFields(logrus.Fields(params)).Debugln(msg)
	}
}
