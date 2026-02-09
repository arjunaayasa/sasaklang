# SasakLang Demo - Kalkulasi
# Contoh operasi aritmatika

tulis("=== Kalkulator SasakLang ===")
tulis("")

# Operasi dasar
gawe a = 10
gawe b = 3

tulis("a =", a)
tulis("b =", b)
tulis("")

tulis("Operasi Dasar:")
tulis("a + b =", a + b)
tulis("a - b =", a - b)
tulis("a * b =", a * b)
tulis("a / b =", a / b)
tulis("a % b =", a % b)
tulis("")

# Operasi perbandingan
tulis("Perbandingan:")
tulis("a > b  =", a > b)
tulis("a < b  =", a < b)
tulis("a == b =", a == b)
tulis("a != b =", a != b)
tulis("a >= b =", a >= b)
tulis("a <= b =", a <= b)
tulis("")

# Konstanta tidak bisa diubah
tetep MAX = 100
tulis("Konstanta MAX =", MAX)

# Operasi kompleks
gawe hasil = (a + b) * 2 - 5
tulis("(a + b) * 2 - 5 =", hasil)

# Negative numbers
gawe negatif = -10
tulis("Nilai negatif:", negatif)
tulis("Negasi:", -negatif)
