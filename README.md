# prambanan

## Deskripsi

adalah pembaca sederhana dari NIK yang tertera pada KTP elektronik. Pembuatan
proyek ini didasari oleh banyaknya permintaan dari *developer* untuk melakukan
validasi dan pengambilan informasi pada KTP elektronik. Jika kamu seorang
*developer*, bantu kami mengembangkan proyek ini lebih kece lagiih! :3
 

## Fitur-fitur

- [ ] Memperoleh informasi provinsi, kota, dan kecamatan  dari NIK
- [ ] Memperoleh informasi jenis kelamin dari NIK
- [ ] Memeroleh informasi tanggal lahir dari NIK
- [ ] Membaca NIK dari foto KTP
- [ ] Menduga foto KTP valid, dilihat dari NIK dan data-data yang berkaitan

## Panduan Pengembangan (untuk *programmer*)

### Lingkungan Kebutuhan

1. `Go` versi `1.11.4` atau terbaru.
2.  `MySql` versi `5.7`.

### Langkah-Langkah
1. *Fork repository* ini;
2. Install dependensi yang diperlukan dengan `make setup`
3. Paket `main` untuk antarmuka web (REST) berada di `cmd/web`.

