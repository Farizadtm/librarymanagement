# librarymanagement
Library management sistem adalah sistem yang dibuat guna membantu staff untuk mendata buku yang mau dipinjamkan atau dikembalikan

## Daftar database yang digunakan
1. Buku
2. Admin
3. Mahasiswa
4. Riwayat Peminjaman
5. Riwayat Pengembalian

## Alur
1. Pertama kita dapat membuat user admin dengan isi

Kolom  | Tipe Data | Keterangan
------------- | ------------- | -------------
Name  | string | nama admin
Username | string | Minimal 4 karakter
Password | string | minimal 8 karakter
Birth_date | string | format YYYY-MM-DD

```json
{
    "name" : "Sulistyo",
    "pass" : "123456",
    "username" : "sulis",
    "birth_date" : "1965-11-09"
}
```

2. Setelah mendaftarkan, gunakan username dan password yang sudah digunakan untuk mendapatkan data dan mengambil data. Namun ada juga beberapa yang tidak diperlukan. Berikut beberapa data yang dapat diakses beserta keterangannya

Master Data Buku 
Method  | Route | Auth
------------- | ------------- | -------------
GET  | /book | false
POST | /book | true
PUT | /book/:id |true
DELETE | /book/:id | true


POST BUKU
```json
{
    "title" : "Sang Pemimpi",
    "author" : "Andrea Hirata",
    "release_year" : "206",
    "category" : "Edukasi"
}
```

Master Data Admin 
Method  | Route | Auth
------------- | ------------- | -------------
GET  | /admin | true
POST | /admin | false
PUT | /admin/:id |true
DELETE | /admin/:id | true


Master Data Student 
Method  | Route | Auth
------------- | ------------- | -------------
GET  | /student | false
POST | /student | true
PUT | /student/:id |true
DELETE | /student/:id | true

POST Student
```json
{
      "name": "Dandi",
      "major": "Teknik Sipil"
}
```

3. Setelah post untuk semua master data seperti buku, admin dan mahasiswa. Maka kemudian kita dapat membuat history peminjaman buku, dan history tersebut dapat dibuat jika semua id tersedia
POST Student

```json
{
    "student_id" : 1,
    "admin_id" : 2,
    "book_id" : 1,
    "duration" : 3
}
```
Diatas ini adalah contoh untuk membuat data peminjaman. Jika data tersebut dibuat maka secara otomatis data ketersediaan buku akan terupdate sehingga ketika ingin meminjamnya sebelum dikembalikan maka itu tidak akan bisa. Dan juga jika buku sedang dipinjam maka POST akan gagal

4. Setelah melakukan peminjaman makan mahasiswa akan mengembalikan buku

```json
{
    "student_id" : 1,
    "admin_id" : 2,
    "book_id" : 1,
}
```
Tidak jauh beda dengan peminjaman, pengembalian datanya seperti diatas dan secara otomatis buku akan kembali tersedia, dan kalaupun buku sedang tidak dipinjamkan maka post gagal
