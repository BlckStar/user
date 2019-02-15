package user

import (
	"bufio"
	"errors"
	"os"
	osuser "os/user"
	"strconv"
	"strings"
)

func GetUser(name string) (User, error) {
	var users = GetAllUsers()
	var returnUser User

	for i := 0; i < len(users); i++ {
		var user = users[i]
		if user.Name == name {
			return user, nil
		}
	}

	return returnUser, errors.New("user " + name + " not found")

}

func GetAllUsers() map[int]User {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var users = make(map[int]User)
	for scanner.Scan() {
		var userParts = strings.Split(scanner.Text(), ":")
		var id, _ = strconv.Atoi(userParts[2])
		var group, _ = osuser.LookupGroupId(userParts[3])

		var user = User{
			Id:       id,
			Name:     userParts[0],
			Fullname: userParts[4],
			Home:     userParts[5],
			Shell:    userParts[6],
			Group:    group.Name,
			System:   IsSystemUserId(id),
		}

		users[id] = user
	}

	return users
}
