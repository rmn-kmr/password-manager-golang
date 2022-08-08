package data

import (
	"errors"
	"log"
	"os"
	"passwordmanager/data/file"
	"passwordmanager/dcrypt"

	"github.com/google/uuid"
)

const fileName = "./data/storage/password.csv"

type PasswordData struct {
	UUID     string `json:"id"`
	Website  string `json:"website"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

var passwordDataCache []PasswordData

// add new entry for password data, Note: duplicate website row will be override
func (ps *PasswordStore) AddNewRecord(newRecord PasswordData) error {
	encryptedPass, err := dcrypt.Encrypt(newRecord.Password, os.Getenv("MY_SECRET"))
	if err != nil {
		log.Print(err)
		return err
	}
	uid := uuid.New().String() // creating Unique ID for each new records
	newRecord.Password = encryptedPass
	newRecord.UUID = uid
	passwordDataCache, err = filterPasswordByWebsite(newRecord.Website)
	if err != nil {
		log.Print(err)
	}
	passwordDataCache = append(passwordDataCache, newRecord)
	err = savePasswordList(passwordDataCache)
	if err != nil {
		return err
	}
	return nil
}

//write password list data
func savePasswordList(data []PasswordData) error {
	var csvRows [][]string
	for _, row := range data {
		csvRows = append(csvRows, []string{row.UUID, row.Website, row.UserName, row.Password})
	}
	err := file.WriteMultiRow(fileName, csvRows)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// fetch all List of password data
func (ps *PasswordStore) GetPasswordList() ([]PasswordData, error) {
	return fetchPasswordData()
}

// fetch all List of password data
func fetchPasswordData() ([]PasswordData, error) {
	if passwordDataCache != nil {
		return passwordDataCache, nil
	}
	rows, err := file.Fetch(fileName)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		passwordData := PasswordData{
			UUID:     row[0],
			Website:  row[1],
			UserName: row[2],
			Password: row[3],
		}
		passwordDataCache = append(passwordDataCache, passwordData)
	}

	return passwordDataCache, nil
}

// helper method for filter password list by website
func filterPasswordByWebsite(website string) ([]PasswordData, error) {
	rows, err := fetchPasswordData()
	if err != nil {
		return nil, err
	}
	var filteredRows []PasswordData
	for _, row := range rows {
		if row.Website != website {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows, nil
}

func (ps *PasswordStore) GetPlainPassword(id string) (string, error) {
	rows, err := fetchPasswordData()
	if err != nil {
		return "", err
	}
	var cipherPass string
	for _, row := range rows {
		if row.UUID == id {
			cipherPass = row.Password
			break
		}
	}
	if cipherPass != "" {
		plainPass, err := dcrypt.Decrypt(cipherPass, os.Getenv("MY_SECRET"))
		if err != nil {
			log.Print(err)
			return "", err
		}
		return plainPass, nil
	}
	return "", errors.New("invalid Id")
}
