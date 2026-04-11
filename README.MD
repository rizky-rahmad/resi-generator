```markdown
# Fleetify - Multi-Step Invoice & Resi Generator

Aplikasi web *Fullstack* berbasis *wizard* (multi-step) untuk mencatat pembuatan Resi Pengiriman dan men-generate Invoice. Dibangun untuk memenuhi spesifikasi Technical Test Junior Fullstack Engineer Fleetify.

##  Tech Stack

**Frontend:**
- Next.js 14 (Pages Router)
- TypeScript (Strict Mode)
- Zustand (State Management + Persist)
- TanStack React Query v5
- TailwindCSS

**Backend:**
- Golang (Go Fiber)
- PostgreSQL
- GORM (Object Relational Mapping)
- JWT Authentication
- Newman (Automated API Testing)

---

##  Prasyarat (Prerequisites)
Pastikan sistem Anda telah terinstal:
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

##  Cara Menjalankan Aplikasi (Zero-Setup Run)

Aplikasi ini menggunakan konfigurasi *Multi-stage Docker Build* dan telah diorkestrasi menggunakan Docker Compose untuk memenuhi syarat **Zero-Setup Run**. Database, Backend, dan Frontend akan berjalan secara otomatis dalam satu jaringan (*network*).

1. **Clone repositori ini:**
   ```bash
   git clone <link-repositori-anda>
   cd resi_generator
   ```

2. **Masuk ke direktori backend (tempat docker-compose berada):**
   ```bash
   cd backend
   ```

3. **Jalankan Docker Compose:**
   ```bash
   docker-compose up --build -d
   ```

4. **Akses Aplikasi:**
   - **Frontend UI:** Buka `http://localhost:3000` di peramban Anda.
   - **Backend API:** Berjalan di `http://localhost:8080` (Internal).

*Catatan: Saat pertama kali dijalankan, backend akan otomatis melakukan Database Migration dan Seeding data awal.*

---

##  Kredensial Uji Coba (Seeded Accounts)

Gunakan akun berikut untuk login ke dalam aplikasi:

| Role | Username | Password | Keterangan |
| :--- | :--- | :--- | :--- |
| **Admin** | `admin` | `admin123` | Memiliki akses penuh (termasuk validasi harga di payload) |
| **Kerani** | `kerani` | `kerani123` | Hanya mengirim `item_id` dan `quantity` |

---

## Fitur & Pencapaian Requirement

Aplikasi ini telah memenuhi seluruh Core Requirements dan Bonus Challenge:

- [x] **Zero-Setup Run:** Orkestrasi penuh via Docker Compose.
- [x] **Multi-stage Dockerfile:** Image Docker teroptimasi dan ringan berbasis Alpine.
- [x] **Auth & Routing:** Endpoint terlindungi oleh JWT Middleware.
- [x] **DB Transaction (ACID):** Penyimpanan Header dan Detail Invoice menggunakan `db.Transaction()` di GORM untuk mencegah data anomali (*Rollback* otomatis jika gagal).
- [x] **Zero-Trust Logic:** Backend **tidak** mempercayai harga dan subtotal dari payload UI. Kalkulasi total harga dihitung ulang murni dari master tabel database.
- [x] **Debounce & Race Condition:** Pencarian barang diimplementasikan menggunakan teknik *Debouncing* dan pembatalan *request* kedaluwarsa (React Query AbortSignal).
- [x] **Print Layout (Bonus):** Implementasi cetak invoice (*Hidden Print Layout*) menggunakan utility classes Tailwind (`print:hidden`, `print:block`) tanpa mengubah struktur DOM utama.
- [x] **Webhook Integration (Bonus):** HTTP POST asinkron (Goroutine) berjalan di *background* setelah invoice sukses dibuat.

---

##  Automated Testing (Newman)
Untuk menjalankan *automated API testing* menggunakan Postman/Newman yang sudah terintegrasi di dalam Docker:

```bash
# Pastikan Anda masih berada di folder 'backend'
docker-compose run --rm newman
```
```

---
