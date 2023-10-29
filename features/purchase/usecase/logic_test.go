package usecase

import (
	"errors"
	"event_ticket/app/mocks"
	"event_ticket/features/purchase"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePurchaseSuccess(t *testing.T) {
	repoData := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(repoData)

	purchaseData := purchase.PurchaseCore{
		UserId:   "123abc",
		EventId:  "abc456",
		Quantity: 2,
	}

	repoData.On("CreatePurchase", purchaseData).Return(1, nil)
	row, err := purchaseUC.CreatePurchase(purchaseData)

	assert.NoError(t, err)
	assert.Equal(t, 1, row)
}

func TestCreatePurchaseEmptyUserIdEventId(t *testing.T){
	repoData := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(repoData)

	purchaseData := purchase.PurchaseCore{
		UserId:   "",
		EventId:  "",
		Quantity: 2,
	}

	_, err := purchaseUC.CreatePurchase(purchaseData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error, Credential data userID and eventID can't empty")
}

func TestCreatePurchaseQuantityError(t *testing.T){
	repoData := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(repoData)

	purchaseData := purchase.PurchaseCore{
		UserId:   "123abc",
		EventId:  "abc123",
	}

	_, err := purchaseUC.CreatePurchase(purchaseData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error, quantity must be a positive integer")
}

func TestCreatePurchaseError(t *testing.T) {
	repoData := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(repoData)

	// Empty EventId
	purchaseData := purchase.PurchaseCore{
		UserId: "123",
	}

	repoData.On("CreatePurchase", purchaseData).Return(0, errors.New("some error"))
	row, err := purchaseUC.CreatePurchase(purchaseData)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error, Credential data userID and eventID can't empty")
	assert.Equal(t, 0, row)
}

func TestDeletePurchaseSuccess(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"

	mockRepo.On("DeletePurchase", purchaseID).Return(nil)
	err := purchaseUC.DeletePurchase(purchaseID)

	assert.NoError(t, err)
}

func TestDeletePurchaseError(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"

	mockRepo.On("DeletePurchase", purchaseID).Return(errors.New("can't delete purchase"))
	err := purchaseUC.DeletePurchase(purchaseID)

	assert.Error(t, err)
	assert.Equal(t, errors.New("can't delete purchase"), err)
}

func TestReadAllPurchaseSuccess(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchasesData := []purchase.PurchaseCore{
		{
			ID: "123abc",
		},
		{
			ID: "abc123",
		},
	}

	mockRepo.On("ReadAllPurchase").Return(purchasesData, nil)
	purchases, err := purchaseUC.ReadAllPurchase()

	assert.NoError(t, err)
	assert.Equal(t, purchasesData, purchases)
}

func TestReadAllPurchaseError(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	mockRepo.On("ReadAllPurchase").Return(nil, errors.New("error get data"))
	purchases, err := purchaseUC.ReadAllPurchase()

	assert.Error(t, err)
	assert.Equal(t, errors.New("error get data"), err)
	assert.Nil(t, purchases)
}

func TestReadSpecificPurchaseSuccess(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"

	purchaseData := purchase.PurchaseCore{
		ID: purchaseID,
	}

	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(purchaseData, nil)
	purchase, err := purchaseUC.ReadSpecificPurchase(purchaseID)

	assert.NoError(t, err)
	assert.Equal(t, purchaseData, purchase)
}

func TestReadSpecificPurchaseError(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"

	expectedError := errors.New("purchase data not found")
	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(purchase.PurchaseCore{}, expectedError)

	purchaseData, err := purchaseUC.ReadSpecificPurchase(purchaseID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, purchaseData, purchase.PurchaseCore{})
}

func TestUpdatePurchaseSuccess(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"
	paymentStatus := "Success"

	data := purchase.PurchaseCore{
		Payment_status: paymentStatus,
	}

	existingPurchase := purchase.PurchaseCore{
		ID:             "123abc",
		Payment_status: "Pending",
	}

	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(existingPurchase, nil)
	mockRepo.On("UpdatePurchase", purchaseID, data).Return(data, nil)

	updatedPurchase, err := purchaseUC.UpdatePurchase(purchaseID, data)

	assert.NoError(t, err)
	assert.Equal(t, data, updatedPurchase)
}

func TestUpdatePurchaseError(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"
	data := purchase.PurchaseCore{
		Payment_status: "Success",
	}

	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(purchase.PurchaseCore{},  errors.New("not found"))
	updatedPurchase, err := purchaseUC.UpdatePurchase(purchaseID, data)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
	assert.Equal(t, purchase.PurchaseCore{}, updatedPurchase)
}

func TestUploadProofSuccess(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"
	proofImageData := "proof_image_data"
	imageSize := int64(1000000)
	
	data := purchase.PurchaseCore{
		Proof_image: proofImageData,
	}

	imageHeader := &multipart.FileHeader{
		Filename: "proof_image.jpg",
		Size:     imageSize,
	}

	existingPurchase := purchase.PurchaseCore{
		ID:          purchaseID,
		Proof_image: "old_proof_image_data",
	}

	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(existingPurchase, nil)
	mockRepo.On("UploadProof", purchaseID, data, imageHeader).Return(data, nil)

	uploadedProof, err := purchaseUC.UploadProof(purchaseID, data, imageHeader)

	assert.NoError(t, err)
	assert.Equal(t, data, uploadedProof)
}

func TestUploadProofError(t *testing.T) {
	mockRepo := new(mocks.PurchaseUseCaseInterface)
	purchaseUC := New(mockRepo)

	purchaseID := "123abc"
	proofImageData := "proof_image_data"
	imageSize := int64(15000000)

	data := purchase.PurchaseCore{
		Proof_image: proofImageData,
	}

	imageHeader := &multipart.FileHeader{
		Filename: "proof_image.jpg",
		Size:     imageSize,
	}
	existingPurchase := purchase.PurchaseCore{
		ID:          purchaseID,
		Proof_image: "old_proof_image_data",
	}

	mockRepo.On("ReadSpecificPurchase", purchaseID).Return(existingPurchase, nil)
	uploadedProof, err := purchaseUC.UploadProof(purchaseID, data, imageHeader)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "image file size should be less than 10 MB")
	assert.Equal(t, purchase.PurchaseCore{}, uploadedProof)
}