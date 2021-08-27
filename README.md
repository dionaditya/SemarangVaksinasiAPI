# API Vaksinasi Semarang

scraped data from http://victori.semarangkota.go.id/

## All data vaccine venue per date

**BASE URL** : `https://vaksinasi-semarang.herokuapp.com`

**URL** : `/api/v1/vaccine-venue`

**Method** : `GET`

**Query** : `tanggal=dd-mmm-yyyy` default: today

**Auth required** : NO


## Success Response

**Code** : `200 OK`

**Content example**

```json
[
    {
        "Aksi": "                            Kuota Telah Terpenuhi\n                                            ",
        "Dosis Vaksinasi": "Dosis 1",
        "Jam Buka": "08:00:00 WIB",
        "Jam Tutup": "09:00:00 WIB",
        "Jenis Vaksinasi": "Sinovac",
        "Keterangan / Persyaratan Khusus": "PESERTA WARGA KOTA SEMARANG USIA 18 TAHUN KE ATAS DITUNJUKAN DENGAN KTP KOTA SEMARANG. HARAP MEMBAWA FOTOKOPI KTP/KK KOTA SEMARANG SAAT LAYANAN VAKSINASI",
        "Kuota": 250,
        "Lokasi Vaksinasi": "BALAIKOTA SEMARANG",
        "Pendaftaran Melalui": "Website",
        "Terisi": 251
    }
]
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
Error
```

## Available data vaccine venue per date

**BASE URL** : `https://vaksinasi-semarang.herokuapp.com`

**URL** : `/api/v1/vaccine-venue/available`

**Method** : `GET`

**Query** : `tanggal=dd-mmm-yyyy` default: today

**Auth required** : NO


## Success Response

**Code** : `200 OK`

**Content example**

```json
[
    {
        "Aksi": "                                                            Mendaftar\n                                                                        ",
        "Dosis Vaksinasi": "Dosis 1",
        "Jam Buka": "07:00:00 WIB",
        "Jam Tutup": "11:00:00 WIB",
        "Jenis Vaksinasi": "Sinopharm",
        "Keterangan / Persyaratan Khusus": "KHUSUS Masyarakat penyandang DISABILITAS & ODGJ (Gangguan Jiwa) di wilayah Kelurahan Tandang, Jangli, Sendangguwo, Kedungmundu, Sambiroto, Mangunharjo (Tembalang), & Kelurahan Sendangmulyo, Syarat :  Membawa FC KTP atau KK",
        "Kuota": 250,
        "Lokasi Vaksinasi": "Puskesmas Kedungmundu",
        "Pendaftaran Melalui": "Website",
        "Terisi": 218
    }
]
```

## Error Response

**Condition** : Unexpected error

**Code** : `200s`

**Content** :

```json
Error
```