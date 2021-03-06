package user

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"time"
)

func AddUser(u *User) error {

	return errors.New("not yet implemented")

	var err error

	if u.Name == "" {
		return errors.New("at least a name must be provided")
	}

	var _, notFoundError = user.Lookup(u.Name)
	if notFoundError == nil {
		return errors.New("user " + u.Name + " does already exist")
	}

	var generatedId, nextUserErr = GetNextUserId(u.System)

	if nextUserErr != nil {
		return nextUserErr
	}

	u.Id = generatedId

	if u.Group == "" {
		u.Group = u.Name
	}

	return err
}

func saveToPasswd(u *User) error {
	var userString = fmt.Sprintf("%s:x:%d:%d:%s:%s:%s\n", u.Name, u.Id, u.Id, u.Fullname, u.Home, u.Shell)

	var passwd, passwdErr = os.OpenFile("/etc/passwd", os.O_APPEND|os.O_WRONLY, 0600)
	if passwdErr != nil {
		return passwdErr
	}

	var _, writeErr = passwd.WriteString(userString)

	passwd.Close()

	return writeErr
}

func saveToShadow(u *User) error {
	var unixTs = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	var doc = strconv.Itoa(int(time.Now().Sub(unixTs).Hours() / 24))
	var userString = fmt.Sprintf("%s:*:%s:0:99999:7:::\n", u.Name, doc)

	var passwd, passwdErr = os.OpenFile("/etc/passwd", os.O_APPEND|os.O_WRONLY, 0600)
	if passwdErr != nil {
		return passwdErr
	}

	var _, writeErr = passwd.WriteString(userString)

	passwd.Close()

	return writeErr
}
