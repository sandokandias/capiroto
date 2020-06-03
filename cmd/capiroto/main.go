package main

import (
	"bytes"
	"fmt"
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
	buf.Write([]byte("ritchie\n"))
	buf.Write([]byte("secret\n"))
	/*for i := 15; i < 4096; i++ {
		buf.WriteByte('a')
	}*/

	fmt.Println("Buffer len:", buf.Len())

	cmd.Stdin = &buf
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}