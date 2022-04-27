package repository

import "karaca/dto"

type IWalletRepository interface {
	Create(*dto.Wallet) error
	Update(*dto.Wallet) error
	GetAll() dto.Wallets
	GetWalletByUsername(string) (*dto.Wallet, error)
}
