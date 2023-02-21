package dummyUser

import userModel "splitwise/user/model"

func CreateDummyUser(allUsers map[int]userModel.User) {
	user1 := userModel.User{UserId: 100, UserName: "vijju", UserBalance: 100}
	user2 := userModel.User{UserId: 101, UserName: "sukariya", UserBalance: 1000}
	user3 := userModel.User{UserId: 102, UserName: "pawan", UserBalance: 2000}
	allUsers[user1.UserId] = user1
	allUsers[user2.UserId] = user2
	allUsers[user3.UserId] = user3
}
