package main

import (
	"github.com/shonayevshyngys/prontopro/rating_service/servers"
	"log"
	"os"
)

var app string

func init() {
	app = os.Getenv("APP")
	switch app {
	case "rating":
		servers.InitRatingServer()
		return
	case "notification":
		servers.InitNotificationService()
		return
	default:
		log.Fatalf("Wrong environment variables, %s does not exist", app)
	}

}
func main() {
	app = os.Getenv("APP")
	switch app {
	case "rating":
		servers.RunRatingService()
	case "notification":
		servers.RunNotificationService()
	default:
		log.Fatalf("Wrong environment variables, %s does not exist", app)
	}
}
