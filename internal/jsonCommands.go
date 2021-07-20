package SecondOfficer

import (
	"encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"os"
	"time"
)

type CommandData struct {
	Name   string   `json:"name"`
	Alias  string   `json:"alias"`
	Args   int      `json:"args"`
	Color  string   `json:"color"`
	Quotes []string `json:"quotes"`
	Images []string `json:"images"`
}

type Command interface {
	GetArgs() int
	GetColor() int
	GetImageURL() string
	GetQuote() string
}

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func (c CommandData) GetArgs() int {
	return c.Args
}

func (c CommandData) GetColor() int {
	if len(c.Color) > 0 {
		i := new(big.Int)
		i.SetString(c.Color, 16)
		return int(i.Int64())
	}

	return -1
}

func (c CommandData) GetImageURL() string {
	if len(c.Images) > 0 {
		return c.Images[random.Intn(len(c.Images))]
	}

	return ""
}

func (c CommandData) GetQuote() string {
	if len(c.Quotes) > 0 {
		return c.Quotes[random.Intn(len(c.Quotes))]
	}

	return ""
}

func ReadCommands(filepath string) (map[string]Command, error) {
	commandsFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var commandsData []CommandData
	err = json.Unmarshal(commandsFile, &commandsData)

	if err != nil {
		return nil, errors.New("Invalid Command Configuration")
	}

	commands := make(map[string]Command)
	for _, command := range commandsData {
		commands[command.Name] = Command(command)
		if len(command.Alias) > 0 {
			commands[command.Alias] = Command(command)
		}
	}

	return commands, nil
}
