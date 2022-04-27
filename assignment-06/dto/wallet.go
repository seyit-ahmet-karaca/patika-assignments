package dto

type Wallet struct {
	Username string `json:"username"`
	Balance int `json:"balance"`
}

type Wallets []*Wallet