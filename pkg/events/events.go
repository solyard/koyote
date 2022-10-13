package events

import (
	"encoding/json"
	"fmt"
	"strconv"

	log "github.com/gookit/slog"
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

func EventMatcher(eventJSON []byte, chatID string) {
	var receivedEventType GitlabEventTypeDetector
	err := json.Unmarshal(eventJSON, &receivedEventType)
	if err != nil {
		log.Error("Cannot unmarshal received event to GitlabTypeDetector structure. Error: ", err)
		return
	}

	event, err := eventComparator(receivedEventType.ObjectKind, eventJSON)
	if err != nil {
		log.Error("Error while compare event with struct", err)
		return
	}

	chatIDInt, err := strconv.Atoi(chatID)
	if err != nil {
		log.Error("Error while convert chat ID to INT: %v", err)
		return
	}

	eventMessage, err := event.TemplateMessage()
	if err != nil {
		log.Error("Error while templating event from received message. Error: ", err)
		return
	}

	telegram.SendEventMessage(int64(chatIDInt), eventMessage)
}

func eventComparator(eventType string, data []byte) (Event, error) {
	switch eventType {
	case "build":
		var gitlabEvent GitlabJobEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "merge_request":
		var gitlabEvent GitlabMergeRequestEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "note":
		var gitlabEvent GitlabNoteEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "pipeline":
		var gitlabEvent GitlabPipelineEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	case "push":
		var gitlabEvent GitlabPushEvent
		err := json.Unmarshal(data, &gitlabEvent)
		return gitlabEvent, err
	default:
		return nil, fmt.Errorf("Unknown event type: %s", eventType)
	}
}
