package hg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
	UsedLetters      []string
}

func StartGame(dictionaryFile string) {
	rand.Seed(time.Now().UnixNano())

	words, err := loadDictionary(dictionaryFile)
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return
	}

	hangmanPositions, err := loadHangmanPositions("hangman.txt")
	if err != nil {
		fmt.Println("Warning: hangman.txt not found, proceeding without visual representation")
	}

	data := HangManData{
		Attempts:         10,
		HangmanPositions: hangmanPositions,
		UsedLetters:      []string{},
	}

	data.ToFind = chooseRandomWord(words)

	data.Word = initializeWord(data.ToFind)

	fmt.Println("Good Luck, you have", data.Attempts, "attempts.")
	fmt.Println(data.Word)

	scanner := bufio.NewScanner(os.Stdin)
	for data.Attempts > 0 && strings.Contains(data.Word, "_") {
		fmt.Print("Choose: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if len(input) == 0 {
			continue
		}

		handlePlayerInput(input, &data)

		fmt.Println(data.Word)

		if !strings.Contains(data.Word, "_") {
			fmt.Println("Congrats!")
			return
		}
	}

	if data.Attempts <= 0 {
		fmt.Println("Game over! The word was:", data.ToFind)
	}
}

func loadDictionary(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(content), "\n")
	var filteredWords []string
	for _, word := range words {
		if len(strings.TrimSpace(word)) > 0 {
			filteredWords = append(filteredWords, strings.TrimSpace(word))
		}
	}

	return filteredWords, nil
}

func loadHangmanPositions(filename string) ([10]string, error) {
	var positions [10]string

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return positions, err
	}

	lines := strings.Split(string(content), "\n")
	positionHeight := 8

	for i := 0; i < 10 && i*positionHeight < len(lines); i++ {
		start := i * positionHeight
		end := start + positionHeight
		if end > len(lines) {
			end = len(lines)
		}
		positions[i] = strings.Join(lines[start:end], "\n")
	}

	return positions, nil
}

func chooseRandomWord(words []string) string {
	if len(words) == 0 {
		return "hello"
	}
	return words[rand.Intn(len(words))]
}

func initializeWord(word string) string {
	mask := make([]rune, len(word))
	for i := range mask {
		mask[i] = '_'
	}

	numToReveal := len(word)/2 - 1
	if numToReveal <= 0 {
		numToReveal = 1
	}

	indices := rand.Perm(len(word))
	for i := 0; i < numToReveal && i < len(indices); i++ {
		idx := indices[i]
		mask[idx] = rune(word[idx])
	}

	return string(mask)
}

func isLetterUsed(letter string, data *HangManData) bool {
	for _, usedLetter := range data.UsedLetters {
		if usedLetter == letter {
			return true
		}
	}
	return false
}

func handlePlayerInput(input string, data *HangManData) {
	if len(input) > 1 {
		handleWordGuess(input, data)
		return
	}

	letter := strings.ToUpper(input)

	if isLetterUsed(letter, data) {
		fmt.Println("You already tried that letter!")
		return
	}

	data.UsedLetters = append(data.UsedLetters, letter)

	word := strings.ToUpper(data.ToFind)
	letterFound := false

	for i, char := range word {
		if string(char) == letter {
			data.Word = data.Word[:i] + letter + data.Word[i+1:]
			letterFound = true
		}
	}

	if !letterFound {
		data.Attempts--
		fmt.Println("Not present in the word,", data.Attempts, "attempts remaining")
		if data.Attempts < 10 && data.Attempts >= 0 && len(data.HangmanPositions[9-data.Attempts]) > 0 {
			fmt.Println(data.HangmanPositions[9-data.Attempts])
		}
	}
}

func handleWordGuess(guess string, data *HangManData) {
	if strings.ToUpper(guess) == strings.ToUpper(data.ToFind) {
		data.Word = strings.ToUpper(data.ToFind)
	} else {
		data.Attempts -= 2
		fmt.Println("Wrong word! You lose 2 attempts,", data.Attempts, "attempts remaining")
		if data.Attempts < 10 && data.Attempts >= 0 && len(data.HangmanPositions[9-data.Attempts]) > 0 {
			fmt.Println(data.HangmanPositions[9-data.Attempts])
		}
	}
}
