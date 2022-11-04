package plog

import (
	"crypto_price_tracker/config"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/getsentry/raven-go"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

type levelType uint

var enabled bool

const (
	levelNone levelType = iota + 1
	levelTrace
	levelDebug
	levelInfo
	levelWarn
	levelError
	levelFatal
	levelAll
	TimeStampFormat = "2006-01-02T15:04:05.000-07:00"
)

var levels = map[string]levelType{
	"NONE":  levelNone,
	"TRACE": levelTrace,
	"DEBUG": levelDebug,
	"INFO":  levelInfo,
	"WARN":  levelWarn,
	"ERROR": levelError,
	"FATAL": levelFatal,
	"ALL":   levelAll,
}

const (
	STR_TYPE_FILE = "FILE"
)

var plogLevel levelType
var logr *logrus.Logger
var path string
var logWriter string

func init() {
	logr = logrus.New()
}

func Init(logType string, enable bool) {
	enabled = enable

	path, _ = filepath.Abs(config.PLOG_LOCATION.Get())

	logWriter = logType
	logTypeStr := strings.ToUpper(logWriter)
	setLogger(logr, path, logTypeStr)

	raven.SetDSN(config.SENTRY_DSN_KEY.Get())
	logLevelStr := strings.ToUpper(config.PLOG_LEVEL.Get())

	plogLevel = levels[logLevelStr]
	if plogLevel == 0 {
		plogLevel = getLevelFromEnvironment()
	}
}

func setLogger(logger *logrus.Logger, logPath string, logType string) {
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
		TimestampFormat:  TimeStampFormat,
	}
	logger.Level = logrus.DebugLevel
	setLogOutput(logger, logPath, logType)
}

func setLogOutput(logger *logrus.Logger, logPath string, logTypeStr string) {

	switch logTypeStr {
	case STR_TYPE_FILE:
		_ = os.Mkdir(logPath, os.ModePerm)
		setFileIO(logger, logPath)
		log_file_scheduler := gocron.NewScheduler()
		log_file_scheduler.Every(1).Day().At("00.00").Do(setFileIO, logger, logPath)
		log_file_scheduler.Start()
	default:
		logger.Out = os.Stdout
	}
}

func getLevelFromEnvironment() levelType {
	if config.IsDevelopment() {
		return levelTrace
	} else if config.IsStaging() {
		return levelInfo
	} else if config.IsProduction() {
		return levelError // Or we can put levelNone
	}

	return levelNone
}

func getFileName() string {
	y, m, d := time.Now().Date()
	dateString := fmt.Sprintf("%d_%d_%d", y, m, d)
	return "log_" + dateString + ".txt"
}

func setFileIO(logger *logrus.Logger, logPath string) {
	file_location := logPath + "/" + getFileName()
	createFile(file_location)
	logger.Out, _ = os.OpenFile(file_location, os.O_RDWR|os.O_APPEND, 0660)
}

func createFile(file_location string) {

	// detect if file exists
	var _, err = os.Stat(file_location)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, _ = os.Create(file_location)
		defer file.Close()
	}
}
