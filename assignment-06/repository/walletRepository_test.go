package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"karaca/dto"
	internalError "karaca/error"
	"karaca/mock"
	"testing"
)

func TestCreateWithoutError(t *testing.T) {
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		Insert("testUser", 0).
		Times(1)

	repository := NewWalletRepository(mockData)

	wallet := &dto.Wallet{Username: "testUser", Balance: 0}
	err := repository.Create(wallet)

	assert.Nil(t, err)
}

func TestCreateWithInvalidDataError(t *testing.T) {
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		Insert(gomock.Any(), gomock.Any()).
		Times(0)

	repository := NewWalletRepository(mockData)

	err := repository.Create(nil)
	_, ok := err.(*internalError.InvalidData)
	assert.True(t, ok)
}

func TestUpdate(t *testing.T) {
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	repository := NewWalletRepository(mockData)

	wallet := &dto.Wallet{Username: "testUser", Balance: 0}
	err := repository.Update(wallet)

	assert.Nil(t, err)
}

func TestUpdatePassNilParameter(t *testing.T) {
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Return(&internalError.UsernameNotFound{}).
		Times(1)

	repository := NewWalletRepository(mockData)

	wallet := &dto.Wallet{Username: "testUser", Balance: 0}
	err := repository.Update(wallet)

	_, ok := err.(*internalError.UsernameNotFound)
	assert.True(t, ok)
}

func TestUpdateNonExistsWallet(t *testing.T) {
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		Update(gomock.Any(), gomock.Any()).
		Times(0)

	repository := NewWalletRepository(mockData)
	err := repository.Update(nil)

	_, ok := err.(*internalError.InvalidData)
	assert.True(t, ok)
}

func TestGetAllWithData(t *testing.T) {
	expectedData := map[string]int{
		"test": 0,
		"test2": 0,
		"test3": 0,
	}

	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		GetAll().
		Return(expectedData).
		Times(1)

	repository := NewWalletRepository(mockData)
	wallets := repository.GetAll()

	assert.Equal(t, len(expectedData), len(wallets))
}

func TestGetAllWithEmptyData(t *testing.T) {
	expectedData := map[string]int{}

	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		GetAll().
		Return(expectedData).
		Times(1)

	repository := NewWalletRepository(mockData)
	wallets := repository.GetAll()

	assert.Equal(t, len(wallets), 0)
}

func TestGetWalletByUsername(t *testing.T) {
	expectedBalance := 0
	expectedUsername := "test"
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		GetByUsername(gomock.Any()).
		Return(0, nil).
		Times(1)

	repository := NewWalletRepository(mockData)
	actualWallets, err := repository.GetWalletByUsername(expectedUsername)

	assert.Nil(t, err)
	assert.Equal(t, expectedBalance, actualWallets.Balance)
	assert.Equal(t, expectedUsername, actualWallets.Username)
}

func TestGetWalletByUsernameNonExistsUsername(t *testing.T) {
	expectedUsername := "test"
	mockData := mock.NewMockIData(gomock.NewController(t))

	mockData.EXPECT().
		GetByUsername(gomock.Any()).
		Return(0, &internalError.UsernameNotFound{}).
		Times(1)

	repository := NewWalletRepository(mockData)
	actualWallets, err := repository.GetWalletByUsername(expectedUsername)
	_, ok := err.(*internalError.UsernameNotFound)

	assert.True(t, ok)
	assert.Nil(t, actualWallets)
}
