package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kjetilnym0en/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Venligst velg convert, average eller exit:")

		if !scanner.Scan() {
			break
		}
		input = scanner.Text()

		switch input {
		case "exit", "q":
			fmt.Println("Avslutter")
			return

		case "convert", "c":
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit")

			// En funksjon som åpner en fil, leser linjer og gjør endringer, og deretter lagrer de endrede linjene i en ny fil.
			yr.ConvertTemperature()

		case "average", "a":
			fmt.Println("Viser gjennomsnitt av temperatur")

			for {
				yr.AverageTemperature()

				var input2 string
				scanjn := bufio.NewScanner(os.Stdin)
				fmt.Println("Vil du tilbake til hovedmenyen? (j/n)")
				for scanjn.Scan() {
					input2 = scanjn.Text()
					if input2 == "j" {
						break
					} else if input2 == "n" {
						break
					}
				}
				if input2 == "j" {
					break
				}
			}
		}
	}
	fmt.Println("Konverteringsprogram avsluttet")
}
