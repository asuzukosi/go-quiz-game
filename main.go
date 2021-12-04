package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	Question string
	Answer   string
}

type Quiz struct {
	Username  string
	Score     int
	Questions []Question
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func NewQuiz(username string, filename string) *Quiz {
	var questions []Question
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("An error occurred while opening file %v", err)
	}

	reader := bufio.NewReader(file)
	sentence, err := Readln(reader)
	for err == nil {
		data := strings.Split(sentence, ";;")

		question := Question{
			Question: data[0],
			Answer:   data[1],
		}

		questions = append(questions, question)
		sentence, err = Readln(reader)
	}
	return &Quiz{
		Username:  username,
		Questions: questions,
		Score:     0,
	}
}

func (q *Quiz) RunQuiz() {
	fmt.Println("Provide the ansers for each question.")
	for index, question := range q.Questions {
		fmt.Printf("Question %d:\n", index+1)
		var response string
		fmt.Println(question.Question)
		fmt.Scanln(&response)

		if strings.ToLower(strings.TrimSpace(response)) == strings.ToLower(strings.TrimSpace(question.Answer)) {
			q.Score += 1
		}

	}
	percentage := float64(q.Score) / float64(len(q.Questions)) * 100
	fmt.Printf("%s your score on this quiz is %.2f%% \n", q.Username, percentage)
}
func main() {
	var name string
	fmt.Println("Enter your username:")
	fmt.Scanln(&name)
	fmt.Printf("Hello world from %s!\n", name)

	var response string
ENTER_RESPONSE:
	fmt.Println("Would you like to play the game: yes or no?")
	fmt.Scanln(&response)

	switch strings.ToLower(response) {
	case "yes":
		fmt.Println("Welcome to the game.")
		quiz := NewQuiz(name, "quiz1.txt")
		quiz.RunQuiz()
		fmt.Println("Thank you for your time")

	case "no":
		fmt.Println("You are not playing the game")

	default:
		fmt.Println("You entered a wrong answer")

		goto ENTER_RESPONSE
	}
}
