package user

import (
    "errors"
)

var users = []User{
    {ID: 1, Name: "Maximiliano Boyzo Sandoval Kuri del pastor Suadero", Email: "max@example.com"},
    {ID: 2, Name: "Romeo Gabriel Boyzo Sandoval", Email: "romeo@example.com"},
}

func GetAllUsers() ([]User, error) {
    return users, nil
}

func GetUserByID(id int) (User, error) {
    for _, u := range users {
        if u.ID == id {
            return u, nil
        }
    }
    return User{}, errors.New("User not found to get")
}

func AddUser(u User) (User, error) {
    u.ID = len(users) + 1
    users = append(users, u)
    return u, nil
}

func UpdateUser(id int, u User) (User, error) {
    for i, usr := range users {
        if usr.ID == id {
            users[i] = u
            users[i].ID = id
            return users[i], nil
        }
    }
    return User{}, errors.New("User not found to update")
}

func DeleteUser(id int) error {
    for i, u := range users {
        if u.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return errors.New("User not found to delete")
}
