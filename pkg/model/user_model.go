package model

type UserSt struct {
	UserId   string `json:"userid"`
	UserName string `json:"username"`
	Pw       string `json:"password"`
}

func GetAllUser() ([]UserSt, error) {
	userModel := []UserSt{}
	rows, err := db.Query("SELECT (userId,userName) FROM users")
	if err != nil {
		return userModel, err
	}
	for rows.Next() {
		var user UserSt
		err := rows.Scan(&user.UserId, &user.UserName)
		if err != nil {
			return userModel, err
		} else {
			userModel = append(userModel, user)
		}
	}
	rows.Close()
	return userModel, err
}

func CreateUser(user *UserSt) error {
	_, err := db.Exec("INSERT INTO users (userName,pw) VALUES (?,?)", user.UserName, user.Pw)
	return err
}

func Login(user *UserSt) (UserSt, error) {
	var userc UserSt
	err := db.QueryRow("SELECT (userName,pw) from users WHERE userName=?", user.UserName).Scan(&userc.UserName, &user.Pw)
	return userc, err
}
