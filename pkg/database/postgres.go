package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/caioLeone/go-arena-api/internal/config"
	_ "github.com/lib/pq"
)

// Connect conecta ao banco de dados PostgreSQL
func Connect(cfg *config.Config) (*sql.DB, error) {
	// Montar DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Abrir conexão
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com PostgreSQL: %w", err)
	}

	// Testar conexão com Ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao PostgreSQL: %w", err)
	}

	log.Println("Conectado ao PostgreSQL com sucesso")

	// Configurar pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

// RunMigrations executa todas as migrations SQL na ordem correta
func RunMigrations(db *sql.DB, migrationsPath string) error {
	// Criar tabela de rastreamento de migrations (se não existir)
	createMigrationsTable := `
    CREATE TABLE IF NOT EXISTS schema_migrations (
        id SERIAL PRIMARY KEY,
        version VARCHAR(255) UNIQUE NOT NULL,
        applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := db.Exec(createMigrationsTable); err != nil {
		return fmt.Errorf("erro ao criar tabela schema_migrations: %w", err)
	}

	// Listar todos os arquivos .up.sql
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("erro ao ler diretório de migrations: %w", err)
	}

	// Filtrar apenas arquivos .up.sql e ordená-los
	var migrationFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	// Executar cada migration
	for _, file := range migrationFiles {
		version := strings.TrimSuffix(file, ".up.sql")

		// Verificar se migration já foi executada
		var exists int
		err := db.QueryRow(
			"SELECT COUNT(*) FROM schema_migrations WHERE version = $1", version).Scan(&exists)

		if err != nil {
			return fmt.Errorf("erro ao verificar migration %s: %w", version, err)
		}

		// Se já foi executada, pular
		if exists > 0 {
			log.Printf("⏭️  Migration já aplicada: %s", version)
			continue
		}

		// Ler conteúdo do arquivo SQL
		filePath := filepath.Join(migrationsPath, file)
		sqlContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("erro ao ler arquivo %s: %w", file, err)
		}

		// Executar SQL
		if _, err := db.Exec(string(sqlContent)); err != nil {
			return fmt.Errorf("erro ao executar migration %s: %w", version, err)
		}

		// Registrar migration como aplicada
		_, err = db.Exec(
			"INSERT INTO schema_migrations (version) VALUES ($1)",
			version,
		)
		if err != nil {
			return fmt.Errorf("erro ao registrar migration %s: %w", version, err)
		}

		log.Printf("Migration aplicada: %s", version)
	}

	log.Println("Todas as migrations foram executadas com sucesso")
	return nil
}

// GetConnection retorna a conexão com o banco (para reutilização)
func GetConnection(cfg *config.Config) (*sql.DB, error) {
	return Connect(cfg)
}
