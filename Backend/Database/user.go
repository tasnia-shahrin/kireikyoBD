package Database

type User struct{
	ID int 	`json:"userid"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string  `json:"password"`
	IsAdmin bool `json:"is_admin"`
}

var users []User
func (u User) Store() User{ //successfully create user and return the user
	
	if u.ID != 0{ //user is already created
		return u
	}
	u.ID = len(users)+1
	users = append(users, u)
	return u
}

//find user
func Find(email,password string) *User{
	for _,u:=range users{
		if u.Email == email && u.Password==password{
			return &u
		}
	}
	return nil //if not found
}