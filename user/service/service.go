package service

import userModel "splitwise/user/model"

func CreateUser(user userModel.User, allUsers map[int]userModel.User) {
	allUsers[user.UserId] = user
}

func GetUser(userId int, allUsers map[int]userModel.User) userModel.User {
	return allUsers[userId]
}
