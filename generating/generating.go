package generating

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	prefix string = "./data/"
)

func createDirIfNotExists(dirname string) error {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		if err := os.Mkdir(dirname, 0755); err != nil {
			return fmt.Errorf("error while directory creating: %w", err)
		}
	}
	return nil
}

func GenerateFile(filename string, nums int) error {
	if err := createDirIfNotExists(prefix); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	file, err := os.Create(prefix + filename)
	if err != nil {
		return fmt.Errorf("unable to create file with name %s: %w", prefix+filename, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < nums; i++ {
		if _, err := writer.WriteString(fmt.Sprintf("%d\n", i)); err != nil {
			return fmt.Errorf("unable to write in file %s: %w", prefix+filename, err)
		}
	}
	return writer.Flush()
}

func ReadDataFromFile(filename string) ([]int, error) {
	file, err := os.Open(prefix + filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %w", prefix+filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("unable to read number from file %s: %w", prefix+filename, err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error in file %s: %w", prefix+filename, err)
	}

	return numbers, nil
}

func ClearGeneratedFiles() error {
	if err := os.RemoveAll(prefix); err != nil {
		return fmt.Errorf("error while clearing data generated files: %w", err)
	}
	return nil
}
