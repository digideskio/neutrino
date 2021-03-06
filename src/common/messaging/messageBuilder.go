package messaging

import (
	"time"

	"github.com/neutrinoapp/neutrino/src/common/models"
)

type MessageBuilder interface {
	Build(op string, origin string, payload models.JSON, options models.Options, t, app, token string) Message
	BuildFromModel(m models.JSON) Message
}

type messageBuilder struct {
}

var b MessageBuilder

func GetMessageBuilder() MessageBuilder {
	if b == nil {
		b = &messageBuilder{}
	}

	return b
}

func (b *messageBuilder) Build(op string, og string, pld models.JSON, opts models.Options, t, app, token string) Message {
	return Message{
		Operation: op,
		Origin:    og,
		Payload:   pld,
		Options:   opts,
		Type:      t,
		App:       app,
		Token:     token,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

func (b *messageBuilder) BuildFromModel(m models.JSON) Message {
	optionsI := m["options"]
	var options models.Options
	if optionsI != nil {
		options = optionsI.(models.Options)
	}

	pldMap := m["pld"]
	pld := models.JSON{}
	if pldMap != nil {
		pld.FromMap(pldMap.(map[string]interface{}))
	}

	return b.Build(
		m["op"].(string),
		m["origin"].(string),
		pld,
		options,
		m["type"].(string),
		m["app"].(string),
		m["token"].(string),
	)
}
