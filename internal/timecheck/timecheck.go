package timecheck

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/pkg/httpreq"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

func IsWorkTime(googleDevApiKey string) bool {
	// Инициализируем часовой пояс (у меня были с этим проблемы на Orange Pi)
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("Ошибка загрузки часового поясв: %v", err)
	}

	currentTime := time.Now().In(location)

	holidays, err := fetchHolidays(googleDevApiKey)
	if err != nil {
		log.Println("Error determining work time: ", err)
	}

	for _, item := range holidays.Items {
		holidayDate, err := time.Parse("2006-01-02", item.Start.Date)
		if err != nil {
			// Fatal потому что мы же не хотим, чтобы приходили оповещения ночью или в праздники
			utils.FatalOnError(err)
		}

		if currentTime.Equal(holidayDate) {
			return false
		}
	}

	if (currentTime.Weekday() >= time.Monday && currentTime.Weekday() <= time.Friday) && (currentTime.Hour() >= 9 && currentTime.Hour() < 19) {
		return true
	}

	return false
}

func fetchHolidays(googleDevApiKey string) (*holidays, error) {
	var holidays holidays

	url, err := utils.ConcatStrings(
		"https://clients6.google.com/calendar/v3/calendars/en.russian%23holiday@group.v.calendar.google.com",
		"/events?calendarId=en.russian%23holiday%40group.v.calendar.google.com&singleEvents=true",
		"&eventTypes=default&eventTypes=focusTime&eventTypes=outOfOffice",
		"&timeZone=Z&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=",
		strconv.Itoa(getCurrentYear()),
		"-01-01T00%3A00%3A00Z&timeMax=",
		strconv.Itoa(getCurrentYear()+1),
		"-01-01T00%3A00%3A00Z&key=",
		googleDevApiKey,
	)
	if err != nil {
		return nil, fmt.Errorf("error concat strings for get holidays request %s", err)
	}

	body, err := httpreq.GetReqBody(url)
	if err != nil {
		return nil, fmt.Errorf("error get holidays data: %s", err)
	}

	err = json.Unmarshal(body, &holidays)
	if err != nil {
		return nil, fmt.Errorf("error encoding body from get issues req %s", err)
	}

	if len(holidays.Items) == 0 {
		return nil, fmt.Errorf("error, an empty response was received")
	}

	return &holidays, nil
}

func getCurrentYear() int {
	return time.Now().Year()
}
