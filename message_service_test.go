package main

import "testing"

func TestThatItCanSendMessage(test *testing.T) {

	messageService := MessageService{}

	message := Message{
		sender:    "daniel",
		recipient: "ezra",
		text:      "hello world",
	}
	messageService.Send(message)
	if len(messageService.messages) != 1 {
		test.Errorf("expected 1 message , got %v", len(messageService.messages))
	}
}

func TestThatTheMessageStoredIsCorrect(test *testing.T) {

	messageService := MessageService{}

	message := Message{

		sender:    "daniel",
		recipient: "ezra",
		text:      "hello world",
	}
	messageService.Send(message)

	if message.sender != messageService.messages[0].sender {
		test.Errorf("expected daniel got %v", messageService.messages[0].sender)
	}

	if messageService.messages[0].recipient != message.recipient {
		test.Errorf("expected recipient %v, got %v",
			message.recipient,
			messageService.messages[0].recipient,
		)
	}

	if messageService.messages[0].text != message.text {
		test.Errorf("expected text %v, got %v",
			message.text,
			messageService.messages[0].text,
		)
	}
}

func TestThatMultipleMessagesCanBeSent(test *testing.T) {

	messageService := MessageService{}

	firstMessage := Message{

		sender: "daniel",

		recipient: "ezra",

		text: "how far",
	}

	secondMessage := Message{

		sender: "daniel",

		recipient: "ezra",

		text: "how far",
	}

	messageService.Send(firstMessage)
	messageService.Send(secondMessage)

	if len(messageService.messages) != 2 {
		test.Errorf("expcected  2 stored messages, got %v", len(messageService.messages))

	}

}

func TestThatTheMultipleMessagesSentAreCorrect(test *testing.T) {

	messageService := MessageService{}
	firstMessage := Message{

		sender: "daniel",

		recipient: "ezra",

		text: "how far",
	}

	secondMessage := Message{

		sender: "mano",

		recipient: "ezra",

		text: "how far",
	}

	messageService.Send(firstMessage)
	messageService.Send(secondMessage)

	if firstMessage.sender != messageService.messages[0].sender {
		test.Errorf("expected daniel got %v", messageService.messages[0].sender)

	}
	if secondMessage.sender != messageService.messages[1].sender {
		test.Errorf("expected mano got %v", messageService.messages[0].sender)

	}

}

func TestGetMessagesBetweenUsers(test *testing.T) {

	messageService := MessageService{}

	messageService.Send(Message{
		sender:    "daniel",
		recipient: "ezra",
		text:      "hello",
	})

	messageService.Send(Message{
		sender:    "daniel",
		recipient: "john",
		text:      "hey john",
	})

	messages := messageService.GetMessages("daniel", "ezra")

	if len(messages) != 1 {
		test.Errorf("expected 1 message, got %d", len(messages))
	}

	if messages[0].text != "hello" {
		test.Errorf("expected hello, got %v", messages[0].text)
	}
}
