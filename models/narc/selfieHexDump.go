package narc

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func GetHexArray() []byte {
	file := "hex"
	output, err := readLines(file)
	if err != nil {
		return []byte("error")
	}
	randNum := randInt(0, 761270)
	return []byte(output[randNum])
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
