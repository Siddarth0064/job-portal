package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	model "job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RDBLayer struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache
type Caching interface {
	AddToCache(ctx context.Context, jid uint, jobData model.Job) error
	GetCacheData(ctx context.Context, jid uint) (string, error)
	//AddtocacheOTP(ctx context.Context, s string, data model.ForgetPass) error
	AddtocacheOTP(ctx context.Context, email string, otp string) error
	GetCacheDataOTP(ctx context.Context, s string) (string, error)
	DeleteCacheotp(ctx context.Context, email string) error
}

func NewRDBLayer(rdb *redis.Client) Caching {
	// if rdb == nil {
	// 	//fmt.Errorf("error in creating the instance of RDB ")
	// 	return nil
	// }
	return &RDBLayer{
		rdb: rdb,
	}
}
func (r *RDBLayer) AddToCache(ctx context.Context, jid uint, jobData model.Job) error {

	jobID := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)

	if err != nil {

		return fmt.Errorf("error in Addtocache jobdata : %w ", err)
	}

	err = r.rdb.Set(ctx, jobID, val, 1*time.Minute).Err()
	return err
}
func (r *RDBLayer) GetCacheData(ctx context.Context, jid uint) (string, error) {

	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := r.rdb.Get(ctx, jobID).Result()

	// if err != nil {
	// 	return "", fmt.Errorf("error in Get Cache : %w", err)
	// }
	return str, err
}

// func (r *RDBLayer) AddtocacheOTP(ctx context.Context,email string, otp string) error {
// 	//jobID := strconv.FormatUint(uint64(jid), 10)
// 	val, err := json.Marshal(data)

// 	if err != nil {

// 		return fmt.Errorf("error in Add to cache otp : %w ", err)
// 	}

// 	err = r.rdb.Set(ctx, otp, val, 3*time.Minute).Err()
// 	return err
// }
// func (r *RDBLayer) GetCacheDataOTP(ctx context.Context, otp string) (string, error) {

// 	//jobID := strconv.FormatUint(uint64(jid), 10)
// 	str, err := r.rdb.Get(ctx, otp).Result()

//		if err != nil {
//			return "", fmt.Errorf("error in Get Cache otp : %w", err)
//		}
//		return str, err
//	}
func (r *RDBLayer) AddtocacheOTP(ctx context.Context, email string, otp string) error {
	err := r.rdb.Set(ctx, email, otp, 3*time.Minute).Err()
	if err != nil {
		log.Error().Msg("error in adding  the otp in cache ")
		return err
	}
	return err
}
func (r *RDBLayer) GetCacheDataOTP(ctx context.Context, email string) (string, error) {
	s, err := r.rdb.Get(ctx, email).Result()
	if err != nil {
		return "", errors.New("error in getting the otp from the cache")
	}
	return s, nil
}

func (r *RDBLayer) DeleteCacheotp(ctx context.Context, email string) error {
	err := r.rdb.Del(ctx, email).Err()
	if err != nil {
		return errors.New("error in deleting cache otp")
	}
	return nil
}
