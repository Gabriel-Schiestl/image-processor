package config

import "fmt"

type dbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewDbConfig(host, user, password, dbname string, port int) *dbConfig {
	return &dbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbname,
	}
}

func (db *dbConfig) ToString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.DbName)
}