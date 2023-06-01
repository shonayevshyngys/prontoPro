package database

type Redis struct {
}

type RedisInterface interface {
	Set(key string, value []byte)
	Get(key string) (string, error)
	RPush(key string, value []byte) (int64, error)
	LRange(key string) ([]string, error)
	Del(key string)
}

func (r *Redis) LRange(key string) ([]string, error) {
	return RedisInstance.LRange(RedisContext, key, 0, -1).Result()
}

func (r *Redis) Set(key string, value []byte) {
	RedisInstance.Set(RedisContext, key, value, 0)
}

func (r *Redis) Get(key string) (string, error) {
	return RedisInstance.Get(RedisContext, key).Result()
}

func (r *Redis) RPush(key string, value []byte) (int64, error) {
	return RedisInstance.RPush(RedisContext, key, value).Result()
}

func (r *Redis) Del(key string) {
	RedisInstance.Del(RedisContext, key)
}
