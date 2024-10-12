package main

import (
	"fmt"
	"log"

	"time"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/redmine"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	for {
		issueList := redmine.GetIssuesList(*cfg)

		fmt.Printf("Получено задач: %v\n", len(issueList.Issues))
		fmt.Println(issueList)
		time.Sleep(600)
	}
}
