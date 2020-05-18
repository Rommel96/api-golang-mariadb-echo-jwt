package models

type Users struct {
	Id       int    `json:"id,primary_key"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func VerifyUser(username string) *Users {
	u := Users{}
	if err := db.Find(&u, Users{Username: username}).Error; err != nil {
		return nil
	}
	return &u
}

func (u *Users) CreateUser() *Users {
	db.NewRecord(&u)
	db.Create(&u)
	return u
}

func GetAllUsers() []Users {
	var users []Users
	db.Find(&users)
	return users
}

func GetUserID(id int) *Users {
	var user Users
	err := db.First(&user, id).Error
	if err != nil {
		return nil
	}
	return &user
}

func (u *Users) UpdateUser(id int) *Users {
	var user Users
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil
	}
	db.Model(&user).Update(u)
	return u
}

func DeleteUser(id int) *Users {
	var user Users
	db.Where("id=?", id).Delete(&user)
	err := db.First(&user, id).Error
	if err != nil {
		return nil
	}
	return &user
}
