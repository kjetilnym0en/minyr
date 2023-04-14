package yr

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kjetilnym0en/funtemps/conv"
)

func ConvertTemperature() {
	if !outputFileDoesNotExist() {
		fmt.Println("Du valgte å ikke opprette filen på nytt")
		return
	}

	File, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer File.Close()

	outputFile, err := createOutputFile()
	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	outputFileWriter := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(File)

	if scanner.Scan() {
		_, err := outputFileWriter.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	for scanner.Scan() {
		// Leser inn en linje fra input-fil, og lagrer i variabelen "line"
		line := scanner.Text()

		// Leser linje fra input-fil, og lager linje output-fil
		_, err := outputFileWriter.WriteString(ReadInputLine(line) + "\n")
		if err != nil {
			panic(err)
		}
	}

	outputFileWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Filen %s med temperaturer i Fahrenheit er opprettet \n", outputFile.Name())
}

// Sjekker om output-filen finnes fra før. Hvis filen ikke finnes fra før, returnerer den true.
func outputFileDoesNotExist() bool {
	if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
		fmt.Print("Filen finnes fra før. Vil du at filen skal opprettes på nytt? (j/n): ")
		var replaceFileInput string
		fmt.Scanln(&replaceFileInput)
		if strings.ToLower(replaceFileInput) == "j" {
			if err := os.Remove("kjevik-temp-fahr-20220318-20230318.csv"); err != nil {
				log.Fatal(err)
			}
			return true
		}
		return false
	}
	return true
}

func createOutputFile() (*os.File, error) {
	outputFilePath := "kjevik-temp-fahr-20220318-20230318.csv"
	if _, err := os.Stat(outputFilePath); err == nil {
		fmt.Printf("File %s already exists. Deleting...\n", outputFilePath)
		if err := os.Remove(outputFilePath); err != nil {
			return nil, fmt.Errorf("could not delete file: %v", err)
		}
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	return outputFile, nil
}

func ReadInputLine(line string) string {
	if line == "" {
		return ""
	}
	fields := strings.Split(line, ";")
	lastField := ""
	if len(fields) > 0 {
		lastField = fields[len(fields)-1]
	}
	convertedField := ""
	if lastField != "" {
		var err error
		convertedField, err = convertLastField(lastField)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return ""
		}
	}
	if convertedField != "" {
		fields[len(fields)-1] = convertedField
	}
	if line[0:7] == "Data er" {
		return "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Kjetil"
	} else {
		return strings.Join(fields, ";")
	}
}

func AverageTemperature() {
	// Åpner input-filen
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// leser alle linjene fra filen
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Ber brukeren om å velge om de vil ha gjennomsnittlig temperatur i celsius eller fahrenheit
	fmt.Println("Velg temperaturenhet (celsius/fahr):")
	var unit string
	fmt.Scan(&unit)

	// Regner ut den gjennomsnittlige temperaturen
	var sum float64
	count := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			log.Fatalf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			log.Fatalf("could not parse temperature in line %d: %s", i, err)
		}

		if unit == "fahr" {
			// Kaller på metode som konverterer temperaturen tilbake til fahrenheit fra celcuis igjen, om det var det brukeren valgte
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}

	if unit == "fahr" {
		average := sum / float64(count)
		average = math.Round(average*100) / 100
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°F\n", average)
	} else {
		average := sum / float64(count)
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°C\n", average)
	}
}

/*
// Funksjon som henter gjennomsnittstemperaturen

	func GetAverageTemperature(filepath string, unit string) (string, error) {
		file, err := os.Open(filepath)
		if err != nil {
			return "", err
		}

		defer file.Close()

		var sum float64
		count := 0
		scanner := bufio.NewScanner(file)
		for i := 0; scanner.Scan(); i++ {
			if i == 0 {
				continue
			}
			fields := strings.Split(scanner.Text(), ";")
			if len(fields) != 4 {
				return "", fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
			}
			if fields[3] == "" {
				continue
			}
			temperature, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return "", fmt.Errorf("could not parse temperature in line %d: %s", i, err)
			}

			if unit == "fahr" {
				temperature = conv.CelsiusToFahrenheit(temperature)
			}
			sum += temperature
			count++
		}
		average := sum / float64(count)
		return fmt.Sprintf("%.2f", average), nil
	}
*/
func GetAverageTemperature(filepath string, unit string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var sum float64
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) != 4 || fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return "", fmt.Errorf("could not parse temperature: %w", err)
		}
		if unit == "fahr" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}
	average := sum / float64(count)
	return fmt.Sprintf("%.2f", average), nil
}

// Funksjon som teller linjer i en fil for yr_test
func LineCounter(File string) int {
	file, err := os.Open(File)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	NumberOfLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			NumberOfLines++
		}
	}
	return NumberOfLines
}

func convertLastField(lastField string) (string, error) {
	// Konverterer det siste tallet på en linje til datatypen float.
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}

	// Kaller på metode som konverterer celsius til fahrenheit
	fahrenheit := conv.CelsiusToFahrenheit(celsius)

	// Konverterer fahrenheit float-verdien tilbake til en string, som var dens opprinnelige type.
	return fmt.Sprintf("%.1f", fahrenheit), nil
}
