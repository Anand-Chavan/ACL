package model

type Acl struct {
	// Id       int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	UserName string `json:"uName" column:"uName"`
	UserId   string `json:"userId" column:"userId"`
	Password string `json:"password" column:"password"`
	UserType string `json:"userType" column:"userType"`
}

func (acl *Acl) Table() string {
	return "users"
}

func (acl *Acl) String() string {
	return Stringify(acl)
}

type Auth struct {
	UserId   string `json:"userId" column:"userId"`
	Password string `json:"password" column:"password"`
}

func (acl *Auth) Table() string {
	return "users"
}

func (acl *Auth) String() string {
	return Stringify(acl)
}

type Key struct {
	UserId     string `json:"userId" column:"userId"`
	SessionKey string `json:"sessionKey" column:"sessionKey"`
}

func (key *Key) Table() string {
	return "usersKey"
}

func (key *Key) String() string {
	return Stringify(key)
}
