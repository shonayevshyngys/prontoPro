package test

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/shonayevshyngys/prontopro/pkg/database"
)

func GetNewMiniRedis() TRedis {
	return TRedis{mRedis: miniredis.NewMiniRedis()}
}

type TRedis struct {
	mRedis         *miniredis.Miniredis
	redisInterface database.RedisInterface
}

func (t *TRedis) Set(key string, value []byte) {
	res := string(value[:])
	err := t.mRedis.Set(key, res)
	if err != nil {
		return
	}
}

func (t *TRedis) Get(key string) (string, error) {
	return t.mRedis.Get(key)
}

func (t *TRedis) RPush(key string, value []byte) (int64, error) {
	str := string(value[:])
	res, err := t.mRedis.RPush(key, str)
	return int64(res), err
}

func (t *TRedis) LRange(key string) ([]string, error) {
	return t.mRedis.List(key)
}

func (t *TRedis) Del(key string) {
	t.mRedis.Del(key)
}
