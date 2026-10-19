package message

import "go_chat/user"

type Message struct {
	Sender    user.User
	Recipient user.User
	Text      string
}
