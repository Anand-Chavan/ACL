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

type Groups struct {
	GroupName        string `json:"groupName" column:"groupName"`
	UserId           string `json:"userId" column:"userId"`
	SessionKey       string `json:"sessionKey" column:"sessionKey"`
	GroupDescription string `json:"groupDescription" column:"groupDescription"`
}

func (grp *Groups) Table() string {
	return "groups"
}

func (grp *Groups) String() string {
	return Stringify(grp)
}

type UserAddToGroup struct {
	GroupName  string `json:"groupName" column:"groupName"`
	UserId     string `json:"userId" column:"userId"`
	SessionKey string `json:"sessionKey" column:"sessionKey"`
}

func (uatg *UserAddToGroup) Table() string {
	return "userGroupMap"
}

func (uatg *UserAddToGroup) String() string {
	return Stringify(uatg)
}

var (
	RootDir string = "/home/sagar/gorootdir"
)

type CreateFileOrFolder struct {
	FilefolderPath  string `json:"filefolderPath" column:"filefolderPath"`
	FilefolderName  string `json:"filefolderName" column:"filefolderName"`
	FilesOrFolderId string `json:"filesOrFolderId" column:"filesOrFolderId"`
	SessionKey      string `json:"sessionKey" column:"sessionKey"`
	UserId          string `json:"userId" column:"userId"`
}

func (filecrt *CreateFileOrFolder) Table() string {
	return "filesfolder"
}

func (filecrt *CreateFileOrFolder) String() string {
	return Stringify(filecrt)
}

type ScanData struct {
	WritePermissionusr int `json:"id"`
	WritePermissiongrp int `json:"id"`
}

type GetFilesFold struct {
	FilefolderPath string `json:"filefolderPath" column:"filefolderPath"`
	SessionKey     string `json:"sessionKey" column:"sessionKey"`
	FilefolderName string `json:"filefolderName" column:"filefolderName"`
	UserId         string `json:"userId" column:"userId"`
}

func (getfile *GetFilesFold) Table() string {
	return "filesfolder"
}

func (getfile *GetFilesFold) String() string {
	return Stringify(getfile)
}

type SelFileFold struct {
	FilefolderPath string `json:"filefolderPath" column:"filefolderPath"`
	FilefolderName string `json:"filefolderName" column:"filefolderName"`
	UserType       string `json:"userType" column:"userType"`
}

type ChangePermission struct {
	FilefolderPath  string `json:"filefolderPath" column:"filefolderPath"`
	FilefolderName  string `json:"filefolderName" column:"filefolderName"`
	UseridOrGroupId string `json:"useridOrGroupId" column:"useridOrGroupId"`
	FilesOrFolderId string `json:"filesOrFolderId" column:"filesOrFolderId"`
	PermissionValue string `json:"permissionValue" column:"permissionValue"`
	WhocallToChange string `json:"whocallToChange" column:"whocallToChange"`
	SessionKey      string `json:"sessionKey" column:"sessionKey"`
}

func (getfile *ChangePermission) Table() string {
	return "filesfolder"
}

func (getfile *ChangePermission) String() string {
	return Stringify(getfile)
}

// type ReadFile struct {
// 	FilefolderPath string `json:"filefolderPath" column:"filefolderPath"`
// 	FilefolderName string `json:"filefolderName" column:"filefolderName"`
// 	UserId         string `json:"userId" column:"userId"`
// }

type NotPermit struct {
	Msg string
}
