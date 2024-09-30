package database

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"github.com/afurgapil/phost/database/internal/config"
)

type Record struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Database struct {
	Records []Record `json:"records"`
	Mutex   sync.Mutex
}

var encryptionKey []byte

func init() {
	keyStr, err := config.LoadConfig("ENCRYPTION_KEY")
	if err != nil {
		log.Fatalf("Error loading encryption key: %v", err)
	}

	encryptionKey = []byte(keyStr)
}
func (db *Database) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		log.Println("The file is empty, a new database will be created.")
		return nil
	}

	decryptedData, err := decrypt(data, encryptionKey)
	if err != nil {
		return err
	}

	return json.Unmarshal(decryptedData, db)
}

func (db *Database) SaveToFile(filename string) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	encryptedData, err := encrypt(data, encryptionKey)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(encryptedData)
	return err
}

func encrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}

func decrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func (db *Database) AddRecord(recordValue string) (int, string) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	id := len(db.Records) + 1
	db.Records = append(db.Records, Record{ID: id, Value: recordValue})
	println("DB:", db, "ID:", id, "Record:", recordValue)
	return id, recordValue
}

func (db *Database) GetRecord(id int) *Record {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	for _, record := range db.Records {
		if record.ID == id {
			return &record
		}
	}
	return nil
}

func (db *Database) DeleteRecord(id int) bool {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	for i, record := range db.Records {
		if record.ID == id {
			db.Records = append(db.Records[:i], db.Records[i+1:]...)
			return true
		}
	}
	return false
}
func (db *Database) ClearRecords() {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	db.Records = []Record{}
}
