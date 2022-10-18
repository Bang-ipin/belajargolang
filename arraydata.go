package arraydata

import (
	"fmt"
	"strings"
)

func ArrayData() {
	x := []string{"SATU", "DUA", "TIGA", "EMPAT"}
	fmt.Println(strings.Trim(fmt.Sprint(x), "[]"))
}
