package core_logs

import (
	"go-microservices/app/config"
	"go-microservices/shared"
	"os"
	"sync"
	"time"

	logrus "github.com/sirupsen/logrus"
)

type LogHelper struct {
	ILogHelper
	logger *logrus.Logger
}

var logHelper *LogHelper
var logHelperOnce sync.Once

func (loghelper *LogHelper) Instance(lg *logrus.Logger) {
	loghelper.logger = lg
}

func GetLogHelper() *LogHelper {
	logHelperOnce.Do(func() {
		Applog := logrus.New()
		Applog.SetFormatter(&logrus.JSONFormatter{})
		Applog.Out = os.Stdout
		Applog.SetReportCaller(false)
		minLevel := config.Get().Log.MinLevel
		Applog.SetLevel(logrus.Level(minLevel))
		Applog.WithField("AppName", config.Get().App.Name)

		logHelper = &LogHelper{}
		logHelper.Instance(Applog)
	})
	return logHelper
}
func (loghelper *LogHelper) GetLogger() *logrus.Logger {
	return logHelper.logger
}
func (loghelper *LogHelper) InfoMessage(context *SessionContext, action string, messsage string) {
	if context == nil {
		context = &SessionContext{}
	}
	logData := shared.AuthSvcLog{
		Log: shared.Action{
			Action:    action,
			Message:   messsage,
			RequestId: context.Id,
		},
	}
	loghelper.WriteLog(logrus.InfoLevel, context, &logData)
}

func (loghelper *LogHelper) SysError(err error) {
	logData := shared.AuthSvcLog{
		Log: shared.Action{
			Message: err.Error(),
		},
	}
	loghelper.WriteLog(logrus.ErrorLevel, nil, &logData)

}

func (loghelper *LogHelper) Error(context *SessionContext, err error) {
	loghelper.ErrorMessage(context, err.Error())
}

func (loghelper *LogHelper) ErrorMessage(context *SessionContext, messsage string) {
	logData := shared.AuthSvcLog{
		Log: shared.Action{
			Message:   messsage,
			RequestId: context.Id,
		},
	}
	loghelper.WriteLog(logrus.ErrorLevel, context, &logData)
}
func (loghelper *LogHelper) WarnMessage(context *SessionContext, messsage string) {
	logData := shared.AuthSvcLog{
		Log: shared.Action{
			Message:   messsage,
			RequestId: context.Id,
		},
	}
	loghelper.WriteLog(logrus.WarnLevel, context, &logData)
}
func (loghelper *LogHelper) WriteFinishLog(lv logrus.Level, context *SessionContext, logData *shared.AuthSvcLog, message string) {
	if len(message) == 0 {
		message = "done"
	}
	logData.Log.Message = message
	logData.Log.Finish()
	loghelper.WriteLog(lv, context, logData)
}

func (loghelper *LogHelper) GetLiteEntry(context *SessionContext, action string) *logrus.Entry {
	entry := loghelper.logger.WithFields(logrus.Fields{
		"Action":  action,
		"StartAt": time.Now().UTC(),
	})
	if context == nil {
		return entry
	}
	if len(context.ClientCode) > 0 {
		entry = entry.WithField("ClientCode", context.ClientCode)
	}
	if len(context.SessionID) > 0 {
		entry = entry.WithField("SessionID", context.SessionID)
	}
	if len(context.ClientVersion) > 0 {
		entry = entry.WithField("ClientVersion", context.ClientVersion)
	}
	if len(context.UserId) > 0 {
		entry = entry.WithField("UserId", context.UserId)
	}
	if len(context.AppVersion) > 0 {
		entry = entry.WithField("AppVersion", context.AppVersion)
	}

	return entry
}

func (loghelper *LogHelper) GetEntry(context *SessionContext, logData *shared.AuthSvcLog) *logrus.Entry {
	entry := loghelper.logger.WithFields(logrus.Fields{
		"Duration":     logData.Log.Duration,
		"IsValidError": logData.Log.IsValidError,
		"StartAt":      logData.Log.StartAt,
	})
	if context == nil {
		return entry
	}
	if len(context.ClientCode) > 0 {
		entry = entry.WithField("ClientCode", context.ClientCode)
	}
	if len(context.SessionID) > 0 {
		entry = entry.WithField("SessionID", context.SessionID)
	}
	if len(context.ClientVersion) > 0 {
		entry = entry.WithField("ClientVersion", context.ClientVersion)
	}
	if len(context.UserId) > 0 {
		entry = entry.WithField("UserId", context.UserId)
	}
	if len(context.AppVersion) > 0 {
		entry = entry.WithField("AppVersion", context.AppVersion)
	}
	if len(logData.Log.RequestId) > 0 {
		entry = entry.WithField("RequestId", logData.Log.RequestId)
	}
	if len(logData.Log.Action) > 0 {
		entry = entry.WithField("Action", logData.Log.Action)
	}
	if len(logData.SubActions) > 0 {
		entry = entry.WithField("SubLogs", logData.SubActions)
	}
	if len(logData.Log.TraceInfo) > 0 {
		entry = entry.WithField("TraceInfo", logData.Log.TraceInfo)
	}
	return entry
}

func (loghelper *LogHelper) WriteLog(lv logrus.Level, context *SessionContext, logData *shared.AuthSvcLog) {
	entry := loghelper.logger.WithFields(logrus.Fields{
		"Duration":     logData.Log.Duration,
		"IsValidError": logData.Log.IsValidError,
		"StartAt":      logData.Log.StartAt,
	})

	if context != nil {
		if len(context.ClientCode) > 0 {
			entry = entry.WithField("ClientCode", context.ClientCode)
		}
		if len(context.SessionID) > 0 {
			entry = entry.WithField("SessionID", context.SessionID)
		}
		if len(context.ClientVersion) > 0 {
			entry = entry.WithField("ClientVersion", context.ClientVersion)
		}
		if len(context.UserId) > 0 {
			entry = entry.WithField("UserId", context.UserId)
		}
		if len(context.AppVersion) > 0 {
			entry = entry.WithField("AppVersion", context.AppVersion)
		}
		if len(logData.Log.RequestId) > 0 {
			entry = entry.WithField("RequestId", logData.Log.RequestId)
		}
		if len(logData.Log.Action) > 0 {
			entry = entry.WithField("Action", logData.Log.Action)
		}
		if len(logData.SubActions) > 0 {
			entry = entry.WithField("SubLogs", logData.SubActions)
		}
		if len(logData.Log.TraceInfo) > 0 {
			entry = entry.WithField("TraceInfo", logData.Log.TraceInfo)
		}
	}

	entry.Log(lv, logData.Log.Message)
}

func (loghelper *LogHelper) StartLog(context *SessionContext, action string, isValidError bool) *shared.AuthSvcLog {

	logdata := shared.AuthSvcLog{
		Log: shared.Action{
			StartAt:      time.Now().UTC(),
			Action:       action,
			RequestId:    context.Id,
			IsValidError: isValidError,
		},
	}
	return &logdata
}
