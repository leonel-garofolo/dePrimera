package config

import (
	"flag"
	"fmt"
)

type Config struct {
	dbUser     string
	dbPswd     string
	dbHost     string
	dbPort     string
	dbName     string
	testDBHost string
	testDBName string
	apiPort    string
	migrate    string
}

func Get() *Config {
	conf := &Config{}
	/*
		MYSQL_PASSWORD=root
		MYSQL_USER=root
		MYSQL_PORT=3306
		MYSQL_HOST=localhost
		MYSQL_DB=de_primera_app
		TEST_DB_HOST=localhost
		TEST_DB_NAME=boilerplatetest
		API_PORT=8081
	*/

	flag.StringVar(&conf.dbUser, "dbuser", "root", "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", "root", "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", "3306", "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", "localhost", "DB host")
	flag.StringVar(&conf.dbName, "dbname", "de_primera_app", "DB name")
	flag.StringVar(&conf.apiPort, "apiPort", "8081", "API Port")
	flag.StringVar(&conf.migrate, "migrate", "up", "specify if we should be migrating DB 'up' or 'down'")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		c.dbUser,
		c.dbPswd,
		dbname,
	)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
}
