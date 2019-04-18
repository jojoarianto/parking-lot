package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jojoarianto/parking-lot/src/handlers"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = CommandRouter(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func CommandRouter(commandStr string) error {

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)

	switch arrCommandStr[0] {
	case "create_parking_lot":
		err := handlers.CreateParkingHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "park":

	case "leave":

	case "status":

	case "slot_numbers_for_cars_with_colour":

	case "slot_number_for_registration_number":

	case "help":

	case "exit":
		os.Exit(0)
	}

	return nil
}
