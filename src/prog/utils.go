package red

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func FlushInputBuffer(reader *bufio.Reader) {
	*reader = *bufio.NewReader(os.Stdin)

	time.Sleep(50 * time.Millisecond)
}

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		ClearScreenFallback()
	}
}

func ClearScreenFallback() {
	fmt.Print("\033[2J\033[H")

	for i := 0; i < 50; i++ {
		fmt.Println()
	}
}
