package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"os"

	"github.com/caioLeone/go-arena-api/internal/config"
	_ "github.com/lib/pq"
)

// CONECTA TUDO AO POSTGRESQL
func Connect(cfg *config.Config) (*sql.DB, error) {
	// Montar DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.BDSSLMode,
	)
	//ABRIR CONEXAO
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexao com o PostgreSQL: %w", err)
	}
	//TESTAR CONEXAO COM PING
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao PostgreSQL: %w", err)
	}
	log.Println("Conectado ao PostgreSQL com Sucesso")
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

// RunMigrations executa todas as Migrations SQL na ordem correta
func RunMigrations(db *sql.DB, migrationsPath string) error {
	//Cria tabela de rastreamento de migrations(se nao existir)
	createMigrationsTable := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			id SERIAL PRIMARY KEY,
			version VARCHAR(255) UNIQUE NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`
	if _, err := db.Exec(createMigrationsTable); err != nil {
		return fmt.Errorf("Erro ao criar tabela scheme_migrations: %w", err)
	}

	//LISTAR TODOS OS ARQUIVOS .UP.SQL
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("Erro ao Ler Diretorio de Migrations: %w", err)
	}

	//FILTRAR APENAS ARQUIVOS UP.SQL E ORDENA-LOS
	var migrationFiles []string
	for _, file := range files{
		if strings.HasSuffix(file.Name(), ".up.sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	//EXECUTAR CADA MIGRATION
	for _, file := range migrationFiles{
		version := strings.TrimSuffix(file, ".up.sql")

		//Verificar se migration ja foi executada
		var exists int 
		err := db.QueryRow(
			"SELECT COUNT(*) FROM scheme_migrations WHERE = $1",
			version,).Scan(&exists)
		if err != nil {
			return fmt.Errorf("Erro ao verificar migration %s: %w", version, err)
		}

		//Se ja foi executado, pular
		if exists > 0{
			log.Printf("Migrations aplicada: %s", version)
			continue
		}

		//LER CONTEUDO DO ARQUIVO SQL
		filepath := filepath.Join(migrationsPath, file)
		sqlContent, err := ioutil.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("Erro ao ler arquivo %s: %w", file, err)
		}

		//EXECUTAR SQL
		if _, err := db.Exec(string(sqlContent)); err != nil {
			return fmt.Errorf("Erro ao executar migration %s: %w", version, err)
		}

		//REGISTRAR MIGRATION COMO APLICADA
		_, err = db.Exec(
			`INSERT INTO schema_migrations (version) VALUES ($1)`, 
			version
		)
		if err != nil {
			return fmt.Errorf("Erro ao Registrar Migration %s: %w", version, err)
		}
		log.Printf("Migration Aplicada: %s", version)
	}

	log.Println("Todas as Migrations Foram Executadas Com Sucesso")
	return nil
}

func GetConnection(cfg *config.Config) (*sql.DB, error) {
	return Connect(cfg)
}
