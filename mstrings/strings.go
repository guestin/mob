package mstrings

import "strings"

func IsBlank(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}

func StrCmp() {

}
