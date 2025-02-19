# Final Project - CV Generator

## 1. Latar Belakang & Tujuan
Aplikasi **CV Generator** ini bertujuan untuk memudahkan pengguna dalam membuat, mengelola, dan menyimpan berbagai versi CV (Curriculum Vitae). Dengan fitur register, login, serta CRUD CV, user dapat:
- Menyesuaikan CV sesuai kebutuhan pekerjaan atau perusahaan berbeda.
- Menyimpan riwayat CV dan memperbarui CV yang sudah ada.
- Menghapus CV yang tidak diperlukan lagi.

## 2. Teknologi yang Digunakan
- **Backend**: Golang (Gin, GORM) + PostgreSQL
- **Frontend**: React (Create React App) + Chakra UI
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)

## 3. Arsitektur Proyek
Struktur folder utama:

### 3.1 Backend
- `models/` berisi definisi struct (User, CV, dsb.).
- `controllers/` berisi handler untuk Auth (Register, Login) dan CRUD CV.
- `routes/` berisi routing (Gin).
- `middlewares/` berisi middleware JWT.
- `database/` berisi inisialisasi koneksi ke PostgreSQL dan AutoMigrate.

### 3.2 Frontend
- `src/components/` berisi komponen umum (Layout, ProtectedRoute, dsb.).
- `src/pages/` berisi halaman (Home, Login, Register, CVList, CVDetail, CVCreate).
- `src/services/api.js` berisi konfigurasi axios (baseURL, interceptors JWT).

## 4. Database & Migrasi
Menggunakan **PostgreSQL**. Jika kamu memakai GORM `AutoMigrate`, cukup jalankan backend, GORM akan membuat tabel.  

**Struktur Tabel Utama**:
1. **users**  
   - `id` (PK, serial)  
   - `name` (varchar)  
   - `email` (varchar unique)  
   - `password_hash` (text)  
   - `created_at`, `updated_at`
2. **cvs**  
   - `id` (PK, serial)  
   - `user_id` (FK → users.id)  
   - `title` (varchar)  
   - `content` (text)  
   - `template` (varchar)  
   - `created_at`, `updated_at`

*(Opsional) Tabel skills dan cv_skills untuk relasi many-to-many.*  

## 5. Daftar API Endpoints (Backend)

Base URL: `http://localhost:8080`

### 5.1 Authentication
1. **POST** `/register`  
   - **Request Body** (JSON):
     ```json
     {
       "name": "string",
       "email": "string",
       "password": "string"
     }
     ```
   - **Response**: `201 Created` jika sukses, atau error (400, 500).

2. **POST** `/login`
   - **Request Body** (JSON):
     ```json
     {
       "email": "string",
       "password": "string"
     }
     ```
   - **Response**: `200 OK` dengan JWT token atau `401 Unauthorized` jika gagal.

### 5.2 CV
Semua endpoint CV membutuhkan header **Authorization: Bearer <token>**.

1. **GET** `/cv`  
   - Mengambil semua CV milik user yang sedang login.
   - **Response**: `200 OK` berisi array CV atau error.

2. **GET** `/cv/:id`
   - Mengambil detail CV by ID (hanya jika user pemilik).
   - **Response**: `200 OK` berisi objek CV atau `404 Not Found` jika tidak ada.

3. **POST** `/cv`
   - Membuat CV baru.
   - **Request Body**:
     ```json
     {
       "title": "string",
       "content": "string",
       "template": "string"
     }
     ```
   - **Response**: `201 Created` atau error.

4. **PUT** `/cv/:id`
   - Mengupdate CV tertentu (hanya pemilik).
   - **Request Body**:
     ```json
     {
       "title": "string",
       "content": "string",
       "template": "string"
     }
     ```
   - **Response**: `200 OK` atau `404 Not Found` jika ID tidak ditemukan.

5. **DELETE** `/cv/:id`
   - Menghapus CV tertentu (hanya pemilik).
   - **Response**: `200 OK` atau `404 Not Found`.

## 6. Flow Interaksi Frontend - Backend

1. **Register** di halaman FE → panggil endpoint `/register` → backend simpan user (hash password) → response sukses.
2. **Login** → panggil endpoint `/login` → backend verifikasi → kembalikan JWT → FE simpan di localStorage.
3. **CRUD CV** → FE panggil endpoint `/cv`, `/cv/:id`, dsb. dengan header `Authorization: Bearer <token>` → backend verifikasi token → DB operation → response.

## 7. Cara Menjalankan (Build & Run)

### 7.1 Setup Database
1. Buat database PostgreSQL (misalnya `cv_db`).
2. Update konfigurasi DB di file `database/db.go` (host, user, password, dbname).

### 7.2 Jalankan Backend
```bash
cd final-project-sa-be
go mod tidy
go run main.go
