package common

import "fmt"

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

const (
	CurrrentUser = "current_user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(request Requester) bool {
	return request.GetRole() == "admin" || request.GetRole() == "mod"
}
