package logger

type MessageType string

var (
	ERROR   MessageType = "ERROR"
	MESSAGE MessageType = "MESSAGE"
)

type LogMessage struct {
	MessageType MessageType
	Message     string
}

func (LogMessage LogMessage) Serialize() map[string]interface{} {
	out := make(map[string]interface{})

	out["messageType"] = LogMessage.MessageType
	out["message"] = LogMessage.Message

	return out
}
