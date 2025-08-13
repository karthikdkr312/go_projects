package flow_controls

import "fmt"

func Switch(day string) {
	switch day {
	case "mon":
		fmt.Println("its Monday")
	case "tues":
		fmt.Println("its TUESDAY")
	case "wednes":
		fmt.Println("its Wednesday")
	case "fri":
		fmt.Println("its friday")
	case "sat":
		fmt.Println("its saturday")
	case "sun":
		fmt.Println("its sunday")
	default:
		fmt.Println("its another day")
	}
}
