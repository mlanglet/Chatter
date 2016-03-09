package client

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	var quit string = "/quit"
	var newLine = '\n'
	var carrigeReturn = '\r'
	var cleanText []byte
	quit += string(carrigeReturn) + string(newLine)
	fmt.Println(quit)
	for {
		text, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		cleanText = formatInputByOS(text)
		sendMessage(cleanText)
	}
}

func sendMessage([]byte) (err error) {
	return err
}

func formatInputByOS(b []byte) (a []byte) {
	// Windows input detected
	if b[len(b)-2] == byte('\r') {
		a = make([]byte, len(b)-2)
	} else { // UNIX input detected
		a = make([]byte, len(b)-1)
	}

	copy(a, b)

	return
}

func printByteSlice(s []byte) {
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(os.Stdout, "%b ", s[i])
	}
	fmt.Fprintln(os.Stdout, "")
}
