package model

type Runtime struct {
	Dev        bool
	APIDev     string
	APIProd    string
	DevOrigin  string
	ProdOrigin string
	Port       string

	//creds
	User      string
	Pass      string
	JWTsecret string
	HashSalt  string
}

type FuncUsage struct {
	RSS     int
	JSONRSS int
	Atom    int
}

var Usage FuncUsage
