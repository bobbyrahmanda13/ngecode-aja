package main

import (
	"fmt"
	"time"
)

func textRunning(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Println()
}

func main() {

	ingat1 := "\nPergi menggunakan BAJU PUTIH \nPulang menggunakan BAJU MERAH \nDia rela di lempar dengan BATU \nDi lempar dengan KAYU \nDi lempar dengan TULANG ONTA \nUNTUK MENYEBARKAN ISLAM tapi \nislam sekarang semakin asing di mata umatNya sendiri, lebih condong dengan BUDAYA BARAT \napakah dirimu tidak malu kepada NABI MUHAMMAD SAW"
	ingat2 := "\nBanyak yang bangga melakukan dosa, dengan dalih saya tidak MUNAFIK"
	ingat3 := "\nBanyak Orang Tua yang MALU ANAKNYA-NYA di nikahi dengan cara yang SEDERHANA \nTapi Tidak pernah MALU Anak-Nya di Pacari secara GRATISSSSS"
	ingat4 := "\nMau tau bangkai BERNYAWA ? \nDia yang mengaku Islam, dan tahu arah Kiblat tapi MALAS melaksanakan Sholat"
	ingat5 := "\nRajin SHOLAT tapi terus MAKSIAT"
	ingat6 := "\n'Barang siapa yang melaksanakan sholat, lantas sholat tersebut tidak mencegah dari perbuatan keji dan mungkar,\nmaka ia hanya akan semakin menjauh dari ALLAH SWT'"
	ingat7 := "\nYang dibesarkan itu IMAN bukan KESOMBONGAN \nYang diluaskan itu SABAR bukan keluhan \nYang diperbaiki itu SHOLAT bukan GAYA HIDUP \nyang diperbaiki itu AKHLAK bukan FISIK \nDan yang perlu dikhawatirkan itu KEMATIAN bukan tentang JODOH"

	textRunning(ingat1)
	textRunning(ingat2)
	textRunning(ingat3)
	textRunning(ingat4)
	textRunning(ingat5)
	textRunning(ingat6)
	textRunning(ingat7)
}
