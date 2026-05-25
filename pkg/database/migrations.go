package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// MigrationsInfo retorna informações sobre as migrations
func MigrationsInfo(migrationsPath string) error {
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("Erro ao ler diretorio: %w", err)
	}

	log.Println("Migrations Encotnradas:")
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			log.Printf(" - %s", file.Name())
		}
	}
	return nil
}

// CreateMigrationsDir cria o diretório de migrations se não existir
func CreateMigrationsDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
