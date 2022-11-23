package dao

var database = map[string]string{
	"zq": "031218",
	"wz": "030126",
}

func AddUser(username, password string) {
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}
