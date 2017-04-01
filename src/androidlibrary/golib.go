package androidlibrary

import (
	"time"
	"log"
)

type JavaCallback interface {
	Heartbeat()
	RelayMessage(msg string, intval int)
}

var callback JavaCallback
var ticker *time.Ticker
var intval int

func RegisterCallback(c JavaCallback) {
	if callback != nil {
		UnregisterCallback()
	}

	callback = c

	log.Print("Callback registered.")

	ticker = time.NewTicker(500 * time.Millisecond)
	go func() {
		for range ticker.C {
			log.Print("Tick.")
			callback.RelayMessage("Tick.", intval)
			callback.Heartbeat()
			intval++
		}
	}()
}

func UnregisterCallback() {
	if callback != nil {
		log.Print("Callback unregistered.")
		ticker.Stop()
		ticker = nil
		callback = nil
	}
}

func Echo(in string) string {
	return "Android said: " + in
}
