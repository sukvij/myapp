package controller

import (
	friendModel "splitwise/friend/model"
	friendService "splitwise/friend/service"
	userModel "splitwise/user/model"

	"github.com/kataras/iris/v12"
)

var allUsers map[int]userModel.User
var allFriends map[int][]friendModel.Friend

func AllApis(app *iris.Application, allFriends1 map[int][]friendModel.Friend, allUsers1 map[int]userModel.User) {
	allFriends = allFriends1
	allUsers = allUsers1
	friendAPI := app.Party("/friend")
	{
		friendAPI.Get("/", allFriendsList)
		friendAPI.Post("/makeFriend", makeFriend)
		friendAPI.Post("/removeFriend", removeFriend)
	}
}

func allFriendsList(ctx iris.Context) {
	result := friendService.AllFriendsList(allFriends, allUsers)
	ctx.JSON(result)
}

func makeFriend(ctx iris.Context) {
	var friends []friendModel.Friend
	ctx.ReadJSON(&friends)
	friendService.MakeFriend(friends, allFriends)
}

func removeFriend(ctx iris.Context) {
	var friends []friendModel.Friend
	ctx.ReadJSON(&friends)
	friendService.RemoveFriend(friends, allFriends)
}
