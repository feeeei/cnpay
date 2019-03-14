package cnpay

import (
	"github.com/objcoding/wxpay"

	"github.com/smartwalle/alipay"
)

type deviceType int

const (
	_ deviceType = iota
	App
	Web
)

type TradePay struct {
	Platform   Platform
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
