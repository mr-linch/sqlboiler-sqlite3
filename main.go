package main

import (
	"github.com/mr-linch/sqlboiler-sqlite3/driver"
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

func main() {
	drivers.DriverMain(&driver.SQLiteDriver{})
}
