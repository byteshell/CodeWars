package kata

func CorrectTail(body string, tail rune) bool {
	if rune(body[len(body) - 1]) == tail {
		return true
	}
  
	return false
}
