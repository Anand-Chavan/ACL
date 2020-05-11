package model

import(
	"fmt"
)

type User struct {
	Name         string `json:"Name" column:"Name"`
	Id            int64  `json:"Id,omitempty" key:"primary" autoincr:"1" column:"Id"`
	Password      string `json:"Password" column:"Password"`
	DateCreation  string `json:"DateCreation,omitempty" key:"primary" autoincr:"1" column:"DateCreation"`
	UserType      string `json:"UserType" column:"UserType"`
}

func (user *User) Table() string {
	return "users"
}

func (user *User) String() string {
	return Stringify(user)
}


type User struct {
	Name         string `json:"Name" column:"Name"`
	Id            int64  `json:"Id,omitempty" key:"primary" autoincr:"1" column:"Id"`
	Password      string `json:"Password" column:"Password"`
	DateCreation  string `json:"DateCreation,omitempty" key:"primary" autoincr:"1" column:"DateCreation"`
	UserType      string `json:"UserType" column:"UserType"`
}

func (user *User) Table() string {
	return "users"
}

func (user *User) String() string {
	return Stringify(user)
}

