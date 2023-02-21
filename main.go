package main

import (
	friendController "splitwise/friend/controller"
	friendDummy "splitwise/friend/dummyFriend"
	friendModel "splitwise/friend/model"

	groupController "splitwise/group/controller"
	groupModel "splitwise/group/model"

	paymentController "splitwise/payment/controller"
	paymentModel "splitwise/payment/model"
	userController "splitwise/user/controller"
	userDummy "splitwise/user/dummyUser"
	userModel "splitwise/user/model"

	"github.com/kataras/iris/v12"
)

var allUsers map[int]userModel.User
var allFriends map[int][]friendModel.Friend
var allGroups map[int][]groupModel.Group
var allPaymentsList map[int][]paymentModel.Payment

func main() {
	allUsers = make(map[int]userModel.User)
	allFriends = make(map[int][]friendModel.Friend)
	allGroups = make(map[int][]groupModel.Group)
	userDummy.CreateDummyUser(allUsers)
	friendDummy.MakeDummyFriends(allFriends)
	app := iris.New()

	userController.AllApis(app, allUsers)
	friendController.AllApis(app, allFriends, allUsers)
	groupController.AllApis(app, allGroups, allFriends, allUsers)
	paymentController.AllApis(app, allPaymentsList)
	app.Listen(":8080")
}
