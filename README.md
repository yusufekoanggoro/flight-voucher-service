
# ✈️ Flight Voucher Service

Sebuah REST API untuk memeriksa dan menghasilkan voucher tempat duduk bagi kru penerbangan, dibuat dengan GoLang dan SQLite, serta dijalankan menggunakan Docker Compose.

## 🔧 Fitur

- `POST /api/check`: Mengecek apakah suatu penerbangan sudah memiliki voucher yang dibuat.
- `POST /api/generate`: Menghasilkan 3 tempat duduk acak untuk kru berdasarkan jenis pesawat.
- Validasi agar tempat duduk tidak duplikat per penerbangan dan tanggal.
- Layout tempat duduk sesuai jenis pesawat.
- Menyimpan data ke SQLite (`data/vouchers.db`).
- Query aman dengan parameterisasi (anti SQL Injection).
- Struktur proyek modular dan mudah dipelihara.

---

## 🚀 Cara Menjalankan (dengan Docker Compose)

### 1. Clone Repository

```bash
git clone https://github.com/yusufekoanggoro/flight-voucher-service
cd flight-voucher-service
```

### 2. Build dan Jalankan dengan Docker Compose

```bash
docker-compose up --build
```

Setelah proses selesai, server akan berjalan di:

```
http://localhost:8080
```

---

## 📂 Struktur Proyek

```
.
├── data/
│   └── vouchers.db             # File database SQLite
│
├── internal/
│   ├── factory/
│   │   ├── base/               # Definisi tipe modul dan konstanta
│   │   ├── interfaces/         # Kontrak untuk dependency injection
│   │   └── module.go           # Registrasi dan inisialisasi modul
│   │
│   ├── infrastructure/
│   │   └── db.go               # Inisialisasi koneksi database
│   │
│   └── modules/
│       └── voucher/
│           ├── delivery/       # Handler REST API, Kafka, Worker, gRPC, dll
│           ├── domain/         # Entitas, DTO, model
│           ├── repository/     # Interaksi ke database
│           ├── usecase/        # Logika bisnis / alur proses
│           └── module.go       # Modul inisialisasi voucher
│
├── utils/
│   └── response.go             # Fungsi utilitas untuk response
│
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 📮 Dokumentasi API

### 🔎 POST `/api/check`

**Request**

```json
{
  "flightNumber": "GA102",
  "date": "2025-07-12"
}
```

**Response**

```json
{
  "exists": true
}
```

---

### 🎫 POST `/api/generate`

**Request**

```json
{
  "name": "Sarah",
  "id": "98123",
  "flightNumber": "ID102",
  "date": "2025-07-12",
  "aircraft": "Airbus 320"
}
```

**Response**

```json
{
  "success": true,
  "seats": ["3B", "7C", "14D"]
}
```

## 🧼 Menghentikan dan Membersihkan Kontainer

```bash
docker-compose down
```

---

## 📞 Kontak

Jika ada pertanyaan atau masukan terkait proyek ini, silakan hubungi melalui email atau GitHub issue.