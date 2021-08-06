package liveness_check

import (
	"log"
	"time"

	"github.com/warungpintar/siera-kube-watch/config"
	"github.com/warungpintar/siera-kube-watch/util"
)

func Ping() {
	if config.GlobalConfig.Livenesscheck.Enabled {

		interval, err := time.ParseDuration(config.GlobalConfig.Livenesscheck.Interval)

		if err != nil {
			log.Fatalf("Error parsing livness check interval value: %v", err)
		}

		log.Printf("Liveness check enabled with interval: %s", interval)

		intervalInSeconds := interval.Seconds()
		for {
			util.PostEvent("This is a dead man's switch mechanism to ensure alert pipeline is working.")
			time.Sleep(time.Duration(intervalInSeconds) * time.Second)
		}
	}
}
