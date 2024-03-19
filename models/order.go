package models

import "time"

type Order struct {
	OrderID          int        `gorm:"column:order_id;primaryKey" json:"order_id"`
	OriginNetwork    string     `gorm:"column:origin_network;not null;primaryKey" json:"origin_network"`
	RecipientAddress string     `gorm:"column:recipient_address;not null;size:42" json:"recipient_address"`
	FromAddress      string     `gorm:"column:recipient_address;not null;size:42" json:"from_address"`
	Amount           string     `gorm:"column:amount;not null" json:"amount"`
	Fee              string     `gorm:"column:fee;not null" json:"fee"`
	Status           string     `gorm:"column:status;not null;default:PENDING" json:"status"`
	Failed           bool       `gorm:"column:failed;not null;default:false" json:"failed"`
	SetOrderTxHash   []byte     `gorm:"column:set_order_tx_hash;not null" json:"set_order_tx_hash"`
	TransferTxHash   []byte     `gorm:"column:transfer_tx_hash" json:"transfer_tx_hash"`
	ClaimTxHash      []byte     `gorm:"column:claim_tx_hash" json:"claim_tx_hash"`
	HerodotusTaskID  string     `gorm:"column:herodotus_task_id" json:"herodotus_task_id"`
	HerodotusBlock   uint64     `gorm:"column:herodotus_block" json:"herodotus_block"`
	HerodotusSlot    []byte     `gorm:"column:herodotus_slot" json:"herodotus_slot"`
	CreatedAt        *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	TransferredAt    *time.Time `gorm:"column:transferred_at" json:"transferred_at"`
	CompletedAt      *time.Time `gorm:"column:completed_at" json:"completed_at"`
}
