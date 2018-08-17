package xsql

import "database/sql"

// ErrNoRowsTo returns 'to' error if err is type of sql.ErrNoRows, err otherwise.
func ErrNoRowsTo(err error, to error) error {
	if err == sql.ErrNoRows {
		return to
	}
	return err
}
