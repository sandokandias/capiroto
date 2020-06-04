package main

import (
	"bytes"
	"os"
	"os/exec"
)

const(
	login = "C:\\Users\\sando\\go\\src\\github.com\\sandokandias\\capiroto\\cmd\\login\\main.go"
	promptui = 	"--promptui"
)

func main() {
	//cmd := exec.Command("go", "run", login)
	cmd := exec.Command("go", "run", login, promptui)

	buf := bytes.Buffer{}

	username := "ritchie\r\n"
	buf.WriteString(username)
	pu := make([]byte, 4096 - len(username))
	for i := 0; i < 4096 - len(username); i++ {
		pu[i] = 97
	}
	buf.Write(pu)

	password := "secret\r\n"
	buf.WriteString(password)
	pp := make([]byte, 4096 - len(password))
	for i := 0; i < 4096 - len(password); i++ {
		pp[i] = 97
	}
	buf.Write(pp)

	cmd.Stdin = &buf
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}