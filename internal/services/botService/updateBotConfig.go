package botService

import (
	"bot_forge_back/internal/services/tgService"
	"bot_forge_back/internal/states"
	"encoding/json"
)

func (s Service) UpdateBotConfig(conf states.BotStates, id string) error {
	marshalledJSON, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	botToken, err := s.r.GetBotTokenByID(id)

	err = s.r.CreateBotConfig(id, marshalledJSON)
	if err != nil {
		return err
	}

	err = tgService.RegNewWebHook(botToken, id)
	if err != nil {
		return err
	}
	return nil
}
