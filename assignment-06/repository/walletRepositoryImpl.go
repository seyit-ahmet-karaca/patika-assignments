package repository

import (
	"karaca/data"
	"karaca/dto"
	internalError "karaca/error"
)

type WalletRepository struct {
	d data.IData
}

func (a *WalletRepository) Create(wallet *dto.Wallet) error {
	if wallet == nil {
		return &internalError.InvalidData{DataField: "wallet"}
	}

	a.d.Insert(wallet.Username, wallet.Balance)
	return nil
}

func (a *WalletRepository) Update(wallet *dto.Wallet) error {
	if wallet == nil {
		return &internalError.InvalidData{DataField: "wallet"}
	}
	if err := a.d.Update(wallet.Username, wallet.Balance); err != nil {
		return err
	}
	return nil
}

func (a *WalletRepository) GetAll() dto.Wallets {
	var wallets = dto.Wallets{}
	for k, v := range a.d.GetAll() {
		wallets = append(wallets, &dto.Wallet{Username: k, Balance: v})
	}
	return wallets
}

func (a *WalletRepository) GetWalletByUsername(username string) (*dto.Wallet, error) {
	if balance, err := a.d.GetByUsername(username); err != nil {
		return nil, err
	} else {
		return &dto.Wallet{Username: username, Balance: balance}, nil
	}
}

func NewWalletRepository(data data.IData) IWalletRepository {
	return &WalletRepository{d: data}
}
