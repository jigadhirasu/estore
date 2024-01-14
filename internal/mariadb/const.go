package mariadb

import (
	"fmt"
	"os"
)

func DBKey(dbname string) string {
	return "ctx_db_" + dbname
}

func TxKey(dbname string) string {
	return "ctx_tx_" + dbname
}

func Name(dbname string) string {
	if dbname == "" {
		return ""
	}

	return fmt.Sprintf("%s_%s", os.Getenv("PREFIX_DBNAME"), dbname)
}
