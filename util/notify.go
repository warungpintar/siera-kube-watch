package util

import (
	"fmt"
	"github.com/warungpintar/siera-kube-watch/config"
	"github.com/warungpintar/siera-kube-watch/model"
	corev1 "k8s.io/api/core/v1"
	"log"
)

func NotifyEvent(event *corev1.Event) {
	if isNamespaceIncluded(event.Namespace) {
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
}

func isNamespaceIncluded(namespace string) bool {
	if config.GlobalConfig.IncludedNamespace == nil || len(config.GlobalConfig.IncludedNamespace) == 0 ||
		isExist(config.GlobalConfig.IncludedNamespace, namespace) {
		return true
	}

	return false
}

func parseEventToMessage(event *corev1.Event) string {
	return fmt.Sprintf("[%s: %s] %s/%s %s %v", event.Type, event.Reason, event.Namespace, event.Name, event.Message, event.LastTimestamp)
}

func isExist(arr []string, element string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}

	return false
}

func postEvent(message string) {
	if config.GlobalConfig.Webhook.Enabled {
		model := model.StdModel{}
		model.New(message)
		err := model.Send(config.GlobalConfig.Webhook.Url)
		if err != nil {
			log.Println(err)
		}
	}

	if config.GlobalConfig.Slack.Enabled {
		model := model.StdModel{}
		model.New(message)
		err := model.Send(config.GlobalConfig.Slack.Url)
		if err != nil {
			log.Println(err)
		}
	}

	if config.GlobalConfig.Telegram.Enabled {
		model := model.StdModel{}
		model.New(message)
		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s", config.GlobalConfig.Telegram.Token, config.GlobalConfig.Telegram.ChatID)
		err := model.Send(url)
		if err != nil {
			log.Println(err)
		}
	}

	if config.GlobalConfig.Workplace.Enabled {
		model := model.WorkplaceModel{}
		model.New(message, config.GlobalConfig.Workplace.ThreadKey)
		url := fmt.Sprintf("https://graph.facebook.com/v3.2/me/messages?access_token=%s&formatting=MARKDOWN", config.GlobalConfig.Workplace.Token)
		err := model.Send(url)
		if err != nil {
			log.Println(err)
		}
	}
}
