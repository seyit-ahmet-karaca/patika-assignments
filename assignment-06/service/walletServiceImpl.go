package service

import (
	"karaca/config"
	"karaca/dto"
	error2 "karaca/error"
	"karaca/repository"
)

type WalletService struct {
	wr repository.IWalletRepository
}

func (as *WalletService) Create(username string) error {
	if _, err := as.wr.GetWalletByUsername(username); err != nil {
		err := as.wr.Create(&dto.Wallet{Username: username, Balance: config.Get().InitialBalanceAmount})
		if err != nil {
			return err
		}
	}
	return nil
}

func (as *WalletService) Update(username string, balanceWithUpdate int) error {
	wallet, err := as.wr.GetWalletByUsername(username)
	if err != nil {
		return err
	}

	if newBalance := wallet.Balance + balanceWithUpdate; newBalance < config.Get().MinimumBalanceAmount {
		return error2.GetInvalidBalance()
	} else {
		wallet.Balance = newBalance
		return as.wr.Update(wallet)
	}
}

func (as *WalletService) GetAll() dto.Wallets {
	return as.wr.GetAll()
}

func (as *WalletService) GetWalletByUsername(username string) (*dto.Wallet, error) {
	wallet, err := as.wr.GetWalletByUsername(username)

	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func NewWalletService(repo repository.IWalletRepository) IWalletService {
	return &WalletService{wr: repo}
}