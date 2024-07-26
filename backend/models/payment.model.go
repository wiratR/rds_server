package models

import (
	"time"

	"github.com/google/uuid"
)

type RequestPaymentBody struct {
	FeeType         string `json:"fee_type"`
	ProductName     string `json:"product_name"`
	ReferURL        string `json:"refer_url"`
	NotifyURL       string `json:"notify_url"`
	MchOrderNo      string `json:"mch_order_no"`
	LocalTotalFee   int    `json:"local_total_fee"`
	Channel         string `json:"channel"`
	ChannelSubAppid string `json:"channel_sub_appid"`
	Attach          string `json:"attach"`
	Appid           string `json:"appid"`
	NonceStr        string `json:"nonce_str"`
	TimeStamp       string `json:"time_stamp"`
	Sign            string `json:"sign"`
}

type PaymentOrderCreate struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FeeType   string     `json:"fee_type"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
