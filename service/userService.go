package service

import (
	"time"
)

const TOKE_EXPIRE_DURATION = time.Hour * 2

type User struct {
	account  string
	password string
}

func createUser() string {
	test := "test+test"
	return test
}
