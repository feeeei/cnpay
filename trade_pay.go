package cnpay

import (
	"fmt"

	"github.com/objcoding/wxpay"
	"github.com/shopspring/decimal"

	"github.com/smartwalle/alipay"
)

type platform int

const (
	_ platform = iota
	Wxpay
	Alipay
)

type deviceType int

const (
	_ deviceType = iota
	App
	Web
)

type TradePay struct {
	Platform   platform
	DeviceType deviceType
	NotifyURL  string
	Subject    string
	TradeNo    string
	Amount     float64
}

func (pay *TradePay) ToAlipay() *alipay.TradePay {
	return &alipay.TradePay{
		NotifyURL:   pay.NotifyURL,
		Subject:     pay.Subject,
		OutTradeNo:  pay.TradeNo,
		TotalAmount: alipayAmount(pay.Amount),
		ProductCode: "QUICK_WAP_WAY",
	}
}

func (pay *TradePay) ToWxpay() wxpay.Params {
	return wxpay.Params{
		"notify_url":   pay.NotifyURL,
		"trade_type":   pay.deviceType(),
		"total_fee":    wxpayAmount(pay.Amount),
		"out_trade_no": pay.TradeNo,
		"body":         pay.Subject,
	}
}

func alipayAmount(amount float64) string {
	return fmt.Sprint(amount)
}

func wxpayAmount(amount float64) string {
	aDecimal := decimal.NewFromFloat(amount)
	bDecimal := decimal.NewFromFloat(100)
	return aDecimal.Mul(bDecimal).Truncate(0).String()
}

func (pay *TradePay) deviceType() string {
	switch pay.DeviceType {
	case App:
		return "APP"
	case Web:
		return "WEB"
	default:
		return ""
	}
}
