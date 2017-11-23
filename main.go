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
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command: ")
		text, _ := reader.ReadString('\n')
		sendCommand(text, s)
	}
}

type CommandRequest struct {
	Command string `json:"command"`
}

func sendCommand(command string, s napping.Session) {
	data := CommandRequest{Command: cleanCommand(command)}

	resp, err := s.Post(EffectsServer+"/command", &data, nil, nil)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%d: %s\n", resp.Status(), resp.RawText())
}

func cleanCommand(command string) string {
	return strings.ToLower(strings.TrimSpace(command))
}
