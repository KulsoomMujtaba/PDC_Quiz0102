package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	check := 0
	for check < 4 {
		fmt.Println("Please Select an Option:")
		fmt.Println("1 - Print Covid19 Cases in Pakistan")
		fmt.Println("2 - Print Covid19 Cases in Soth Korea")
		fmt.Println("3 - Print Covid19 Cases in France")
		fmt.Println("4 - Print my personalized message to Corona Virus")
		fmt.Println("0 - Exit")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare("1", text) == 0 {
			fmt.Println("Cases in Pakistan = 2300")
			check++
		}

		if strings.Compare("2", text) == 0 {
			fmt.Println("Cases in South Korea = 1200")
			check++
		}

		if strings.Compare("3", text) == 0 {
			fmt.Println("Cases in France = 40000")
			check++
		}

		if strings.Compare("4", text) == 0 {
			fmt.Println("Qurantine khatam karo")
			check++
		}

		if strings.Compare("0", text) == 0 {
			if check == 4 {
				break
			}

			if check != 4 {
				fmt.Println("Review all options to exit")
			}
		}
	}
}
