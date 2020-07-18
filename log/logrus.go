package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type LogrusLogger struct {
	log    logrus.Ext1FieldLogger
	prefix string
	fields Fields
}

func NewLogrusLogger(logrus logrus.Ext1FieldLogger, prefix string, fields Fields) *LogrusLogger {
	return &LogrusLogger{
		log:    logrus,
		prefix: prefix,
		fields: fields,
	}
}

func NewDefaultLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	logger.Formatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	}

	return NewLogrusLogger(logger, "main", nil)
}

func (l *LogrusLogger) Print(args ...interface{}) {
 }

func (l *LogrusLogger) Printf(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Trace(args ...interface{}) {
 }

func (l *LogrusLogger) Tracef(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Debug(args ...interface{}) {
 }

func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Info(args ...interface{}) {
 }

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Warn(args ...interface{}) {
 }

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Error(args ...interface{}) {
 }

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
 }

func (l *LogrusLogger) Fatal(args ...interface{}) {
 }

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.prepareEntry().Fatalf(format, args...)
}

func (l *LogrusLogger) Panic(args ...interface{}) {
 }

func (l *LogrusLogger) Panicf(format string, args ...interface{}) {
 }

func (l *LogrusLogger) WithPrefix(prefix string) Logger {
	return NewLogrusLogger(l.log, prefix, l.Fields())
}

func (l *LogrusLogger) Prefix() string {
	return l.prefix
}

func (l *LogrusLogger) WithFields(fields Fields) Logger {
	return NewLogrusLogger(l.log, l.Prefix(), l.Fields().WithFields(fields))
}

func (l *LogrusLogger) Fields() Fields {
	return l.fields
}

func (l *LogrusLogger) prepareEntry() *logrus.Entry {
	return l.log.
		WithFields(logrus.Fields(l.Fields())).
		WithField("prefix", l.Prefix())
}
