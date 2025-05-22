package main

import (
	"log"
	"time"

	"services/internal/province/service"

	"github.com/robfig/cron/v3"
)

func main() {
	// Setting up to wok with UTC
	c := cron.New(cron.WithLocation(time.UTC))
	_, err := c.AddFunc("0 14 * * *", func() {
		log.Println("Daily Nuke Time!")
		err := service.UpdateDestroymentRound()
		if err != nil {
			log.Printf("Nuke updating error: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Adding Cron job error: %v", err)
	}

	c.Start()

	select {}
}
