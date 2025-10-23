package iteracao

func Repetir(target string, times int) string {
	var result string

	for i:= 0; i < times; i++ {
		result += target
	}
	
	return result
}