package dtos

import (
	"time"
	"yab-explorer/domain/models"
)

type OrderDto struct {
	OrderId          int        `json:"order_id"`
	OriginNetwork    string     `json:"origin_network"`
	RecipientAddress string     `json:"recipient_address"`
	FromAddress      string     `json:"from_address"`
	Amount           string     `json:"amount"`
	Fee              string     `json:"fee"`
	Status           string     `json:"status"`
	Failed           bool       `json:"failed"`
	SetOrderTxHash   []byte     `json:"set_order_tx_hash"`
	TransferTxHash   []byte     `json:"transfer_tx_hash"`
	ClaimTxHash      []byte     `json:"claim_tx_hash"`
	HerodotusTaskId  string     `json:"herodotus_task_id"`
	HerodotusBlock   uint64     `json:"herodotus_block"`
	HerodotusSlot    []byte     `json:"herodotus_slot"`
	CreatedAt        *time.Time `json:"created_at"`
	TransferredAt    *time.Time `json:"transferred_at"`
	CompletedAt      *time.Time `json:"completed_at"`
}

func OrderToDto(order models.Order) OrderDto {
	return OrderDto{
		OrderId:          order.OrderId,
		OriginNetwork:    order.OriginNetwork,
		RecipientAddress: order.RecipientAddress,
		FromAddress:      order.FromAddress,
		Amount:           order.Amount,
		Fee:              order.Fee,
		Status:           order.Status,
		Failed:           order.Failed,
		SetOrderTxHash:   order.SetOrderTxHash,
		TransferTxHash:   order.TransferTxHash,
		ClaimTxHash:      order.ClaimTxHash,
		HerodotusTaskId:  order.HerodotusTaskId,
		HerodotusBlock:   order.HerodotusBlock,
		HerodotusSlot:    order.HerodotusSlot,
		CreatedAt:        order.CreatedAt,
		TransferredAt:    order.TransferredAt,
		CompletedAt:      order.CompletedAt,
	}
}
