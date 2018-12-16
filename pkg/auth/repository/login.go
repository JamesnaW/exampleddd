package repository

import "exampleddd"

// Get userid by access token
func (r repo) Get(accessToken string) (bool, int, error) {
	if accessToken == "" {
		return false, 0, nil
	}
	if accessToken == "1" {
		return true, 1, nil
	}
	return false, 0, nil
}

// Check username and compare user password
// user: admin
// password: admin
// hash: $2a$04$k1Zr0PKXlgLSSRWFfLC1xOyTS46ezNh6KG7nmyEv4cYszkj1ajPRK
func (r repo) Check(user, pass string) (bool, int, error) {
	if user != "admin" {
		return false, 0, nil
	}
	if !exampleddd.ComparePasswords("$2a$04$k1Zr0PKXlgLSSRWFfLC1xOyTS46ezNh6KG7nmyEv4cYszkj1ajPRK", []byte(pass)) {
		return false, 0, nil
	}
	return true, 1, nil
}
