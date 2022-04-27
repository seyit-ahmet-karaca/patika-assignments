package error

type WalletAlreadyExists struct {
	Message string `json:"errorMessage"`
}

var walletAlreadyExists = &WalletAlreadyExists{Message: "Wallet already exists"}

func (*WalletAlreadyExists) Error() string {
	return walletAlreadyExists.Message
}
