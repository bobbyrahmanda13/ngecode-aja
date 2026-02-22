package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Meminta input pesan (seperti read -p)
	fmt.Print("Git Message : ")
	gitMessage, _ := reader.ReadString('\n')
	gitMessage = strings.TrimSpace(gitMessage)

	// 2. Menjalankan git add .
	exec.Command("git", "add", ".").Run()

	// 3. Menjalankan git commit -m "pesan"
	cmd := exec.Command("git", "commit", "-m", gitMessage)

	// Menghubungkan output git ke terminal kita agar terlihat hasilnya
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 4. Menjalankan git push
	exec.Command("git", "push").Run()

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Gagal commit: %v\n", err)
	} else {
		fmt.Println("Berhasil commit!")
	}
}
