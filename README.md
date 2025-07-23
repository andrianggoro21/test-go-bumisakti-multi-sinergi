# Golang Skill Test

## Knowledge Test dan Logic Test
ada di dalam folder docs : Jawaban Knowledge Test dan Logic Test - Andri Anggoro.pdf

## ✅ Versi Go

Dikembangkan menggunakan **Go 1.24.2**.

## 📦 Struktur File

skill-test-golang/
├── main.go // Semua jawaban soal dalam satu file

## 🧪 Isi dan Penjelasan

Berikut ini adalah jawaban dari masing-masing soal yang diimplementasikan dalam `main.go`:

1. **Pengecekan Tipe Data (Type Verification)**  
   Menggunakan `interface{}` (`any`) dan `type switch` untuk memverifikasi tipe data (int, string, float64, dll).

2. **Function Closure**  
   Fungsi `counter()` mengembalikan fungsi lain yang mengakses variabel `count` dari lingkup luar (closure).

3. **Class Inheritance (Struct Embedding)**  
   Struct `Student` meng-embed `Person`, memungkinkan `Student` mewarisi properti dan method `Person`.

4. **Permutasi**  
   Menggunakan rekursi dan teknik backtracking untuk mencetak semua kemungkinan susunan elemen array.

5. **Penghentian Goroutine**  
   Menggunakan channel untuk mengirim sinyal agar goroutine berhenti dengan aman.

## 🛠️ Langkah Install & Run

### 1. Unzip project

```bash
unzip skill-test-golang.zip
cd skill-test-golang
```

### 2. Run main.go

```bash
go run main.go
```
