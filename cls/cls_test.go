package cls

import (
	"log"
	"testing"
)

func TestCall(t *testing.T) {
	LoadENV()
	_str := `มีเงินเข้า1,000.00บ.จาก KTB x7336 MR. CHAIYASART 01/01/20@01:00`
	log.Println(Call("TTB", _str))
	// log.Println(LoadRegexBank())

}
