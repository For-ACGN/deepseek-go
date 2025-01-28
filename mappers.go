package deepseek

import (
	"errors"

	"github.com/cohesion-org/deepseek-go/constants"
	handlers "github.com/cohesion-org/deepseek-go/handlers"
)

var validRoles = map[string]bool{
	constants.ChatMessageRoleUser:      true,
	constants.ChatMessageRoleAssistant: true,
	constants.ChatMessageRoleSystem:    true,
}

func MapMessageToChatCompletionMessage(m handlers.Message) (ChatCompletionMessage, error) {
	if m.Role == "" {
		return ChatCompletionMessage{}, errors.New("message role cannot be empty")
	}

	if m.Content == "" {
		return ChatCompletionMessage{}, errors.New("message content cannot be empty")
	}
	if !validRoles[m.Role] {
		return ChatCompletionMessage{}, errors.New("invalid role: %s. Valid roles are can be found in official deepseek documentation")
	}

	return ChatCompletionMessage{
		Role:    m.Role,
		Content: m.Content,
	}, nil
}
