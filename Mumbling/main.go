package kata

var last string

func Accum(s string) string {
	last = ""
	strChar := []byte(s)
	for i, v := range strChar {
		if int(v) < 97 {
			strChar[i] = byte(v + 32)
		}
	}
	for i := 0; i < len(strChar); i++ {
		last = last + string(strChar[i]-32)
		for j := 0; j < i; j++ {
			last = last + string(strChar[i])
		}
		if i+1 != len(strChar) {
			last = last + "-"
		}
	}

	return last

}
