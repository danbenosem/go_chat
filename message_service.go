package main

type MessageService struct {
	messages []Message
}

func (service *MessageService) Send(message Message) {

	service.messages = append(service.messages, message)

}

func (service *MessageService) GetMessages(sender string, recipient string) []Message {

	var userMessages []Message

	for _, message := range service.messages {
		if sender == message.sender && recipient == message.recipient {

			userMessages = append(userMessages, message)
		}

	}
	return userMessages
}
