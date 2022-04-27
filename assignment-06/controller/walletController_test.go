package controller

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"karaca/dto"
	error2 "karaca/error"
	"karaca/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	expectedUsername := "testUser"

	mockService := mock.NewMockIWalletService(mockController)
	mockService.EXPECT().
		Create(expectedUsername).
		Return(nil).
		Times(1)

	controller := NewWalletController(mockService)

	r := httptest.NewRequest(http.MethodPut, "/"+expectedUsername, http.NoBody)
	w := httptest.NewRecorder()
	controller.Handle(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestCreateWithInvalidData(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	givenUsername := "testUser"
	expectedErr := &error2.InvalidData{}

	service := mock.NewMockIWalletService(mockController)
	service.
		EXPECT().
		Create(givenUsername).
		Return(expectedErr).
		Times(1)

	controller := NewWalletController(service)

	r := httptest.NewRequest(http.MethodPut, "/"+givenUsername, http.NoBody)
	w := httptest.NewRecorder()
	controller.Handle(w, r)

	actualErr := &error2.InvalidData{}
	json.Unmarshal(w.Body.Bytes(), actualErr)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))

}

func TestGetAll(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	serviceReturn := []*dto.Wallet{
		{Username: "test1", Balance: 0},
		{Username: "test2", Balance: 0},
		{Username: "test3", Balance: 0},
		{Username: "test4", Balance: 0},
		{Username: "test5", Balance: 0},
	}

	service := mock.NewMockIWalletService(mockController)
	service.
		EXPECT().
		GetAll().
		Return(serviceReturn).
		Times(1)

	controller := NewWalletController(service)

	r := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	w := httptest.NewRecorder()
	controller.Handle(w, r)

	actualWallets := []*dto.Wallet{}
	json.Unmarshal(w.Body.Bytes(), &actualWallets)

	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, serviceReturn, actualWallets)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestGet(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUsername := "test"
	expectedWallet := &dto.Wallet{Username: testUsername, Balance: 0}

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		GetWalletByUsername(testUsername).
		Return(expectedWallet, nil).
		Times(1)

	controller := NewWalletController(service)

	r := httptest.NewRequest(http.MethodGet, "/"+testUsername, http.NoBody)
	w := httptest.NewRecorder()
	controller.Handle(w, r)

	actualWallet := &dto.Wallet{}
	json.Unmarshal(w.Body.Bytes(), actualWallet)

	assert.Equal(t, expectedWallet, actualWallet)
	assert.Equal(t, expectedWallet.Balance, actualWallet.Balance)
	assert.Equal(t, expectedWallet.Username, actualWallet.Username)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestGetNonExistsUser(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUsername := "test"
	expectedErr := &error2.UsernameNotFound{}

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		GetWalletByUsername(testUsername).
		Return(nil, expectedErr).
		Times(1)

	controller := NewWalletController(service)

	r := httptest.NewRequest(http.MethodGet, "/"+testUsername, http.NoBody)
	w := httptest.NewRecorder()
	controller.Handle(w, r)

	actualErr := &error2.UsernameNotFound{}
	json.Unmarshal(w.Body.Bytes(), actualErr)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestUpdateByUsername(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUser := "testUser"
	testBalance := 10

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		Update(testUser, testBalance).
		Return(nil).
		Times(1)

	controller := NewWalletController(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(`{"balance": `+strconv.Itoa(testBalance)+`}`)))
	w := httptest.NewRecorder()

	controller.Handle(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestUpdateByUsernameNonExistUsername(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUser := "testUser"
	testBalance := 10
	expectedError := &error2.UsernameNotFound{}

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		Update(testUser, testBalance).
		Return(expectedError).
		Times(1)

	controller := NewWalletController(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(`{"balance": `+strconv.Itoa(testBalance)+`}`)))
	w := httptest.NewRecorder()

	controller.Handle(w, r)

	actualUsernameError := &error2.UsernameNotFound{}
	json.Unmarshal(w.Body.Bytes(), actualUsernameError)

	assert.Equal(t, expectedError, actualUsernameError)
	assert.Equal(t, w.Result().StatusCode, http.StatusInsufficientStorage)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestUpdateByUsernameLessThenMinimumBalance(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUser := "testUser"
	testBalance := 10
	expectedError := &error2.InvalidBalance{}

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		Update(testUser, testBalance).
		Return(expectedError).
		Times(1)

	controller := NewWalletController(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(`{"balance": `+strconv.Itoa(testBalance)+`}`)))
	w := httptest.NewRecorder()

	controller.Handle(w, r)

	actualUsernameError := &error2.InvalidBalance{}
	json.Unmarshal(w.Body.Bytes(), actualUsernameError)

	assert.Equal(t, expectedError, actualUsernameError)
	assert.Equal(t, w.Result().StatusCode, http.StatusInsufficientStorage)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestUpdateByUsernameWrongJsonStructure(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testUser := "testUser"
	testBalance := 10
	expectedError := &error2.InvalidBalance{}

	service := mock.NewMockIWalletService(mockController)
	service.EXPECT().
		Update(testUser, testBalance).
		Return(expectedError).
		Times(0)

	controller := NewWalletController(service)
	r := httptest.NewRequest(http.MethodPost, "/"+testUser, bytes.NewBuffer([]byte(`{"balance"}`)))
	w := httptest.NewRecorder()

	controller.Handle(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}