package models

type Wallet struct {
	PublicKey string  `json:"public_key"`
	User      User    `json:"user"`
	Balance   float32 `json:"balance"`
	UpdatedAt string  `json:"updated_at"`
}
