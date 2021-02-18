package utils

func CheckBitSetVar(mybool bool) int8 {
	if mybool {
		return 1
	}
	return 0 //you just saved youself an else here!
}
