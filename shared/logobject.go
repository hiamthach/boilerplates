package shared

import "time"

type Action struct {
	RequestId    string    `json:"RequestId,omitempty" bson:"RequestId,omitempty"`
	Duration     int64     `json:"Duration" bson:"Duration"`
	Message      string    `json:"Message,omitempty" bson:"Message,omitempty"`
	Action       string    `json:"Action,omitempty" bson:"Action,omitempty"`
	IsValidError bool      `json:"IsValidError,omitempty" bson:"IsValidError,omitempty"`
	Error        error     `json:"Error,omitempty" bson:"Error,omitempty"`
	StartAt      time.Time `json:"StartAt,omitempty" bson:"StartAt,omitempty"`
	TraceInfo    string    `json:"TraceInfo,omitempty" bson:"TraceInfo,omitempty"`
}

type SysError struct {
	RequestId string `json:"RequestId,omitempty" bson:"RequestId,omitempty"`
	Message   error  `json:"Message,omitempty" bson:"Message,omitempty"`
	Reporter  string `json:"Reporter,omitempty" bson:"Reporter,omitempty"`
}

type AuthSvcLog struct {
	Log        Action        `json:"LogObj,omitempty" bson:"LogObj,omitempty"`
	SubActions []LightAction `json:"SubActions,omitempty" bson:"SubActions,omitempty"`
}

type LightAction struct {
	Duration int64     `json:"Duration" bson:"Duration"`
	Action   string    `json:"Action,omitempty" bson:"Action,omitempty"`
	Message  string    `json:"Message,omitempty" bson:"Message,omitempty"`
	startAt  time.Time //`json:"StartAt,omitempty" bson:"StartAt,omitempty"`
}

func (action *Action) Finish() {
	action.Duration = time.Since(action.StartAt).Milliseconds()
}

func (action *LightAction) Finish() {
	action.Duration = time.Since(action.startAt).Milliseconds()
}

func (action *Action) Start() {
	action.StartAt = time.Now().UTC()
}

func (action *LightAction) Start() {
	action.startAt = time.Now().UTC()
}

func (logObjs *AuthSvcLog) AddSubAction(action string, startTime time.Time) {
	log := LightAction{
		startAt: startTime,
		Action:  action,
	}
	log.Duration = time.Since(log.startAt).Milliseconds()
	logObjs.SubActions = append(logObjs.SubActions, log)
}
