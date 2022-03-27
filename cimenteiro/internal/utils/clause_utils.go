package utils

func RemoveBrackets(str string) string {
	if len(str) < 1 {
		return str
	}
	var start, end uint = 0, uint(len(str))
	if str[0] == '(' {
		start = 1
		if str[len(str)-1] == ')' {
			end = uint(len(str) - 1)
		}
	}
	return str[start:end]
}
