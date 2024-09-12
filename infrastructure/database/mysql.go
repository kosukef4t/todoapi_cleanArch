package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql" // MySQLドライバー
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // ファイルソース
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

// LoadConfig reads and parses the config.yaml file
func LoadConfig() (*Config, error) {
	path := filepath.Join("config", "config.yaml")
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &config, nil
}

// Initialize sets up the database connection and runs migrations
func Initialize() (*sql.DB, error) {

	// Load the configuration
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("could not load config: %w", err)
	}

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	// Ping the database to ensure the connection is valid
	if err = db.Ping(); err != nil {
		log.Printf("Ping error: %v", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run database migrations
	if err := runMigrations(dsn); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

func runMigrations(dsn string) error {
	// データベース接続を初期化
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// マイグレーションファイルの絶対パスを取得
	absolutePath, err := filepath.Abs("./migrations")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// マイグレーションソースを初期化
	migrationSource := fmt.Sprintf("file://%s", filepath.ToSlash(absolutePath))

	// MySQLドライバーのインスタンスを作成
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize mysql driver: %w", err)
	}

	// マイグレーションを実行
	m, err := migrate.NewWithDatabaseInstance(migrationSource, "mysql", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("Database schema is already up-to-date. No migrations were applied")
		} else {
			fmt.Println("fail to run migrations: %w", err)
		}
	} else {
		fmt.Println("Migrations applied successfully")
	}
	return nil
}
