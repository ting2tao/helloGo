package strings

import strings2 "strings"

var stringBuffer strings2.Builder

func Append(strings ...string) string {
	for _, str := range strings {
		stringBuffer.WriteString(str)
	}
	return stringBuffer.String()
}

//
func Implode(strArr []string) string {
	for _, str := range strArr {
		stringBuffer.WriteString(str)
	}
	return stringBuffer.String()
}

func Explode(str string, sep string) []string {

	return strings2.Split(str, sep)
}
