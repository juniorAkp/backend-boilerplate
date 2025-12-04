package job

import (
	"github.com/hibiken/asynq"
	zerolog "github.com/jackc/pgx-zerolog"
	"github.com/juniorAkp/backend-boilerplate/internal/config"
)

type JobService struct {
	Client *asynq.Client
	server *asynq.Server
	logger *zerolog.Logger
}

func NewJobService(lg *zerolog.Logger, cfg *config.Config) *JobService {
	redisAddr := cfg.Redis.Address

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			//abstract the concurrency and queues to config file later
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	return &JobService{
		Client: client,
		server: server,
		logger: lg,
	}
}
