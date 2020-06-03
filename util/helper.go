package util

import (
	"fmt"
	"github.com/warungpintar/siera-kube-watch/config"
	corev1 "k8s.io/api/core/v1"
)

func NotifyEvent(event *corev1.Event) {
	if !isExist(config.GlobalConfig.ExcludedReasons, event.Reason) {
		if event.Type != NORMAL {
			message := parseEventToMessage(event)
			postEvent(message)
		} else {
			if isExist(config.GlobalConfig.IncludedReasons, event.Reason) {
				message := parseEventToMessage(event)
				postEvent(message)
			}
		}
	}
}

func parseEventToMessage(event *corev1.Event) string {
	return fmt.Sprintf("[%s: %s] %s/%s %s", event.Type, event.Reason, event.Namespace, event.Name, event.Message)
}

func isExist(arr []string, element string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}

	return false
}
