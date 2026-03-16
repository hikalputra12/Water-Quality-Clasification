# Water Quality Classification Project

Proyek ini merupakan sistem klasifikasi untuk menentukan tingkat keamanan air minum berdasarkan berbagai parameter kimia dan mikrobiologi. Proyek ini mencakup seluruh siklus pengembangan model, mulai dari analisis data (EDA) hingga deployment model ke dalam aplikasi sederhana.

## 📌 Deskripsi Proyek

Tujuan utama dari proyek ini adalah untuk memprediksi apakah air aman untuk dikonsumsi (`is_safe`) berdasarkan 20 atribut kualitas air yang berbeda. Model dikembangkan menggunakan algoritma **Decision Tree Classifier** yang dilatih menggunakan dataset kualitas air yang komprehensif.

## 📂 Struktur Repositori

```text
.
├── data/
│   ├── waterQuality1.csv           # Dataset mentah untuk pelatihan
│   └── database_description.txt    # Penjelasan kolom pada dataset
├── training_model/
│   └── training_model.ipynb        # Proses pembersihan data & pelatihan model
├── water_quality/
│   ├── main.py                     # Aplikasi utama untuk interface prediksi
│   └── model/
│       └── model_dt_final.joblib   # Model yang sudah dilatih (serialisasi)
├── requirements.txt                # Daftar library Python yang dibutuhkan
└── README.md                       # Dokumentasi proyek
```

## 📊 Dataset Detail

Dataset ini mencakup berbagai parameter kualitas air yang dikategorikan sebagai berikut:

### **Fitur (Fitur Kimia & Biologi)**
| Kategori | Parameter |
| :--- | :--- |
| **Logam Berat** | Aluminium, Arsenic, Barium, Cadmium, Chromium, Copper, Lead, Mercury, Selenium, Silver, Uranium |
| **Zat Kimia & Mineral** | Ammonia, Chloramine, Fluoride, Nitrates, Nitrites, Perchlorate |
| **Organisme (Mikrobiologi)** | Bacteria, Viruses |
| **Radioaktif** | Radium |

### **Target (Label)**
Kolom target untuk prediksi adalah `is_safe`:
* **0**: Tidak Aman (Not Safe)
* **1**: Aman (Safe)

## 🛠️ Persiapan Lingkungan (Setup)

Pastikan Anda sudah menginstal **Python** (versi 3.8 atau lebih baru disarankan) di sistem Anda. Ikuti langkah-langkah di bawah ini untuk menjalankan proyek secara lokal:

### **1. Clone Repositori**
Buka terminal atau command prompt, lalu jalankan perintah berikut untuk mengunduh repositori:
```bash
git clone [https://github.com/hikalputra12/water-quality-clasification.git](https://github.com/hikalputra12/water-quality-clasification.git)
cd water-quality-clasification
```

### **2. Buat Virtual Environment (Opsional/Disarankan)**
Untuk menjaga agar library tidak bentrok dengan proyek lain:

```bash
# Windows
python -m venv venv
venv\Scripts\activate

# macOS/Linux
python3 -m venv venv
source venv/bin/activate
```
# **3. Instal Library yang Dibutuhkan**
Gunakan pip untuk menginstal semua dependensi yang tertera di dalam file requirements.txt:

```bash
pip install -r requirements.txt
```

## **🚀 Penggunaan**
### **1. Pelatihan Model**
Jika Anda ingin melihat proses analisis data dan cara model dilatih, buka file notebook di folder training_model/:
```bash 
jupyter notebook training_model/training_model.ipynb
```
### **2. Menjalankan Aplikasi Prediksi**
Untuk menjalankan aplikasi utama yang menggunakan model yang sudah dilatih untuk melakukan prediksi:
```bash
python water_quality/main.py
```
## 🧠 Metodologi Model
### 1. Preprocessing: Menangani nilai yang hilang, menghapus baris yang tidak valid (seperti nilai '#NUM!' pada kolom ammonia dan is_safe).

### 2. Model: Menggunakan beberapa model tetapi model yang di pilih yaitu Decision Tree Classifier.

### 3. Evaluasi: Model dievaluasi menggunakan metrik akurasi, presisi, dan recall untuk memastikan klasifikasi air aman/tidak aman dilakukan dengan benar.

### 4. Deployment: Model disimpan menggunakan joblib agar dapat dimuat kembali secara instan tanpa perlu pelatihan ulang.