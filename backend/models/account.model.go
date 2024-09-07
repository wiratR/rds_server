package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AccountToken   string     `gorm:"type:varchar(100);not null"`
	AccountType    string     `gorm:"type:varchar(50);not null"`
	Status         string     `gorm:"type:varchar(50);not null"`
	Balance        int        `gorm:"type:int;not null"`
	BlockFlag      int        `gorm:"type:int;not null"`
	LastEntrySpId  int        `gorm:"type:int;not null"`
	LastEntryLocId int        `gorm:"type:int;not null"`
	LastEntryTime  *time.Time `gorm:"type:timestamp"`
	Active         bool       `gorm:"not null;default:false"`
	// Foreign key, unique constraint for one-to-one relationship
	UserID *uuid.UUID `gorm:"type:uuid;unique;not null"`
	// Association
	User         *User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime"`
	TxnHistories []TxnHistory `gorm:"foreignKey:AccountID"` // One-to-many relationship with TxnHistory
}

type AccoutCreateInput struct {
	AccountType string `json:"account_type" validate:"required"`
	UserID      string `json:"user_id" validate:"required"`
}

type AccountResponse struct {
	ID             uuid.UUID `json:"id,omitempty"`
	AccountToken   string    `json:"account_token,omitempty"`
	AccountType    string    `json:"account_type,omitempty"`
	Status         string    `json:"status,omitempty"`
	Balance        int       `json:"balance,omitempty"`
	BlockFlag      int       `json:"block_flag,omitempty"`
	LastEntrySpId  int       `json:"last_entry_sp_id,omitempty"`
	LastEntryLocId int       `json:"last_entry_loc_id,omitempty"`
	LastEntryTime  time.Time `json:"last_entry_time"`
	Active         bool      `json:"active,omitempty"`
	// UserID         uuid.UUID    `json:"user_id"`
	// User               UserResponse         `json:"user"`
	TxnHistoriesDetail []TxnHistoryResponse `json:"txn_histories_detail"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
}

func FilterAccountRecord(account *Account, txnHistoriesDetail []TxnHistoryResponse) AccountResponse {
	var lastEntryTime time.Time
	if account.LastEntryTime != nil {
		lastEntryTime = *account.LastEntryTime
	}

	return AccountResponse{
		ID:                 *account.ID,
		AccountToken:       account.AccountToken,
		AccountType:        account.AccountType,
		Status:             account.Status,
		Balance:            account.Balance,
		BlockFlag:          account.BlockFlag,
		LastEntrySpId:      account.LastEntrySpId,
		LastEntryLocId:     account.LastEntryLocId,
		LastEntryTime:      lastEntryTime,
		Active:             account.Active,
		CreatedAt:          account.CreatedAt,
		UpdatedAt:          account.UpdatedAt,
		TxnHistoriesDetail: txnHistoriesDetail,
	}
}
