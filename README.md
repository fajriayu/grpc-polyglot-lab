
# LAPORAN PRAKTIKUM
## Komunikasi Inter-service dengan gRPC (Polyglot)

**Nama:** Fajri Ayu  
**NIM:** 2024903430026  
**Kelas:** TRKJ 3C  
**Mata Kuliah:** Paradigma Sistem Untuk IT  
**Program Studi:** Teknologi Rekayasa Komputer Jaringan  
**Tahun:** 2026

## KATA PENGANTAR

Puji syukur ke hadirat Allah SWT karena atas rahmat dan karunia-Nya, laporan praktikum berjudul "Komunikasi Inter-service dengan gRPC (Polyglot)" ini dapat diselesaikan. Laporan ini disusun untuk memahami komunikasi antar-service menggunakan gRPC dengan beberapa bahasa pemrograman, yaitu Go, Python, dan Node.js.

Praktikum dilakukan mulai dari menyiapkan lingkungan pengembangan, membuat kontrak API menggunakan Protocol Buffers, membangun server gRPC menggunakan Go, serta membuat client menggunakan Python dan Node.js. Pengujian dilakukan untuk memastikan komunikasi antara client dan server berjalan dengan baik, termasuk penambahan field nomor telepon pada respons.

Penulis menyadari bahwa laporan ini masih memiliki kekurangan. Oleh karena itu, kritik dan saran yang membangun diharapkan agar laporan berikutnya dapat menjadi lebih baik. Semoga laporan ini bermanfaat bagi pembaca.

## BAB I — PENDAHULUAN

### 1.1 Latar Belakang

Dalam arsitektur microservices, layanan yang menggunakan bahasa pemrograman berbeda perlu bertukar data melalui kontrak yang konsisten. gRPC memungkinkan suatu program memanggil metode pada layanan lain menggunakan kontrak Protocol Buffers (Protobuf).

Praktikum ini menerapkan komunikasi lintas bahasa dengan Go sebagai server gRPC, serta Python dan Node.js sebagai client. Kontrak API didefinisikan dalam file `user.proto`.

### 1.2 Tujuan Praktikum

- Memahami Schema-First Development menggunakan Protocol Buffers.
- Membuat server gRPC menggunakan Go.
- Membuat client Python dan Node.js yang terhubung ke server yang sama.
- Menambahkan field `phone` pada respons.
- Memahami static code generation dan dynamic runtime loading.
- Mengunggah proyek ke GitHub.

### 1.3 Dasar Teori

gRPC merupakan framework komunikasi antar-service yang menggunakan Protocol Buffers sebagai format kontrak. Metode `GetUserProfile` menggunakan pola unary RPC, yaitu satu request menghasilkan satu response.

Go dan Python menggunakan kode hasil generasi, sedangkan Node.js membaca definisi `.proto` saat program berjalan melalui `@grpc/proto-loader`.

## BAB II — ALAT DAN BAHAN

Perangkat dan teknologi yang digunakan:

| Komponen | Versi/Konfigurasi |
|---|---|
| Sistem Operasi | Windows |
| Git | 2.51.0.windows.1 |
| Python | 3.14.3 |
| Node.js | v24.14.0 |
| npm | 11.9.0 |
| Go | go1.27.2 windows/amd64 |
| Protocol Buffers | libprotoc 36.2 |
| gRPC Server | Go, port 50051 |
| Client | Python dan Node.js |

## BAB III — LANGKAH KERJA DAN IMPLEMENTASI

### 3.1 Persiapan dan Struktur Proyek

Proyek dibuat dalam direktori `grpc-polyglot-lab`. Git digunakan untuk mengelola perubahan kode, sedangkan `.gitignore` digunakan untuk mengecualikan direktori dependensi dan file sementara.

Struktur proyek:

```text
grpc-polyglot-lab/
├── .gitignore
├── README.md
├── user.proto
├── go-server/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── pb/
├── python-client/
│   ├── client.py
│   ├── user_pb2.py
│   └── user_pb2_grpc.py
├── node-client/
│   ├── client.js
│   ├── package.json
│   └── package-lock.json
└── images/
```

**Gambar 1. Struktur direktori proyek dan persiapan Git**

![Struktur proyek](images/01_persiapan_dan_struktur_proyek.png)

### 3.2 Membuat Kontrak API `user.proto`

File `user.proto` mendefinisikan pesan request, response, dan service yang digunakan oleh server serta client. Pada tugas evaluasi, ditambahkan field `phone` dengan nomor field 3.

```proto
syntax = "proto3";

package user;

option go_package = "go-server/pb";

message UserRequest {
    int32 user_id = 1;
}

message UserResponse {
    string name = 1;
    string email = 2;
    string phone = 3;
}

service UserService {
    rpc GetUserProfile(UserRequest) returns (UserResponse);
}
```

**Gambar 2. Isi user.proto setelah penambahan field phone**

![User proto phone](images/02_user_proto_phone.png)

### 3.3 Implementasi gRPC Server dengan Go

Server Go mengimplementasikan metode `GetUserProfile` dan berjalan pada port 50051. Ketika menerima request dengan `user_id = 1`, server mengembalikan nama, email, dan nomor telepon.

Data pengujian:
- Nama: Fajri Ayu
- Email: Ayu@universitas.ac.id
- Phone: 081234567890

Server dijalankan menggunakan perintah:

```bash
go run main.go
```

**Gambar 3. Server Go berjalan pada port 50051**

![Server Go berjalan](images/03_server_go_berjalan.png)

### 3.4 Implementasi Client Python

Client Python menggunakan library `grpcio` serta modul hasil generasi Protobuf. Client terhubung ke `localhost:50051` dan memanggil metode `GetUserProfile` dengan `user_id = 1`.

Perintah untuk menjalankan client:

```bash
python client.py
```

Hasil pengujian:

```text
Nama  : Fajri Ayu
Email : Ayu@universitas.ac.id
Phone : 081234567890
```

**Gambar 4. Hasil pengujian client Python**

![Hasil client Python](images/04_client_python_phone.png)

### 3.5 Implementasi Client Node.js

Client Node.js menggunakan paket `@grpc/grpc-js` dan `@grpc/proto-loader`. File `user.proto` dibaca saat runtime sehingga tidak memerlukan file stub JavaScript hasil generasi secara terpisah.

Perintah untuk menjalankan client:

```bash
node client.js
```

Hasil pengujian diharapkan menampilkan nama, email, dan nomor telepon yang dikirim server.

**Gambar 5. Hasil pengujian client Node.js**

![Hasil client Node.js](images/05_client_nodejs_phone.png)

### 3.6 Pengujian Komunikasi Inter-service

Pengujian dilakukan dengan menjalankan server Go terlebih dahulu, kemudian menjalankan client Python dan Node.js pada terminal terpisah. Kedua client terhubung ke server yang sama melalui `localhost:50051`.

**Gambar 6. Log server setelah menerima request dari client**

![Log server](images/06_log_request_server.png)

### 3.7 Tugas Evaluasi: Penambahan Field Phone

Tugas evaluasi dilakukan dengan menambahkan `string phone = 3;` ke dalam pesan `UserResponse`. Perubahan diterapkan pada kontrak API, server Go, serta client Python dan Node.js. Kode hasil generasi untuk Go dan Python diperbarui agar mendukung field baru.

Nomor telepon pengujian yang digunakan adalah `081234567890`.

**Gambar 7. Bukti hasil evaluasi setelah field phone ditambahkan**

![Evaluasi penambahan phone](images/07_evaluasi_phone.png)

### 3.8 Analisis Static Code Generation dan Dynamic Runtime Loading

Go menggunakan kode hasil generasi Protobuf dan gRPC yang dibuat sebelum program dijalankan. Python menggunakan pendekatan serupa melalui `grpc_tools.protoc`. Jika kontrak berubah, kode hasil generasi harus diperbarui.

Node.js menggunakan `@grpc/proto-loader` untuk membaca file `.proto` ketika program berjalan. Pendekatan ini tidak membutuhkan stub JavaScript hasil generasi sebelumnya, tetapi file skema harus tersedia pada lokasi yang benar saat runtime.

## BAB IV — ANALISIS HASIL PRAKTIKUM

Praktikum menunjukkan bahwa file `user.proto` berfungsi sebagai kontrak bersama antara server Go dan client Python serta Node.js. Ketiga komponen menggunakan struktur request dan response yang sama meskipun dikembangkan dengan bahasa pemrograman berbeda.

Server Go berhasil dijalankan pada port 50051. Pengujian client Python menunjukkan data nama, email, dan nomor telepon. Client Node.js juga disiapkan untuk menampilkan ketiga field tersebut melalui kontrak yang sama.

Penambahan field `phone` memperlihatkan pentingnya memperbarui implementasi secara konsisten ketika kontrak API berubah. Praktikum juga memperlihatkan perbedaan pendekatan static code generation pada Go dan Python dengan dynamic runtime loading pada Node.js.

Proyek praktikum kemudian disimpan dan diunggah ke repository GitHub.

**Gambar 8. Repository proyek di GitHub**

![Repository GitHub](images/08_github_repository.png)

## BAB V — KESIMPULAN

Praktikum komunikasi inter-service dengan gRPC (Polyglot) menerapkan pendekatan Schema-First Development menggunakan Protocol Buffers. Server Go melayani request `GetUserProfile`, sedangkan client Python dan Node.js menggunakan service yang sama.

Penambahan field `phone` menunjukkan bahwa perubahan kontrak harus diterapkan secara konsisten pada server dan client. Praktikum juga memberikan pemahaman mengenai perbedaan static code generation dan dynamic runtime loading. Seluruh kode proyek disimpan dalam repository GitHub agar dapat diakses dan diperiksa.

## REPOSITORY GITHUB

Repository proyek praktikum dapat diakses melalui alamat berikut:

https://github.com/fajriayu/grpc-polyglot-lab
