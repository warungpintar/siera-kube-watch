package util

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func NotifyEvent(event *corev1.Event) {
	if event.Type != NORMAL {
		message := parseEventToMessage(event)
		postEvent(message)
	} else {
		if event.Reason == SCALING_REPLICAT_SET || event.Reason == STARTED || event.Reason == KILLING {
			message := parseEventToMessage(event)
			postEvent(message)
		}
	}
}

func parseEventToMessage(event *corev1.Event) string {
	return fmt.Sprintf("[%s: %s] %s/%s %s", event.Type, event.Reason, event.Namespace, event.Name, event.Message)
}
