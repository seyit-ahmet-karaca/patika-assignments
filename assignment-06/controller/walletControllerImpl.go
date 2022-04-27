package controller

import (
	"encoding/json"
	"karaca/dto"
	error2 "karaca/error"
	"karaca/service"
	"net/http"
)

type WalletController struct {
	service service.IWalletService
}

func (a *WalletController) Create(w http.ResponseWriter, r *http.Request) {
	err := a.service.Create(r.URL.Path[1:])

	w.Header().Add("content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(error2.ParseError(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *WalletController) UpdateByUsername(w http.ResponseWriter, r *http.Request) {
	var walletUpdate = &dto.WalletUpdateRequest{}

	w.Header().Add("content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(walletUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(error2.ParseError(err))
		return
	}

	updateErr := a.service.Update(r.URL.Path[1:], walletUpdate.Balance)

	if updateErr != nil {
		w.WriteHeader(http.StatusInsufficientStorage)
		w.Write(error2.ParseError(updateErr))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *WalletController) GetAll(w http.ResponseWriter, r *http.Request) {
	allWallets := a.service.GetAll()
	w.Header().Add("content-type", "application/json")

	walletJson, err := json.Marshal(allWallets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(error2.ParseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(walletJson)
}

func (a *WalletController) Get(w http.ResponseWriter, r *http.Request) {
	walletInfo, err := a.service.GetWalletByUsername(r.URL.Path[1:])

	w.Header().Add("content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(error2.ParseError(err))
		return
	}

	walletJson, err := json.Marshal(walletInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(error2.ParseError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(walletJson)
}

func (a *WalletController) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		if len(r.URL.Path) == 1 {
			a.GetAll(w, r)
		} else {
			a.Get(w, r)
		}
	case http.MethodPost:
		a.UpdateByUsername(w, r)
	case http.MethodPut:
		a.Create(w, r)
	default:
		return
	}
}

func NewWalletController(service service.IWalletService) IController {
	return &WalletController{service: service}
}
