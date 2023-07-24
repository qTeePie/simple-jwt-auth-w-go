package data

//Struct for user data
type user struct {
	email    string
	username string
	//server will not know plaintext password, encrypted on client
	passwordhash string
	fullname     string
	createdAt    string
	//role is 0 => standard user
	role int
}

//Faking database
var userList = []user{
	{
		email:        "issi@gmail.com",
		username:     "issichik",
		passwordhash: "hashedme1",
		fullname:     "Issi Chik",
		createdAt:    "1631600786",
		role:         1,
	},
	{
		email:        "checkers@example.com",
		username:     "checkers",
		passwordhash: "hashedme2",
		fullname:     "Checker S",
		createdAt:    "1631600837",
		role:         0,
	},
}

//Get user by unique email
func GetUserObject(email string) (user, bool) {
	for _, user := range userList {
		if user.email == email {
			//Match!
			return user, true
		}
	}

	//No match, return false
	return user{}, false
}

//Function belonging to user struct
func (u *user) ValidatePasswordHash(pswdhash string) bool {
	return u.passwordhash == pswdhash
}

//Adds new users
func AddNewUserObject(email string, username, string, passwordhash string, fullname string, role int) bool {
	//declare new object
	newUser := user{
		email:        email,
		passwordhash: passwordhash,
		username:     username,
		fullname:     fullname,
		role:         role,
	}

	//Check if user exists
	for _, u := range userList {
		if u.email == email || u.username == username {
			return false
		}
	}

	userList = append(userList, newUser)
	return true
}
