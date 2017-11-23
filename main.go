package main

import (
	"bufio"
	"os"
	"fmt"
	"net/http"
	"strings"
	"github.com/jmcvetta/napping"
)

var EffectsServer = "http://localhost:8080"

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		EffectsServer = argsWithoutProg[0]
	}

	s := napping.Session{}
	h := &http.Header{}
	h.Set("Content-Type", "application/json")
	s.Header = h

	for {

		fmt.Print("Enter command: ")
		text := getStdIn()
		sendCommand(text, s)
	}
}

func getStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return cleanCommand(text)
}

type CommandRequest struct {
	Command string `json:"command"`
}

func sendCommand(command string, s napping.Session) {
	if command == "shutdown code" {
		command += handleShutdownCommand()
	}

	data := CommandRequest{Command: cleanCommand(command)}

	resp, err := s.Post(EffectsServer+"/command", &data, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d: %s\n", resp.Status(), resp.RawText())
}

func cleanCommand(command string) string {
	return strings.ToLower(strings.TrimSpace(command))
}

func handleShutdownCommand() string {
	fmt.Println(">>> Warning, too many wrong guesses will cause problems!")
	fmt.Print("Enter code: ")
	return getStdIn()
}
