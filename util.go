package substack

import (
	"bufio"
	"os"
	"strings"
)

func loadEnv(filepath string) (map[string]string, error) {
	apiKeys := make(map[string]string)
	file, err := os.Open(filepath)
	if err != nil {
		return apiKeys, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
			apiKeys[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return apiKeys, err
	}
	return apiKeys, nil
}
