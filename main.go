package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	go func() {
		tick := time.Tick(1 * time.Second)
		for range tick {
			name, err := GetActivityName()
			if err != nil {
				log.Println(err)
			}

			a := NewActivity(name)

			err = a.Store()
			if err != nil {
				log.Println(err)
			}
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, os.Kill)
	<-s
}
