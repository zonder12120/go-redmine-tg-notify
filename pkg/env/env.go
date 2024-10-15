package env

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadEnv() error {
	file, err := os.Open(".env")
	if os.IsNotExist(err) {
		return fmt.Errorf("no .env file found, using system environment variables: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		line = strings.ReplaceAll(line, "\"", "")

		keyValue := strings.SplitN(line, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid line in .env file: %s", line)
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %s", err)
	}

	return nil
}

func GetSliceIntFromEnv(key string) []int {
	str := os.Getenv(key)

	str = strings.TrimSpace(str)

	if str == "" {
		return nil
	}

	strSlice := strings.Split(str, ",")

	intSlice := make([]int, len(strSlice))

	for i, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}

		intSlice[i] = num
	}

	return intSlice
}
