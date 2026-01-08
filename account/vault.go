package account

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/blue-script/password/encrypter"
	"github.com/blue-script/password/output"
	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Not successful unmarshal file data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	deleteIndex := -1
FindIndex:
	for index, value := range vault.Accounts {
		isMatched := value.Url == url
		if isMatched {
			deleteIndex = index
			break FindIndex
		}
	}

	if deleteIndex == -1 {
		color.Red("Not found index")
		return false
	}

	vault.Accounts = slices.Delete(vault.Accounts, deleteIndex, deleteIndex+1)

	vault.save()
	return true
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var results []Account
	for _, acc := range vault.Accounts {
		isMatched := checker(acc, str)
		if isMatched {
			results = append(results, acc)
		}
	}
	return results
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError("Not successful marshal account")
		return
	}
	vault.db.Write(data)
}
