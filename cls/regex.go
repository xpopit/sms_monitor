package cls

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

// type Ban
type Bank struct {
	Amount   string
	Bank     string
	Name     string
	Datetime time.Time
}

func Call(BankName, SMS string) (b *Bank) {
	//
	// log.Println()
	_b := LoadRegexBank(BankName)
	// สร้าง regular expression
	fmt.Println(_b)
	re := regexp2.MustCompile(_b.Regex, regexp2.None)

	// ค้นหาข้อความที่ตรงกับ regular expression
	match, _ := re.FindStringMatch(SMS)
	if match != nil {
		// log.Println("nil")
		amount := match.GroupByNumber(1).String()
		bank := match.GroupByNumber(2).String()
		name := match.GroupByNumber(3).String()
		// จัดการกับจำนวนเงิน
		amount = strings.Replace(amount, "บ", "", -1)
		amount = strings.Replace(amount, ",", "", -1)
		amountFloat, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("Error parsing amount:", err)
			return
		}

		// แปลงและจัดรูปแบบจำนวนเงิน
		formattedAmount := fmt.Sprintf("%.2f", amountFloat)

		layout := "02/01/06@15:04"

		// แปลงข้อความเป็น time.Time
		datetime, err := time.Parse(layout, match.GroupByNumber(4).String())
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		// datetime :=
		b = &Bank{
			Amount:   formattedAmount,
			Bank:     bank,
			Name:     name,
			Datetime: datetime,
		}
		// แสดงผลข้อมูลที่ดึงออกมา
		// fmt.Println("Amount:", formattedAmount)
		// fmt.Println("Bank:", bank)
		// fmt.Println("Name:", name)
		// fmt.Println("Date and Time:", datetime)
	} else {
		fmt.Println("No match found")
	}
	return
}
