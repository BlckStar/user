package user

type User struct {
	Id       int    `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	Group    string `json:"group" form:"group" query:"group"`
	Fullname string `json:"fullname" form:"fullname" query:"fullname"`
	Home     string `json:"home" form:"home" query:"home"`
	Shell    string `json:"shell" form:"shell" query:"shell"`
	System   bool   `json:"system" form:"system" query:"system"`
}
