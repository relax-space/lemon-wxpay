package wxpay

type ReqBaseDto struct {
	AppId    string `json:"appid,omitempty" query:"appid"`
	SubAppId string `json:"sub_appid,omitempty" query:"sub_appid"`
	MchId    string `json:"mch_id,omitempty" query:"mch_id"`
	SubMchId string `json:"sub_mch_id,omitempty" query:"sub_mch_id"`
	NonceStr string `json:"nonce_str,omitempty" query:"nonce_str"`

	Sign string `json:"sign,omitempty" query:"sign"`
}

type ReqPayDto struct {
	*ReqBaseDto
	DeviceInfo string `json:"device_info,omitempty" query:"device_info"`
	Body       string `json:"body,omitempty" query:"body"`
	Detail     string `json:"detail,omitempty" query:"detail"`
	Attach     string `json:"attach,omitempty" query:"attach"`
	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`

	TotalFee       int64  `json:"total_fee,omitempty" query:"total_fee"` //int64
	FeeType        string `json:"fee_type,omitempty" query:"fee_type"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" query:"spbill_create_ip"`
	GoodsTag       string `json:"goods_tag,omitempty" query:"goods_tag"`
	LimitPay       string `json:"limit_pay,omitempty" query:"limit_pay"`

	TimeStart  string       `json:"time_start,omitempty" query:"time_start"`
	TimeExpire string       `json:"time_expire,omitempty" query:"time_expire"`
	AuthCode   string       `json:"auth_code,omitempty" query:"auth_code"` //10,11,12,13,14,15
	SceneInfo  SceneInfoDto `json:"scene_info,omitempty" query:"scene_info"`
}

type ReqPrepayDto struct {
	*ReqBaseDto
	DeviceInfo string `json:"device_info,omitempty" query:"device_info"`
	SignType   string `json:"sign_type,omitempty" query:"sign_type"`
	Body       string `json:"body,omitempty" query:"body"`
	Detail     string `json:"detail,omitempty" query:"detail"`
	Attach     string `json:"attach,omitempty" query:"attach"`

	OutTradeNo     string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	FeeType        string `json:"fee_type,omitempty" query:"fee_type"`
	TotalFee       int64  `json:"total_fee,omitempty" query:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" query:"spbill_create_ip"`
	TimeStart      string `json:"time_start,omitempty" query:"time_start"`

	TimeExpire string `json:"time_expire,omitempty" query:"time_expire"`
	GoodsTag   string `json:"goods_tag,omitempty" query:"goods_tag"`
	NotifyUrl  string `json:"notify_url,omitempty" query:"notify_url"`
	TradeType  string `json:"trade_type,omitempty" query:"trade_type"`
	ProductId  string `json:"product_id,omitempty" query:"product_id"`

	LimitPay  string       `json:"limit_pay,omitempty" query:"limit_pay"`
	OpenId    string       `json:"openid,omitempty" query:"openid"`
	SubOpenId string       `json:"sub_openid,omitempty" query:"sub_openid"`
	SceneInfo SceneInfoDto `json:"scene_info,omitempty" query:"scene_info"`
}

type ReqQueryDto struct {
	*ReqBaseDto
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id"`
}

type ReqRefundDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	OutRefundNo   string `json:"out_refund_no,omitempty" query:"out_refund_no"`
	TotalFee      int64  `json:"total_fee,omitempty" query:"total_fee"`   //int64
	RefundFee     int64  `json:"refund_fee,omitempty" query:"refund_fee"` //int64

	RefundFeeType string `json:"refund_fee_type,omitempty" query:"refund_fee_type"`
	RefundDesc    string `json:"refund_desc,omitempty" query:"refund_desc"`
	RefundAccount string `json:"refund_account,omitempty" query:"refund_account"`
}
type ReqReverseDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no"`
}
type ReqRefundQueryDto struct {
	*ReqBaseDto
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	OutRefundNo   string `json:"out_refund_no,omitempty" query:"out_refund_no"`
	RefundId      string `json:"refund_id,omitempty" query:"refund_id"`
	Offset        int    `json:"offset,omitempty" query:"offset"`
}

type ReqBillDto struct {
	*ReqBaseDto
	BillDate string `json:"bill_date,omitempty" query:"bill_date"` //20181217
	BillType string `json:"bill_type,omitempty" query:"bill_type"`
	TarType  string `json:"tar_type,omitempty" query:"tar_type"`
}

type ReqCustomerDto struct {
	Key              string `json:"key,omitempty" query:"key"`
	CertPathName     string `json:"cert_path_name,omitempty" query:"cert_path_name"`
	CertPathKey      string `json:"cert_path_key,omitempty" query:"cert_path_key"`
	RootCa           string `json:"root_ca,omitempty" query:"root_ca"`
	UnifiedNotifyUrl string `json:"unified_notify_url,omitempty" query:"unified_notify_url"`
}

//response

type RespDto struct {
	ReturnCode string `json:"return_code,omitempty" query:"return_code"`
	ReturnMsg  string `json:"return_msg,omitempty" query:"return_msg"`
	AppId      string `json:"appid,omitempty" query:"appid"`
	SubAppId   string `json:"sub_appid,omitempty" query:"sub_appid"`
	MchId      string `json:"mch_id,omitempty" query:"mch_id"`

	SubMchId   string `json:"sub_mch_id,omitempty" query:"sub_mch_id"`
	NonceStr   string `json:"nonce_str,omitempty" query:"nonce_str"`
	Sign       string `json:"sign,omitempty" query:"sign"`
	ResultCode string `json:"result_code,omitempty" query:"result_code"`
	ErrCode    string `json:"err_code,omitempty" query:"err_code"`

	ErrCodeDes string `json:"err_code_des,omitempty" query:"err_code_des"`
}

type RespPayDto struct {
	DeviceInfo     string `json:"device_info,omitempty" query:"device_info"`
	OpenId         string `json:"openid,omitempty" query:"openid"`
	IsSubscribe    string `json:"is_subscribe,omitempty" query:"is_subscribe"`
	SubOpenid      string `json:"sub_openid,omitempty" query:"sub_openid"`
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe"`

	TradeType   string `json:"trade_type,omitempty" query:"trade_type"`
	BankType    string `json:"bank_type,omitempty" query:"bank_type"`
	FeeType     string `json:"fee_type,omitempty" query:"fee_type"`
	TotalFee    int64  `json:"total_fee,omitempty" query:"total_fee"` //int64
	CashFeeType string `json:"cash_fee_type,omitempty" query:"cash_fee_type"`

	CashFee            int64  `json:"cash_fee,omitempty" query:"cash_fee"`                         //int64
	SettlementTotalFee int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee"` //int64
	CouponFee          int64  `json:"coupon_fee,omitempty" query:"coupon_fee"`                     //int64
	TransactionId      string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo         string `json:"out_trade_no,omitempty" query:"out_trade_no"`

	Attach  string `json:"attach,omitempty" query:"attach"`
	TimeEnd string `json:"time_end,omitempty" query:"time_end"` //yyyyMMddHHmmss
}

type RespQueryDto struct {
	DeviceInfo     string `json:"device_info,omitempty" query:"device_info"`
	OpenId         string `json:"openid,omitempty" query:"openid"`
	IsSubscribe    string `json:"is_subscribe,omitempty" query:"is_subscribe"`
	SubOpenid      string `json:"sub_openid,omitempty" query:"sub_openid"`
	SubIsSubscribe string `json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe"`

	TradeType  string `json:"trade_type,omitempty" query:"trade_type"`
	TradeState string `json:"trade_state,omitempty" query:"trade_state"`
	BankType   string `json:"bank_type,omitempty" query:"bank_type"`
	Detail     string `json:"detail,omitempty" query:"detail"`       //serviceProvidor
	TotalFee   int64  `json:"total_fee,omitempty" query:"total_fee"` //int64

	FeeType            string `json:"fee_type,omitempty" query:"fee_type"`
	SettlementTotalFee int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee"` //int64
	CashFee            int64  `json:"cash_fee,omitempty" query:"cash_fee"`                         //int64
	CashFeeType        string `json:"cash_fee_type,omitempty" query:"cash_fee_type"`
	CouponFee          int64  `json:"coupon_fee,omitempty" query:"coupon_fee"` //int64

	CouponCount int64 `json:"coupon_count,omitempty" query:"coupon_count"` //int64

	CouponType0 string `json:"coupon_type_0,omitempty" query:"coupon_type_0"`
	CouponId0   string `json:"coupon_id_0,omitempty" query:"coupon_id_0"`
	CouponFee0  int64  `json:"coupon_fee_0,omitempty" query:"coupon_fee_0"` //int64

	CouponType1 string `json:"coupon_type_1,omitempty" query:"coupon_type_1"`
	CouponId1   string `json:"coupon_id_1,omitempty" query:"coupon_id_1"`
	CouponFee1  int64  `json:"coupon_fee_1,omitempty" query:"coupon_fee_1"` //int64

	CouponType2 string `json:"coupon_type_2,omitempty" query:"coupon_type_2"`
	CouponId2   string `json:"coupon_id_2,omitempty" query:"coupon_id_2"`
	CouponFee2  int64  `json:"coupon_fee_2,omitempty" query:"coupon_fee_2"` //int64

	CouponType3 string `json:"coupon_type_3,omitempty" query:"coupon_type_3"`
	CouponId3   string `json:"coupon_id_3,omitempty" query:"coupon_id_3"`
	CouponFee3  int64  `json:"coupon_fee_3,omitempty" query:"coupon_fee_3"` //int64

	CouponType4 string `json:"coupon_type_4,omitempty" query:"coupon_type_4"`
	CouponId4   string `json:"coupon_id_4,omitempty" query:"coupon_id_4"`
	CouponFee4  int64  `json:"coupon_fee_4,omitempty" query:"coupon_fee_4"` //int64

	TransactionId  string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	Attach         string `json:"attach,omitempty" query:"attach"`
	TimeEnd        string `json:"time_end,omitempty" query:"time_end"`
	TradeStateDesc string `json:"trade_state_desc,omitempty" query:"trade_state_desc"`
}

type RespRefundDto struct {
	DeviceInfo    string `json:"device_info,omitempty" query:"device_info"`
	TransactionId string `json:"transaction_id,omitempty" query:"transaction_id"`
	OutRefundNo   string `json:"out_refund_no,omitempty" jsoqueryn:"out_refund_no"`
	RefundId      string `json:"refund_id,omitempty" query:"refund_id"`

	RefundFee            int64  `json:"refund_fee,omitempty" query:"refund_fee"`                       //int64
	SettlementRefund_Fee int64  `json:"settlement_refund_fee,omitempty" query:"settlement_refund_fee"` //int64
	TotalFee             int64  `json:"total_fee,omitempty" query:"total_fee"`                         //int64
	SettlementTotalFee   int64  `json:"settlement_total_fee,omitempty" query:"settlement_total_fee"`   //int64
	FeeType              string `json:"fee_type,omitempty" query:"fee_type"`

	CashFee           int64  `json:"cash_fee,omitempty" query:"cash_fee"`                       //int64
	CashRefundFee     int64  `json:"cash_refund_fee,omitempty" query:"cash_refund_fee"`         //int64
	CouponRefundFee   int64  `json:"coupon_refund_fee,omitempty" query:"coupon_refund_fee"`     //int64
	CouponRefundCount int64  `json:"coupon_refund_count,omitempty" query:"coupon_refund_count"` //int64
	CouponType0       string `json:"coupon_type_0,omitempty" query:"coupon_type_0"`
	CouponRefundId0   string `json:"coupon_refund_id_0,omitempty" query:"coupon_refund_id_0"`
	CouponRefundFee0  int64  `json:"coupon_refund_fee_0,omitempty" query:"coupon_refund_fee_0"` //int64

	CouponType1      string `json:"coupon_type_1,omitempty" query:"coupon_type_1"`
	CouponRefundId1  string `json:"coupon_refund_id_1,omitempty" query:"coupon_refund_id_1"`
	CouponRefundFee1 int64  `json:"coupon_refund_fee_1,omitempty" query:"coupon_refund_fee_1"` //int64

	CouponType2      string `json:"coupon_type_2,omitempty" query:"coupon_type_2"`
	CouponRefundId2  string `json:"coupon_refund_id_2,omitempty" query:"coupon_refund_id_2"`
	CouponRefundFee2 int64  `json:"coupon_refund_fee_2,omitempty" query:"coupon_refund_fee_2"` //int64

	CouponType3      string `json:"coupon_type_3,omitempty" query:"coupon_type_3"`
	CouponRefundId3  string `json:"coupon_refund_id_3,omitempty" query:"coupon_refund_id_3"`
	CouponRefundFee3 int64  `json:"coupon_refund_fee_3,omitempty" query:"coupon_refund_fee_3"` //int64

	CouponType4      string `json:"coupon_type_4,omitempty" query:"coupon_type_4"`
	CouponRefundId4  string `json:"coupon_refund_id_4,omitempty" query:"coupon_refund_id_4"`
	CouponRefundFee4 int64  `json:"coupon_refund_fee_4,omitempty" query:"coupon_refund_fee_4"` //int64

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`
}
type RespReverseDto struct {
	Recall string `json:"recall,omitempty" query:"recall"`
}

type RespPrepayDto struct {
	TradeType string `json:"trade_type,omitempty" query:"trade_type"`
	PrepayId  string `json:"prepay_id,omitempty" query:"prepay_id"`
	CodeUrl   string `json:"code_url,omitempty" query:"code_url"`
}

type RespRefundQueryDto struct {
	TransactionId      string `json:"transaction_id;omitempty" query:"transaction_id"`
	OutTradeNo         string `json:"out_trade_no;omitempty" query:"out_trade_no"`
	TotalFee           int64  `json:"total_fee;omitempty" query:"total_fee"`
	SettlementTotalFee int64  `json:"settlement_total_fee;omitempty" query:"settlement_total_fee"`
	FeeType            string `json:"fee_type;omitempty" query:"fee_type"`

	CashFee          int64 `json:"cash_fee;omitempty" query:"cash_fee"`
	RefundCount      int64 `json:"refund_count;omitempty" query:"refund_count"`
	TotalRefundCount int64 `json:"total_refund_count;omitempty" query:"total_refund_count"`

	OutRefundNo1 string `json:"out_refund_no_1;omitempty" query:"out_refund_no_1"`
	OutRefundNo2 string `json:"out_refund_no_2;omitempty" query:"out_refund_no_2"`
	OutRefundNo3 string `json:"out_refund_no_3;omitempty" query:"out_refund_no_3"`
	OutRefundNo4 string `json:"out_refund_no_4;omitempty" query:"out_refund_no_4"`

	RefundId1 string `json:"refund_id_1;omitempty" query:"refund_id_1"`
	RefundId2 string `json:"refund_id_2;omitempty" query:"refund_id_2"`
	RefundId3 string `json:"refund_id_3;omitempty" query:"refund_id_3"`
	RefundId4 string `json:"refund_id_4;omitempty" query:"refund_id_4"`

	RefundChannel1 string `json:"refund_channel_1;omitempty" query:"refund_channel_1"`
	RefundChannel2 string `json:"refund_channel_2;omitempty" query:"refund_channel_2"`
	RefundChannel3 string `json:"refund_channel_3;omitempty" query:"refund_channel_3"`
	RefundChannel4 string `json:"refund_channel_4;omitempty" query:"refund_channel_4"`

	RefundFee1 int64 `json:"refund_fee_1;omitempty" query:"refund_fee_1"`
	RefundFee2 int64 `json:"refund_fee_2;omitempty" query:"refund_fee_2"`
	RefundFee3 int64 `json:"refund_fee_3;omitempty" query:"refund_fee_3"`
	RefundFee4 int64 `json:"refund_fee_4;omitempty" query:"refund_fee_4"`

	SettlementRefundFee1 int64 `json:"settlement_refund_fee_1;omitempty" query:"settlement_refund_fee_1"`
	SettlementRefundFee2 int64 `json:"settlement_refund_fee_2;omitempty" query:"settlement_refund_fee_2"`
	SettlementRefundFee3 int64 `json:"settlement_refund_fee_3;omitempty" query:"settlement_refund_fee_3"`
	SettlementRefundFee4 int64 `json:"settlement_refund_fee_4;omitempty" query:"settlement_refund_fee_4"`

	CouponType01 string `json:"coupon_type_0_1;omitempty" query:"coupon_type_0_1"`
	CouponType02 string `json:"coupon_type_0_2;omitempty" query:"coupon_type_0_2"`
	CouponType03 string `json:"coupon_type_0_3;omitempty" query:"coupon_type_0_3"`
	CouponType04 string `json:"coupon_type_0_4;omitempty" query:"coupon_type_0_4"`

	CouponRefundFee1 int64 `json:"coupon_refund_fee_1;omitempty" query:"coupon_refund_fee_1"`
	CouponRefundFee2 int64 `json:"coupon_refund_fee_2;omitempty" query:"coupon_refund_fee_2"`
	CouponRefundFee3 int64 `json:"coupon_refund_fee_3;omitempty" query:"coupon_refund_fee_3"`
	CouponRefundFee4 int64 `json:"coupon_refund_fee_4;omitempty" query:"coupon_refund_fee_4"`

	CouponRefundCount1 int64 `json:"coupon_refund_count_1;omitempty" query:"coupon_refund_count_1"`
	CouponRefundCount2 int64 `json:"coupon_refund_count_2;omitempty" query:"coupon_refund_count_2"`
	CouponRefundCount3 int64 `json:"coupon_refund_count_3;omitempty" query:"coupon_refund_count_3"`
	CouponRefundCount4 int64 `json:"coupon_refund_count_4;omitempty" query:"coupon_refund_count_4"`

	CouponRefundId01 string `json:"coupon_refund_id_0_1;omitempty" query:"coupon_refund_id_0_1"`
	CouponRefundId02 string `json:"coupon_refund_id_0_2;omitempty" query:"coupon_refund_id_0_2"`
	CouponRefundId03 string `json:"coupon_refund_id_0_3;omitempty" query:"coupon_refund_id_0_3"`
	CouponRefundId04 string `json:"coupon_refund_id_0_4;omitempty" query:"coupon_refund_id_0_4"`

	CouponRefundFee01 int64 `json:"coupon_refund_fee_0_1;omitempty" query:"coupon_refund_fee_0_1"`
	CouponRefundFee02 int64 `json:"coupon_refund_fee_0_2;omitempty" query:"coupon_refund_fee_0_2"`
	CouponRefundFee03 int64 `json:"coupon_refund_fee_0_3;omitempty" query:"coupon_refund_fee_0_3"`
	CouponRefundFee04 int64 `json:"coupon_refund_fee_0_4;omitempty" query:"coupon_refund_fee_0_4"`

	RefundStatus1 string `json:"refund_status_1;omitempty" query:"refund_status_1"`
	RefundStatus2 string `json:"refund_status_2;omitempty" query:"refund_status_2"`
	RefundStatus3 string `json:"refund_status_3;omitempty" query:"refund_status_3"`
	RefundStatus4 string `json:"refund_status_4;omitempty" query:"refund_status_4"`

	RefundAccount1 string `json:"refund_account_1;omitempty" query:"refund_account_1"`
	RefundAccount2 string `json:"refund_account_2;omitempty" query:"refund_account_2"`
	RefundAccount3 string `json:"refund_account_3;omitempty" query:"refund_account_3"`
	RefundAccount4 string `json:"refund_account_4;omitempty" query:"refund_account_4"`

	RefundRecvAccout1 string `json:"refund_recv_accout_1;omitempty" query:"refund_recv_accout_1"`
	RefundRecvAccout2 string `json:"refund_recv_accout_2;omitempty" query:"refund_recv_accout_2"`
	RefundRecvAccout3 string `json:"refund_recv_accout_3;omitempty" query:"refund_recv_accout_3"`
	RefundRecvAccout4 string `json:"refund_recv_accout_4;omitempty" query:"refund_recv_accout_4"`

	RefundSuccessTime1 string `json:"refund_success_time_1;omitempty" query:"refund_success_time_1"`
	RefundSuccessTime2 string `json:"refund_success_time_2;omitempty" query:"refund_success_time_2"`
	RefundSuccessTime3 string `json:"refund_success_time_3;omitempty" query:"refund_success_time_3"`
	RefundSuccessTime4 string `json:"refund_success_time_4;omitempty" query:"refund_success_time_4"`
}

type NotifyDto struct {
	ReturnCode string `xml:"return_code,omitempty" json:"return_code,omitempty"  query:"return_code"`
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"  query:"return_msg"`
	AppId      string `xml:"appid,omitempty" json:"appid,omitempty" query:"appid"`
	MchId      string `xml:"mch_id,omitempty" json:"mch_id,omitempty" query:"mch_id"`
	SubAppId   string `xml:"sub_appid,omitempty" json:"sub_appid,omitempty" query:"sub_appid"`

	SubMchId   string `xml:"sub_mch_id,omitempty" json:"sub_mch_id,omitempty" query:"sub_mch_id"`
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty" query:"device_info"`
	NonceStr   string `xml:"nonce_str,omitempty" json:"nonce_str,omitempty" query:"nonce_str"`
	Sign       string `xml:"sign,omitempty" json:"sign,omitempty" query:"sign"`
	ResultCode string `xml:"result_code,omitempty" json:"result_code,omitempty" query:"result_code"`

	ErrCode     string `xml:"err_code,omitempty" json:"err_code,omitempty" query:"err_code"`
	ErrCodeDes  string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty" query:"err_code_des"`
	OpenId      string `xml:"openid,omitempty" json:"openid,omitempty" query:"openid"`
	IsSubscribe string `xml:"is_subscribe,omitempty" json:"is_subscribe,omitempty"  query:"is_subscribe"`
	SubOpenId   string `xml:"sub_openid,omitempty" json:"sub_openid,omitempty"  query:"sub_openid"`

	SubIsSubscribe string `xml:"sub_is_subscribe,omitempty" json:"sub_is_subscribe,omitempty" query:"sub_is_subscribe"`
	TradeType      string `xml:"trade_type,omitempty" json:"trade_type,omitempty" query:"trade_type"`
	BankType       string `xml:"bank_type,omitempty" json:"bank_type,omitempty" query:"bank_type"`
	TotalFee       int64  `xml:"total_fee,omitempty" json:"total_fee,omitempty" query:"total_fee"` //int64
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty" query:"fee_type"`

	CashFee            int64  `xml:"cash_fee,omitempty" json:"cash_fee,omitempty" query:"cash_fee"` //int64
	CashFeeType        string `xml:"cash_fee_type,omitempty" json:"cash_fee_type,omitempty" query:"cash_fee_type"`
	SettlementTotalFee int64  `xml:"settlement_total_fee,omitempty" json:"settlement_total_fee,omitempty" query:"settlement_total_fee"` //int64
	CouponFee          int64  `xml:"coupon_fee,omitempty" json:"coupon_fee,omitempty" query:"coupon_fee"`                               //int64
	CouponCount        int64  `xml:"coupon_count,omitempty" json:"coupon_count,omitempty" query:"coupon_count"`                         //int64

	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty" query:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty" query:"out_trade_no"`
	Attach        string `xml:"attach,omitempty" json:"attach,omitempty" query:"attach"`
	TimeEnd       string `xml:"time_end,omitempty" json:"time_end,omitempty" query:"time_end"`
}

type SceneInfoDto struct {
	Id       string `json:"id,omitempty" query:"id"`
	Name     string `json:"name,omitempty" query:"name"`
	AreaCode string `json:"area_code,omitempty" query:"area_code"`
	Address  string `json:"address,omitempty" query:"address"`
}
