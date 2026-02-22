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

	fmt.Print("Git Message : ")
	// 1. Meminta input pesan (seperti read -p)
	gitMessage, _ := reader.ReadString('\n')
	gitMessage = strings.TrimSpace(gitMessage)

	// 2. Menjalankan git add .
	exec.Command("git", "add", ".").Run()

	// 3. Menjalankan git commit -m "pesan"
	cmd := exec.Command("git", "commit", "-m", gitMessage)

	// Menghubungkan output git ke terminal kita agar terlihat hasilnya
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Gagal commit: %v\n", err)
	} else {
		fmt.Println("Berhasil commit!")
	}

	fmt.Println("Sedang melakukan push...")
	pushCmd := exec.Command("git", "push")
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr

	errPush := pushCmd.Run()

	if errPush != nil {
		fmt.Printf("Gagal push: %v\n", errPush)
	} else {
		fmt.Println("Berhasil push ke remote!")
	}
}
