package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	switch runtime.GOOS {

	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

	case "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	default:
		panic("OS not supported")
	}
}

func printHeader(attempts int, success int, error int) {
	fmt.Println(" ================== ")
	fmt.Println(" == NUMBER GUESS == ")
	fmt.Println(" ================== ")
	fmt.Printf("\n Attempts: %d", attempts)
	fmt.Printf("\n Successes %d \n Errors %d \n", success, error)
}

func printInstuctions() {
	fmt.Println(" This is a golang minigame  ")
	fmt.Println(" Choose a number between 0 and 100 ")
	fmt.Println(" ================== ")
}

func confirmRestart(hadSuccess bool) bool {
	fmt.Println("")
	if hadSuccess {
		fmt.Println("You gessed, congrats!")
	} else {
		fmt.Println("You failed, Better luck next time")
	}
	fmt.Println("")

	var userConfirmation string

	fmt.Println("Type 'y' to play again.")
	fmt.Print("> ")
	fmt.Scan(&userConfirmation)

	if userConfirmation == "y" {
		return true
	}

	return false
}

func getUserResponse() (int, error) {
	var response int
	fmt.Print("> ")
	_, error := fmt.Scan(&response)
	if error != nil {
		return 0, errors.New("error while processing your output, try Again")
	}

	return response, nil
}

func checkUserReponse(userResponse int) bool {
	correctAnswer := rand.IntN(101)
	return userResponse == correctAnswer
}

func main() {
	shouldPlayAgain := true

	attempts, successes, errors := 0, 0, 0

	for shouldPlayAgain {
		clearTerminal()
		printHeader(attempts, successes, errors)

		printInstuctions()

		response, error := getUserResponse()
		if error != nil {
			fmt.Println(error.Error())
			return
		}

		hadSuccess := checkUserReponse(response)
		shouldPlayAgain = confirmRestart(hadSuccess)

		attempts++
		if hadSuccess {
			successes++
		} else {
			errors++
		}
	}
}
