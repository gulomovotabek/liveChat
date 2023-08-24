package models

import (
	"math/rand"
	"time"
)

type ReadGroup struct {
	Id           int
	ClientsCount int
}

type Group struct {
	Id      int
	Clients []*Client
	Public  bool
}

func CreateGroup(creator Client, isPublic bool) Group {
	return Group{
		Id:      randomIdGenerator(),
		Clients: []*Client{&creator},
		Public:  isPublic,
	}
}

func randomIdGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10000000000-1000000000) + 1000000000
}

func GetItemByID(groups []*Group, id int) (*Group, bool) {
	for _, item := range groups {
		if item.Id == id {
			return item, true
		}
	}
	return &Group{}, false
}
