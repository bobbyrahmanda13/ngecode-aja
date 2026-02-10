package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// Target yang ingin dicek
	target := "google.com"

	fmt.Printf("Memulai monitoring jaringan ke: %s\n", target)
	fmt.Println("Tekan Ctrl+C untuk menghentikan.")
	fmt.Println("-------------------------------------------")

	for {
		// Membuat context dengan timeout 5 detik
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

		fmt.Printf("[%s] Menjalankan ping...\n", time.Now().Format("15:04:05"))

		// Menjalankan perintah ping sistem
		// -c 5 (Linux/macOS) mengirim 5 paket.
		// Jika di Windows, gunakan "ping", "-n", "5", target
		cmd := exec.CommandContext(ctx, "ping", "-c", "5", target)

		// Menangkap output
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("Peringatan: Jaringan bermasalah atau timeout. Error: %v\n", err)
		} else {
			fmt.Println(string(out))
		}

		// Membersihkan context
		cancel()

		fmt.Println("Menunggu sebelum pengulangan berikutnya...")
		fmt.Println("-------------------------------------------")

		// Jeda singkat sebelum memulai loop lagi agar tidak membebani sistem
		time.Sleep(50 * time.Second)
	}
}
