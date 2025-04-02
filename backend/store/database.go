package store

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error){
	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=medical sslmode=disable")

	if err != nil{
		return nil, fmt.Errorf("failed to open database: %v", err)	
	}

	fmt.Println("database connected successfully")
	return db, nil
	
}

func MigrateFS (db *sql.DB, migrationFS fs.FS, dir string) error {
	//get working dir
	currentDir, err := os.Getwd()
	if err != nil{
		return fmt.Errorf("error getting working directory: %w", err)
	}

	//construct full path
	fullpath := filepath.Join(currentDir, dir)
	log.Printf("current directory is: %s", currentDir)
	log.Printf("full path dir: %s", fullpath)

	//verify if directory exists
	if _, err := os.Stat(fullpath); os.IsNotExist(err){
		log.Printf("error getting directory %s", fullpath)
		return fmt.Errorf("error getting directory %s", fullpath)
	}

	err = goose.SetDialect("postgres")
	if err != nil{
		return fmt.Errorf("set dialect: %w", err)
	}
	goose.SetVerbose(true)
	err = goose.Up(db, fullpath)
	if err != nil{
		return fmt.Errorf("goose up failed %w", err)
	}
	return nil
}
