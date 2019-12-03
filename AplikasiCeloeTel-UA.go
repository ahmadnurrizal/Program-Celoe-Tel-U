package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

const N int = 1000

/////////////////////////////////
type tabData [N]dataMaha
type tabJawaban [N]string
type dataMaha struct {
	nama, nim string
	jawaban   tabJawaban
	nilai     int
}

///////////////////////////////
type tabForum [N]dataForum
type forum [N]string
type dataForum struct {
	nama  forum
	komen forum
}

////////////////////////////////
type tabAksi [N]aksi
type aksi struct {
	input, edit, hapus string
	jumSoal            int
}

func scanner() string { // NEXT DATA
	var nextData string
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	nextData = scan.Text()
	return nextData
}

func clear() { // MEMBERSIHKAN LAYAR CMD
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

// FUNCTION INPUT SOAL [ DOSEN ] ===========================================================================================
func input(T *tabAksi, jumSoal *int) {
	var i int
	if *jumSoal == 0 {
		*jumSoal = 1
		i = 1
	} else {
		i = *jumSoal + 1
	}
	fmt.Print(i, ". ")
	fmt.Scan(&T[i].input)
	T[i].input = T[i].input + " " + scanner()
	for i < N-1 && T[i].input != "selesai " {
		i++
		fmt.Print(i, ". ")
		fmt.Scan(&T[i].input)
		T[i].input = T[i].input + " " + scanner()
	}
	*jumSoal = i - 1
}

// FUNCTION EDIT SOAL [ DOSEN ] ====================================================================================
func edit(T *tabAksi) {
	var nomor int

	fmt.Println("\n*** TEKAN '0' JIKA TELAH SELESAI")
	fmt.Print("NOMOR PILIHAN : ")
	fmt.Scanln(&nomor)
	for nomor != 0 {
		fmt.Print("SOAL BARU : ")
		fmt.Scan(&T[nomor].input)
		fmt.Println("")
		T[nomor].input = T[nomor].input + " " + scanner()
		fmt.Print("NOMOR PILIHAN : ")
		fmt.Scanln(&nomor)
	}
}

// FUNCTION HAPUS SOAL [ DOSEN ] ===================================================================================
func hapus(T *tabAksi) {
	var nomor int
	fmt.Println("\n*** TEKAN '0' JIKA TELAH SELESAI")
	fmt.Print("NOMOR PILIHAN : ")
	fmt.Scanln(&nomor)
	for nomor != 0 {
		T[nomor].input = " "
		fmt.Print("NOMOR PILIHAN : ")
		fmt.Scanln(&nomor)
	}
}

// FUNCTION TAMPILKAN SOAL [ DOSEN ]  ==========================================================================
func tampilkanSoal(T *tabAksi, jumSoal int) {
	var i int = 1
	for i <= jumSoal {
		fmt.Println("TUGAS ", i, ":", T[i].input)
		i++
	}
}

// FUNCTION INPUT NILAI [ DOSEN ]  ==============================================================
func inputNilai(U *tabData, jumMaha int) {
	var i int
	i = 1
	for i <= jumMaha {
		fmt.Print(i, ". ", U[i].nama, " ", U[i].nim)
		fmt.Println(" ")
		i = i + 1
	}
	fmt.Println("\nPILIH 'NOMOR MAHASISWA' UNTUK MENGINPUT NILAINYA \n" +
		"*** TEKAN '0' JIKA TELAH SELESAI\n")
	for i != 0 {
		fmt.Print("NOMOR MAHASISWA : ")
		fmt.Scanln(&i)
		if i != 0 {
			fmt.Print("NILAI ", U[i].nama, " : ")
			fmt.Scanln(&U[i].nilai)
		}
	}
}

// FUNCTION INPUT FORUM [ DOSEN & MAHASISWA]  ================================================================
func inputForum(F *tabForum, mainUser, dosen string, jumKomen *int, tipeUser int) {
	var i int
	if *jumKomen == 0 {
		*jumKomen = 0
		i = 1
	} else {
		i = *jumKomen + 1
	}
	if tipeUser == 1 {
		F[i].nama[i] = dosen
	} else if tipeUser == 2 {
		F[i].nama[i] = mainUser
	}
	fmt.Print(F[i].nama[i], " : ")
	fmt.Scan(&F[i].komen[i])
	F[i].komen[i] = F[i].komen[i] + " " + scanner()
	if F[i].komen[i] != "selesai " {
		for F[i].komen[i] != "selesai " {
			i = i + 1
			if tipeUser == 1 {
				F[i].nama[i] = dosen
			} else if tipeUser == 2 {
				F[i].nama[i] = mainUser
			}
			fmt.Print(F[i].nama[i], " : ")
			fmt.Scan(&F[i].komen[i])
			F[i].komen[i] = F[i].komen[i] + " " + scanner()
			*jumKomen = i - 1
		}
	}

}

// FUNCTION TAMPILKAN FORUM [ DOSEN & MAHASISWA ]  ==============================================================
func tampilkanForum(F *tabForum, jumKomen int) {
	var i int

	fmt.Println("==================================== \n" +
		"======== SILAHKAN BERDISKUSI ======= \n" +
		"============== FORUM =============== \n" +
		"\n*** KETIK 'selesai' JIKA TELAH SELESAI")
	if jumKomen == 0 {
		fmt.Println("\nMAAF, BELUM ADA YANG BERDISKUSI \n" +
			"SILAHKAN KOMEN UNTUK MEMULAI DISKUSI\n")
	} else {
		i = 1
		for i <= jumKomen {
			fmt.Println(F[i].nama[i], ":", F[i].komen[i])
			i = i + 1
		}
	}
}

// TAMPILKAN NILAI [ DOSEN & MAHASISWA ] ============================================
func tampilkanNilai(U tabData, jumMaha int) {
	var i, pass, max int
	var temp dataMaha
	i = 1
	pass = 1
	for pass <= jumMaha-1 {
		max = pass
		i = pass + 1
		for i <= jumMaha {
			if U[max].nilai < U[i].nilai {
				max = i
			}
			i = i + 1
		}
		temp = U[pass]
		U[pass] = U[max]
		U[max] = temp
		pass = pass + 1
	}
	i = 1
	fmt.Println("==================================== \n" +
		"============ DAFTAR NILAI ========== \n" +
		"============= MAHASISWA ============ \n")
	for i <= jumMaha {
		fmt.Println(i, ".", U[i].nama, " ", U[i].nim, ":", U[i].nilai)
		i = i + 1
	}

}

// JAWAB [ MAHASISWA ] ============================================================================
func jawab(U *tabData, jumSoal, jumMaha, idxMaha int, markMaha bool) {
	var i int

	fmt.Print("\n*** TEKAN '0' JIKA SELESAI \n " + "PILIH NOMOR SOAL : ")
	fmt.Scanln(&i)
	if markMaha == false {
		for i != 0 && i <= jumSoal {
			fmt.Print("JAWABAN : ")
			fmt.Scan(&U[jumMaha].jawaban[i])
			fmt.Println("")
			U[jumMaha].jawaban[i] = U[jumMaha].jawaban[i] + " " + scanner()
			fmt.Print("PILIH SOAL : ")
			fmt.Scanln(&i)
		}
	} else {
		for i != 0 && i <= jumSoal {
			fmt.Print("JAWABAN : ")
			fmt.Scan(&U[idxMaha].jawaban[i])
			fmt.Println("")
			U[idxMaha].jawaban[i] = U[idxMaha].jawaban[i] + " " + scanner()
			fmt.Print("PILIH SOAL : ")
			fmt.Scanln(&i)
		}
	}
	clear()
}

// FUNCTION TAMPILKAN USER
func user(U *tabData, jumMaha int) {
	var i int
	i = 1
	fmt.Println("*** DAFTAR NAMA MAHASISWA ***")
	for i <= jumMaha {
		fmt.Println(i, ".", U[i].nama, " ", U[i].nim)
		i = i + 1
	}
}

//  FUNCTION TAMPILKAN JAWABAN
func jawaban(U *tabData, jumSoal int) {
	var i, nomor int

	fmt.Print("\nPILIH NOMOR MAHASISWA : ")
	fmt.Scanln(&nomor)
	for nomor != 0 {
		i = 1
		for i <= jumSoal {
			if U[nomor].jawaban[i] != "" {
				fmt.Println(i, ".", U[nomor].jawaban[i])
			} else {
				fmt.Println(i, ". BELUM TERJAWAB")
			}
			i = i + 1
		}
		fmt.Print("PILIH NOMOR MAHASISWA : ")
		fmt.Scanln(&nomor)
	}
}

func main() {

	var (
		userT tabData
		userQ tabData
		tugas tabAksi
		quis  tabAksi
		forum tabForum
	)
	var (
		jumTugas, jumQuis, tipeUser        int
		jumMaha, jumKomen, idxMaha         int
		jumUser, tipe, aksi, i, markDosen  int
		mainUser, next, nim, dosen, matkul string
		markMaha                           bool
	)
	clear()

	fmt.Println("========================================================================== \n" +
		"================= SELAMAT DATANG DI PROGRAM Celoe Tel-U ==================\n" +
		"========================================================================== \n")

	fmt.Println("======================================= \n" +
		"=========== SILAHKAN LOGIN ============ \n" +
		"======================================= \n")
	fmt.Println("LOGIN SEBAGAI : \n" + "1. DOSEN \n" + "2. MAHASISWA \n")
	fmt.Println("*** TEKAN '999' JIKA INGIN KELUAR DARI PROGRAM")
	fmt.Print("===> ")
	fmt.Scanln(&tipeUser)
	if tipeUser == 2 {
		clear()
		fmt.Println("============ MASUKKAN USERNAME DAN NIM MAHASISWA ============")
		fmt.Print("|| USERNAME : ")
		fmt.Scan(&mainUser)
		mainUser = mainUser + " " + scanner()
		fmt.Print("|| NIM : ")
		fmt.Scanln(&nim)
	} else if tipeUser == 1 {
		clear()
		if markDosen == 0 {
			fmt.Println("============ MASUKKAN NAMA DOSEN ============")
			fmt.Print("NAMA DOSEN : ")
			fmt.Scan(&dosen)
			dosen = dosen + " " + scanner()
			markDosen = markDosen + 1
			fmt.Print("MATA KULIAH : ")
			fmt.Scan(&matkul)
			matkul = matkul + " " + scanner()

		}
		mainUser = "dosen "

	} else if tipeUser == 999 {
		mainUser = "00 "
	}
	for mainUser != "00 " && nim != "00 " {
		clear()
		jumUser = jumUser + 1 // untuk forum
		forum[jumUser].nama[i] = mainUser

		if mainUser != "dosen " {
			jumMaha = jumMaha + 1
			userT[jumMaha].nama = mainUser
			userQ[jumMaha].nama = mainUser
			userT[jumMaha].nim = nim
			userQ[jumMaha].nim = nim

			i = 1
			markMaha = false
			for i <= jumMaha && !markMaha && jumMaha != 1 {
				if userT[i-1].nim == nim {
					markMaha = true
					jumMaha = jumMaha - 1
					idxMaha = i - 1
				}
				i = i + 1
			}
			i = 1
		}
		tipe = 0
		for tipe != 5 {
			if mainUser == "dosen " {

				fmt.Println("\n========= PILIH AKSI [ DOSEN ] ========== \n" +
					"1. TUGAS \n" + "2. QUIS \n" +
					"3. FORUM \n" + "4. RIWAYAT LOGIN \n" + "5. LOGOUT")
			} else {
				fmt.Println("\n======== PILIH AKSI [ MAHASISWA ] ======= \n" +
					"1. TUGAS \n" + "2. QUIS \n" +
					"3. FORUM \n" + "4. RIWAYAT LOGIN \n" + "5. DOSEN & MATKUL\n" + "6. LOGOUT")
			}
			fmt.Print("\n ===> ")
			fmt.Scanln(&tipe)
			clear()
			aksi = 0
			// DOSEN %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
			// TUGAS ==================================================
			if mainUser == "dosen " {
				if tipe == 1 {

					for aksi != 8 {
						fmt.Println("\n========= PILIH AKSI [ DOSEN ] ========== \n" +
							"1. TAMBAH SOAL TUGAS\n" + "2. EDIT SOAL TUGAS\n" + "3. HAPUS SOAL TUGAS \n" +
							"4. LIHAT SOAL TUGAS \n" + "5. LIHAT JAWABAN TUGAS MAHASISWA \n" + "6. INPUT NILAI TUGAS \n" +
							"7. LIHAT NILAI TUGAS MAHASISWA \n" + "8. KEMBALI")
						fmt.Print("===> ")
						fmt.Scanln(&aksi)
						clear()
						if aksi == 1 {
							fmt.Println("\n==================================== \n" +
								"====== SILAHKAN MENAMBAH SOAL ====== \n" +
								"============== TUGAS =============== \n")
							input(&tugas, &jumTugas) // Menambah Soal Tugas

						} else if aksi == 2 {
							fmt.Println("\n========= DAFTAR TUGAS [ DOSEN ] ========== \n")
							tampilkanSoal(&tugas, jumTugas)
							fmt.Println("==================================== \n" +
								"====== SILAHKAN MENGEDIT SOAL ====== \n" +
								"============== TUGAS =============== \n")
							edit(&tugas) // Edit Soal Tugas

						} else if aksi == 3 {
							fmt.Println("\n========= DAFTAR TUGAS [ DOSEN ] ========== \n")
							tampilkanSoal(&tugas, jumTugas)
							fmt.Println("==================================== \n" +
								"====== SILAHKAN MENGEDIT SOAL ====== \n" +
								"============== TUGAS =============== \n")
							hapus(&tugas) // Fungsi Hapus Tugas

						} else if aksi == 6 {
							fmt.Println("\n====== DAFTAR NAMA MAHASISWA [ DOSEN ] ======= \n")
							inputNilai(&userT, jumMaha) // Input Nilai
						} else if aksi == 5 {
							if jumMaha == 0 {
								fmt.Println("MAAF, BELUM ADA MAHASISWA YANG MENJAWAB SOAL :(")
							} else {
								user(&userT, jumMaha)
								jawaban(&userT, jumTugas) // tampilkan jawaban
							}

						} else if aksi == 7 {
							tampilkanNilai(userT, jumMaha)
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
						} else if aksi == 4 {
							fmt.Println("\n========= DAFTAR TUGAS [ DOSEN ] ========== \n")
							tampilkanSoal(&tugas, jumTugas)
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
						}
						clear()

					}
					// QUIS  ==================================================================
				} else if tipe == 2 {
					for aksi != 8 {
						fmt.Println("\n========= PILIH AKSI [ DOSEN ] ========== \n" +
							"1. TAMBAH SOAL QUIS\n" + "2. EDIT SOAL QUIS\n" + "3. HAPUS SOAL QUIS \n" +
							"4. LIHAT SOAL QUIS \n" + "5. LIHAT JAWABAN QUIS MAHASISWA \n" + "6. INPUT NILAI QUIS \n" +
							"7. LIHAT NILAI QUIS MAHASISWA \n" + "8. KEMBALI")
						fmt.Print("====> ")
						fmt.Scanln(&aksi)
						clear()
						if aksi == 1 {
							fmt.Println("\n==================================== \n" +
								"====== SILAHKAN MENAMBAH SOAL ====== \n" +
								"=============== QUIS =============== \n" +
								"*** KETIK 'selesai' JIKA TELAH SELESAI \n")
							input(&quis, &jumQuis) // Menambah Soal Quis

						} else if aksi == 2 {
							fmt.Println("\n========= DAFTAR QUIS [ DOSEN ] ========== \n")
							tampilkanSoal(&quis, jumQuis)
							fmt.Println("\n==================================== \n" +
								"====== SILAHKAN MENGEDIT SOAL ====== \n" +
								"=============== QUIS =============== \n")
							edit(&quis) // Edit soal Quis

						} else if aksi == 3 {
							fmt.Println("\n========= DAFTAR QUIS [ DOSEN ] ========== \n")
							tampilkanSoal(&quis, jumQuis)
							fmt.Println("==================================== \n" +
								"====== SILAHKAN MENGHAPUS SOAL ====== \n" +
								"=============== QUIS =============== \n")
							hapus(&quis) // Fungsi Hapus Quis

						} else if aksi == 5 {
							fmt.Println("\n==================================== \n" +
								"========= JAWABAN MAHASISWA ======== \n" +
								"=============== QUIS =============== \n")
							if jumMaha == 0 {
								fmt.Println("MAAF, BELUM ADA MAHASISWA YANG MENJAWAB SOAL :(")
							} else {
								user(&userQ, jumMaha)
								jawaban(&userQ, jumQuis) // tampilkan jawaban
							}
						} else if aksi == 6 {
							fmt.Println("\n====== DAFTAR NAMA MAHASISWA [ DOSEN ] ======= \n")
							inputNilai(&userQ, jumMaha) // Input Nilai
						} else if aksi == 7 {
							tampilkanNilai(userQ, jumMaha) // TAMPILKAN NILAI
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
						} else if aksi == 4 {
							fmt.Println("\n========= DAFTAR QUIS [ DOSEN ] ========== \n")
							tampilkanSoal(&quis, jumQuis)
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
						}
						clear()

					}
					// FORUM  =======================================================================
				} else if tipe == 3 {
					tampilkanForum(&forum, jumKomen)
					inputForum(&forum, mainUser, dosen, &jumKomen, tipeUser)
				} else if tipe == 4 {
					i = 1
					fmt.Println("\n===== RIWAYAT LOGIN =====")
					for i <= jumMaha {
						fmt.Println(i, ".", userT[i].nama, " ", userT[i].nim)
						i = i + 1
					}
					fmt.Println("JUMLAH MAHASISWA :", jumMaha)
					fmt.Print("\n|| TEKAN 'ENTER' >>>")
					fmt.Scanln(&next)
				}

				clear()
				aksi = 0
				// MAHASISWA  %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
			} else {
				// JAWAB TUGAS
				if tipe == 1 {
					for aksi != 3 {
						fmt.Println("\n========== PILIH AKSI [ MAHASISWA ] ========= \n" +
							"1. JAWAB SOAL \n" +
							"2. LIHAT NILAI \n" + "3. KEMBALI")
						fmt.Print("\n====> ")
						fmt.Scanln(&aksi)
						clear()
						if aksi == 1 {
							fmt.Println("\n========= DAFTAR TUGAS [ MAHASISWA ] ========== \n")
							tampilkanSoal(&tugas, jumTugas)
							jawab(&userT, jumTugas, jumMaha, idxMaha, markMaha) // Menjawab Soal Tugas
						} else if aksi == 2 {
							tampilkanNilai(userT, jumMaha) // Tampilkan Nilai
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
							clear()
						}
					}
					// JAWAB QUIS
				} else if tipe == 2 {
					for aksi != 3 {
						fmt.Println("\n========== PILIH AKSI [ MAHASISWA ] ========= \n" +
							"1. JAWAB SOAL \n" +
							"2. LIHAT NILAI \n" + "3. KEMBALI ")
						fmt.Print("\n====> ")
						fmt.Scanln(&aksi)
						clear()
						if aksi == 1 {
							fmt.Println("\n========= DAFTAR TUGAS [ MAHASISWA ] ========== \n")
							tampilkanSoal(&quis, jumQuis)
							jawab(&userQ, jumQuis, jumMaha, idxMaha, markMaha) // Menjawab Soal Quis

						} else if aksi == 2 {
							tampilkanNilai(userQ, jumMaha) // Tampilkan Nilai
							fmt.Print("\n|| TEKAN 'ENTER' >>>")
							fmt.Scanln(&next)
							clear()
						}
					}
					// FORUM MAHASISWA--------------------------------------------------------------------------------------------
				} else if tipe == 3 {
					tampilkanForum(&forum, jumKomen)
					inputForum(&forum, mainUser, dosen, &jumKomen, tipeUser)
				} else if tipe == 4 {
					i = 1
					fmt.Println("\n===== RIWAYAT LOGIN =====")
					for i <= jumMaha {
						fmt.Println(i, ".", userT[i].nama, " ", userT[i].nim)
						i = i + 1
					}
					fmt.Println("JUMLAH MAHASISWA :", jumMaha)
					fmt.Print("\n|| TEKAN 'ENTER' >>>")
					fmt.Scanln(&next)

				} else if tipe == 5 {
					fmt.Println("NAMA DOSEN : ", dosen)
					fmt.Println("MATA KULIAH : ", matkul)
					fmt.Print("\n|| TEKAN 'ENTER' >>>")
					fmt.Scanln(&next)
					tipe = tipe + 2 // optional

				} else if tipe == 6 {
					tipe = tipe - 1
				}
				clear()
			}
		}
		fmt.Println("======================================= \n" +
			"=========== SILAHKAN LOGIN ============ \n" +
			"======================================= \n")
		fmt.Println("LOGIN SEBAGAI : \n" + "1. DOSEN \n" + "2. MAHASISWA \n")
		fmt.Println("*** TEKAN '999' JIKA INGIN KELUAR DARI PROGRAM")
		fmt.Print("===> ")
		fmt.Scanln(&tipeUser)
		if tipeUser == 2 {
			clear()
			fmt.Println("============ MASUKKAN USERNAME DAN NIM MAHASISWA ============")
			fmt.Print("|| USERNAME : ")
			fmt.Scan(&mainUser)
			mainUser = mainUser + " " + scanner()
			fmt.Print("|| NIM : ")
			fmt.Scanln(&nim)
		} else if tipeUser == 1 {
			clear()
			if markDosen == 0 {
				fmt.Println("============ MASUKKAN NAMA DOSEN ============")
				fmt.Print("NAMA DOSEN : ")
				fmt.Scan(&dosen)
				dosen = dosen + " " + scanner()
				markDosen = markDosen + 1
			}
			mainUser = "dosen "

		} else if tipeUser == 999 {
			mainUser = "00 "
		}
	}
	clear()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n                                        ==================================")
	fmt.Println("                                        ========== TERIMA KASIH ==========")
	fmt.Println("                                        ================================== \n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n|| TEKAN 'ENTER' >>>")
	fmt.Scanln(&next)
	clear()
}
