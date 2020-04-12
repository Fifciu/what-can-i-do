package models

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email   string    `json:"email"`
	Fullname   string    `json:"fullname"`
	Provider string `json:"provider"`
	Flags uint `json:"flags"`
}

func (user User) TableName() string {
	return "users"
}

func GetUserByEmail(userEmail string) *User {
	user := &User{}

	GetDB().Table("users").Select("*").Where("email = ?", userEmail).First(user)

	return user
}

func GetUserById(userId string) *User {
	user := &User{}

	GetDB().Table("users").Select("*").Where("id = ?", userId).First(user)

	return user
}

func (user *User) Save(email string, fullname string, provider string) bool {
	existingUser := &User{}
	GetDB().Table("users").Select("*").Where("email = ?", email).First(existingUser)

	if existingUser.Email == email {
		return false
	}

	user.Email = email
	user.Fullname = fullname
	user.Provider = provider
	user.Flags = 0
	GetDB().Create(user)
	return true
}