package shared

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/status"
)

type CustomErr struct {
	error
	Code    int32
	Message string
}

func (e *errorCodes) InternalErr(f string, a ...interface{}) error {
	return status.Errorf(http.StatusInternalServerError, fmt.Sprintf(f, a...))
}
func (e *errorCodes) BadRequestErr(f string, a ...interface{}) error {
	return status.Errorf(http.StatusBadRequest, fmt.Sprintf(f, a...))
}

type errorCodes struct {
	OK                  CustomErr
	Unauthorized        CustomErr
	UnKnownError        CustomErr
	NoContent           CustomErr
	DataExisted         CustomErr
	NetworkError        CustomErr
	ParamError          CustomErr
	RequestExpired      CustomErr
	InvalidClientHeader CustomErr
	InvalidUserIdHeader CustomErr
}

var ErrorCodes = &errorCodes{
	OK:                  CustomErr{Code: 200, Message: "success"},
	Unauthorized:        CustomErr{Code: 401, Message: "unauthorized"},
	NoContent:           CustomErr{Code: 400, Message: "no content"},
	InvalidClientHeader: CustomErr{Code: 422, Message: "invalid client header"},
	InvalidUserIdHeader: CustomErr{Code: 421, Message: "invalid user id"},
	ParamError:          CustomErr{Code: 423, Message: "param error"},
	RequestExpired:      CustomErr{Code: 424, Message: "request has expired"},
	NetworkError:        CustomErr{Code: 500, Message: "cannot create account missing Param "},
	UnKnownError:        CustomErr{Code: 499, Message: "unknown error"},
}

func (e *errorCodes) IsUnknown(err error) bool {
	return e.GetCode(err) == 499
}

func (e *errorCodes) Create(code int32, message string) CustomErr {
	r := CustomErr{
		Code:    code,
		Message: message,
	}
	return r
}
func (e *errorCodes) CreateInValidParam(paramName string) CustomErr {
	r := CustomErr{
		Code:    456,
		Message: fmt.Sprintf("%s %s", "invalid param ", paramName),
	}
	return r
}

func (e *errorCodes) GetCode(err error) int32 {
	state, _ := status.FromError(err)
	return int32(state.Code())
}
