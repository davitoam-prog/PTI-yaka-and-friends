package main

import (
	"fmt"
	"strings"
)

func main() {
	// Data kurs mata uang (hardcoded)
	// 1 USD = ... mata uang lain
	kursMataUang := map[string]float64{
		"USD": 1.0,     // Base currency
		"IDR": 15500.0, // 1 USD = 15,500 IDR
		"EUR": 0.92,    // 1 USD = 0.92 EUR
		"GBP": 0.79,    // 1 USD = 0.79 GBP
		"JPY": 148.0,   // 1 USD = 148 JPY
		"SGD": 1.35,    // 1 USD = 1.35 SGD
		"MYR": 4.68,    // 1 USD = 4.68 MYR
	}

	fmt.Println("=== CONVERTER MATA UANG SEDERHANA ===")
	fmt.Println("Mata uang yang tersedia:")

	// Menampilkan semua mata uang yang tersedia
	for kode := range kursMataUang {
		fmt.Printf("- %s\n", kode)
	}
	fmt.Println()

	// Perulangan utama program
	for {
		var jumlah float64
		var dari, ke string

		// Input jumlah uang
		fmt.Print("Masukkan jumlah uang: ")
		fmt.Scan(&jumlah)

		if jumlah <= 0 {
			fmt.Println("ERROR: Jumlah harus lebih dari 0!")
			continue
		}

		// Input mata uang asal
		fmt.Print("Masukkan kode mata uang asal (contoh: USD): ")
		fmt.Scan(&dari)
		dari = strings.ToUpper(dari)

		// Validasi mata uang asal
		validDari := false
		for kode := range kursMataUang {
			if kode == dari {
				validDari = true
				break
			}
		}

		if !validDari {
			fmt.Printf("ERROR: Mata uang %s tidak tersedia!\n", dari)
			continue
		}

		// Input mata uang tujuan
		fmt.Print("Masukkan kode mata uang tujuan: ")
		fmt.Scan(&ke)
		ke = strings.ToUpper(ke)

		// Validasi mata uang tujuan
		validKe := false
		for kode := range kursMataUang {
			if kode == ke {
				validKe = true
				break
			}
		}

		if !validKe {
			fmt.Printf("ERROR: Mata uang %s tidak tersedia!\n", ke)
			continue
		}

		// Proses konversi
		hasil := 0.0

		if dari == ke {
			// Jika mata uang sama
			hasil = jumlah
		} else if dari == "USD" {
			// Konversi dari USD ke mata uang lain
			hasil = jumlah * kursMataUang[ke]
		} else if ke == "USD" {
			// Konversi ke USD dari mata uang lain
			hasil = jumlah / kursMataUang[dari]
		} else {
			// Konversi antara dua mata uang non-USD
			// Langkah 1: Konversi ke USD dulu
			dalamUSD := jumlah / kursMataUang[dari]
			// Langkah 2: Konversi dari USD ke mata uang tujuan
			hasil = dalamUSD * kursMataUang[ke]
		}

		// Menampilkan hasil konversi
		fmt.Println("\n=== HASIL KONVERSI ===")
		fmt.Printf("%.2f %s = %.2f %s\n\n", jumlah, dari, hasil, ke)

		// Menanyakan apakah ingin konversi lagi
		var pilihan string
		fmt.Print("Ingin konversi lagi? (y/n): ")
		fmt.Scan(&pilihan)

		if strings.ToLower(pilihan) != "y" {
			fmt.Println("Terima kasih telah menggunakan Converter Mata Uang!")
			break
		}

		fmt.Println()
	}
}
