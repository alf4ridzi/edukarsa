# EduKarsa

Platform Pendidikan Terpadu untuk Guru & Siswa

Edukarsa adalah platform pendidikan yang dirancang untuk memudahkan guru dalam mengajar, mengevaluasi, dan memahami kebutuhan belajar siswa secara menyeluruh.


## Fitur Utama

- Manajemen pembelajaran terpadu
- Sistem evaluasi dan penilaian
- Kolaborasi guru dan siswa
- Aplikasi mobile yang responsif

---

## Teknologi

- **Mobile**: React Native
- **Backend**: Golang
- **Database**: PostgreSQL
- **Containerization**: Docker

---

## Instalasi

### Sebelum menjalankan

Pastikan Anda telah menginstal:
- [Git](https://git-scm.com/)
- [Node.js](https://nodejs.org/) (untuk Mobile App)
- [Golang](https://go.dev/doc/install) (untuk Backend Native)
- [Docker](https://www.docker.com/) (opsional, untuk Backend Docker)

### Clone Repository

```bash
git clone https://github.com/alf4ridzi/edukarsa
cd edukarsa
```

---

## Mobile App

### 1. Masuk ke direktori app

```bash
cd app
```

### 2. Setup Environment Variables

```bash
cp env.example .env
```

> **Catatan**: Jangan lupa edit file `.env` sesuai dengan konfigurasi Anda

### 3. Install Dependencies

```bash
npm install
```

### 4. Jalankan Aplikasi

```bash
npm start
```

---

## Backend

### 1. Masuk ke direktori backend

```bash
cd backend
```

### 2. Setup Environment Variables

```bash
cp env.example .env
```

> **Catatan**: Jangan lupa edit file `.env` sesuai dengan konfigurasi Anda

---

### Opsi 1: Menggunakan Docker (Recommended)

```bash
docker compose up --build -d
```

---

### Opsi 2: Native (Tanpa Docker)

#### 1. Install Golang

Jika belum terinstall, download dan install dari: https://go.dev/doc/install

#### 2. Install Air (Live Reload)

```bash
go install github.com/air-verse/air@latest
```

#### 3. Tambahkan Golang ke PATH

Pastikan Golang sudah ditambahkan ke PATH environment variable Anda.

**Linux/MacOS:**
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

**Windows (PowerShell):**
```powershell
$env:Path += ";$(go env GOPATH)\bin"
```

#### 4. Jalankan Backend

```bash
air
```

---

## Lisensi

Project ini dilisensikan di bawah [MIT License](LICENSE)
