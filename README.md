Terima kasih kepada Bapak Ariaseta Setia Alam dan Muhammad Zuhrul Umam.

Structure folder

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```js
/ (Root Project)
├── go.mod                      # File manajemen dependensi [1]
├── /cmd                        # Aplikasi utama (main application) [2]
│   └── /myapp                  # Nama direktori sesuai nama file eksekusi [2]
├── /internal                   # Kode aplikasi dan library privat (tidak bisa diimpor luar) [2], [3]
│   ├── /app                    # Kode spesifik aplikasi (opsional) [4]
│   └── /pkg                    # Kode library yang dibagi antar komponen internal (opsional) [4]
├── /pkg                        # Kode library publik yang aman digunakan aplikasi eksternal [2], [5]
├── /vendor                     # Dependensi aplikasi (jika menggunakan vendor) [6]
├── /api                        # Spesifikasi OpenAPI/Swagger, skema JSON, protokol [7]
├── /web                        # Aset web statis, template server-side, SPAs [7]
├── /configs                    # Template file konfigurasi atau konfigurasi default [7]
├── /init                       # Konfigurasi sistem init (systemd, dll) [7]
├── /scripts                    # Skrip untuk build, instalasi, analisis [8]
├── /build                      # Pemaketan dan Integrasi Berkelanjutan (CI) [8]
│   ├── /package                # Skrip paket dan konfigurasi cloud/container [8]
│   └── /ci                     # Skrip dan konfigurasi CI [9]
├── /deployments                # Konfigurasi IaaS, PaaS, orkestrasi (Docker/K8s) [9]
├── /test                       # Pengujian eksternal tambahan [10]
│   └── /testdata               # Data pengujian (diabaikan oleh Go compiler) [10]
├── /docs                       # Dokumentasi desain dan pengguna [11]
├── /tools                      # Tools pendukung proyek [11]
├── /examples                   # Contoh aplikasi atau library publik [11]
├── /third_party                # Tools eksternal atau kode yang di-fork [11]
├── /githooks                   # Git hooks [12]
├── /assets                     # Gambar, logo, dan aset repositori lainnya [12]
└── /website                    # Data situs web proyek
```
