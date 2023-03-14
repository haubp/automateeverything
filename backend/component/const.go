package component

const CurrentUser = "user"

type Requester interface {
	GetUserName() string
}
