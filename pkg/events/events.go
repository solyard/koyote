package events

import (
	"encoding/json"
	"fmt"

	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/redis"
	"github.com/koyote/pkg/telegram"
)

type Event interface {
	TemplateMessage() (string, error)
}

func (e GitlabJobEvent) TemplateMessage() (string, error) {
	result, err := templateJobEventMessage(e, "job.tpl")
	if err != nil {
		return "", err
	}

	return result, nil
}

func (e GitlabMergeRequestEvent) TemplateMessage() (string, error) {
	result, err := templateMREventMessage(e, "merge_request.tpl")
	if err != nil {
		return "", err
	}

	return result, nil
}

func (e GitlabNoteEvent) TemplateMessage() (string, error) {
	result, err := templateNoteEventMessage(e, "note.tpl")
	if err != nil {
		return "", err
	}

	return result, nil
}

func (e GitlabPipelineEvent) TemplateMessage() (string, error) {
	result, err := templatePipelineEventMessage(e, "pipeline.tpl")
	if err != nil {
		return "", err
	}

	return result, nil
}

func (e GitlabPushEvent) TemplateMessage() (string, error) {
	result, err := templatePushEventMessage(e, "push.tpl")
	if err != nil {
		return "", err
	}

	return result, nil
}

func EventMatcher(eventJSON []byte, chatID string) error {
	var receivedEventType GitlabEventTypeDetector
	err := json.Unmarshal(eventJSON, &receivedEventType)
	if err != nil {
		return err
	}

	event, err := eventComparator(receivedEventType.ObjectKind, eventJSON)
	if err != nil {
		return err
	}

	eventMessage, err := event.TemplateMessage()
	if err != nil {
		return err
	}

	err = telegram.SendEventMessage(chatID, eventMessage)
	if err != nil && config.GlobalAppConfig.Redis.Enabled {
		redis.PublishEventToRedisChannel(fmt.Sprintf("chatID:%v|message:%v", chatID, eventMessage))
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func eventComparator(eventType string, data []byte) (Event, error) {
	var event Event
	switch eventType {
	case "build":
		event = &GitlabJobEvent{}
	case "merge_request":
		event = &GitlabMergeRequestEvent{}
	case "note":
		event = &GitlabNoteEvent{}
	case "pipeline":
		event = &GitlabPipelineEvent{}
	case "push":
		event = &GitlabPushEvent{}
	default:
		return nil, fmt.Errorf("Unknown event type: %s", eventType)
	}

	return event, json.Unmarshal(data, &event)
}
