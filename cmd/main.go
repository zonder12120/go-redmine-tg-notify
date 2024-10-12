package main

import (
	"fmt"
	"log"

	"github.com/zonder12120/go-redmine-tg-notify/internal/redmine"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic()
	}

	issueList := redmine.GetIssuesList(*cfg)

	fmt.Println(issueList)
}
