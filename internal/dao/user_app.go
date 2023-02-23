// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/xm-chentl/goframedemo/internal/dao/internal"
)

// internalUserAppDao is internal type for wrapping internal DAO implements.
type internalUserAppDao = *internal.UserAppDao

// userAppDao is the data access object for table user_app.
// You can define custom methods on it to extend its functionality as you wish.
type userAppDao struct {
	internalUserAppDao
}

var (
	// UserApp is globally public accessible object for table user_app operations.
	UserApp = userAppDao{
		internal.NewUserAppDao(),
	}
)

// Fill with you ideas below.