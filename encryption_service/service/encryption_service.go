package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/mergermarket/go-pkcs7"
	"github.com/rysmaadit/finantier_test/encryption_service/common/errors"
	"github.com/rysmaadit/finantier_test/encryption_service/config"
	"github.com/rysmaadit/finantier_test/encryption_service/contract"

	log "github.com/sirupsen/logrus"
)

type encryptionService struct {
	appConfig *config.Config
}

type EncryptionServiceInterface interface {
	Encrypt(req *contract.GetDailyTimeSeriesStockResponse) (*contract.EncryptedDataResponse, error)
}

func NewEncryptionService(appConfig *config.Config) *encryptionService {
	return &encryptionService{
		appConfig: appConfig,
	}
}

func (s *encryptionService) Encrypt(req *contract.GetDailyTimeSeriesStockResponse) (*contract.EncryptedDataResponse, error) {
	key := []byte(s.appConfig.EncryptionKey)
	reqBytes, _ := json.Marshal(req)

	cipherText, err := aes256CBCEncryptor(key, reqBytes)

	if err != nil {
		log.Errorln(err)
		return nil, fmt.Errorf("error encrypt request payload")
	}

	response := &contract.EncryptedDataResponse{Data: cipherText}
	return response, nil
}

func aes256CBCEncryptor(key, message []byte) ([]byte, error) {
	plainText, err := pkcs7.Pad(message, aes.BlockSize)

	if err != nil {
		return nil, fmt.Errorf(`plainText: "%s" has error`, plainText)
	}

	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return nil, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return cipherText, nil
}

func aes256CBCDecryptor(key, encryptedMessage []byte) ([]byte, error) {
	cipherText, _ := hex.DecodeString(fmt.Sprintf("%x", encryptedMessage))

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return cipherText, nil
}
