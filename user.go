package user;

type User struct {
	id int `json:"id" form:"id" query:"id"`
	name string `json:"name" form:"name" query:"name"`
	group string `json:"group" form:"group" query:"group"`
	fullname string `json:"fullname" form:"fullname" query:"fullname"`
	home string `json:"home" form:"home" query:"home"`
	shell string `json:"shell" form:"shell" query:"shell"`
	system bool `json:"system" form:"system" query:"system"`
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u User) Id() int {
	return u.id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u User) Name() string {
	return u.name
}

func (u *User) SetGroup(group string) {
	u.group = group
}

func (u User) Group() string {
	return u.group
}

func (u *User) SetFullname(fullname string) {
	u.fullname = fullname
}

func (u User) Fullname() string {
	return u.fullname
}

func (u *User) SetHome(home string) {
	u.home = home
}

func (u User) Home() string {
	if u.home == "" {
		u.home = "/home/" + u.Name()
	}
	return u.home
}

func (u *User) SetShell(shell string) {
	u.shell = shell
}

func (u User) Shell() string {
	if u.shell == "" {
		u.shell = "/bin/false"
	}
	return u.shell
}

func (u *User) SetSystem(system bool) {
	u.system = system
}

func (u User) System() bool {
	return u.system
}
