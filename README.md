<h1 align="center">SasakLang</h1>

<p align="center">
  <img src="sasaklang.png" alt="SasakLang Logo" width="200"/>
</p>

<p align="center">
  <strong>Bahasa Pemrograman Berbasis Bahasa Sasak (Lombok)</strong>
</p>

<p align="center">
  <em>"Belajar coding sambil melestarikan budaya lokal"</em>
</p>

**SasakLang** adalah bahasa pemrograman esoterik (esolang) berbasis interpreter yang menggunakan kosa kata Bahasa Sasak dari Lombok. Dibuat untuk edukasi, melatih logika berpikir, sekaligus melestarikan budaya lokal lewat coding.

## ✨ Fitur Utama

- **Sederhana** → Sintaks mirip Python/Go, mudah dipelajari pemula
- **Sasak 100%** → Keyword menggunakan Bahasa Sasak (`gawe`, `lamun`, `selame`, `fungsi`, dll)
- **Ringan & Cepat** → Dibangun dengan Go, tanpa dependency berat
- **Lengkap** → Variabel, I/O, Control flow, Functions, Array, Map

## 📦 Instalasi

### Build dari Source

Pastikan Go 1.22+ sudah terinstall:

```bash
git clone https://github.com/arjunaayasa/sasaklang.git
cd sasaklang
go build -o sasaklang ./cmd/sasaklang
./sasaklang version
```

## 🚀 Cara Pakai

```bash
# Masuk REPL (interactive mode)
./sasaklang

# Jalankan file
./sasaklang run examples/hello.ssk
```

## 📚 Kamus Syntax

### Keywords & Tipe Data

| Sasak | English | Kegunaan |
|-------|---------|----------|
| `gawe` | let/var | Deklarasi variabel |
| `tetep` | const | Deklarasi konstanta |
| `lamun` | if | Kondisional If |
| `endah` | else | Kondisional Else |
| `selame` | while | Perulangan While |
| `ojok` | for | Perulangan For |
| `fungsi` | function | Definisi Fungsi |
| `tulakan` | return | Mengembalikan nilai |
| `mentelah` | break | Keluar dari loop |
| `lanjutan` | continue | Lanjut iterasi berikutnya |
| `kenak` | true | Boolean True |
| `salak` | false | Boolean False |
| `ndarak` | null | Nilai Null/Kosong |

### Operators

| Operator | Sasak | Kegunaan |
|----------|-------|----------|
| `+ - * / %` | - | Aritmatika |
| `== != > >= < <=` | - | Perbandingan |
| `&&` | `ance` | Logika DAN |
| `\|\|` | `atau` | Logika ATAU |
| `!` | `ndek` | Logika BUKAN (NOT) |
| `=` | - | Assignment |

## 🔧 Fungsi Bawaan (Built-in)

| Fungsi | Deskripsi |
|--------|-----------|
| `cetak(...args)` | Cetak ke layar (println) |
| `isik(prompt?)` | Baca input dari pengguna |
| `belong(x)` | Panjang string atau array (length) |
| `jenis(x)` | Cek tipe data variable |
| `waktu()` | Unix timestamp saat ini |
| `tedem(ms)` | Jeda eksekusi (sleep) |
| `acak(max)` | Angka acak 0 s.d max-1 |
| `sorong(arr, val)` | Tambah item ke array (push) |
| `bait(col, key)` | Ambil nilai dari array/map (get) |
| `ngatur(col, key, val)` | Set nilai di array/map (set) |

## 💻 Contoh Kode

### Hello World & Input
```sasak
cetak("Halo Dunia!")
gawe nama = isik("Aran side sai? ")
cetak("Halo,", nama)
```

### Percabangan
```sasak
gawe nilai = 85
lamun (nilai >= 80) {
    cetak("Lulus!")
} endah {
    cetak("Coba malik.")
}
```

### Perulangan & Array
```sasak
gawe angka = [1, 2, 3]
ojok (gawe i = 0; i < belong(angka); i = i + 1) {
    cetak("Angka:", bait(angka, i))
}

# Pakai 'sorong' untuk nambah data
gawe angkaBaru = sorong(angka, 4)
```

### Fungsi
```sasak
fungsi tambah(a, b) {
    tulakan a + b
}
cetak(tambah(5, 10))
```

## 🎨 VS Code Extension

Extension untuk syntax highlighting dan snippet telah tersedia di Visual Studio Code Marketplace.

### Cara Install:
1.  Buka **VS Code**.
2.  Masuk ke menu **Extensions** (Ctrl+Shift+X).
3.  Cari **"SasakLang"**.
4.  Klik **Install**.

### Install Manual (.vsix):
Jika tidak ketemu di marketplace, Anda bisa install menggunakan file installer yang tersedia di [Releases](https://github.com/arjunaayasa/sasaklang/releases):
1.  Download file `.vsix` terbaru.
2.  Di VS Code, buka menu Extensions -> Klik titik tiga (`...`) di pojok kanan atas.
3.  Pilih **"Install from VSIX..."**.
4.  Pilih file yang sudah didownload.


## 🤝 Kontribusi

Kontribusi sangat diterima! Bantu kami melestarikan budaya Sasak melalui kode.

1.  **Fork** repository ini.
2.  Buat branch fitur baru (`git checkout -b fitur-keren`).
3.  Commit perubahan Anda (`git commit -m 'Menambahkan fitur keren'`).
4.  Push ke branch (`git push origin fitur-keren`).
5.  Buat **Pull Request**.

Jangan lupa jalankan test sebelum commit:
```bash
go test ./...
```

## 📄 Lisensi

MIT
