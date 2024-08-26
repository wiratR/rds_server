package models

import (
	"time"

	"github.com/google/uuid"
)

type TxnHistory struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	TxnRefId        string    `gorm:"type:varchar(255);not null" validate:"required"`
	TxnTypeId       int       `gorm:"not null" validate:"required"`
	TxnAmount       int       `gorm:"not null" validate:"required"`
	SpId            int       `gorm:"not null" validate:"required"`
	LocEntryId      int       `gorm:"not null" validate:"required"`
	LocExitId       int       `gorm:"not null" validate:"required"`
	EquipmentNumber string    `gorm:"type:varchar(255);not null" validate:"required"`
	AccountID       uuid.UUID `gorm:"type:uuid;not null"`                             // Foreign key to Account
	Account         Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Association
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

type TxnCreateInput struct {
	TxnRefId        string `json:"txn_ref_id,omitempty"`
	TxnTypeId       int    `json:"txn_type_id,omitempty"`
	TxnAmount       int    `json:"txn_amount,omitempty"`
	SpId            int    `json:"sp_id,omitempty"`
	LocEntryId      int    `json:"loc_entry_id,omitempty"`
	LocExitId       int    `json:"loc_exit_id,omitempty"`
	EquipmentNumber string `json:"equipment_number,omitempty"`
}

type TxnHistoryResponse struct {
	TxnRefId        string    `json:"txn_ref_id,omitempty"`
	TxnAmount       int       `json:"txn_amount,omitempty"`
	TxnDetail       TxnDetail `json:"txn_detail,omitempty"`
	SpDetail        SpDetail  `json:"sp_detail,omitempty"`
	LocEntryDetail  LocDetail `json:"loc_entry_detail,omitempty"`
	LocExitDetail   LocDetail `json:"loc_exit_detail,omitempty"`
	EquipmentNumber string    `json:"equipment_number,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TxnDetail struct {
	TxnTypeId   int    `json:"txn_type_id,omitempty"`
	TxnTypeName string `json:"txn_type_name,omitempty"`
}

type SpDetail struct {
	SpId   int    `json:"sp_id,omitempty"`
	SpName string `json:"sp_name,omitempty"`
}

type LineDetail struct {
	LineId   int    `json:"line_id,omitempty"`
	LineName string `json:"line_name,omitempty"`
}

type LocDetail struct {
	LocId      int        `json:"loc_id,omitempty"`
	LocName    string     `json:"loc_name,omitempty"`
	LineDetail LineDetail `json:"line_detail,omitempty"`
}
