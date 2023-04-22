package app

import (
	"DMAPI/app/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var CFG *config.Config
var DB *sqlx.DB

func init() {
	CFG = config.InitCfg()

	spew.Dump(CFG)

	conn := `
	          host=` + CFG.Db.Host + `
	        dbname=` + CFG.Db.DbName + `
			   user=` + CFG.Db.UserName + `
	       sslmode=disable
			   port=` + CFG.Db.Port + `
			password=` + CFG.Db.Password + `
	`
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	DB = db

}
