package handler

import (
	"log" // Tambahkan ini untuk mencetak error ke terminal
	"model-ml/dto"
	"model-ml/ml"

	"github.com/gin-gonic/gin"
)

type WaterHandler struct {
	Engine *ml.MLEngine
}

func (h *WaterHandler) Predict(c *gin.Context) {
	var req dto.WaterQualityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah", "detail": err.Error()})
		return
	}

	result, err := h.Engine.Predict(req.ToSlice())
	if err != nil {
		// 1. Tampilkan error asli ONNX di terminal server
		log.Printf("🔥 ERROR DARI ONNX RUNTIME: %v\n", err)

		// 2. Tampilkan error asli di response (Postman / cURL)
		c.JSON(500, gin.H{
			"error":  "Gagal prediksi",
			"detail": err.Error(), // Ini kunci untuk tahu penyakitnya!
		})
		return
	}

	c.JSON(200, gin.H{"is_safe": result,
		"message": func() string {
			if result == 1 {
				return "Air aman untuk dikonsumsi"
			}
			return "Air tidak aman untuk dikonsumsi"
		}(),
	})
}
