package log

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type contextType string

const logGroupIDConst = "loggroupid"
const logGroupIDField = contextType(logGroupIDConst)

var (
	singleton     sync.Once
	defaultLogger logrusImpl
	useFile       bool
	path          string
	name          string
	maxAge        int
	format        logrus.Formatter
)

type logrusImpl struct {
	theLogger *logrus.Logger
}

func setFile(pathFile, nameFile string, maxAgeInDays int) {
	path = pathFile
	name = nameFile
	maxAge = maxAgeInDays
	useFile = true
}

func getLogImpl() Logger {
	singleton.Do(func() {

		format = &nested.Formatter{
			FieldsOrder:      []string{"opid", "func"},
			TimestampFormat:  time.RFC3339,
			HideKeys:         true,
			NoColors:         true,
			NoFieldsColors:   false,
			NoFieldsSpace:    true,
			ShowFullLevel:    true,
			NoUppercaseLevel: false,
			TrimMessages:     true,
		}

		defaultLogger = logrusImpl{theLogger: logrus.New()}
		defaultLogger.theLogger.SetFormatter(format)

		if viper.GetBool("production") {
			defaultLogger.theLogger.Out = ioutil.Discard
		}

		if !useFile {
			return
		}

		writer, _ := rotatelogs.New(
			fmt.Sprintf("%s/logs/%s.log.%s", path, name, "%Y%m%d"),
			rotatelogs.WithLinkName(fmt.Sprintf("%s/logs/%s.log", path, name)),
			rotatelogs.WithMaxAge(time.Duration(maxAge*24)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(1*24)*time.Hour),
		)

		defaultLogger.theLogger.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  writer,
				logrus.ErrorLevel: writer,
			},
			defaultLogger.theLogger.Formatter,
		))

	})

	return &defaultLogger
}

func (x *logrusImpl) Info(ctx context.Context, message string, args ...interface{}) {
	logMessage := fmt.Sprintf(message, args...)
	x.theLogger.Info(logMessage)
}

func (x *logrusImpl) Error(ctx context.Context, message string, args ...interface{}) {
	logMessage := fmt.Sprintf(message, args...)
	x.theLogger.Error(logMessage)
}
