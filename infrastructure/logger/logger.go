package logger

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"sternx/config"
)

type ShortStringer interface {
	ShortString() string
}

var globalInst *logger

type logger struct {
	config *config.Config
	subs   map[string]*SubLogger
	writer io.Writer
}

type SubLogger struct {
	logger zerolog.Logger
	name   string
	obj    fmt.Stringer
}

func getLoggersInst() *logger {
	if globalInst == nil {

		globalInst = &logger{
			subs:   make(map[string]*SubLogger),
			writer: zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"},
		}
		log.Logger = zerolog.New(globalInst.writer).With().Timestamp().Logger()
	}

	return globalInst
}

func InitGlobalLogger(conf *config.Config) {
	if globalInst != nil {
		return
	}

	writers := []io.Writer{}
	// console writer
	writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})

	globalInst = &logger{
		config: conf,
		subs:   make(map[string]*SubLogger),
		writer: io.MultiWriter(writers...),
	}
	log.Logger = zerolog.New(globalInst.writer).With().Timestamp().Logger()

	lvl, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		Warn("invalid default log level", "error", err)
	}
	log.Logger = log.Logger.Level(lvl)
}

func addFields(event *zerolog.Event, keyvals ...interface{}) *zerolog.Event {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "!MISSING-VALUE!")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			key = "!INVALID-KEY!"
		}
		///
		value := keyvals[i+1]
		switch v := value.(type) {
		case fmt.Stringer:
			if isNil(v) {
				event.Any(key, v)
			} else {
				event.Stringer(key, v)
			}
		case ShortStringer:
			if isNil(v) {
				event.Any(key, v)
			} else {
				event.Str(key, v.ShortString())
			}
		case error:
			event.AnErr(key, v)
		case []byte:
			event.Str(key, hex.EncodeToString(v))
		default:
			event.Any(key, v)
		}
	}

	return event
}

func NewSubLogger(name string, obj fmt.Stringer) *SubLogger {
	inst := getLoggersInst()
	sl := &SubLogger{
		logger: zerolog.New(inst.writer).With().Timestamp().Logger(),
		name:   name,
		obj:    obj,
	}

	lvlStr := inst.config.LogLevel

	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		Warn("invalid log level", "error", err, "name", name)
	}
	sl.logger = sl.logger.Level(lvl)

	inst.subs[name] = sl

	return sl
}

func (sl *SubLogger) logObj(event *zerolog.Event, msg string, keyvals ...interface{}) {
	if sl.obj != nil {
		event = event.Str(sl.name, sl.obj.String())
	}

	addFields(event, keyvals...).Msg(msg)
}

func (sl *SubLogger) Trace(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Trace(), msg, keyvals...)
}

func (sl *SubLogger) Debug(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Debug(), msg, keyvals...)
}

func (sl *SubLogger) Info(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Info(), msg, keyvals...)
}

func (sl *SubLogger) Warn(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Warn(), msg, keyvals...)
}

func (sl *SubLogger) Error(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Error(), msg, keyvals...)
}

func (sl *SubLogger) Fatal(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Fatal(), msg, keyvals...)
}

func (sl *SubLogger) Panic(msg string, keyvals ...interface{}) {
	sl.logObj(sl.logger.Panic(), msg, keyvals...)
}

func Trace(msg string, keyvals ...interface{}) {
	addFields(log.Trace(), keyvals...).Msg(msg)
}

func Debug(msg string, keyvals ...interface{}) {
	addFields(log.Debug(), keyvals...).Msg(msg)
}

func Info(msg string, keyvals ...interface{}) {
	addFields(log.Info(), keyvals...).Msg(msg)
}

func Warn(msg string, keyvals ...interface{}) {
	addFields(log.Warn(), keyvals...).Msg(msg)
}

func Error(msg string, keyvals ...interface{}) {
	addFields(log.Error(), keyvals...).Msg(msg)
}

func Fatal(msg string, keyvals ...interface{}) {
	addFields(log.Fatal(), keyvals...).Msg(msg)
}

func Panic(msg string, keyvals ...interface{}) {
	addFields(log.Panic(), keyvals...).Msg(msg)
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return reflect.ValueOf(i).IsNil()
	}

	return false
}

func (sl *SubLogger) SetObj(obj fmt.Stringer) {
	sl.obj = obj
}
