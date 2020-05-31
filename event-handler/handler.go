package event_handler

import (
	"github.com/warungpintar/siera-kube-watch/util"
	corev1 "k8s.io/api/core/v1"
)

func OnAddEvent(obj interface{}) {
	event := obj.(*corev1.Event)
	util.NotifyEvent(event)
}

func OnDeleteEvent(obj interface{}) {
	event := obj.(*corev1.Event)
	util.NotifyEvent(event)
}

func OnUpdateEvent(oldObj, newObj interface{}) {
	event := newObj.(*corev1.Event)
	util.NotifyEvent(event)
}
