package main
import "fmt"

type date struct {
	hari  int
	bulan int
	tahun int
}

type Assessment struct {
	Tanggal       date
	MoodUtama     string
	Skor          int
	SkorMoodUtama int
	Catatan       string
}

type dataUser [1000]Assessment

var dataAssessment dataUser
var jumlahData int

func main() {
	isiDummySebulan(&dataAssessment, &jumlahData)
	mainMenu()
}

// Fungsi utama untuk menampilkan menu by Cathya
func mainMenu() {
	var pilihan int
	for pilihan != 7 {
		fmt.Println("------------------------------")
		fmt.Println(" 		M A I N  M E N U       ")
		fmt.Println("------------------------------")
		fmt.Println("1. Tampilkan Semua Data")
		fmt.Println("2. Cari Data Berdasarkan Tanggal")
		fmt.Println("3. Tambah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Dapatkan Rekomendasi")
		fmt.Println("6. Cari Data Berdasarkan Skor")
		fmt.Println("7. Keluar")
		fmt.Println("------------------------------")
		fmt.Print("Masukkan pilihan Anda (1-6): ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			// Mengurutkan data berdasarkan tanggal
			selectionSortTanggal(&dataAssessment, jumlahData)
			fmt.Println("\n--- Data Asli (Sebelum Diurutkan) ---")
			tampilkanDataTabel(dataAssessment, jumlahData)
			var sortPilihan int
			for sortPilihan != 3 {
				fmt.Println("\n--- Ingin Mengurutkan? ---")
				fmt.Println("1. Urutkan Ascending (Skor dari kecil ke besar)")
				fmt.Println("2. Urutkan Descending (Skor dari besar ke kecil)")
				fmt.Println("3. Kembali ke Menu Utama")
				fmt.Print("Masukkan pilihan (1/2/3): ")
				fmt.Scanln(&sortPilihan)

				if sortPilihan == 1 {
					insertionSortData(&dataAssessment, jumlahData, true)
					fmt.Println("\n--- Data Setelah Diurutkan Ascending ---")
					tampilkanDataTabel(dataAssessment, jumlahData)
				} else if sortPilihan == 2 {
					insertionSortData(&dataAssessment, jumlahData, false)
					fmt.Println("\n--- Data Setelah Diurutkan Descending ---")
					tampilkanDataTabel(dataAssessment, jumlahData)
				} else if sortPilihan != 3 {
					fmt.Println("Pilihan tidak valid.")
				}
			}
			isiDummySebulan(&dataAssessment, &jumlahData)
		} else if pilihan == 2 {
			var d date
			fmt.Print("Masukkan tanggal (dd mm yyyy): ")
			fmt.Scan(&d.hari, &d.bulan, &d.tahun)
			cariDataBerdasarkanTanggal(dataAssessment, jumlahData, d)
			kembaliKeMenu()
		} else if pilihan == 3 {
			var d date
			var mood string
			var skor, skorMood int
			var catatan string

			fmt.Print("Tanggal (dd mm yyyy): ")
			fmt.Scan(&d.hari, &d.bulan, &d.tahun)
			fmt.Print("Mood utama: ")
			fmt.Scan(&mood)
			fmt.Print("Skor umum (0-10): ")
			fmt.Scan(&skor)
			fmt.Print("Skor mood utama (0-10): ")
			fmt.Scan(&skorMood)
			fmt.Print("Catatan (boleh kosong): ")
			fmt.Scanln()
			fmt.Scanln(&catatan)

			tambahAssessment(d, skor, skorMood, mood, catatan)
			kembaliKeMenu()
		} else if pilihan == 4 {
			var d date
			fmt.Print("Masukkan tanggal yang akan dihapus (dd mm yyyy): ")
			fmt.Scan(&d.hari, &d.bulan, &d.tahun)
			hapusData(d)
			kembaliKeMenu()
		} else if pilihan == 5 {
			var skor int
			fmt.Print("Masukkan skor self-assessment (0-10): ")
			fmt.Scan(&skor)
			fmt.Println("Rekomendasi:", beriRekomendasi(skor))
			kembaliKeMenu()
		} else if pilihan == 6 {
			var skorCari int
  			fmt.Print("Masukkan skor yang ingin dicari: ")
    		fmt.Scan(&skorCari)

    		//pastikan data telah diurutkan (terurut naik)
    		insertionSortData(&dataAssessment, jumlahData, true)

    		indeks := binarySearchSkor(dataAssessment, jumlahData, skorCari)
    		if indeks != -1 {
        		fmt.Println("Data dengan skor ditemukan:")
        		tampilkanSatuData(dataAssessment[indeks])
    		} else {
        		fmt.Println("Skor tidak ditemukan dalam data.")
    		}
    		kembaliKeMenu()
		} else if pilihan == 7 {
   			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Tambahan: kembali ke menu dengan enter
func kembaliKeMenu() {
	fmt.Print("\nTekan ENTER untuk kembali ke menu utama...")
	fmt.Scanln()
	fmt.Scanln()
}

// Fungsi isi dummy by Wijdan
func isiDummySebulan(a *dataUser, jumlah *int) {
	moodList := [30]string {
		"Senang", "Senang", "Netral", "Cemas", "Cemas",
		"Sedih", "Sedih", "Senang", "Netral", "Marah",
		"Takut", "Netral", "Senang", "Senang", "Cemas",
		"Sedih", "Netral", "Marah", "Takut", "Senang",
		"Netral", "Cemas", "Sedih", "Senang", "Marah",
		"Takut", "Senang", "Sedih", "Marah", "Senang",
	}

	*jumlah = 0
	for i := 1; i <= 30; i++ {
		a[*jumlah] = Assessment{
			Tanggal:       date{i, 4, 2025},
			MoodUtama:     moodList[i-1],
			SkorMoodUtama: (i*2)%10 + 1,
			Skor:          (i*3)%10 + 1,
			Catatan:       "Tidak ada catatan",
		}
		*jumlah++
	}
	a[28].Catatan = "Betmut parah gw, dikasi tgjwb situ malah ngilang"
}

// Fungsi tampilkan dalam bentuk tabel by Wijdan
func tampilkanDataTabel(a dataUser, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Printf("| %-10s | %-10s | %-5s | %-10s | %-40s |\n", "Tanggal", "Mood", "Skor", "Skor Mood", "Catatan")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		t := a[i].Tanggal
		fmt.Printf("| %02d/%02d/%04d | %-10s | %-5d | %-10d | %-40s |\n",
			t.hari, t.bulan, t.tahun,
			a[i].MoodUtama,
			a[i].Skor,
			a[i].SkorMoodUtama,
			ringkasCatatan(a[i].Catatan, 40))
	}
	fmt.Println("-------------------------------------------------------------------------------------------")
}

func ringkasCatatan(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}

// Tambah data baru by Nada
func tambahAssessment(t date, skor int, skorMood int, mood string, catatan string) {
	if jumlahData < 1000 {
		dataAssessment[jumlahData] = Assessment{
			Tanggal:       t,
			Skor:          skor,
			SkorMoodUtama: skorMood,
			MoodUtama:     mood,
			Catatan:       catatan,
		}
		jumlahData++
		fmt.Println("Assessment berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas data penuh.")
	}
}

// Cari data (tanggal) by Wijdan
// sementara kami buat ini dengan metode sequential search nanti akan kami optimasi
// dengan binary search
func cariDataBerdasarkanTanggal(a dataUser, jumlah int, d date) {
	found := false
	for i := 0; i < jumlah; i++ {
		if a[i].Tanggal == d {
			tampilkanSatuData(a[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan untuk tanggal tersebut.")
	}
}

// Tampilkan satu data by Wijdan
func tampilkanSatuData(a Assessment) {
	fmt.Printf("Tanggal: %02d/%02d/%d\n", a.Tanggal.hari, a.Tanggal.bulan, a.Tanggal.tahun)
	fmt.Println("Mood Utama:", a.MoodUtama)
	fmt.Printf("Skor Mood Utama: %d/10\n", a.SkorMoodUtama)
	fmt.Printf("Skor Umum: %d/10\n", a.Skor)
	fmt.Println("Catatan:", a.Catatan)
	fmt.Println("--------------------------")
}

// Hapus data by Nada
func hapusData(t date) {
	idx := -1
	for i := 0; i < jumlahData; i++ {
		if dataAssessment[i].Tanggal == t {
			idx = i
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahData-1; i++ {
		dataAssessment[i] = dataAssessment[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus.")
}

// Fungsi rekomendasi (masih tahap pengembangan) by Cathya
func beriRekomendasi(skor int) string {
	switch {
	case skor <= 3:
		return "Luangkan waktu untuk istirahat."
	case skor <= 6:
		return "Cobalah teknik pernapasan atau meditasi."
	case skor <= 8:
		return "Pertahankan keseimbangan dan tetap waspada terhadap stres."
	default:
		return "Kondisi Anda baik, lanjutkan dengan aktivitas positif <3"
	}
}

// Fungsi insertion sort untuk mengurutkan data berdasarkan skor by Wijdan
func insertionSortData(a *dataUser, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		temp := a[i]
		j := i - 1
		if ascending {
			for j >= 0 && a[j].Skor > temp.Skor {
				a[j+1] = a[j]
				j--
			}
		} else {
			for j >= 0 && a[j].Skor < temp.Skor {
				a[j+1] = a[j]
				j--
			}
		}
		a[j+1] = temp
	}
}

//binarySearchSkor mencari skor tertentu dalam data yang telah diurutkan by Nada
// nanti mungkin akan kami buat fitur pencarian skor mingguan dengan sequential search
// ditukar dengan fitur cari data tanggal yang akan beralih ke binary search
func binarySearchSkor(a dataUser, jumlah int, targetSkor int) int {
    left := 0
    right := jumlah - 1

    for left <= right {
        mid := (left + right) / 2
        if a[mid].Skor == targetSkor {
            return mid // Skor ditemukan
        } else if a[mid].Skor < targetSkor {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1 //skor tidak ditemukan
}

// Fungsi selection sort untuk mengurutkan data berdasarkan tanggal by Wijdan
func selectionSortTanggal(a *dataUser, jumlah int) {
	var temp Assessment

	for i := 0; i < jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlah; j++ {
			if lebihAwal(a[j].Tanggal, a[minIdx].Tanggal) {
				minIdx = j
			}
		}
		temp = a[i]
		a[i] = a[minIdx]
		a[minIdx] = temp
	}
}

// Fungsi pembanding: true jika t1 lebih awal dari t2
func lebihAwal(t1, t2 date) bool {
	if t1.tahun < t2.tahun {
		return true
	} else if t1.tahun == t2.tahun {
		if t1.bulan < t2.bulan {
			return true
		} else if t1.bulan == t2.bulan {
			return t1.hari < t2.hari
		}
	}
	return false
}

// Coming soon: Fungsi Implementasi Nilai Ekstrim by Cathya
// Comming soon: Fungsi untuk menampilkan grafik by ??