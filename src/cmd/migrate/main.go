package main

import (
	"database/sql"
	"github.com/oktopriima/marvel/bootstrap/http"
	"time"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/oktopriima/marvel/pkg/config"
	"log"
	"os"
	"strings"
)

func main() {
	c := http.NewBootstrap()

	err := c.Invoke(migration)
	if err != nil {
		log.Printf("error while run migration.\n message : %v", err)
	}
	return
}

func migration(cfg config.AppConfig) error {
	arg := os.Args[1:]
	if len(arg) == 0 {
		return fmt.Errorf("not enough argument. use './migration migration-help' to see available command")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", cfg.Mysql.MigrationDirectory),
		cfg.Mysql.Database,
		driver,
	)
	if err != nil {
		return err
	}

	switch strings.ToLower(arg[0]) {
	case "up":
		return up(m)
	case "down":
		return down(m)
	case "create":
		return create(cfg.Mysql.MigrationDirectory)
	case "version":
		return version(m)
	case "help":
		return help()
	default:
		return fmt.Errorf("unsupported argument")
	}
}

func up(m *migrate.Migrate) error {
	err := m.Up()
	if err != nil {
		return err
	}
	return nil
}

func down(m *migrate.Migrate) error {
	err := m.Down()
	if err != nil {
		return err
	}
	return nil
}

func create(dir string) error {
	arg := os.Args[1:]
	if len(arg) < 2 {
		return fmt.Errorf("not enough argument, add file migration name after 'create'. create init-users")
	}

	steps := []string{
		"up", "down",
	}

	var errMsg []string
	for _, step := range steps {
		filename := fmt.Sprintf("%s_%s.%s.sql", time.Now().Format("20060102150405"), arg[1], step)

		fmt.Println("creating migration file")
		fmt.Println(filename)

		_, err := os.Create(fmt.Sprintf("%s/%s", dir, filename))
		if err != nil {
			errMsg = append(errMsg, err.Error())
		}
	}

	if len(errMsg) > 0 {
		return fmt.Errorf("error while create migrations\n%s", strings.Join(errMsg, "\n"))
	}

	return nil
}

func version(m *migrate.Migrate) error {
	u, _, err := m.Version()
	if err != nil {
		return err
	}

	fmt.Printf("your current migration version is on %d\n", u)

	return nil
}

func help() error {
	b, err := os.ReadFile("migration-help")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
