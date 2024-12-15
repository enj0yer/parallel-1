package visualizing

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func trimParams(params []string) []string {
	var trimmed []string = make([]string, 0)
	for _, param := range params {
		trimmed = append(trimmed, strings.Trim(param, " "))
	}

	return trimmed
}

func resolveMode(tokens []string) {
	if tokens[0] == "se" {

	}
}

func VisualizeDatFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, ",")
		return fmt.Errorf("not implemented: %v", params)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading from file: %w", err)
	}
	return nil
}
