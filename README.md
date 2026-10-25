LAPORAN 
KOMUNIKASI INTER-SERVICE DENGAN gRPC (POLYGLOT)
 
Ditulis Oleh :
NAMA	: FAJRI AYU
NIM		: 2024903430026
Kelas		: TRKJ 3C

 

TEKNOLOGI REKAYASA KOMPUTER JARINGAN
TEKNOLOGI INFORMASI dan KOMPUTER
POLITEKNIK NEGERI LHOKSEUMAWE
TAHUN 2026
 
LEMBAR PENGESAHAN

No.pratikum				: 02/TIK/TRKJ-3C/Paradigma Sistem Untuk IT
Judul					: Komunikasi Inter-Service Dengan Grpc (Polyglot)
Nama					: Fajri Ayu
NIM					: 2024903430026
Kelas					: TRKJ-3C
Jurusan				: Teknologi Informasi Dan Komputer
Prodi 					: Teknologi Rekayasa Komputer Dan Jaringan
Tanggal pratikum			: 19 Oktober 2026
Tanggal penyerahan			: 19 Oktober 2026
Repository github			: https://github.com/fajriayu/grpc-polyglot-lab




          Dosen pengampu,


 
Muhammad Reza Zulman, S.ST.,M.Sc
NIP. 199205012022031005
 

KATA PENGANTAR
Puji syukur ke hadirat Allah SWT karena atas rahmat dan karunia-Nya, laporan praktikum berjudul “Komunikasi Inter-service dengan gRPC (Polyglot)” ini dapat diselesaikan. Laporan ini disusun sebagai dokumentasi pelaksanaan praktikum yang bertujuan untuk memahami komunikasi antar-service menggunakan gRPC dengan menerapkan beberapa bahasa pemrograman (polyglot), yaitu Go, Python, dan Node.js.
Praktikum dilakukan secara bertahap, dimulai dari menyiapkan lingkungan pengembangan dan perangkat yang dibutuhkan, membuat definisi layanan menggunakan Protocol Buffers melalui file user.proto, membangun server gRPC menggunakan bahasa Go, serta membuat client menggunakan Python dan Node.js. Selanjutnya, dilakukan pengujian komunikasi antara client dan server untuk memastikan bahwa layanan dapat berjalan dengan baik. Praktikum ini juga mencakup penambahan field phone pada pesan respons, regenerasi kode yang diperlukan, serta pengujian untuk memastikan data nama, email, dan nomor telepon dapat diterima oleh client.
Penulis menyadari bahwa laporan praktikum ini masih memiliki kekurangan dan dapat disempurnakan. Oleh karena itu, kritik dan saran yang membangun sangat diharapkan agar penyusunan laporan berikutnya dapat menjadi lebih baik. Semoga laporan ini dapat memberikan manfaat bagi pembaca serta menjadi dokumentasi yang jelas mengenai proses pelaksanaan dan hasil praktikum komunikasi antar-service menggunakan gRPC.
 
DAFTAR ISI
LEMBAR PENGESAHAN	0
KATA PENGANTAR	1
DAFTAR ISI	2
BAB I PENDAHULUAN	3
1.1 Latar Belakang	3
1.2 Tujuan Praktikum	3
1.3 Dasar Teori	3
BAB II ALAT DAN BAHAN	4
BAB III LANGKAH KERJA DAN IMPLEMENTASI	5
3.1 Persiapan dan struktur proyek	5
3.2 Membuat kontrak API user.proto	6
3.3 Implementasi gRPC server dengan Go	7
3.4 Implementasi client Python	7
3.5 Implementasi client Node.js	8
3.6 Pengujian komunikasi inter-service	9
3.7 Tugas evaluasi: penambahan field phone	10
3.8 Analisis static code generation dan dynamic runtime loading	11
BAB IV ANALISIS HASIL PRAKTIKUM	12
BAB V KESIMPULAN	12

 
BAB I PENDAHULUAN
1.1 Latar Belakang
Dalam arsitektur microservices, layanan yang dibuat menggunakan bahasa pemrograman berbeda perlu bertukar data melalui kontrak yang konsisten. gRPC (Remote Procedure Call) memungkinkan suatu program memanggil metode pada layanan lain dengan menggunakan kontrak Protocol Buffers (Protobuf). Dengan cara ini, server dan client dapat dikembangkan menggunakan bahasa berbeda tetapi tetap menggunakan definisi request, response, dan service yang sama.
Praktikum ini menerapkan komunikasi lintas bahasa (polyglot) dengan Go sebagai gRPC server, serta Python dan Node.js sebagai client. Kontrak API didefinisikan pada berkas user.proto. Pada tugas evaluasi, struktur response diperluas dengan field phone agar perubahan kontrak dapat diuji pada seluruh komponen.
1.2 Tujuan Praktikum
•	Memahami pendekatan Schema-First Development menggunakan Protocol Buffers.
•	Memastikan perangkat pengembangan memiliki Git, Python, Node.js, Go, dan compiler protoc.
•	Membuat gRPC server menggunakan bahasa Go.
•	Membuat client Python dan Node.js yang memanggil service Go yang sama.
•	Menambahkan field phone pada UserResponse, membangkitkan ulang kode yang diperlukan, dan menampilkan data telepon.
•	Mendokumentasikan perbedaan static code generation dan dynamic runtime loading serta mengunggah proyek ke GitHub.
1.3 Dasar Teori
gRPC adalah framework RPC untuk komunikasi antarlayanan. Dalam praktikum ini, komunikasi menggunakan Protocol Buffers sebagai format kontrak dan HTTP/2 sebagai transport. File .proto berperan sebagai single source of truth yang mendefinisikan pesan dan metode service. Tooling protobuf menghasilkan kode pendukung agar client dan server menggunakan format data yang kompatibel.
Pola GetUserProfile yang digunakan adalah unary RPC: client mengirim satu UserRequest dan server mengembalikan satu UserResponse. Go dan Python menggunakan kode hasil generasi (static code generation), sedangkan client Node.js membaca definisi .proto ketika runtime melalui @grpc/proto-loader (dynamic runtime loading). 

BAB II ALAT DAN BAHAN
Perangkat yang digunakan adalah laptop Windows dengan Command Prompt (CMD) dan Visual Studio Code/Notepad untuk mengelola berkas proyek.
Komponen	Versi/konfigurasi yang terverifikasi
Git	2.51.0.windows.1
Python	3.14.3
Node.js	v24.14.0
npm	11.9.0
Go	go1.27.2 windows/amd64
Protocol Buffers Compiler	libprotoc 36.2
gRPC Server	Go, port TCP 50051
Client	Python dan Node.js
Repository	https://github.com/fajriayu/grpc-polyglot-lab

 

BAB III LANGKAH KERJA DAN IMPLEMENTASI
3.1 Persiapan dan struktur proyek
Proyek dibuat pada direktori lokal grpc-polyglot-lab. Git diinisialisasi untuk mencatat perubahan kode, dan file .gitignore digunakan untuk mengecualikan direktori dependensi dan file sementara.
grpc-polyglot-lab/
├── .gitignore
├── README.md
├── user.proto
├── go-server/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── pb/
│       ├── user.pb.go
│       └── user_grpc.pb.go
├── python-client/
│   ├── client.py
│   ├── user_pb2.py
│   └── user_pb2_grpc.py
└── node-client/
    ├── client.js
    ├── package.json
    └── package-lock.json
File .gitignore yang digunakan berisi aturan untuk tidak memasukkan node_modules/, __pycache__/, .venv/, dan file executable *.exe ke riwayat Git.
 
Gambar 1. Struktur direktori proyek dan persiapan Git.
3.2 Membuat kontrak API user.proto
Kontrak mendefinisikan pesan UserRequest dengan user_id bertipe int32, pesan UserResponse, dan service UserService dengan RPC GetUserProfile. Untuk memenuhi tugas evaluasi, ditambahkan field phone dengan nomor field 3. Perubahan kontrak harus tercermin pada implementasi server dan client.
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
Catatan: nilai go_package disesuaikan dengan struktur modul Go yang digunakan pada proyek. Nama package dan path impor harus konsisten dengan konfigurasi go.mod.
 
Gambar 2. Isi user.proto yang telah ditambahkan field phone.
3.3 Implementasi gRPC server dengan Go
Server Go mengimplementasikan metode GetUserProfile dan mendengarkan koneksi TCP pada port 50051. Ketika client meminta user_id = 1, server mengembalikan data profil Fajri Ayu. Untuk ID lain, server mengembalikan nama 'User Tidak Ditemukan' dengan email dan phone kosong. Kode Protobuf dan gRPC untuk Go ditempatkan di direktori go-server/pb.
// Ringkasan logika server
if req.GetUserId() == 1 {
    return &pb.UserResponse{
        Name:  "Fajri Ayu",
        Email: "Ayu@universitas.ac.id",
        Phone: "081234567890",
    }, nil
}
 
return &pb.UserResponse{
    Name:  "User Tidak Ditemukan",
    Email: "",
    Phone: "",
}, nil
Server dijalankan dari folder go-server menggunakan perintah:
go run main.go
Output yang tercatat saat server dijalankan menunjukkan server gRPC berjalan pada port 50051.
 Gambar 3. Terminal yang menampilkan server Go berjalan di port 50051.
3.4 Implementasi client Python
Client Python menggunakan library grpcio dan modul hasil generasi user_pb2 serta user_pb2_grpc. Client membuka channel ke localhost:50051, membuat stub UserService, kemudian memanggil GetUserProfile dengan user_id = 1. Setelah field phone ditambahkan, modul Python dibangkitkan ulang menggunakan grpc_tools.protoc.
python -m grpc_tools.protoc --proto_path=../ --python_out=. --grpc_python_out=. ../user.proto
# Ringkasan client.py
with grpc.insecure_channel("localhost:50051") as channel:
    stub = user_pb2_grpc.UserServiceStub(channel)
    response = stub.GetUserProfile(user_pb2.UserRequest(user_id=1))
    print("Nama  :", response.name)
    print("Email :", response.email)
print("Phone :", response.phone)

Hasil uji Python yang berhasil:
Nama  : Fajri Ayu
Email : Ayu@universitas.ac.id
Phone : 081234567890
 
Gambar 4. Hasil client Python

3.5 Implementasi client Node.js
Client Node.js menggunakan paket @grpc/grpc-js dan @grpc/proto-loader. Berbeda dari Go dan Python, client Node.js membaca file user.proto pada saat program berjalan sehingga tidak perlu menghasilkan file JavaScript stub secara terpisah. Properti request ditulis sebagai userId karena opsi keepCase bernilai false.
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
 
const PROTO_PATH = '../user.proto';
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: false,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});
 
const userProto = grpc.loadPackageDefinition(packageDefinition).user;
const client = new userProto.UserService(
    'localhost:50051',
    grpc.credentials.createInsecure()
);
 
client.GetUserProfile({ userId: 1 }, (error, response) => {
    if (error) {
        console.error('Gagal mengambil data:', error.message);
        return;
    }
    console.log('Nama  :', response.name);
    console.log('Email :', response.email);
    console.log('Phone :', response.phone);
});
Client Node.js dijalankan dari direktori node-client dengan perintah node client.js, sementara server Go tetap berjalan.
 
Hasil client Node.js menampilkan nama, email, dan phone.
3.6 Pengujian komunikasi inter-service
Pengujian dilakukan dengan menjalankan server Go terlebih dahulu, lalu menjalankan client Python dan Node.js dalam terminal terpisah. Kedua client mengirim permintaan ke alamat localhost:50051 dan menerima data dari implementasi server yang sama. Hasil Python yang tercatat berhasil menampilkan tiga field. Pengujian Node.js juga disiapkan untuk menampilkan tiga field sesuai kontrak; tempelkan screenshot terminal aktual sebagai bukti.
Komponen	Peran	Hasil yang diharapkan
Go server	Menerima RPC dan mengembalikan UserResponse	Server berjalan di port 50051
Python client	Memanggil GetUserProfile(user_id=1)	Nama, email, phone tercetak
Node.js client	Memanggil RPC melalui skema yang dibaca runtime	Nama, email, phone tercetak
 Gambar 6. Log server setelah menerima request dari client 
3.7 Tugas evaluasi: penambahan field phone
Tugas evaluasi mandiri pada modul meminta penambahan string phone = 3 ke UserResponse, kompilasi ulang kode untuk bahasa yang memerlukan code generation, serta penyesuaian server dan client agar nomor telepon ikut dicetak. Implementasi ini menambahkan phone pada user.proto, server Go, client Python, dan client Node.js. Kode Python dan Go yang dihasilkan ulang disimpan pada direktori masing-masing.
Pada proyek ini, nomor telepon uji adalah 081234567890. Nomor tersebut merupakan data contoh yang ditanam langsung pada server untuk keperluan praktikum, bukan hasil pembacaan dari basis data eksternal.
 Gambar 7. Bukti hasil evaluasi setelah field phone ditambahkan.
3.8 Analisis static code generation dan dynamic runtime loading
Pada Go, file user.proto dikompilasi menggunakan protoc dengan plugin Go dan gRPC sehingga menghasilkan user.pb.go dan user_grpc.pb.go. Pendekatan ini termasuk static code generation karena kode pendukung dibuat sebelum program dijalankan. Python menggunakan grpc_tools.protoc untuk menghasilkan user_pb2.py dan user_pb2_grpc.py dengan prinsip yang sama. Jika kontrak berubah, file hasil generasi perlu diperbarui agar program memahami field terbaru.
Node.js dalam proyek ini memakai @grpc/proto-loader untuk membaca user.proto saat runtime. Pendekatan dynamic runtime loading tidak memerlukan stub JavaScript yang dihasilkan sebelumnya, sehingga alur awal lebih sederhana. Namun, file skema harus tersedia pada path yang benar ketika aplikasi berjalan, dan perubahan kontrak perlu tetap dijaga kompatibilitasnya dengan server. Ketiga komponen tetap menggunakan kontrak yang sama sehingga field dan metode RPC dapat dipahami lintas bahasa. 
BAB IV ANALISIS HASIL PRAKTIKUM
Berdasarkan proses pengerjaan, kontrak user.proto menjadi acuan bersama bagi server Go dan dua client yang berbeda bahasa. Pemisahan kontrak dari implementasi membantu menjaga konsistensi bentuk request dan response. Metode GetUserProfile merupakan unary RPC karena setiap pemanggilan mengirim satu request dan menerima satu response.
Server Go berhasil dijalankan pada port 50051. Hasil yang telah tercatat dari client Python menunjukkan nama Fajri Ayu, email Ayu@universitas.ac.id, dan nomor telepon 081234567890. Client Node.js menggunakan skema yang sama dan kode untuk mencetak field phone telah disiapkan. Screenshot aktual dari terminal Node.js perlu dilampirkan untuk mendokumentasikan bukti pengujian pada komputer.
Penambahan field phone memperlihatkan bahwa perubahan kontrak perlu diterapkan secara konsisten. Pada Go dan Python, file hasil generasi diperbarui setelah skema berubah. Pada Node.js, loader membaca skema ketika runtime, sehingga field baru tersedia selama file .proto yang digunakan sudah diperbarui. Dengan demikian, hasil praktikum mendukung konsep komunikasi inter-service polyglot dengan kontrak bersama.
Proyek juga berhasil diunggah ke GitHub. Perintah git push menghasilkan branch master -> master dan pesan bahwa branch lokal melacak origin/master. Repository dapat digunakan untuk menyimpan kode sumber dan dokumentasi.
BAB V KESIMPULAN
Praktikum komunikasi inter-service dengan gRPC (Polyglot) telah menerapkan pendekatan Schema-First Development dengan Protocol Buffers. Sebuah kontrak user.proto mendefinisikan UserRequest, UserResponse, dan UserService. Server Go melayani permintaan GetUserProfile pada port 50051, sedangkan client Python dan Node.js menggunakan service yang sama meskipun dibangun dengan bahasa pemrograman berbeda.
Pada tugas evaluasi, field phone ditambahkan ke UserResponse dan implementasi server serta client disesuaikan untuk menampilkan nomor telepon. Pengujian Python tercatat berhasil menampilkan nama, email, dan phone. Perbedaan pendekatan code generation juga dipelajari: Go dan Python menggunakan kode yang dibangkitkan sebelum runtime, sedangkan Node.js memuat definisi .proto saat runtime. Terakhir, proyek berhasil di-commit dan di-push ke repository GitHub.
