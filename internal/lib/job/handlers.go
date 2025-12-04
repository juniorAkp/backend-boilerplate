package job

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/juniorAkp/backend-boilerplate/internal/config"
	"github.com/juniorAkp/backend-boilerplate/internal/lib/email"
	"github.com/rs/zerolog"
)

var emailClient *email.Client

func (sj *JobService) InitHandlers(config *config.Config, logger *zerolog.Logger) {
	emailClient = email.NewClient(config, logger)
}

func (js *JobService) handleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var p WelcomeEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal welcome email payload: %w", err)
	}

	js.logger.Info().
		Str("type", "welcome").
		Str("to", p.To).
		Msg("Processing welcome email task")

	err := emailClient.SendWelcomeEmail(
		p.To,
		p.FirstName,
	)
	if err != nil {
		js.logger.Error().
			Str("type", "welcome").
			Str("to", p.To).
			Err(err).
			Msg("Failed to send welcome email")
		return err
	}

	js.logger.Info().
		Str("type", "welcome").
		Str("to", p.To).
		Msg("Successfully sent welcome email")
	return nil
}
