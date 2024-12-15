package main

import (
	"fmt"
	"parallel-1/generating"
	"parallel-1/measuring"
	"parallel-1/processing"
	"strconv"
	"strings"
)

type Command struct {
	name   string
	runner func(args []string) (string, error)
	params string
}

func (c Command) String() string {
	return strings.Join([]string{c.name, c.params}, " ")
}

func (c Command) Run(args []string) (string, error) {
	if output, err := c.runner(args); err != nil {
		return "", fmt.Errorf("error while running command %s: %w", c.name, err)
	} else {
		return output, nil
	}
}

var Commands map[string]Command

func init() {
	Commands = make(map[string]Command)
	Commands["list"] = Command{
		name: "list",
		runner: func(args []string) (string, error) {
			if len(args) > 0 {
				return "", fmt.Errorf("to many arguments provided")
			}
			var buffer []string
			for _, command := range Commands {
				buffer = append(buffer, "    "+command.String())
			}
			return strings.Join(buffer, "\n"), nil
		},
		params: "",
	}

	Commands["gen"] = Command{
		name:   "gen",
		params: "<filename> <size>",
		runner: func(args []string) (string, error) {
			if len(args) != 2 {
				return "", fmt.Errorf("expected 2 arguments, provided %d", len(args))
			}
			filename := args[0]
			size, err := strconv.Atoi(args[1])
			if err != nil || size <= 0 {
				return "", fmt.Errorf("size must be a positive number: %w", err)
			}
			if err := generating.GenerateFile(filename, size); err != nil {
				return "", err
			} else {
				return fmt.Sprintf("File %s with %d elements was successfully generated", filename, size), nil
			}
		},
	}

	Commands["clear"] = Command{
		name:   "clear",
		params: "",
		runner: func(args []string) (string, error) {
			if len(args) > 0 {
				return "", fmt.Errorf("to many arguments provided")
			}
			if err := generating.ClearGeneratedFiles(); err != nil {
				return "", fmt.Errorf("error while clearing files: %w", err)
			}
			return "Files was successfully cleared", nil
		},
	}

	Commands["seq"] = Command{
		name:   "seq",
		params: "<filename> <applier>",
		runner: func(args []string) (string, error) {
			if len(args) != 2 {
				return "", fmt.Errorf("expected 2 arguments, provided %d", len(args))
			}
			filename := args[0]
			applierName := args[1]
			applier, ok := processing.IntAppliers[applierName]
			if !ok {
				return "", fmt.Errorf("applier with name %s is not exists", applierName)
			}
			data, err := generating.ReadDataFromFile(filename)
			if err != nil {
				return "", fmt.Errorf("error while reading data from file: %w", err)
			}

			converter := processing.NewConverter(data, applier)
			result := measuring.MeasureTime(func() {
				converter.ProcessSequentially()
			}, measuring.Seq, 0, applierName, len(data))
			return result.Log(), nil
		},
	}

	Commands["sim"] = Command{
		name:   "sim",
		params: "<filename> <applier> [,threads]",
		runner: func(args []string) (string, error) {
			if len(args) < 3 {
				return "", fmt.Errorf("expected at least 3 arguments, provided %d", len(args))
			}
			filename := args[0]
			applierName := args[1]
			applier, ok := processing.IntAppliers[applierName]
			if !ok {
				return "", fmt.Errorf("applier with name %s is not exists", applierName)
			}
			data, err := generating.ReadDataFromFile(filename)
			if err != nil {
				return "", fmt.Errorf("error while reading data from file: %w", err)
			}

			threads := make([]int, 0)
			for _, arg := range args[2:] {
				parsed, err := strconv.Atoi(arg)
				if err != nil {
					return "", fmt.Errorf("error while parsing threads arguments: %w", err)
				}
				threads = append(threads, parsed)
			}

			converter := processing.NewConverter(data, applier)
			results := make([]string, 0)
			for _, threadCount := range threads {
				result := measuring.MeasureTime(func() {
					converter.ProcessSimultaneously(threadCount)
				}, measuring.Sim, threadCount, applierName, len(data))
				results = append(results, result.Log())
			}
			return strings.Join(results, "\n"), nil
		},
	}
}
