package main

import (
	"fmt"
	"log"

	"github.com/warungpintar/siera-kube-watch/config"
	liveness_check "github.com/warungpintar/siera-kube-watch/liveness-check"

	eventHandler "github.com/warungpintar/siera-kube-watch/event-handler"
	"k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//setup kubernetesConfig so the client can connect to kube cluster
	kubernetesConfig, err := clientcmd.BuildConfigFromFlags("", "")

	// uncomment this to test on on your local
	// kubernetesConfig, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		log.Fatalf("Error parsing kubernetes config: %v", err)
	}

	err = config.GlobalConfig.Load()
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(kubernetesConfig)
	if err != nil {
		log.Fatalf("Error create kubernetes client set: %v", err)
	}

	log.Println("Client set created")

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(clientset, 0)
	eventInformer := kubeInformerFactory.Core().V1().Events().Informer()

	eventStopper := make(chan struct{})
	defer runtime.HandleCrash()

	eventInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    eventHandler.OnAddEvent,
		DeleteFunc: eventHandler.OnDeleteEvent,
		UpdateFunc: eventHandler.OnUpdateEvent,
	})

	go liveness_check.Ping()

	go eventInformer.Run(eventStopper)

	if !cache.WaitForCacheSync(eventStopper, eventInformer.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for event caches to sync"))
		return
	}
	<-eventStopper
}
