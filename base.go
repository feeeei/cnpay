package cnpay

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Platform int

const (
	_ Platform = iota
	Wxpay
	Alipay
)

func alipayAmount(amount float64) string {
	return fmt.Sprint(amount)
}

func wxpayAmount(amount float64) string {
	aDecimal := decimal.NewFromFloat(amount)
	bDecimal := decimal.NewFromFloat(100)
	return aDecimal.Mul(bDecimal).Truncate(0).String()
}
