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
		if len(cmdString) <= 1 {
			// for handling empty command
			continue
		}
		err = CommandRouter(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// CommandRouter is orchestra of command input
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
		err := handlers.CarParkingHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "leave":
		err := handlers.CarLeaveParkingHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "status":
		err := handlers.StatusParkingSlotHandler()
		if err != nil {
			return err
		}
	case "registration_numbers_for_cars_with_colour":
		err := handlers.GetPlatNoCarByColourHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "slot_numbers_for_cars_with_colour":
		err := handlers.GetSlotByCarColorHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "slot_number_for_registration_number":
		err := handlers.GetSlotsByCarPlatNoHandler(arrCommandStr)
		if err != nil {
			return err
		}
	case "help":

	case "exit":
		os.Exit(0)
	}

	return nil
}
