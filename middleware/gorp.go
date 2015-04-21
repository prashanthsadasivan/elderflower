package middleware

import (
	"database/sql"
	"elderflower/config"
	"elderflower/models"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

func DatabaseMiddleware() gin.HandlerFunc {
	Db, err := sql.Open(config.Get("driver"), config.Get("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	Dbm := &gorp.DbMap{Db: Db, Dialect: gorp.PostgresDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.SoleUser{}).SetKeys(true, "UserId")
	setColumnSizes(t, map[string]int{
		"Number":   15,
		"GcmId":    2000,
		"QrSecret": 35,
	})

	t = Dbm.AddTable(models.SMSMessage{}).SetKeys(true, "MessageId")
	setColumnSizes(t, map[string]int{
		"Num":     15,
		"Message": 2000,
	})
	Dbm.TraceOn("[gorp]", log.New(ioutil.Discard, "INFO  ", log.Ldate|log.Ltime|log.Lshortfile))
	Dbm.CreateTables()
	log.Println("initialized")
	return func(c *gin.Context) {
		txn, err := Dbm.Begin()
		if err != nil {
			panic(err)
		}
		c.Set("txn", txn)
		c.Next()
		if c.LastError() != nil {
			if err := txn.Rollback(); err != nil && err != sql.ErrTxDone {
				panic(err)
			}
		} else {
			if err := txn.Commit(); err != nil && err != sql.ErrTxDone {
				panic(err)
			}
		}
	}
}
