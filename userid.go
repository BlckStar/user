package user

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const IdRangesFile = "/etc/adduser.conf"
const SystemIdRangeMinKey = "FIRST_SYSTEM_UID"
const SystemIdRangeMaxKey = "LAST_SYSTEM_UID"

const UserIdRangeMinKey = "FIRST_UID"
const UserIdRangeMaxKey = "LAST_UID"

var systemIdMin int
var systemIdMax int
var userIdMin int
var userIdMax int

func IsSystemUserId(id int) bool {
	if systemIdMin == 0 || systemIdMax == 0 || userIdMin == 0 || userIdMax == 0 {
		calculateSystemRanges()
	}

	return systemIdMin < id && systemIdMax > id
}

func GetNextUserId(system bool) (int, error) {
	if systemIdMin == 0 || systemIdMax == 0 {
		calculateSystemRanges()
	}

	var users = GetAllUsers()
	var i, min, max, maxId int

	if system {
		min = systemIdMin
		max = systemIdMax
	} else {
		min = userIdMin
		max = userIdMax
	}

	var maxIdErr error
	i = 1
	for {
		if min+i > max {
			maxIdErr = errors.New("no Id left")
		}
		if users[min+i].Id == 0 || min+i > max {
			maxId = min + i
			break
		}
		i++
	}

	if maxIdErr != nil {
		return 0, maxIdErr
	}

	return maxId, nil

}

func calculateSystemRanges() {
	file, err := os.Open(IdRangesFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()

		if strings.Contains(line, SystemIdRangeMinKey) {
			var lineParts = strings.Split(line, "=")
			systemIdMin, _ = strconv.Atoi(lineParts[1])
		}

		if strings.Contains(line, SystemIdRangeMaxKey) {
			var lineParts = strings.Split(line, "=")
			systemIdMax, _ = strconv.Atoi(lineParts[1])
		}

		if strings.Contains(line, UserIdRangeMinKey) {
			var lineParts = strings.Split(line, "=")
			userIdMin, _ = strconv.Atoi(lineParts[1])
		}

		if strings.Contains(line, UserIdRangeMaxKey) {
			var lineParts = strings.Split(line, "=")
			userIdMax, _ = strconv.Atoi(lineParts[1])
		}
	}
}
