package exceptions

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var gormErrors = map[error]int{
	gorm.ErrRecordNotFound:                400,
	gorm.ErrInvalidTransaction:            500,
	gorm.ErrNotImplemented:                501,
	gorm.ErrMissingWhereClause:            400,
	gorm.ErrUnsupportedRelation:           400,
	gorm.ErrPrimaryKeyRequired:            400,
	gorm.ErrModelValueRequired:            400,
	gorm.ErrModelAccessibleFieldsRequired: 400,
	gorm.ErrSubQueryRequired:              400,
	gorm.ErrInvalidData:                   400,
	gorm.ErrUnsupportedDriver:             500,
	gorm.ErrRegistered:                    400,
	gorm.ErrInvalidField:                  400,
	gorm.ErrEmptySlice:                    400,
	gorm.ErrDryRunModeUnsupported:         500,
	gorm.ErrInvalidDB:                     500,
	gorm.ErrInvalidValue:                  400,
	gorm.ErrInvalidValueOfLength:          400,
	gorm.ErrPreloadNotAllowed:             400,
	gorm.ErrDuplicatedKey:                 409,
	gorm.ErrForeignKeyViolated:            409,
	gorm.ErrCheckConstraintViolated:       400,
}

func NewDatabaseException(err error) *BaseException {
	for gormError, statusCode := range gormErrors {
		if errors.Is(err, gormError) {
			return &BaseException{
				StatusCode: statusCode,
				Message:    err.Error(),
			}
		}
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		// case 1146: // Table doesn't exist
		// 	return &BaseException{StatusCode: 500, Message: mysqlErr.Message}
		case 1062: // Duplicate entry
			return &BaseException{StatusCode: 409, Message: mysqlErr.Message}
		default:
			return &BaseException{StatusCode: 500, Message: mysqlErr.Error()}
		}
	}

	return &BaseException{StatusCode: 500, Message: err.Error()}
}
