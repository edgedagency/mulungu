package logger

import (
	"strings"

	"golang.org/x/net/context"

	gLog "google.golang.org/appengine/log"
)

//Debugf sends logs to google appengine logging
func Debugf(ctx context.Context, tag, format string, args ...interface{}) {
	gLog.Debugf(ctx, "["+strings.ToUpper(tag)+"] "+format, args)
}

//Errorf sends logs to google appengine logging
func Errorf(ctx context.Context, tag, format string, args ...interface{}) {
	gLog.Errorf(ctx, "["+strings.ToUpper(tag)+"] "+format, args)
}

//Infof sends logs to google appengine logging
func Infof(ctx context.Context, tag, format string, args ...interface{}) {
	gLog.Infof(ctx, "["+strings.ToUpper(tag)+"] "+format, args)
}

//Criticalf sends logs to google appengine logging
func Criticalf(ctx context.Context, tag, format string, args ...interface{}) {
	gLog.Criticalf(ctx, "["+strings.ToUpper(tag)+"] "+format, args)
}

//Warningf sends logs to google appengine logging
func Warningf(ctx context.Context, tag, format string, args ...interface{}) {
	gLog.Warningf(ctx, "["+strings.ToUpper(tag)+"] "+format, args)
}
