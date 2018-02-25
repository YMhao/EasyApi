package swagger

func Summary(str string) string {
	r := []rune(str)
	len := len(r)
	if len > 20 {
		r = r[:20]
	}
	return string(r)
}
