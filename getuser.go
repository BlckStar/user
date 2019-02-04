package user

import(
	"os"
	"bufio"
	"strings"
	"strconv"
	"errors"
	osuser "os/user"
)

func GetUser(name string) (User, error) {
	var users = GetAllUsers();
	var returnUser User

    for i := 0; i< len(users); i++ {
        var user = users[i];
        if user.Name() == name {
            return user, nil
        }
	}
	
	return returnUser, errors.New("user " + name + " not found")

}

func GetAllUsers() map[int]User{
	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner:= bufio.NewScanner(file)
	var users = make(map[int]User);
	for scanner.Scan(){
		var userParts = strings.Split(scanner.Text(), ":")
		var id, _ = strconv.Atoi(userParts[2]);
		var group, _ = osuser.LookupGroupId(userParts[3])
	
		var user = User{}
		user.SetId(id)
		user.SetName(userParts[0])
		user.SetFullname(userParts[4])
		user.SetHome(userParts[5])
		user.SetShell(userParts[6])
		user.SetGroup(group.Name)
		user.SetSystem(IsSystemUserId(user.Id()))

		users[user.Id()] = user
	}

	return users;

}
