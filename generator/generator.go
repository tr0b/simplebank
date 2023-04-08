package generator

import (
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"

	db "github.com/tr0b/simplebank/db/sqlc"
)

type Secret struct {
	Phrase string
}

func GenerateAccount(currency string, owner string) (db.Account, error) {
	a := db.Account{}
	err := faker.FakeData(&a)
	a.Currency = currency
	a.Owner = owner
	if err != nil {
		return a, err
	}

	return a, nil
}

func GenerateTransfer() (db.Transfer, error) {
	t := db.Transfer{}
	err := faker.FakeData(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func GenerateUser() (db.User, error) {
	u := db.User{}
	err := faker.FakeData(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}

func GenerateSecret() (Secret, error) {
	s := Secret{}
	err := faker.FakeData(&s, options.WithRandomStringLength(32))

	if err != nil {
		return s, err
	}

	return s, nil
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
