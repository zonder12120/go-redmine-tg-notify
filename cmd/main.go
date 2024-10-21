package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zonder12120/go-redmine-tg-notify/internal/config"
	"github.com/zonder12120/go-redmine-tg-notify/internal/notify"
	"github.com/zonder12120/go-redmine-tg-notify/internal/redmine"
	"github.com/zonder12120/go-redmine-tg-notify/internal/timecheck"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

// Дефолтный таймаут запроса обновлений
const defaultTimeout = 60 * time.Second

func main() {

	// Срез номеров задач, которые будут игнорироваться при оповещениях
	var ignoredIssues = []int{
		71060,
	}

	// Инициализирует мапу задач, которые будут игнорироваться (для скорости обнаружения совпадений)
	ignoredIssuesMap := redmine.InitIgnoredIssuesMap(ignoredIssues)

	// Инициализирует .env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err)
	}

	err = cfg.CheckAfterInit()
	utils.FatalOnError(err)

	rmClient := redmine.NewClient(cfg.RedmineBaseURL, cfg.RedmineAPIKey, cfg.ProjectsID)

	// Вывод в консоль всех имеющихся проектов, их id и соответствующего имени для конфига
	// Начинается вывод списка с оглавления "Projects List:"
	rmClient.GetProjectsList()

	oldIssueList, err := rmClient.GetIssuesList()
	if err != nil {
		log.Println(err)
	}

	oldIssuesMap := redmine.MakeMapIssuesList(oldIssueList)

	rmClient.AddJournalsIssuesMap(oldIssuesMap)

	if timecheck.IsWorkTime(cfg.TimeZone) {
		notify.SendMessage("Бот запущен")
		notify.SendMessage(fmt.Sprintf("Бот работает каждые %v", defaultTimeout))
	}

	log.Println("The bot is running")

	log.Println("Initialisation new tasks...")

	time.Sleep(defaultTimeout)

	for {
		newIssueList, err := rmClient.GetIssuesList()
		if err != nil {
			log.Println(err)
		}

		newIssuesMap := redmine.MakeMapIssuesList(newIssueList)

		rmClient.AddJournalsIssuesMap(newIssuesMap)

		notify.Updates(oldIssuesMap, newIssuesMap, ignoredIssuesMap)

		log.Printf("ITERATION IS OVER\n\n\n")

		oldIssuesMap = newIssuesMap

		time.Sleep(defaultTimeout)
	}

}
