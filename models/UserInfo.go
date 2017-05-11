package models

type UserInfo struct {
	Id int64 `json:"Id,string,omitempty"`
	FullName string
	Email string
	Password string
	Address string
	Telephone string
}
