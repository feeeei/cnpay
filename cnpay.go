package cnpay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/objcoding/wxpay"
	"github.com/smartwalle/alipay"
)

var wxClient *wxpay.Client

var alipayClient *alipay.AliPay

func InitWxpay(appID, mchID, apiKey, certPath string, debug bool) {
	account := wxpay.NewAccount(appID, mchID, apiKey, debug)
	account.SetCertData(certPath)
	wxClient = wxpay.NewClient(account)
}

func InitAlipay(appID, aliPublicKey, privateKey string, debug bool) {
	alipayClient = alipay.New(appID, aliPublicKey, privateKey, debug)
}

func Pay(tradePay *TradePay) (map[string]string, error) {
	switch tradePay.Platform {
	case Alipay:
		return AlipayPay(tradePay)
	case Wxpay:
		return WxpayPay(tradePay)
	default:
		return nil, fmt.Errorf("Unknow platform")
	}
}

func Refund(refund *TradeRefund) (interface{}, error) {
	switch refund.Platform {
	case Alipay:
		return alipayClient.TradeRefund(*refund.ToAlipay())
	case Wxpay:
		resp, err := wxClient.Refund(refund.ToWxpay())
		if err != nil {
			return nil, err
		}
		result := map[string]string(resp)
		return result, nil
	default:
		return nil, fmt.Errorf("Unknow platform")
	}
}

func AlipayPay(tradePay *TradePay) (map[string]string, error) {
	ali := tradePay.ToAlipay()
	resp := make(map[string]string)
	switch tradePay.DeviceType {
	case App:
		pay := alipay.AliPayTradeAppPay{TradePay: *ali}
		if results, err := alipayClient.TradeAppPay(pay); err == nil {
			resp["orderString"] = results
			return resp, nil
		} else {
			return nil, err
		}
	case Web:
		pay := alipay.AliPayTradePagePay{TradePay: *ali}
		if url, err := alipayClient.TradePagePay(pay); err == nil {
			resp["orderString"] = url.String()
			return resp, nil
		} else {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknow device type")
	}
}

func WxpayPay(tradePay *TradePay) (map[string]string, error) {
	params, err := wxClient.UnifiedOrder(tradePay.ToWxpay())
	if err == nil {
		return translateWxpayAppResult(tradePay, params), nil
	}
	return nil, err
}

func translateWxpayAppResult(tradePay *TradePay, params wxpay.Params) map[string]string {
	if tradePay.DeviceType == App {
		p := make(wxpay.Params)
		p["appid"] = params["appid"]
		p["partnerid"] = params["mch_id"]
		p["noncestr"] = params["nonce_str"]
		p["prepayid"] = params["prepay_id"]
		p["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
		p["package"] = "Sign=WXPay"
		p["sign"] = wxClient.Sign(p)
		return map[string]string(p)
	} else {
		return map[string]string(params)
	}
}
