package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucapette/tracker"
)

func main() {
	go func() {
		tick := time.Tick(1 * time.Second)
		for range tick {
			name, err := tracker.GetActivityName()
			if err != nil {
				log.Println(err)
			}

			a := tracker.NewActivity(name)

			err = a.Store()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	<-s
}
