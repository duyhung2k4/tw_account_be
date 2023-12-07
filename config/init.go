package config

import (
	"flag"

	"github.com/go-chi/jwtauth/v5"
)

func init() {
	loadEnv()

	//jwt
	tokenAuth = jwtauth.New("HS256", []byte("key"), nil)

	var flagMigrateDB bool

	flag.BoolVar(&flagMigrateDB, "db", false, "migrate table")
	flag.Parse()

	connect(flagMigrateDB)
}
