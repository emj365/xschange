package models

// User data
type User struct {
	GoodAmount uint
	Balance    uint
}

// UserList contain users
type UserList []*User
