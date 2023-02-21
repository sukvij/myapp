package dummyfriend

import friendModel "splitwise/friend/model"

func MakeDummyFriends(allFrieds map[int][]friendModel.Friend) {
	friend1 := friendModel.Friend{UserId: 100}
	friend2 := friendModel.Friend{UserId: 101}
	allFrieds[100] = append(allFrieds[100], friend2)
	allFrieds[101] = append(allFrieds[101], friend1)
}
