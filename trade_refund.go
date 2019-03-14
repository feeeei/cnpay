package cnpay

import (
	"github.com/objcoding/wxpay"
	"github.com/smartwalle/alipay"
)

type TradeRefund struct {
	Platform     Platform
	TradeNo      string  //商户订单号
	RefundNo     string  //商户退单号（aliapy可不传）
	TotalAmount  float64 //订单总金额（alipay可不传）
	RefundAmount float64 //退款金额
	RefundReason string  //退款原因（选填）
}

func (refund *TradeRefund) ToWxpay() wxpay.Params {
	return wxpay.Params{
		"out_trade_no":  refund.TradeNo,
		"out_refund_no": refund.RefundNo,
		"total_fee":     wxpayAmount(refund.TotalAmount),
		"refund_fee":    wxpayAmount(refund.RefundAmount),
		"refund_desc":   refund.RefundReason,
	}
}

func (refund *TradeRefund) ToAlipay() *alipay.AliPayTradeRefund {
	return &alipay.AliPayTradeRefund{
		TradeNo:      refund.TradeNo,
		OutRequestNo: refund.RefundNo,
		RefundAmount: alipayAmount(refund.RefundAmount),
		RefundReason: refund.RefundReason,
	}
}
