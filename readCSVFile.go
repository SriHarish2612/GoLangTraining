package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type movieData struct {
	movieType string
	director  string
	cast      string
	country   string
	listed_in string
}

func main() {

	fmt.Println("Enter 1 for type")
	fmt.Println("Enter 2 for listed_in")
	fmt.Println("Enter 3 for type and country")
	fmt.Printf("Enter Your Choice:")
	var userInput string
	fmt.Scanln(&userInput)
	data := readCSVFile()

	switch userInput {
	case "1":
		fmt.Print("Enter type:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userEnteredtype := scanner.Text()
		showType(data, userEnteredtype)
	case "2":
		fmt.Print("Enter listed_in:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		listed_in := scanner.Text()
		displaylisted_in(data, listed_in)
	case "3":
		fmt.Print("Enter type:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userEnteredtype := scanner.Text()
		fmt.Print("Enter country:")
		scanner.Scan()
		country := scanner.Text()
		displayTypeAndCountry(data, userEnteredtype, country)
	default:
		fmt.Println("Wrong choice")
	}
}

func readCSVFile() [][]string {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}

func showType(data [][]string, param string) {
	for _, line := range data {
		if line[1] == param {
			emp := movieData{
				movieType: line[1],
				director:  line[3],
				cast:      line[4],
			}
			fmt.Println(emp.movieType + " " + emp.director + " " + emp.cast)
		}
	}
}

func displaylisted_in(data [][]string, param string) {
	for _, line := range data {
		if line[10] == param {
			emp := movieData{
				movieType: line[1],
				director:  line[3],
				cast:      line[4],
				listed_in: line[10],
			}
			fmt.Println(emp.movieType + " " + emp.director + " " + emp.listed_in)
		}
	}
}

func displayTypeAndCountry(data [][]string, param string, param2 string) {
	for _, line := range data {
		if line[1] == param && line[5] == param2 {
			emp := movieData{
				movieType: line[1],
				director:  line[3],
				cast:      line[4],
				listed_in: line[10],
				country:   line[5],
			}
			fmt.Println(emp.country + " " + emp.director + " " + emp.cast)
		}
	}
}
