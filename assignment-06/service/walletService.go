package service

import "karaca/dto"

type IWalletService interface {
	Create(string) error
	Update(string, int) error
	GetAll() dto.Wallets
	GetWalletByUsername(string) (*dto.Wallet, error)
}
