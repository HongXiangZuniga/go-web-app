package authorize

type RedisRepository interface {
	SetHash(hash, email string) error
	Authorize(hash string) (*string, error)
}
