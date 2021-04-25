package dao

import (
	_ "database/sql"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("dao: record not found")
)

func Dao() error {
	return errors.Wrapf(ErrNotFound, "dao options masage")
}
