package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// Struktur data untuk menampung satu baris sparepart
type Sparepart struct {
	Kode  string
	Nama  string
	Harga string
	Tipe  string
	Ket   string
}

func main() {
	// --- BAGIAN 1: MENANGKAP INPUT DARI TERMINAL ---

	// Kita siapkan 'bendera' (flags) agar bisa mengetik: -q="sesuatu"
	filePtr := flag.String("file", "spareparts.csv", "Nama file CSV database")
	kunciPtr := flag.String("q", "", "Kata kunci yang mau dicari (Wajib)")
	tipePtr := flag.String("by", "all", "Cari di mana? (pilihan: 'kode', 'nama', 'all')")

	flag.Parse() // Perintah untuk memproses input user

	// Cek apakah user lupa memasukkan kata kunci (-q)
	if *kunciPtr == "" {
		fmt.Println("⚠️  Error: Anda belum memasukkan kata kunci pencarian.")
		fmt.Println("👉 Contoh: go run main.go -q \"nama\"")
		os.Exit(1)
	}

	// --- BAGIAN 2: BACA FILE CSV ---

	// Buka file
	file, err := os.Open(*filePtr)
	if err != nil {
		fmt.Printf("❌ Gagal membuka file %s: %v\n", *filePtr, err)
		os.Exit(1)
	}
	defer file.Close()

	// Baca isi file
	reader := csv.NewReader(file)
	dataMentah, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("❌ Gagal membaca isi CSV: %v\n", err)
		os.Exit(1)
	}

	// --- BAGIAN 3: PROSES PENCARIAN ---

	var hasilPencarian []Sparepart
	kataKunci := strings.ToLower(*kunciPtr) // Ubah ke huruf kecil semua biar gampang ketemu

	// Loop baris per baris (mulai index 1 untuk melewati judul kolom)
	for i, baris := range dataMentah {
		if i == 0 {
			continue
		} // Lewati baris judul (Header)

		// Ambil data dari kolom CSV
		// baris[0] = Kode, baris[1] = Nama, baris[2] = Harga, baris[3] = Stok
		if len(baris) >= 4 {
			kode := baris[0]
			nama := baris[1]
			harga := baris[2]
			tipe := baris[3]

			// Siapkan variabel penanda ketemu atau tidak
			ketemu := false

			modePencarian := strings.ToLower(*tipePtr)

			// Logika pencarian
			switch modePencarian {
			case "kode":
				if strings.Contains(strings.ToLower(kode), kataKunci) {
					ketemu = true
				}
			case "nama":
				if strings.Contains(strings.ToLower(nama), kataKunci) {
					ketemu = true
				}
			case "harga":
				if strings.Contains(strings.ToLower(harga), kataKunci) {
					ketemu = true
				}
			case "tipe":
				if strings.Contains(strings.ToLower(tipe), kataKunci) {
					ketemu = true
				}
			default: // "all"
				if strings.Contains(strings.ToLower(kode), kataKunci) ||
					strings.Contains(strings.ToLower(nama), kataKunci) ||
					strings.Contains(strings.ToLower(harga), kataKunci) ||
					strings.Contains(strings.ToLower(tipe), kataKunci) {
					ketemu = true
				}
			}

			// Jika ketemu, masukkan ke list hasil
			if ketemu {
				s := Sparepart{Kode: kode, Nama: nama}
				if len(baris) > 2 {
					s.Harga = baris[2]
				}
				if len(baris) > 3 {
					s.Tipe = baris[3]
				}
				hasilPencarian = append(hasilPencarian, s)
			}
		}
	}

	// --- BAGIAN 4: TAMPILKAN HASIL ---

	if len(hasilPencarian) == 0 {
		fmt.Printf("Mencari '%s'... Tidak ditemukan.\n", *kunciPtr)
		return
	}
	fmt.Printf("✅ Ditemukan %d barang dengan kata kunci '%s':\n\n", len(hasilPencarian), *kunciPtr)

	// Gunakan Tabwriter supaya tabelnya lurus rapi
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	// HEADER
	fmt.Fprintln(w, "KODE\tNAMA BARANG\tHARGA\tTIPE\tKET")
	// data
	fmt.Fprintln(w, "----\t-----------\t-----\t--------\t--------")

	for _, brg := range hasilPencarian {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", brg.Kode, brg.Nama, brg.Harga, brg.Tipe, brg.Ket)
	}
	w.Flush() // Cetak ke layar
}
