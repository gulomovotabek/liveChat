package handler

const (
	CreateGroup = iota
	JoinGroup
	LeaveGroup
	ListGroup

	SendMessage

	SetUsername
)

type InnerMessage struct {
	MessageType    int
	GroupId        int
	SenderUsername string
	Context        string
}

type Message struct {
	Header MessageHeader
	Body   interface{}
}

type MessageHeader struct {
	MessageType int
	Length      int
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
