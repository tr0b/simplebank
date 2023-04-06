package generator

import (
	"github.com/go-faker/faker/v4"

	db "github.com/tr0b/simplebank/db/sqlc"
)

func GenerateAccount() (db.Account, error) {
	a := db.Account{}
	err := faker.FakeData(&a)
	if err != nil {
		return a, err
	}

	return a, nil
}

// func GenerateCreateAccountParams() (db.CreateAccountParams, error) {
// 	p := db.CreateAccountParams{}
// 	err := faker.FakeData(&p)
// 	if err != nil {
// 		return p, err
// 	}
//
// 	return p, nil
// }

// func GenerateCreateEntryParams() (db.CreateEntryParams, error) {
// 	p := db.CreateEntryParams{}
// 	err := faker.FakeData(&p)
// 	if err != nil {
// 		return p, err
// 	}
//
// 	return p, nil
// }
