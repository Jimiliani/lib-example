package lib_example

import (
	"lib_example/generated"
)

type messageProcessor interface {
	ProcessMessage(msg generated.MyMessage)
}

type Library struct {
	messageProcessor messageProcessor
}

func NewLibrary(messageProcessor messageProcessor) Library {
	return Library{messageProcessor: messageProcessor}
}
