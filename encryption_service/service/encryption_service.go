package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/rysmaadit/finantier_test/encryption_service/config"
	"github.com/rysmaadit/finantier_test/encryption_service/contract"
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
	key, _ := hex.DecodeString(s.appConfig.EncryptionKey)
	payloadBytes, err := json.Marshal(req)

	if err != nil {
		log.Errorln("error marshal payload: ", err)
		return nil, err
	}

	r, err := aes.NewCipher(key)

	if err != nil {
		log.Errorln("error encrypt payload: ", err)
		return nil, err
	}

	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Errorln("error read nonce: ", err)
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(r)

	if err != nil {
		log.Errorln("error read nonce: ", err)
		return nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, payloadBytes, nil)

	response := &contract.EncryptedDataResponse{Data: ciphertext}
	return response, nil
}
