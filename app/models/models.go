package models

import (
	"rebuymaster/config"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
	"github.com/wawandco/ox/pkg/buffalotools"
)

var (
	// DB returns the DB connection for a given connection name
	db = buffalotools.DatabaseProvider(config.FS())
)

// DB returns the DB connection for the current environment.
func DB() *pop.Connection {
	return db(envy.Get("GO_ENV", "development"))
}

// FindDB allows to pull a connection by the name of it
// this comes handy when you have multiple databases to read
// from.
func FindDB(connname string) *pop.Connection {
	return db(connname)
}