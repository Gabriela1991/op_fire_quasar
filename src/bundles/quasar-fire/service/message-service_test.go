package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	messageTestSrv MessageService = NewMessageService()
)

func getMultipleMessage() [][]string {
	var messages [][]string
	var m1 = []string{"este", " ", " ", "mensaje", " "}
	var m2 = []string{" ", "es", " ", " ", "secreto"}
	var m3 = []string{" ", "es", "un", "mensaje", "secreto"}
	messages = append(messages, m1)
	messages = append(messages, m2)
	messages = append(messages, m3)
	return messages
}

func TestGetMultipleMessage_Ok(t *testing.T) {
	response := messageTestSrv.GetMessage(getMultipleMessage()...)
	assert.NotNil(t, response)
}

func TestGetSimpleMessage_Ok(t *testing.T) {
	var m1 = []string{"este", " ", " ", "mensaje", " "}
	response := messageTestSrv.GetMessage(m1)
	assert.NotNil(t, response)
}
