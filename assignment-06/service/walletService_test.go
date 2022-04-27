package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"karaca/config"
	"karaca/dto"
	error2 "karaca/error"
	"karaca/mock"
	"log"
	"os"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mock.NewMockIWalletRepository(mockController)
	username := "test"
	mockRepository.
		EXPECT().
		GetWalletByUsername(username).
		Return(nil, &error2.UsernameNotFound{}).
		Times(1)

	mockRepository.
		EXPECT().
		Create(&dto.Wallet{Username: username, Balance: config.Get().InitialBalanceAmount}).
		Return(nil).
		Times(1)

	service := NewWalletService(mockRepository)
	err := service.Create(username)

	assert.Nil(t, err)
}

func TestCreateExistsWallet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := mock.NewMockIWalletRepository(mockController)
	username := "test"
	mockRepository.
		EXPECT().
		GetWalletByUsername(username).
		Return(nil, nil).
		Times(1)

	mockRepository.
		EXPECT().
		Create(&dto.Wallet{Username: username, Balance: config.Get().InitialBalanceAmount}).
		Return(nil).
		Times(0)

	service := NewWalletService(mockRepository)
	err := service.Create(username)

	assert.Nil(t, err)
}


func TestUpdate(t *testing.T) {
	testUsername := "test"
	testWalletBalance := 100
	testBalanceUpdate := -50

	wallet := &dto.Wallet{Username: testUsername, Balance: testWalletBalance}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)

	repository.
		EXPECT().
		GetWalletByUsername(testUsername).
		Return(wallet, nil).
		Times(1)

	wallet.Balance = testBalanceUpdate + testWalletBalance
	repository.
		EXPECT().
		Update(wallet).
		Return(nil).
		Times(1)

	service := NewWalletService(repository)
	err := service.Update(testUsername, testBalanceUpdate)

	assert.Nil(t, err)
}

func TestUpdateNonExistWallet(t *testing.T) {
	testUsername := "test"
	testWalletBalance := 100
	testBalanceUpdate := -50

	wallet := &dto.Wallet{Username: testUsername, Balance: testWalletBalance}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)

	repository.
		EXPECT().
		GetWalletByUsername(testUsername).
		Return(nil, &error2.UsernameNotFound{}).
		Times(1)

	wallet.Balance = testBalanceUpdate + testWalletBalance
	repository.
		EXPECT().
		Update(wallet).
		Return(nil).
		Times(0)

	service := NewWalletService(repository)
	err := service.Update(testUsername, testBalanceUpdate)

	_, ok := err.(*error2.UsernameNotFound)

	assert.True(t, ok)
}

func TestUpdateInvalidBalance(t *testing.T) {
	testUsername := "test"
	testWalletBalance := 100
	testBalanceUpdate := -250

	wallet := &dto.Wallet{Username: testUsername, Balance: testWalletBalance}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)

	repository.
		EXPECT().
		GetWalletByUsername(testUsername).
		Return(wallet, nil).
		Times(1)

	wallet.Balance = testBalanceUpdate + testWalletBalance
	repository.
		EXPECT().
		Update(wallet).
		Return(nil).
		Times(0)

	service := NewWalletService(repository)
	err := service.Update(testUsername, testBalanceUpdate)

	_, ok := err.(*error2.InvalidBalance)
	assert.True(t, ok)
}

func TestGetAll(t *testing.T) {
	expectedWallets := dto.Wallets{
		{Username: "test", Balance: 0},
		{Username: "test1", Balance: 0},
		{Username: "test2", Balance: 0},
		{Username: "test3", Balance: 0},
	}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)
	repository.
		EXPECT().
		GetAll().
		Return(expectedWallets).
		Times(1)

	service := NewWalletService(repository)
	actualWallets := service.GetAll()

	assert.Equal(t, expectedWallets, actualWallets)
	assert.Equal(t, len(expectedWallets), len(actualWallets))
}

func TestGetAllEmptyData(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	expectedWallets := dto.Wallets{}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)
	repository.
		EXPECT().
		GetAll().
		Return(expectedWallets).
		Times(1)

	service := NewWalletService(repository)
	actualWallets := service.GetAll()

	assert.Equal(t, expectedWallets, actualWallets)
	assert.Equal(t, len(expectedWallets), len(actualWallets))
}

func TestGetWalletByUsername(t *testing.T) {
	testUsername := "testUsername"
	testBalance := 0

	expectedWallet := &dto.Wallet{Username: testUsername, Balance: testBalance}

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)
	repository.
		EXPECT().
		GetWalletByUsername(testUsername).
		Return(expectedWallet, nil).
		Times(1)

	service := NewWalletService(repository)
	actualWallet, err := service.GetWalletByUsername(testUsername)

	assert.Nil(t, err)
	assert.Equal(t, expectedWallet.Username, actualWallet.Username)
	assert.Equal(t, expectedWallet.Balance, actualWallet.Balance)
}

func TestGetWalletByUsernameUsernameNotFound(t *testing.T) {
	testUsername := "testUsername"

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	repository := mock.NewMockIWalletRepository(mockController)
	repository.
		EXPECT().
		GetWalletByUsername(testUsername).
		Return(nil, &error2.UsernameNotFound{}).
		Times(1)

	service := NewWalletService(repository)
	actualWallet, err := service.GetWalletByUsername(testUsername)

	_, ok := err.(*error2.UsernameNotFound)

	assert.True(t, ok)
	assert.Nil(t, actualWallet)
}
