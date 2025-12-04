package job

import (
	"github.com/hibiken/asynq"
	"github.com/juniorAkp/backend-boilerplate/internal/config"
	"github.com/rs/zerolog"
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

func (js *JobService) Start() error {

	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskWelcome, js.handleWelcomeEmailTask)

	js.logger.Info().Msg("Starting Job Service...")

	if err := js.server.Start(mux); err != nil {
		js.logger.Error().Err(err).Msg("Could not start Job Service")
		return err
	}

	return nil
}

func (js *JobService) Stop() {
	js.logger.Info().Msg("Stopping Job Service...")
	js.server.Stop()
	js.Client.Close()
}
