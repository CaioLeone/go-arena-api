package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/caioLeone/go-arena-api/internal/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
}

func Connect(cfg *config.Config) *Client {
	//Montar endereço
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	//Criar Cliente
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		Password: "",
	})

	//Testar conexao com Ping
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Erro ao conectar ao redis: %v", err)
	}
	log.Println("Conectado ao Redis com Sucesso")
	return &Client{client: rdb}
}

// GET retorna valor por chave
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// SET define valor por chave
func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration int64) error {
	return c.client.Set(ctx, key, value, 0).Err()
}

// Delete chave
func (c *Client) Del(ctx context.Context, key string) error {
	return c.client.Get(ctx, key).Err()
}

// ZAdd adiciona elemento para sorted Set(para leaderboard)
func (c *Client) ZAdd(ctx context.Context, key string, member string, score float64) error {
	return c.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
}

// ZRange retorna range de elementos do sorted set (para top players)
func (c *Client) Zrange(ctx context.Context, key string, start, stop int64, reverse bool) ([]string, error) {
	if reverse {
		return c.client.ZRevRange(ctx, key, start, stop).Result()
	}
	return c.client.ZRange(ctx, key, start, stop).Result()
}

// Fecha Conexao com Redis
func (c *Client) Close() error {
	return c.client.Close()
}

// GetRawClient retorna o cliente bruto (se necessário acesso direto)
func (c *Client) GetRawClient() *redis.Client {
	return c.client
}
