package isbot

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Amount of bots that the regex should detect from the list
// If detection falls below this, the test fails
const threshold = 0.90

func TestCheckRegex(t *testing.T) {
	f, err := os.Open("user-agents-bots.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var trueCount, falseCount float64
	var falses []string

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ua := sc.Text()
		isBot := CheckRegex(ua)
		if isBot {
			trueCount++
			continue
		}
		falseCount++
		if falseCount < 20 {
			falses = append(falses, ua)
		}
	}

	detectionRate := (trueCount) / (trueCount + falseCount)
	fmt.Printf("Detection rate: %.2f, Threshold: %.2f\n", detectionRate, threshold)

	if detectionRate < threshold {
		t.Fatalf("True: %.0f, False: %.0f\n%s", trueCount, falseCount,
			strings.Join(falses, "\n"))
	}
}
