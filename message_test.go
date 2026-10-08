package main

import (
	"testing"
)

func TestMessageCreation(test *testing.T) {

	message := Message{
		sender:    "daniel",
		recipient: "ezra",
		text:      "how far",
	}

	if message.sender != "daniel" {
		test.Errorf("expected sender daniel, got %v ", message.sender)
	}

	if message.recipient != "ezra" {
		test.Errorf("expected recipent ezra, got %v ", message.recipient)
	}

	if message.text != "how far" {
		test.Errorf("expected text how far, got %v ", message.text)
	}
}
