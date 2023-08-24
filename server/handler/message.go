package handler

const (
	CreateGroup = iota
	JoinGroup
	LeaveGroup
	ListGroup

	SendMessage

	SetUsername
)

type Message struct {
	Header MessageHeader
	Body   interface{}
}

type MessageHeader struct {
	MessageType int
	Length      int // todo: not implemented
}

type SendMessageDTO struct {
	Message string
}

type JoinGroupDTO struct {
	GroupId int
}

type SetUsernameMessage struct {
	Username string
}
