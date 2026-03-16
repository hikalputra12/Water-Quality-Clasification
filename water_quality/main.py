from fastapi import FastAPI, HTTPException
from typing import Dict, Any
import pandas as pd
import joblib
import uvicorn

# Inisialisasi Aplikasi FastAPI
app = FastAPI(
    title="Water Quality Prediction API", 
    description="API Inference untuk mendeteksi apakah air aman dikonsumsi berdasarkan kandungan kimianya.",
    version="1.0"
)

# Muat Artifacts (Dijalankan sekali saat server menyala)
try:
    # Taruh file .joblib ini di folder yang sama dengan file main.py ini
    model = joblib.load('/home/haikal/Water-Quality-Clasification/water_quality/model/model_dt_final.joblib') 
    print("✅ Model Kualitas Air berhasil dimuat ke memori!")
except Exception as e:
    print(f"❌ Gagal memuat file joblib: {e}")

# Endpoint Prediksi
@app.post("/predict")
def predict_water_quality(data: Dict[str, Any]):
    """
    Endpoint ini menerima JSON berisi fitur-fitur kimia air (seperti aluminium, ammonia, dll),
    dan mengembalikan prediksi status air (Aman/Tidak Aman).
    """
    try:
        # Konversi: JSON (Dictionary) -> Pandas DataFrame (1 Baris)
        df_input = pd.DataFrame([data])
        
        # Prediksi: Model menebak kelas (biasanya 0 = Tidak Aman, 1 = Aman)
        prediction = model.predict(df_input)
        
        # Menerjemahkan hasil tebakan angka menjadi teks agar mudah dipahami Front-End/User
        hasil_angka = int(prediction[0])
        status_air = "Aman" if hasil_angka == 1 else "Tidak Aman"
        
        # Respon API
        return {
            "status": "success",
            "prediction_code": hasil_angka,
            "description": status_air
        }
        
    except Exception as e:
        raise HTTPException(status_code=400, detail=f"Terjadi kesalahan pemrosesan data: {str(e)}")

# Menjalankan file ini langsung dengan `python main.py`
if __name__ == "__main__":
    uvicorn.run("main:app", host="localhost", port=8000, reload=True)