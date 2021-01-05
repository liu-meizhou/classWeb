package DB

import (
	"github.com/beego/beego/v2/client/orm"
)

func init() {
	// set default database
	orm.RegisterDataBase("postgres", "postgres", "postgresql://postgres:123456@42.193.143.9:5432/postgres?sslmode=disable&charset=utf8")

	// create table
	orm.RunSyncdb("postgres", false, true)
}
