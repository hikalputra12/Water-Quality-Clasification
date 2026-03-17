package main

import (
	"log"
	"model-ml/handler"
	"model-ml/ml"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Engine
	engine, err := ml.NewMLEngine(
		"/home/haikal/Water-Quality-Clasification/deploy-go/model/model_anda.onnx",
		"/home/haikal/Water-Quality-Clasification/deploy-go/lib/libonnxruntime.so",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	//Setup Handler
	h := &handler.WaterHandler{Engine: engine}

	//Setup Router
	r := gin.Default()
	r.POST("/predict", h.Predict)

	r.Run(":8080")
}
