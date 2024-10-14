package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/notify"
	"github.com/zonder12120/go-redmine-tg-notify/internal/redmine"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

const defaultTimeout = 6 * time.Second

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	err = cfg.CheckAfterInit()
	utils.FatalOnError(err)

	rmClient := redmine.NewClient(cfg.RedmineBaseURL, cfg.RedmineAPIKey)

	oldIssueList, err := rmClient.GetIssuesList()
	if err != nil {
		log.Println(err)
	}

	oldIssuesMap := redmine.MakeMapIssuesList(oldIssueList)

	rmClient.AddJournalsIssuesMap(oldIssuesMap)

	log.Printf("The bot is running")
	notify.Notify("Бот запущен")

	log.Printf("Initialisation old tasks... (%v)", defaultTimeout)
	notify.Notify(fmt.Sprintf("Бот работает каждые %v", defaultTimeout))

	time.Sleep(defaultTimeout)

	for {
		newIssueList, err := rmClient.GetIssuesList()
		if err != nil {
			log.Println(err)
		}

		newIssuesMap := redmine.MakeMapIssuesList(newIssueList)

		rmClient.AddJournalsIssuesMap(newIssuesMap)

		rmClient.NotifyUpdates(oldIssuesMap, newIssuesMap)

		log.Printf("ITERATION IS OVER\n\n\n")

		oldIssuesMap = newIssuesMap

		time.Sleep(defaultTimeout)
	}
}
