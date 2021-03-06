package cnpay

import (
	"fmt"
	"net/url"

	"github.com/objcoding/wxpay"
	"github.com/smartwalle/alipay"
)

func WxpayCallback(body *[]byte) (map[string]string, error) {
	params := wxpay.XmlToMap(string(*body))
	if !wxClient.ValidSign(params) {
		return nil, fmt.Errorf("签名失败")
	}
	result := map[string]string(params)
	if params["return_code"] != "SUCCESS" {
		return result, fmt.Errorf("支付失败")
	}
	return result, nil
}

func AlipayCallback(form url.Values) (*alipay.TradeNotification, error) {
	return getAlipayTradeNotification(form)
}

func Notification(platform Platform, isSuccess bool) string {
	switch platform {
	case Wxpay:
		noti := wxpay.Notifies{}
		if !isSuccess {
			return noti.NotOK("faild")
		}
		return noti.OK()
	case Alipay:
		if !isSuccess {
			return "faild"
		}
		return "success"
	default:
		return ""
	}
}

func getAlipayTradeNotification(req url.Values) (*alipay.TradeNotification, error) {
	if ok, _ := alipayClient.VerifySign(req); !ok {
		return nil, fmt.Errorf("签名失败")
	}
	noti := &alipay.TradeNotification{}
	noti.AppId = req.Get("app_id")
	noti.AuthAppId = req.Get("auth_app_id")
	noti.NotifyId = req.Get("notify_id")
	noti.NotifyType = req.Get("notify_type")
	noti.NotifyTime = req.Get("notify_time")
	noti.TradeNo = req.Get("trade_no")
	noti.TradeStatus = req.Get("trade_status")
	noti.TotalAmount = req.Get("total_amount")
	noti.ReceiptAmount = req.Get("receipt_amount")
	noti.InvoiceAmount = req.Get("invoice_amount")
	noti.BuyerPayAmount = req.Get("buyer_pay_amount")
	noti.SellerId = req.Get("seller_id")
	noti.SellerEmail = req.Get("seller_email")
	noti.BuyerId = req.Get("buyer_id")
	noti.BuyerLogonId = req.Get("buyer_logon_id")
	noti.FundBillList = req.Get("fund_bill_list")
	noti.Charset = req.Get("charset")
	noti.PointAmount = req.Get("point_amount")
	noti.OutTradeNo = req.Get("out_trade_no")
	noti.OutBizNo = req.Get("out_biz_no")
	noti.GmtCreate = req.Get("gmt_create")
	noti.GmtPayment = req.Get("gmt_payment")
	noti.GmtRefund = req.Get("gmt_refund")
	noti.GmtClose = req.Get("gmt_close")
	noti.Subject = req.Get("subject")
	noti.Body = req.Get("body")
	noti.RefundFee = req.Get("refund_fee")
	noti.Version = req.Get("version")
	noti.SignType = req.Get("sign_type")
	noti.Sign = req.Get("sign")
	noti.PassbackParams = req.Get("passback_params")
	noti.VoucherDetailList = req.Get("voucher_detail_list")

	return noti, nil
}
