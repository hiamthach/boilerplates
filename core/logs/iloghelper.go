package core_logs

import (
	"go-microservices/shared"

	"github.com/sirupsen/logrus"
)

type SessionContext struct {
	Id             string `json:"Id,omitempty" bson:"Id,omitempty"`
	UserAgent      string `json:"UserAgent,omitempty" bson:"UserAgent,omitempty"`
	ClientCode     string `json:"ClientCode,omitempty" bson:"ClientCode,omitempty"`
	UserId         string `json:"UserId,omitempty" bson:"UserId,omitempty"`
	UserName       string `json:"UserName,omitempty" bson:"UserName,omitempty"`
	SessionID      string `json:"SessionID,omitempty" bson:"SessionID,omitempty"`
	ClientIP       string `json:"ClientIp,omitempty" bson:"ClientIp,omitempty"`
	AppVersion     string `json:"AppVersion,omitempty" bson:"AppVersion,omitempty"`
	ClientVersion  string `json:"ClientVersion,omitempty" bson:"ClientVersion,omitempty"`
	ClientPlatform string `json:"ClientPlatform,omitempty" bson:"ClientPlatform,omitempty"`
	ClientBundle   string `json:"ClientBundle,omitempty" bson:"ClientBundle,omitempty"`
	ShardID        int32  `json:"ShardID,omitempty" bson:"ShardID,omitempty"`
	IsAdmin        bool   `json:"IsAdmin,omitempty" bson:"IsAdmin,omitempty"`
}

type ILogHelper interface {
	Instance(lg *logrus.Logger)
	GetLogger() *logrus.Logger
	GetEntry(context *SessionContext, logData *shared.AuthSvcLog) *logrus.Entry
	GetLiteEntry(context *SessionContext, action string) *logrus.Entry
	InfoMessage(context *SessionContext, action string, messsage string)
	ErrorMessage(context *SessionContext, messsage string)
	Error(context *SessionContext, err error)
	SysError(err error)
	WarnMessage(context *SessionContext, messsage string)
	WriteLog(lv logrus.Level, context *SessionContext, logData *shared.AuthSvcLog)
	WriteFinishLog(lv logrus.Level, context *SessionContext, logData *shared.AuthSvcLog, message string)
	StartLog(context *SessionContext, action string, isValidError bool) *shared.AuthSvcLog
	Duration(log shared.Action)
}
