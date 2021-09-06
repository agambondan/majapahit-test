package main

import (
	"errors"
	"fmt"
	"test-majapahit/model"
)

type Users []model.User

func main() {
	var userr = model.User{
		Id:                 1,
		Username:           "",
		Password:           "johny",
		Email:              "johny@gmail.com",
		AlamatKTP:          "selangkangan",
		Pekerjaan:          "customer service",
		NamaLengkap:        "johny johny",
		PendidikanTerakhir: "S3",
		NomorTelepon:       "452319128321",
	}
	var users Users
	user, user1 := model.InitUser()
	fmt.Println(users.createUser(&users, user))
	fmt.Println(users.createUser(&users, user1))
	fmt.Println(users.createUser(&users, userr))
	fmt.Println(users.getUserById(user1.Id))
	//fmt.Println(users.updateUserById(users, userr, user.Id))
	//fmt.Println()
	fmt.Println(users.deleteUserId(users, user.Id))
	//fmt.Println(users)
	fmt.Println(users.login(user.Username, user.Email, user.Password))
	fmt.Println(users.login(user1.Username, user1.Email, user1.Password))

}

func (u *Users) createUser(users *Users, user model.User) (Users, map[string]string) {
	validation := user.Validation("create")
	if len(validation) != 0 {
		return *users, validation
	} else {
		*users = append(*users, user)
	}
	return *users, nil
}

func (u Users) getUsers() Users {
	return u
}

func (u Users) getUserByUsername(username string) (model.User, error) {
	var err error
	var user model.User
	for i := 0; i < len(u); i++ {
		if u[i].Username == username {
			user = u[i]
		}
	}
	if user.Username == "" {
		err = errors.New("user tidak ditemukan")
	}
	return user, err
}

func (u Users) getUserById(id int) (model.User, error) {
	var err error
	var user model.User
	for i := 0; i < len(u); i++ {
		if u[i].Id == id {
			user = u[i]
		}
	}
	if user.Username == "" {
		err = errors.New("user tidak ditemukan")
	}
	return user, err
}

func (u Users) updateUserById(users Users, user model.User, id int) (model.User, error) {
	userById, err := u.getUserById(id)
	if err != nil {
		return userById, err
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			copy(users[i:], users[i+1:])
			users[len(users)-1] = user
			users = users[:len(users)-1]
		}
	}
	return user, nil
}

func (u Users) deleteUserId(users Users, id int) (Users, error) {
	_, err := u.getUserById(id)
	if err != nil {
		return users, err
	}
	for i := 0; i < len(u); i++ {
		if u[i].Id == id {
			copy(users[i:], users[i+1:])
			users[len(users)-1] = model.User{}
			users = users[:len(users)-1]
			//users = append(users[:i], users[i+1:])
		}
	}
	return users, nil
}

func (u Users) login(username, email, password string) string {
	getUser, err := u.getUserByUsername(username)
	if err != nil {
		return err.Error()
	}
	if (getUser.Username == username || getUser.Email == email) && getUser.Password == password {
		return "Success"
	} else {
		return "Failed"
	}
}
