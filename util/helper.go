package util

import (
	"fmt"
	"github.com/warungpintar/siera-kube-watch/config"
	corev1 "k8s.io/api/core/v1"
	"log"
)

func NotifyEvent(event *corev1.Event) {
	if !isExist(config.GlobalConfig.ExcludedReasons, event.Reason) {
		if event.Type != NORMAL {
			message := parseEventToMessage(event)
			log.Println(message)
			postEvent(message)
		} else {
			if event.Reason == SCALING_REPLICAT_SET || event.Reason == STARTED || event.Reason == KILLING {
				message := parseEventToMessage(event)
				log.Println(message)
				postEvent(message)
			}
		}
	}
}

func parseEventToMessage(event *corev1.Event) string {
	return fmt.Sprintf("[%s] %s/%s %s", event.Reason, event.Namespace, event.Name, event.Message)
}

func NotifyPod(pod *corev1.Pod) {
	message := fmt.Sprintf("%s/%s Phase: %s Reason: %s Message: %s", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name, pod.Status.Phase, pod.Status.Reason, pod.Status.Message)
	log.Println(message)
}

func isExist(arr []string, element string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}

	return false
}