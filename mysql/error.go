package mysql

import (
	"errors"
	"fmt"
)

var (
	ErrBadConn       = errors.New("connection was bad")
	ErrMalformPacket = errors.New("Malform packet error")

	ErrTxDone      = errors.New("sql: Transaction has already been committed or rolled back")
	ErrBadPkgLen   = errors.New("bad packet length")
	ErrPktSync     = errors.New("packet sync error")
	ErrPktSyncMul  = errors.New("packet sync mul error")
	ErrPktTooLarge = errors.New("packet to large")
)

type SqlError struct {
	Code    uint16
	Message string
	State   string
}

func (e *SqlError) Error() string {
	return fmt.Sprintf("ERROR %d (%s): %s", e.Code, e.State, e.Message)
}

//default mysql error, must adapt errname message format
func NewDefaultError(errCode uint16, args ...interface{}) *SqlError {
	e := new(SqlError)
	e.Code = errCode

	if s, ok := MySQLState[errCode]; ok {
		e.State = s
	} else {
		e.State = DEFAULT_MYSQL_STATE
	}

	if format, ok := MySQLErrName[errCode]; ok {
		e.Message = fmt.Sprintf(format, args...)
	} else {
		e.Message = fmt.Sprint(args...)
	}

	return e
}

func NewError(errCode uint16, message string) *SqlError {
	e := new(SqlError)
	e.Code = errCode

	if s, ok := MySQLState[errCode]; ok {
		e.State = s
	} else {
		e.State = DEFAULT_MYSQL_STATE
	}

	e.Message = message

	return e
}
