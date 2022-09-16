package errorenum

import "todoapps/shared/model/apperror"

const (
	SomethingError                apperror.ErrorType = "ER0000 something error"
	MessageMustNotEmpty           apperror.ErrorType = "ER0001 message must not empty"
	Random6CharLengthNotSatisfied apperror.ErrorType = "ER0002 random6 char length not satisfied"
	CannotCheckTheDoneTodoItem    apperror.ErrorType = "ER0003 cannot check the done todo item"
)
