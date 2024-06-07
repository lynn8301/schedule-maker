package Uservices

func Login(username string, password string) error {
	return Auth(username, password)
}
