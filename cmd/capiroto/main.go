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
	pad(len(username), &buf)

	password := "secret\r\n"
	buf.WriteString(password)
	pad(len(password), &buf)

	buf.Write([]byte{14, 13, 10})
	pad(3, &buf)

	cmd.Stdin = &buf
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}

func pad(siz int, buf *bytes.Buffer) {
	pu := make([]byte, 4096-siz)
	for i := 0; i < 4096-siz; i++ {
		pu[i] = 97
	}
	buf.Write(pu)
}