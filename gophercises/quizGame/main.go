package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Color codes
const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func main() {
	filename := flag.String("file", "", "Enter the path of csv file containing problems")
	timeout := flag.Int("timeout", 30, "Timer for the quiz in quiz.")
	flag.Parse()

	if *filename == "" {
		exit(fmt.Sprintf("Pass filename through flag 'file'"))
	}

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		exit(fmt.Sprintf("Error opening file %v : %v\n", *filename, err))
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Error parsing file %v : %v\n", *filename, err))
	}

	problems := parseLines(lines)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

problemLoop:
	for _, p := range problems {
		fmt.Printf("%v = ", p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println(string(colorRed), "\nTimeout")
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

		}
	}

	correctAns := fmt.Sprintf("Correct: %d", correct)
	wrongAns := fmt.Sprintf("Wrong: %d", len(problems)-correct)
	fmt.Println(string(colorGreen), correctAns)
	fmt.Println(string(colorRed), wrongAns)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}
