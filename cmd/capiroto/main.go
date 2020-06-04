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

	// first input text
	username := "ritchie\r\n"
	buf.WriteString(username)
	pad(len(username), &buf)

	// second input text
	password := "secret\r\n"
	buf.WriteString(password)
	pad(len(password), &buf)

	// select example: 106 for No, 107 for Yes
	buf.Write([]byte{106, 13, 10})
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