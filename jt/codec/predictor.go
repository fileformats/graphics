package codec

type PredictorType byte

const (
	Lag1 PredictorType = iota
	Lag2
	Stride1
	Stride2
	StripIndex
	Ramp
	Xor1
	Xor2
	NULL
)

func UnpackResidual(residuals []int32, predictor PredictorType) {
	var predicted int32
	var residualsLength = len(residuals)
	for i := 0; i < residualsLength; i++ {
		if i < 4 {
			residuals[i] = residuals[i]
		} else {
			v1 := residuals[i-1]
			v2 := residuals[i-2]
			//v3 := residuals[i - 3];
			v4 := residuals[i-4]
			predicted = 0
			switch predictor {
			case Lag1:
				predicted = v1
			case Xor1:
				predicted = v1
			case Lag2:
				predicted = v2
			case Xor2:
				predicted = v2
			case Stride1:
				predicted = v1 + (v1 - v2)
			case Stride2:
				predicted = v2 + (v2 - v4)
			case StripIndex:
				if v2-v4 < 8 && v2-v4 > -8 {
					predicted = v2 + (v2 - v4)
				} else {
					predicted = v2 + 2
				}
			case Ramp:
				predicted = int32(i)
			}

			if predictor == Xor1 || predictor == Xor2 {
				residuals[i] = residuals[i] ^ predicted
			} else {
				residuals[i] = residuals[i] + predicted
			}
		}
	}
}
