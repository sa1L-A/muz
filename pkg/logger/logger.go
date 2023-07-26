package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
)

var (
	worker         *zerolog.Logger
	workerFileName string
)

func init() {
	newLog, err := createLogger()
	if err != nil {
		panic("Init worker failed!")
	}
	worker = newLog
}

// Init Logger
func createLogger() (*zerolog.Logger, error) {

	now := time.Now()
	logFileName := now.Format("2006-01-02") + ".log"
	if workerFileName == logFileName {
		return worker, nil
	}

	var l zerolog.Level

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 1

	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	multi := zerolog.MultiLevelWriter(os.Stdout, src)
	logger := zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &logger, nil
}

// Debug -.
func Debug(message string, args ...interface{}) {
	worker.Debug().Msg(fmt.Sprintf(message, args...))
}

// Info -.
func Info(message string, args ...interface{}) {
	worker.Info().Msg(fmt.Sprintf(message, args...))
}

// Warn -.
func Warn(message string, args ...interface{}) {
	worker.Warn().Msg(fmt.Sprintf(message, args...))
}

// Error -.
func Error(message string, args ...interface{}) {
	if worker.GetLevel() == zerolog.DebugLevel {
		worker.Debug().Msg(fmt.Sprintf(message, args...))
	}
	worker.Error().Msg(fmt.Sprintf(message, args...))
}

// Fatal -.
func Fatal(message string, args ...interface{}) {

	worker.Fatal().Msg(fmt.Sprintf(message, args...))
	os.Exit(1)
}
