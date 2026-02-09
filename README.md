# SasakLang

<p align="center">
  <strong>Bahasa Pemrograman Berbasis Bahasa Sasak (Lombok)</strong>
</p>

<p align="center">
  <em>"Belajar coding sambil melestarikan budaya lokal"</em>
</p>

**SasakLang** adalah bahasa pemrograman esoterik (esolang) berbasis interpreter yang menggunakan kosa kata Bahasa Sasak dari Lombok. Dibuat untuk edukasi, melatih logika berpikir, sekaligus melestarikan budaya lokal lewat coding.

## ✨ Fitur Utama

- **Sederhana** → Sintaks mirip Python/Go, mudah dipelajari pemula
- **Sasak 100%** → Keyword menggunakan Bahasa Sasak (gawe, yen, salama, pungsi, dll)
- **Ringan & Cepat** → Dibangun dengan Go, tanpa dependency berat
- **Lengkap** → Variabel, I/O, Control flow, Functions, Array, Map

## 📦 Instalasi

### Quick Install (Recommended)

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/arjunaayasa/sasaklang/main/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/arjunaayasa/sasaklang/main/scripts/install.ps1 | iex
```

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
sasaklang

# Jalankan file
sasaklang run examples/hello.sl
sasaklang examples/hello.sl    # shortcut

# Cek versi
sasaklang version
```

## 📚 Kamus Syntax

### Keywords

| Sasak | English | Kegunaan |
|-------|---------|----------|
| `gawe` | let/var | Deklarasi variabel |
| `tetep` | const | Deklarasi konstanta |
| `tulis` | print | Cetak output |
| `tanya` | input | Baca input |
| `yen` | if | Kondisional |
| `neng` | else | Alternatif kondisi |
| `salama` | while | Perulangan while |
| `kanggo` | for | Perulangan for |
| `pungsi` | function | Definisi fungsi |
| `balik` | return | Return value |

### Literals

| Sasak | English | Nilai |
|-------|---------|-------|
| `bener` | true | Boolean true |
| `salah` | false | Boolean false |
| `kosong` | null | Null value |

### Operators

| Operator | Kegunaan |
|----------|----------|
| `+ - * / %` | Aritmatika |
| `== != > >= < <=` | Perbandingan |
| `&& \|\| !` | Logika |
| `=` | Assignment |

## 💻 Contoh Kode

### Hello World
```
# Hello World di SasakLang
tulis("Selamat datang di SasakLang!")

gawe nama = "Lombok"
tulis("Halo,", nama)
```

### Variabel dan Konstanta
```
gawe x = 10       # variabel (bisa diubah)
tetep PI = 3      # konstanta (tidak bisa diubah)

x = x + 1         # OK
# PI = 4          # Error: tidak bisa mengubah konstanta
```

### Kondisional
```
gawe nilai = 85

yen (nilai >= 90) {
    tulis("Grade: A")
} neng {
    yen (nilai >= 80) {
        tulis("Grade: B")
    } neng {
        tulis("Grade: C")
    }
}
```

### Perulangan
```
# While loop
gawe i = 5
salama (i > 0) {
    tulis(i)
    i = i - 1
}

# For loop
kanggo (gawe j = 1; j <= 5; j = j + 1) {
    tulis("Iterasi ke-", j)
}
```

### Fungsi
```
pungsi tambah(a, b) {
    balik a + b
}

pungsi faktorial(n) {
    yen (n <= 1) {
        balik 1
    } neng {
        balik n * faktorial(n - 1)
    }
}

tulis(tambah(5, 3))      # Output: 8
tulis(faktorial(5))       # Output: 120
```

### Array dan Map
```
# Array
gawe angka = [1, 2, 3, 4, 5]
tulis(angka[0])           # Output: 1
tulis(panjang(angka))     # Output: 5

# Map
gawe orang = {"nama": "Amaq", "umur": 30}
tulis(orang["nama"])      # Output: Amaq
```

## 🔧 Fungsi Bawaan (Built-in)

| Fungsi | Kegunaan |
|--------|----------|
| `tulis(...args)` | Cetak output dengan spasi, diakhiri newline |
| `tanya(prompt?)` | Baca input dari user, return string |
| `panjang(x)` | Panjang string atau array |
| `jenis(x)` | Nama tipe data sebagai string |
| `waktu()` | Unix timestamp (int64) |
| `dorong(arr, item)` | Tambah item ke array (return array baru) |
| `pertama(arr)` | Elemen pertama array |
| `akhir(arr)` | Elemen terakhir array |

## 🎨 VS Code Extension

Extension untuk syntax highlighting tersedia di folder `vscode-extension/`:

```bash
# Copy ke VS Code extensions folder
cp -r vscode-extension ~/.vscode/extensions/sasaklang
```

Restart VS Code dan file `.sl` akan memiliki syntax highlighting.

## 🧪 Testing

```bash
# Jalankan semua test
go test ./...

# Test dengan verbose
go test -v ./...
```

## 📁 Struktur Proyek

```
sasaklang/
├── cmd/sasaklang/main.go      # CLI entry point
├── pkg/sasaklang/
│   ├── token/token.go         # Token definitions
│   ├── lexer/lexer.go         # Lexical analyzer
│   ├── ast/ast.go             # Abstract Syntax Tree
│   ├── parser/parser.go       # Parser (Pratt parsing)
│   ├── object/object.go       # Runtime objects
│   ├── evaluator/evaluator.go # Tree-walking evaluator
│   ├── builtins/builtins.go   # Built-in functions
│   ├── errors/errors.go       # Error handling
│   └── repl/repl.go           # REPL
├── examples/                   # Contoh kode
├── scripts/                    # Installer scripts
├── vscode-extension/           # VS Code extension
└── README.md
```

## 🤝 Kontribusi

Kontribusi sangat diterima!

1. Fork repo ini
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahan (`git commit -m 'Tambah fitur baru'`)
4. Push ke branch (`git push origin fitur-baru`)
5. Buat Pull Request

## 📄 Lisensi

Dirilis di bawah [MIT License](LICENSE). Silakan digunakan, dimodifikasi, dan disebarluaskan.

---

<p align="center">
  Made with ❤️ for Sasak culture preservation
</p>
