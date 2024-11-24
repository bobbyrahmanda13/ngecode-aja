package main

import (
	"fmt"
	"time"
)

func textRunning2(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println()
}

func main() {

	ingat1 := "\n....\n\nKita Punya Cinta,\ntapi ALLAH punya ATURAN\nALLAH SWT melarang kita untuk PACARAN\ntapi ALLAH TIDAK MELARANG kita untuk MENCINTAI seseorang\n\n....\n"

	textRunning2(ingat1)
}
