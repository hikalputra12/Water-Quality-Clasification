package dto

type WaterQualityRequest struct {
	Aluminium   float32 `json:"aluminium"`
	Ammonia     float32 `json:"ammonia"`
	Arsenic     float32 `json:"arsenic"`
	Barium      float32 `json:"barium"`
	Cadmium     float32 `json:"cadmium"`
	Chloramine  float32 `json:"chloramine"`
	Chromium    float32 `json:"chromium"`
	Copper      float32 `json:"copper"`
	Flouride    float32 `json:"flouride"`
	Bacteria    float32 `json:"bacteria"`
	Viruses     float32 `json:"viruses"`
	Lead        float32 `json:"lead"`
	Nitrates    float32 `json:"nitrates"`
	Nitrites    float32 `json:"nitrites"`
	Mercury     float32 `json:"mercury"`
	Perchlorate float32 `json:"perchlorate"`
	Radium      float32 `json:"radium"`
	Selenium    float32 `json:"selenium"`
	Silver      float32 `json:"silver"`
	Uranium     float32 `json:"uranium"`
}

// ToSlice mengubah struct menjadi slice float32 dengan urutan yang benar untuk ONNX
func (w *WaterQualityRequest) ToSlice() []float32 {
	return []float32{
		w.Aluminium, w.Ammonia, w.Arsenic, w.Barium, w.Cadmium,
		w.Chloramine, w.Chromium, w.Copper, w.Flouride, w.Bacteria,
		w.Viruses, w.Lead, w.Nitrates, w.Nitrites, w.Mercury,
		w.Perchlorate, w.Radium, w.Selenium, w.Silver, w.Uranium,
	}
}
