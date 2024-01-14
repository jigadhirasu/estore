package mariadb

import (
	"estore/internal/util"
	"fmt"
	"os"
)

func T1ParseDatabaseDSN(dbname string) string {
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MARIADB_USER"),
		os.Getenv("MARIADB_PASS"),
		os.Getenv("MARIADB_HOST"),
		Name(dbname),
	)

	return dns
}

var F1ParseDatabaseDSN = util.TransferToFunc1(T1ParseDatabaseDSN)
