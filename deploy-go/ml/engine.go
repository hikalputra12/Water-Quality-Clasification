package ml

import (
	ort "github.com/yalue/onnxruntime_go"
)

type MLEngine struct {
	session *ort.AdvancedSession
	input   *ort.Tensor[float32]
	output  *ort.Tensor[int64]
}

func NewMLEngine(modelPath, libPath string) (*MLEngine, error) {
	ort.SetSharedLibraryPath(libPath)

	if err := ort.InitializeEnvironment(); err != nil {
		return nil, err
	}

	inputShape := ort.NewShape(1, 20)
	inputTensor, _ := ort.NewEmptyTensor[float32](inputShape)

	outputShape := ort.NewShape(1) // 1 baris, 1 label
	outputTensor, _ := ort.NewEmptyTensor[int64](outputShape)

	// INI YANG DIUBAH: "label" menjadi "output_label"
	sess, err := ort.NewAdvancedSession(modelPath,
		[]string{"float_input"},
		[]string{"output_label"}, // <--- PERBAIKANNYA DI SINI
		[]ort.ArbitraryTensor{inputTensor},
		[]ort.ArbitraryTensor{outputTensor}, nil)

	if err != nil {
		return nil, err
	}

	return &MLEngine{session: sess, input: inputTensor, output: outputTensor}, nil
}

func (e *MLEngine) Predict(data []float32) (int64, error) {
	copy(e.input.GetData(), data)
	if err := e.session.Run(); err != nil {
		return -1, err
	}
	return e.output.GetData()[0], nil
}

func (e *MLEngine) Close() {
	e.session.Destroy()
	e.input.Destroy()
	e.output.Destroy()
	ort.DestroyEnvironment()
}
