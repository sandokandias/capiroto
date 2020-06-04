package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

var prompt bool

func init() {
	flag.BoolVar(&prompt, "promptui", false, "read stdin with promptui")
}

func main() {
	flag.Parse()

	if prompt {
		runPrompt()
	} else {
		runDefault()
	}

}

func runDefault() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type the username: ")
	username, _ := reader.ReadString('\n')
	fmt.Println("Read:", username)

	fmt.Println("Type the password: ")
	pass, _ := reader.ReadString('\n')
	fmt.Println("Read:", pass)
}

func runPrompt()  {

	stdin := os.Stdin

	pu := promptui.Prompt{
		Label: "Username",
		Stdin: stdin,
	}
	username, err := pu.Run()
	if err != nil {
		log.Fatalf("Username prompt failed %v\n", err)
	} else {
		fmt.Println("Informed:", username)
	}

	pp := promptui.Prompt{
		Label: "Password",
		Stdin: stdin,
	}
	pass, err := pp.Run()
	if err != nil {
		log.Fatalf("Password prompt failed %v\n", err)
	} else{
		fmt.Println("Informed:", pass)
	}
}
