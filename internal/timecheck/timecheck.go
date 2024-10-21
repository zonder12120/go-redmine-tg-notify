package timecheck

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	httpreq "github.com/zonder12120/go-redmine-tg-notify/pkg/httpreq"
	"github.com/zonder12120/go-redmine-tg-notify/pkg/utils"
)

func IsWorkTime(timeZone string) bool {
	// Инициализируем часовой пояс (у меня были с этим проблемы на Orange Pi)
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Fatalf("Time zone loading error: %v", err)
	}

	currentTime := time.Now().In(location)

	if currentTime.Hour() >= 9 && currentTime.Hour() < 19 {
		year := currentTime.Year()

		month := currentTime.Month()

		day := currentTime.Day()

		flag, err := checkDayOff(year, int(month), day)
		utils.FatalOnError(err)

		switch flag {
		case 0:
			return true
		case 1:
			return false
		default:
			return true
		}
	}

	return false
}

func checkDayOff(year, month, day int) (int, error) {
	url, err := utils.ConcatStrings(
		fmt.Sprintf("https://isdayoff.ru/api/getdata?year=%d", year),
		fmt.Sprintf("&month=%d", month),
		fmt.Sprintf("&day=%d", day),
	)
	if err != nil {
		return 0, fmt.Errorf("error concat strings for get holidays request %s", err)
	}

	body, err := httpreq.GetReqBody(url)
	if err != nil {
		return 0, fmt.Errorf("error get holidays data: %s", err)
	}

	bodyStr := strings.TrimSpace(string(body))

	dayOffFlag, err := strconv.Atoi(bodyStr)
	if err != nil {
		return 0, fmt.Errorf("convert string body to int: %s", err)
	}

	return dayOffFlag, nil
}
