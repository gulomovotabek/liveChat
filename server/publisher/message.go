package publisher

type InnerMessage struct {
	MessageType    int
	GroupId        int
	SenderUsername string
	Context        string
}
