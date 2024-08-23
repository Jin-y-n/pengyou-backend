package constant

// const in redis
const (
	MessageTypeText          = 1
	MessageTypeImage         = 2
	MessageTypeFile          = 3
	MessageTypeFileRequest   = 4
	MessageTypeDisconnect    = 5
	MessageTypeCutChat       = 6
	MessageTypeEstablishChat = 7
)

// const in frontend, backend interactive
const (
	UnknownMessageType   = 0
	TextMessage          = 1
	FileRequestMessage   = 2
	FriendRequestMessage = 3
	EstablishChatMessage = 4
	CutChatMessage       = 5
)
