package main

import (
	"fmt"
	"karaca/config"
	"karaca/controller"
	"karaca/data"
	"karaca/repository"
	"karaca/service"
	"net/http"
)

func main() {
	fmt.Println(config.Get().InitialBalanceAmount)

	walletRepository := repository.NewWalletRepository(data.NewData())
	walletService := service.NewWalletService(walletRepository)
	walletController := controller.NewWalletController(walletService)
	http.HandleFunc("/", walletController.Handle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}


