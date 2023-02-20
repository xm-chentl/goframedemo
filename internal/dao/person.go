// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/xm-chentl/goframedemo/internal/dao/internal"
)

// internalPersonDao is internal type for wrapping internal DAO implements.
type internalPersonDao = *internal.PersonDao

// personDao is the data access object for table person.
// You can define custom methods on it to extend its functionality as you wish.
type personDao struct {
	internalPersonDao
}

var (
	// Person is globally public accessible object for table person operations.
	Person = personDao{
		internal.NewPersonDao(),
	}
)

// Fill with you ideas below.
