# 💧 Water Quality Classification Project

Proyek ini adalah Proyek ini adalah sarana latihan untuk memahami siklus pengembangan Machine Learning dan integrasi sistem backend (Python & Go). Proyek ini telah diperbarui untuk mendukung sistem inference berperforma tinggi menggunakan **FastAPI (Python)** dan **ONNX Runtime (Go)**.

## 📌 Deskripsi Proyek

Tujuan utama proyek ini adalah memprediksi status keamanan air (is_safe) berdasarkan 20 atribut kualitas air. Model inti dikembangkan menggunakan **Decision Tree Classifier** yang telah dioptimalkan untuk menangani data yang tidak seimbang (imbalanced data).

## 📂 Struktur Repositori
```text
.
├── data/
│   ├── waterQuality1.csv           # Dataset mentah
│   └── database_description.txt    # Penjelasan ambang batas (threshold) bahaya
├── training_model/
│   ├── training_model.ipynb        # Notebook EDA & Pelatihan Model
│   └── requirements.txt            # Library khusus pelatihan (scikit-learn, pandas, dll)
├── water_quality/                  # REST API (Python)
│   ├── main.py                     # Implementasi FastAPI
│   └── model/
│       ├── model_dt_final.joblib   # Serialisasi model Decision Tree
│       └── preprocessor.joblib     # Scaler/Preprocessor data
├── deploy-go/                      # High Performance Inference (Go)
│   ├── main.go                     # API Server menggunakan Gin/Fiber
│   ├── ml/engine.go                # Wrapper ONNX Runtime untuk Go
│   └── model/
│       └── model_anda.onnx         # Model dalam format lintas platform (ONNX)
└── README.md                       # Dokumentasi utama
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

### **2. Backend Python (FastAPI)**
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

## **4. Backend Go (Inference)**
Pastikan Anda memiliki onnxruntime library terpasang di sistem Linux Anda

```bash
cd deploy-go
go mod tidy
go run main.go
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

### **3. Penggunaan API 
Endpoint Prediksi (POST)
Kirimkan data mentah kimia air dalam format JSON ke http://localhost:8000/predict untuk python dan  http://localhost:8080/predict untuk golang.
Contoh Body Request:
```text
{
  "aluminium": 1.65, "ammonia": 9.08, "arsenic": 0.04, "barium": 2.85,
  "cadmium": 0.007, "chloramine": 0.35, "chromium": 0.83, "copper": 0.17,
  "flouride": 0.05, "bacteria": 0.2, "viruses": 0, "lead": 0.054,
  "nitrates": 16.08, "nitrites": 1.13, "mercury": 0.007, "perchlorate": 37.75,
  "radium": 6.78, "selenium": 0.08, "silver": 0.34, "uranium": 0.02
}
```
## 🧠 Metodologi Model
1. Preprocessing: Menangani nilai yang hilang, menghapus baris yang tidak valid (seperti nilai '#NUM!' pada kolom ammonia dan is_safe).

2. Model: Menggunakan beberapa model tetapi model yang di pilih yaitu Decision Tree Classifier.

3. Evaluasi: Model dievaluasi menggunakan metrik akurasi, presisi, dan recall untuk memastikan klasifikasi air aman/tidak aman dilakukan dengan benar.

4. Deployment: Model disimpan menggunakan joblib agar dapat dimuat kembali secara instan tanpa perlu pelatihan ulang.

### 🚀 Perbandingan Performa (Benchmark)

Sebagai bagian dari latihan optimasi sistem, dilakukan pengujian kecepatan *request* (*latency*) untuk membandingkan performa antara *backend* Python (FastAPI) dan *backend* Go (ONNX Runtime). Pengujian dilakukan dengan mengirimkan satu permintaan prediksi (*inference*):

| Runtime | Average Request Latency | Keterangan |
| :--- | :--- | :--- |
| **Python (FastAPI)** | 45 ms | Cocok untuk pengembangan cepat dan riset |
| **Go (ONNX Runtime)** | 20 ms | **2.25x lebih cepat**, ideal untuk sistem produksi |

---

### 📝 Analisis Latihan
* **Efisiensi**: Implementasi menggunakan **Go** dan **ONNX Runtime** menunjukkan pengurangan *latency* yang signifikan (lebih dari 50%) dibandingkan dengan FastAPI.
* **Skalabilitas**: Kecepatan **20 ms** pada Go memberikan keunggulan dalam menangani *concurrency* yang lebih tinggi, yang sangat ideal untuk skenario sistem produksi nyata.
* **Tujuan**: Perbandingan ini dilakukan untuk memahami karakteristik performa dari berbagai *stack* teknologi yang berbeda dalam memproses model *Machine Learning* yang sama.

> **Catatan**: Hasil ini didasarkan pada lingkungan pengujian lokal dan ditujukan murni untuk keperluan latihan dan pembelajaran teknis.
