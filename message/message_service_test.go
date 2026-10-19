package message

import (
	"go_chat/user"
	"testing"
)

func TestThatItCanSendMessage(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}

	message := Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello world",
	}

	messageService.Send(message)

	if len(messageService.messages) != 1 {
		test.Errorf("expected 1 message, got %v", len(messageService.messages))
	}
}

func TestThatTheMessageStoredIsCorrect(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}

	message := Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello world",
	}

	messageService.Send(message)

	if message.Sender != messageService.messages[0].Sender {
		test.Errorf("expected daniel, got %v", messageService.messages[0].Sender.Name)
	}

	if messageService.messages[0].Recipient != message.Recipient {
		test.Errorf("expected recipient %v, got %v",
			message.Recipient.Name,
			messageService.messages[0].Recipient.Name,
		)
	}

	if messageService.messages[0].Text != message.Text {
		test.Errorf("expected text %v, got %v",
			message.Text,
			messageService.messages[0].Text,
		)
	}
}

func TestThatMultipleMessagesCanBeSent(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}

	firstMessage := Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "how far",
	}

	secondMessage := Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "how far",
	}

	messageService.Send(firstMessage)
	messageService.Send(secondMessage)

	if len(messageService.messages) != 2 {
		test.Errorf("expected 2 stored messages, got %v", len(messageService.messages))
	}
}

func TestThatTheMultipleMessagesSentAreCorrect(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	mano := user.User{Name: "mano"}
	ezra := user.User{Name: "ezra"}

	firstMessage := Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "how far",
	}

	secondMessage := Message{
		Sender:    mano,
		Recipient: ezra,
		Text:      "how far",
	}

	messageService.Send(firstMessage)
	messageService.Send(secondMessage)

	if firstMessage.Sender != messageService.messages[0].Sender {
		test.Errorf("expected daniel, got %v", messageService.messages[0].Sender.Name)
	}

	if secondMessage.Sender != messageService.messages[1].Sender {
		test.Errorf("expected mano, got %v", messageService.messages[1].Sender.Name)
	}
}

func TestGetMessagesBetweenUsersFromOneDirection(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}
	john := user.User{Name: "john"}

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello",
	})

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: john,
		Text:      "hey john",
	})

	messages := messageService.GetMessages(daniel, ezra)

	if len(messages) != 1 {
		test.Errorf("expected 1 message, got %d", len(messages))
	}

	if messages[0].Text != "hello" {
		test.Errorf("expected hello, got %v", messages[0].Text)
	}
}

func TestThatCommunicationHappensOnBothSides(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello",
	})

	messageService.Send(Message{
		Sender:    ezra,
		Recipient: daniel,
		Text:      "hey",
	})

	danielMessages := messageService.GetMessages(daniel, ezra)
	ezraMessages := messageService.GetMessages(ezra, daniel)

	if len(danielMessages) != 2 {
		test.Errorf("expected 2 messages, got %d", len(danielMessages))
	}

	if len(ezraMessages) != 2 {
		test.Errorf("expected 2 messages, got %d", len(ezraMessages))
	}
}

func TestGetMessagesDoesNotReturnOtherConversations(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}
	john := user.User{Name: "john"}

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello",
	})

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: john,
		Text:      "hey john",
	})

	messages := messageService.GetMessages(daniel, ezra)

	if len(messages) != 1 {
		test.Errorf("expected 1 message, got %d", len(messages))
	}

	if messages[0].Text != "hello" {
		test.Errorf("expected hello, got %v", messages[0].Text)
	}
}

func TestMessagesAreReturnedInOrder(test *testing.T) {

	messageService := MessageService{}

	daniel := user.User{Name: "daniel"}
	ezra := user.User{Name: "ezra"}

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "hello",
	})

	messageService.Send(Message{
		Sender:    ezra,
		Recipient: daniel,
		Text:      "hey",
	})

	messageService.Send(Message{
		Sender:    daniel,
		Recipient: ezra,
		Text:      "how are you?",
	})

	messages := messageService.GetMessages(daniel, ezra)

	if messages[0].Text != "hello" {
		test.Errorf("expected hello, got %v", messages[0].Text)
	}

	if messages[1].Text != "hey" {
		test.Errorf("expected hey, got %v", messages[1].Text)
	}

	if messages[2].Text != "how are you?" {
		test.Errorf("expected how are you?, got %v", messages[2].Text)
	}
}
