package logger

import (
    "os"
    "path/filepath"
    "runtime"
    "sync"

    log "github.com/sirupsen/logrus"
)

type Level uint32

func (level Level) String() string {
    switch level {
    case DebugLevel:
        return "debug"
    case InfoLevel:
        return "info"
    case WarnLevel:
        return "warning"
    case ErrorLevel:
        return "error"
    case FatalLevel:
        return "fatal"
    case PanicLevel:
        return "panic"
    }

    return "unknown"
}

const (
    PanicLevel Level = iota
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
)

type singleton struct {
    AppName      string
    Level        log.Level
    NativeLogger bool
}

type LoggerType int

const (
    TextLogger LoggerType = iota
    JSONLogger
)

const (
    timeFormat = "2006-01-02 15:04:05.000-0700"
)

var (
    formatterText = &log.TextFormatter{}
    formatterJSON = &log.JSONFormatter{}
)

var instance *singleton
var once sync.Once

func (m *singleton) NewText() {
    formatterText.DisableSorting = false
    formatterText.FullTimestamp = true
    formatterText.DisableColors = true
    formatterText.TimestampFormat = timeFormat
    log.SetFormatter(formatterText)

}

func (m *singleton) NewJSON() {
    formatterJSON.TimestampFormat = timeFormat
    log.SetFormatter(formatterJSON)
}

func (m *singleton) GetLogger() *log.Entry {
    return log.WithFields(
        log.Fields{
            "appName": m.AppName,
        },
    )
}

func New(appName string, level Level, t LoggerType) {
    once.Do(func() {
        level, _ := log.ParseLevel(level.String())
        instance = &singleton{AppName: appName, Level: level}
        log.SetOutput(os.Stdout)
        log.SetLevel(instance.Level)
        if t == TextLogger {
            instance.NewText()
        } else {
            instance.NewJSON()
        }
    })
}

func getLogger() *log.Entry {
    _, path, line, ok := runtime.Caller(2)
    if ok {
        fields := make(map[string]interface{}, 2)
        _, file := filepath.Split(path)
        fields["file"] = file
        fields["line"] = line
        return instance.GetLogger().WithFields(fields)
    }
    return instance.GetLogger()
}

func Debug(args ...interface{}) {
    getLogger().Debug(args...)
}

func Print(args ...interface{}) {
    getLogger().Print(args...)
}

func Info(args ...interface{}) {
    getLogger().Info(args...)
}

func Warn(args ...interface{}) {
    getLogger().Warn(args...)
}

func Warning(args ...interface{}) {
    getLogger().Warning(args...)
}

func Error(args ...interface{}) {
    getLogger().Error(args...)
}

func Fatal(args ...interface{}) {
    getLogger().Fatal(args...)
}

func Panic(args ...interface{}) {
    getLogger().Panic(args...)
}

func Debugf(format string, args ...interface{}) {
    getLogger().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
    getLogger().Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
    getLogger().Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
    getLogger().Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
    getLogger().Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
    getLogger().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
    getLogger().Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
    getLogger().Panicf(format, args...)
}

func Debugln(args ...interface{}) {
    getLogger().Debugln(args...)
}

func Infoln(args ...interface{}) {
    getLogger().Infoln(args...)
}

func Println(args ...interface{}) {
    getLogger().Println(args...)
}

func Warnln(args ...interface{}) {
    getLogger().Warnln(args...)
}

func Warningln(args ...interface{}) {
    getLogger().Warningln(args...)
}

func Errorln(args ...interface{}) {
    getLogger().Errorln(args...)
}

func Fatalln(args ...interface{}) {
    getLogger().Fatalln(args...)
}

func Panicln(args ...interface{}) {
    getLogger().Panicln(args...)
}
