package mysql

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

// uniqueViolation is the error number returned by MySQL for unique violation.
const uniqueViolation = 1062

// dataOutOfRange is the error number returned by MySQL due to data being out of range.
const dataOutOfRange = 1264

// dataTooLong is the error number returned by MySQL due to data being too long.
const dataTooLong = 1406

// constraintViolationDelete is the error number returned by MySQL for constraint violation while deleting rows.
const constraintViolationDelete = 1451

// constraintViolationInsertAndUpdate is the error number returned by MySQL for constraint violation while inserting or updating rows.
const constraintViolationInsertAndUpdate = 1452

// ErrDBDuplicateEntry is the error returned when the database has a duplicated entry
var ErrDBDuplicateEntry = errors.New("duplicate entry")

// ErrDBNoRows is the error returned when the database does not find rows.
var ErrDBNoRows = errors.New("duplicate entry")

// ErrDBConflict is the error returned when the database finds a constraint conflict.
var ErrDBConflict = errors.New("conflict between constraints")

// ErrDBValueExceeded is the error returned when the database finds a constraint conflict.
var ErrDBValueExceeded = errors.New("attribute value exceeded")

// CheckError checks if the passed error originates from MySQL, if it does, it parses it into a generic database error for the application.
func CheckError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrDBNoRows
	}

	var mysqlerr *mysql.MySQLError
	ok := errors.As(err, &mysqlerr)

	if ok {
		switch mysqlerr.Number {
		case uniqueViolation:
			return ErrDBDuplicateEntry
		case constraintViolationInsertAndUpdate, constraintViolationDelete:
			return ErrDBConflict
		case dataOutOfRange, dataTooLong:
			return ErrDBValueExceeded
		}
	}

	return err
}
