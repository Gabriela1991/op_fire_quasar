package service

import (
	"strings"
)

type messageService struct{}

type MessageService interface {
	GetMessage(messages ...[]string) (msg string)
}

func NewMessageService() MessageService {
	return &messageService{}
}

func (*messageService) GetMessage(messages ...[]string) (msg string) {
	var messageCon []string = concatenateArrays(messages)
	return putMessage(unique(messageCon), messages)
}

func unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func concatenateArrays(msgArrays [][]string) []string {
	var messages []string
	for _, msg := range msgArrays {
		messages = append(messages, msg...)
	}
	return messages
}

func putMessage(message []string, msgArrays [][]string) string {
	finalMessage := make([]string, len(message), len(message))
	index := len(message)
	for _, phrase := range message {
		if strings.Compare(" ", phrase) != 0 {
			for _, msg := range msgArrays {
				for i, m := range msg {
					if strings.Compare(" ", phrase) != 0 && strings.Compare(phrase, m) == 0 && i < index {
						index = i
					}
				}
			}
			finalMessage[index] = phrase
			index = len(message)
		}
	}
	return strings.Join(finalMessage, " ")
}
