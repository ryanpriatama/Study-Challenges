# STUDY CHALLENGES
[![Go reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)

_API_ quiz game _untuk aplikasi Study Challenges_

# Table of Content

- [Introduction](#introduction)
  - [Summary](#summary)
  - [Subjects](#subjects)
- [Features](#features)
- [How to Use](#how-to-use)
  - [Get Started]((#get-started))
  - [API Documentation](#api-documentation)

***
# Introduction
### **Summary**
Study Challenges adalah sebuah aplikasi berbasis kuis dalam bidang pendidikan, di dalam aplikasi ini seorang user dapat mengerjakan satu atau lebih set soal yang berisi 5 nomor soal. Satu set soal di ambil secara acak oleh sistem, namun materi pelajaran dan tingkat kesulitan dapat dipilih sesuai keinginan user.

Untuk dapat mengakses fitur, user harus memiliki akun terlebih dahulu. Disini terdapat dua role untuk masing-masing user, antara lain : Admin dan User pengguna. Masing-masing role dapat menginisiasi soal, namun Admin memiliki wewenang khusus untuk memberikan ijin publish soal dari soal yang sudah user inisiasi sebelumnya, sehingga apabila Admin belum memberikan izin publish, maka soal dari user belum dapat di akses oleh pengguna lainnya.

### **Subjects**
Materi tantangan yang dapat di pilih oleh user antara lain :
| Subject | Category ID | Level |
| --- | --- | --- |
| Kimia | Category 1 | 1, 2, 3 |
| Fisika | Category 2 | 1, 2, 3 |
| Biologi | Category 3 | 1, 2, 3 |
| Matematika | Category 4 | 1, 2, 3 |
| Bahasa Inggris | Category 5 | 1, 2, 3 |
| Ekonomi | Category 6 | 1, 2, 3 |
| Geografi | Category 7| 1, 2, 3 |
***
# Features

| No. | Features | Role | Keterangan |
| --- | --- | --- | --- |
| 1. | Register. | User, Admin. | |
| 2. | Login. | User, Admin. | |
| 3. | Memilih dan mengerjakan Set Soal random berdasarkan kategori dan tingkat kesulitan. | User. | Satu set berisi 5 nomor soal. |
| 4. | Melihat perolehan poin dari pengerjaan set soal. | User. | |
| 5. | Memiliki Badge yang diperoleh dari kalkulasi poin keseluruhan. | User. | Badge antara lain : Bronze, Silver, dan Gold. |
| 6. | Melihat Leaderboard. | User. | Ranking 1 s.d 10 |
| 7. | Menginisiasi soal baru, lalu dapat di akses hanya jika Admin mengizinkan _publish_. | User. | Admin dapat meng-_accept_/ me-_reject_ status soal. |
| 8. | Menginisiasi soal baru, dan langsung dapat di akses oleh setiap user. | Admin. | |
| 9. | Menerima atau menolak inisiasi soal | Admin. | |

***

# How to Use
## **Get Started**
### Berikut langkah-langkah yang dapat diikuti untuk mengakses API Study Challenges.
### 1. Instalasi [Postman](https://www.postman.com/downloads/ "postman-download").
### 2. Register User
Anda dapat meng-copy link berikut ini:
```
http://13.208.193.149:8080/users/signup
```
Agar dapat mengakses set soal, pengguna harus memiliki akun terlebih dahulu, method yang dipakai adalah method `POST`. Silakan copy link di atas lalu masukkan pada `request` postman anda.

pada tab `body` pilih `form-data`, dan masukkan key berupa `nama` , `email` dan `password`, pastikan ketiga key ini memiliki centang di sebelah kiri, lalu kita dapat mengisi `form-data` dengan data yang kita miliki.

example :
| KEY | VALUE |
| --- | --- |
| nama | my name |
| email | myname@myemail.com |
| password | mysecret |

Kita dapat mengirimkan informasi yang sudah dilengkapi sebelumnya dengan menekan tombol `send`. Jika muncul pesan seperti berikut: 
```json
{
    "data": {
        "Email": "myname@myemail.com",
        "ID": 4,
        "Name": "my name",
        "Rank": "bronze",
        "Total Poin": 0
    },
    "message": "new user added"
}
```
Selamat! Anda sudah berhasil mendaftarkan akun anda. (Keterangan: ID akan berbeda-beda setiap user)
### 3. Login User
Silakan copy link di bawah ini:
```
http://13.208.193.149:8080/users/login
```
Satu langkah lagi sebelum anda dapat mengakses set soal, yaitu `login` sebagai user. Silakan masukkan link yang sudah di copy sebelumnya ke `request` postman anda, lalu pilih method `POST`, berikutnya pada tab `body` dapat dipilih tab `form-data`, masukkan `Key` berupa `email` dan `password` yang sudah berhasil anda daftarkan sebelumnya. Berikut contoh form pengisian pada `form-data`:
| KEY | VALUE |
| --- | --- |
| email | myname@myemail.com |
| password | mysecret |

Apabila Login berhasil, maka akan tampil pesan pada `Response` Postman dengan format berikut:

```
{
    "message": "Welcome",
    "users": {
        "ID": 4,
        "Name": "my name",
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Mjk5NDIwNDcsInJvbGUiOiJ1c2VyIiwidXNlcklkIjo0fQ.QucoMKsx6qWTRDEfpG_fdHrwY70J0JMieVKHJEdg5Cg"
    }
}  
```
Copy `Token` dan `ID` yang diberikan, token dan Id tersebut merupakan informasi yang krusial untuk dapat mengakses fitur-fitur di dalam aplikasi Study Challenges.
### 4. Menyelesaikan Tantangan
Setelah login berhasil dilakukan, users akan mendapatkan `Token`, silakan copy token dan pilih tab `Authorization`, lalu pada bagian `Type` pilih `Bearer Token`, selanjutnya isi bagian `Token` dengan token tersebut.
Langkah berikutnya adalah memasukkan request pada Postman (menggunakan method `POST`) dengan endpoint sebagai berikut:
```
http://13.208.193.149:8080/users/:user_id/soal
```
example :
```
http://13.208.193.149:8080/users/4/soal
```
pada contoh diatas, user dengan id 4 ingin membuat sebuah set soal baru. Selanjutnya anda dapat menekan tab `Body` dan memilih `form-data`, silakan pilih [kategori](#subjects "kategori soal") soal beserta tingkat kesulitan yang diinginkan, lalu isi pada form-data postman, berikut sebagai contoh:
| KEY | VALUE |
| --- | --- |
| CategoryID | 5 |
| KesulitanID | 2 |

dari format di atas, maka terlihat user memilih pelajaran Bahasa Inggris dengan tingkat kesulitan 2.
Sistem akan bekerja untuk menyusun set soal, apabila proses tersebut berhasil, akan muncul pesan dengan format berikut:
```json
{
    "data": {
        "ID": 8,
        "Kesulitan": "Medium",
        "Mata Pelajaran": "Bahasa Inggris",
        "UserID": 4
    },
    "message": "new set soal added"
}
```
bersama pesan di atas, pada `Response` juga akan ditampilkan ID dari set soal yang baru saja disusun oleh sistem, ID tersebut dapat kita gunakan untuk lanjut ke tahap berikutnya, yaitu menampilkan set soal.

Agar dapat melihat set soal yang sebelumnya sudah di pilih oleh sistem, user dapat mengikuti format request berikut untuk kemudian dimasukkan ke dalam request Postman dengan method `GET`.
```
http://13.208.193.149:8080/users/:user_id/soal/:set_soal_id
```
example :
```    
http://13.208.193.149:8080/users/4/soal/8
```
dari contoh tersebut terlihat bahwa user ID 4 mengirimkan `request` untuk mendapatkan data set soal dengan ID 8. Apabila proses tersebut berhasil, maka akan muncul `Response` dengan format berikut (contoh set soal):
```json
{
    "data": [
        {
            "Soal": "A: “Do you like the game?”\nB: “[…....]”",
            "Soal_id": 76,
            "Pilihan_A": "The dinner is really delightful.",
            "Pilihan_B": "I do enjoy this game.",
            "Pilihan_C": "I am pleased with all the food.",
            "Pilihan_D": "Terrible!"
        },
        {
            "Soal": "“Some students are going to arrange a reunion party at school next month.” The passive is......",
            "Soal_id": 84,
            "Pilihan_A": " A reunion party are going to arrange by some student at school next month.",
            "Pilihan_B": "A reunion party is going to arrange by some student at school next month.",
            "Pilihan_C": "A reunion party is going to arrange by some students at school next month.",
            "Pilihan_D": "A reunion party is going to be arranged by some students at school next month."
        },
        {
            "Soal": "“Robby will buy John a new cloth.” The passive is….",
            "Soal_id": 82,
            "Pilihan_A": "John will be bought by Robby a new cloth.",
            "Pilihan_B": "John would be bought a new cloth by Robby.",
            "Pilihan_C": "A new cloth will be bought by Robby for John.",
            "Pilihan_D": "A new cloth will be bought by John for Robby."
        },
        {
            "Soal": "“SBY and Budiono won the general election.” The passive is .....",
            "Soal_id": 81,
            "Pilihan_A": "SBY and Budiono are won by the general election.",
            "Pilihan_B": "The general election is won by SBY and Budiono.",
            "Pilihan_C": "The general election was won by SBY and Budiono.",
            "Pilihan_D": "The general election were won by SBY and Budiono."
        },
        {
            "Soal": "A: “Will you pick me up this evening?”;  B: \"[..........]”",
            "Soal_id": 80,
            "Pilihan_A": "sure you can.",
            "Pilihan_B": "Of course please help yourself.",
            "Pilihan_C": "No, I don’t think I can.",
            "Pilihan_D": "Sure I will."
        }
    ],
    "message": "All questions show successfully"
}
```
Silakan rekam jawaban yang ingin anda kirimkan berdasarkan soal yang tertera.

Untuk langkah berikutnya silakan ikuti format endpoint berikut:
```
http://13.208.193.149:8080/users/:user_id/soal/:set_soal_id
```
Apabila anda ingin mengirimkan jawaban dari set soal, anda dapat mengikuti format endpoint di atas lalu memasukkan `Request` ke dalam `Postman` dengan method PUT. Berikut contoh format link:
```
http://13.208.193.149:8080/users/4/soal/8
```
lalu pada tab `Body` anda dapat memilih `form-data`, silakan isi jawaban yang anda yakini benar. Berikut contoh pengisian data:
| KEY | VALUE |
| --- | --- |
| Soal_1 | b |
| Soal_2 | d |
| Soal_3 | c |
| Soal_4 | c |
| Soal_5 | d |

Setelah menekan tombol `send`, jika berhasil maka akan muncul `Response` dengan format sebagai berikut :
```json
{
    "data": {
        "Soal Salah": null,
        "Total Score": 15
    },
    "message": "Question Answered"
}
```

### 5. Langkah Selanjutnya.
Untuk mengakses fitur lainnya dari aplikasi Study Challenge, mohon kunjungi [API Documentation](#api-documentation "api documentation")
***
## **API Documentation**
Berikut merupakan Endpoint yang dapat dipergunakan untuk mengakses fitur dalam aplikasi Study Challenges : [klik disini](https://app.swaggerhub.com/apis/ryanpriatama/studychallanges/1#/ "Study-challenges-endpoint")

***
