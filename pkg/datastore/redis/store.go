package redis

import (
	"crypto/tls"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/datastore"
	"github.com/srijanaravali/spinnaker-servicebroker/pkg/service"
)

type store struct {
	redisClient  *redis.Client
	prefix       string
	instanceList string
	bindingList  string
}

// NewStore returns a new Redis-based implementation of the Store interface
func NewStore(config Config) (datastore.DataStore, error) {
	redisOpts := &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password:   config.RedisPassword,
		DB:         config.RedisDB,
		MaxRetries: 5,
	}
	if config.RedisEnableTLS {
		redisOpts.TLSConfig = &tls.Config{
			ServerName: config.RedisHost,
		}
	}
	return &store{
		redisClient:  redis.NewClient(redisOpts),
		prefix:       config.RedisPrefix,
		instanceList: wrapKey(config.RedisPrefix, "instances"),
		bindingList:  wrapKey(config.RedisPrefix, "bindings"),
	}, nil
}

func (s *store) WriteInstance(instance service.ServiceInstance) error {
	key := s.getInstanceKey(instance.ID)
	json, err := instance.ToJSON()
	if err != nil {
		return err
	}
	pipeline := s.redisClient.TxPipeline()
	pipeline.Set(key, json, 0)

	pipeline.SAdd(s.instanceList, key)
	_, err = pipeline.Exec()
	if err != nil {
		return fmt.Errorf(
			`error writing instance "%s": %s`,
			instance.ID,
			err,
		)
	}
	return err
}

func (s *store) GetInstance(instanceID string) (service.ServiceInstance, bool, error) {
	key := s.getInstanceKey(instanceID)
	strCmd := s.redisClient.Get(key)
	if err := strCmd.Err(); err == redis.Nil {
		return service.ServiceInstance{}, false, nil
	} else if err != nil {
		return service.ServiceInstance{}, false, err
	}
	bytes, err := strCmd.Bytes()
	if err != nil {
		return service.ServiceInstance{}, false, err
	}
	instance, err := service.NewInstanceFromJSON(bytes)
	if err != nil {
		return instance, false, err
	}
	return instance, err == nil, err
}

func (s *store) DeleteInstance(instanceID string) (bool, error) {
	instance, ok, err := s.GetInstance(instanceID)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	key := s.getInstanceKey(instanceID)
	pipeline := s.redisClient.TxPipeline()
	pipeline.Del(key)

	pipeline.SRem(s.instanceList, key)
	_, err = pipeline.Exec()
	if err != nil {
		return false, fmt.Errorf(
			`error deleting instance "%s": %s`,
			instance.ID,
			err,
		)
	}
	return true, nil
}

func (s *store) getInstanceKey(instanceID string) string {
	return wrapKey(s.prefix, fmt.Sprintf("instances:%s", instanceID))
}

func (s *store) TestConnection() error {
	return s.redisClient.Ping().Err()
}

func wrapKey(prefix, key string) string {
	if prefix != "" {
		return fmt.Sprintf("%s:%s", prefix, key)
	}
	return key
}
