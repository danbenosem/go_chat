package message

import "go_chat/user"

type MessageService struct {
	messages []Message
}

func (service *MessageService) Send(message Message) {

	service.messages = append(service.messages, message)

}

func (service *MessageService) GetMessages(sender user.User, recipient user.User) []Message {
	{

		var userMessages []Message

		for _, message := range service.messages {
			if sender == message.Sender && recipient == message.Recipient {

				userMessages = append(userMessages, message)

			}

			if recipient == message.Sender && sender == message.Recipient {
				userMessages = append(userMessages, message)
			}

		}
		return userMessages
	}
}
