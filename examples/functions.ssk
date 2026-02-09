# SasakLang Demo - Functions
# Contoh definisi dan pemanggilan fungsi

tulis("=== Fungsi di SasakLang ===")
tulis("")

# Fungsi sederhana
pungsi sapa(nama) {
    tulis("Halo,", nama + "!")
}

sapa("Lombok")
sapa("Sasak")

tulis("")

# Fungsi dengan return
pungsi tambah(a, b) {
    balik a + b
}

pungsi kali(a, b) {
    balik a * b
}

tulis("tambah(5, 3) =", tambah(5, 3))
tulis("kali(4, 7) =", kali(4, 7))

tulis("")

# Fungsi rekursif - Fibonacci
pungsi fibonacci(n) {
    yen (n <= 1) {
        balik n
    } neng {
        balik fibonacci(n - 1) + fibonacci(n - 2)
    }
}

tulis("Fibonacci sequence:")
kanggo (gawe i = 0; i <= 10; i = i + 1) {
    tulis("fib(", i, ") =", fibonacci(i))
}

tulis("")

# Fungsi rekursif - Factorial
pungsi faktorial(n) {
    yen (n <= 1) {
        balik 1
    } neng {
        balik n * faktorial(n - 1)
    }
}

tulis("Factorial:")
kanggo (gawe i = 1; i <= 5; i = i + 1) {
    tulis(i, "! =", faktorial(i))
}

tulis("")

# Higher-order function style
pungsi operasi(a, b, op) {
    balik op(a, b)
}

gawe hasil = operasi(10, 5, tambah)
tulis("operasi(10, 5, tambah) =", hasil)

hasil = operasi(10, 5, kali)
tulis("operasi(10, 5, kali) =", hasil)

tulis("")

# Array dan Map
tulis("Array dan Map:")
gawe angka = [1, 2, 3, 4, 5]
tulis("Array:", angka)
tulis("Elemen pertama:", angka[0])
tulis("Panjang array:", panjang(angka))

gawe orang = {"nama": "Amaq", "umur": 30}
tulis("Map:", orang)
tulis("Nama:", orang["nama"])

tulis("")

# Built-in functions
tulis("Built-in Functions:")
tulis("panjang('hello') =", panjang("hello"))
tulis("jenis(42) =", jenis(42))
tulis("jenis('teks') =", jenis("teks"))
tulis("jenis(bener) =", jenis(bener))
tulis("waktu() =", waktu())
