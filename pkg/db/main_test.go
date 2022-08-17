package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/jonss/testcontainers-go-wrapper/pg"
)

func TestMain(m *testing.M) {
	cfg := pg.PostgresCfg{
		ImageName: pg.DefaultImage,
		Password:  "a_secret_password",
		UserName:  "a_user",
		DbName:    "db_name",
	}
	pgInfo, err := pg.Container(context.Background(), cfg)
	defer pgInfo.TearDown()

	if err != nil {
		panic(fmt.Sprintf("pg.Container(); unexpected error %v", err))
	}

	sqlDB, err := sql.Open("postgres", pgInfo.DbURL)
	if err != nil {
		panic(fmt.Sprintf("sql.Open(); error %v", err))
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("sql.DB.Ping(); error %v", err))
	}
}