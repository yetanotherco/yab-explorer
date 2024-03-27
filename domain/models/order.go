package models

import "time"

type Order struct {
	OrderId          int        `gorm:"column:order_id;primaryKey"`
	OriginNetwork    string     `gorm:"column:origin_network;not null;primaryKey"`
	RecipientAddress string     `gorm:"column:recipient_address;not null;size:42"`
	FromAddress      string     `gorm:"column:from_address;not null;size:42"`
	Amount           string     `gorm:"column:amount;not null"`
	Fee              string     `gorm:"column:fee;not null"`
	Status           string     `gorm:"column:status;not null;default:PENDING"`
	Failed           bool       `gorm:"column:failed;not null;default:false"`
	SetOrderTxHash   []byte     `gorm:"column:set_order_tx_hash;not null"`
	TransferTxHash   []byte     `gorm:"column:transfer_tx_hash"`
	ClaimTxHash      []byte     `gorm:"column:claim_tx_hash"`
	HerodotusTaskId  string     `gorm:"column:herodotus_task_id"`
	HerodotusBlock   uint64     `gorm:"column:herodotus_block"`
	HerodotusSlot    []byte     `gorm:"column:herodotus_slot"`
	CreatedAt        *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	TransferredAt    *time.Time `gorm:"column:transferred_at"`
	CompletedAt      *time.Time `gorm:"column:completed_at"`
}
