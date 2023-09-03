package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/samber/lo"
	tele "gopkg.in/telebot.v3"
)

func (s *Service) initSysHandlers(allCommands string) {
	s.bot.Handle("/help", func(c tele.Context) error {
		return c.Send("Вот, что я могу:\n" + allCommands)
	})
	s.bot.Handle("/start", func(c tele.Context) error {
		return c.Send(
			"Привет! Я могу помочь тебе управлять твоими финансами 😎\n" +
				"Просто отправь мне команду /help, чтобы узнать, что я могу сделать",
		)
	})

}

func genCommands(commands *map[string]string) string {
	return strings.Join(
		lo.MapToSlice(*commands, func(text, description string) string {
			return fmt.Sprintf("%s - %s", text, description)
		}), "\n")
}

func (s *Service) initCommands() (string, error) {
	defer s.logger.Infof("Bot has %d commands", len(s.endpoints))

	{
		// find present in bot but not currently registered commands, delete them
		currentCommands, err := s.bot.Commands()
		if err != nil {
			return "", err
		}

		var commandsToDelete []interface{}
		lo.ForEach(currentCommands, func(command tele.Command, _ int) {
			cmdWithPrefix := "/" + command.Text
			if _, ok := s.endpoints[cmdWithPrefix]; !ok {
				s.logger.Warnf("Command %s not registered", command.Text)
				commandsToDelete = append(commandsToDelete, command)
			} else {
				// command registered
				if command.Description != s.endpoints[cmdWithPrefix] {
					s.logger.Warnf("Command %s description mismatch: %s != %s", command.Text, command.Description, s.endpoints[cmdWithPrefix])
					commandsToDelete = append(commandsToDelete, command)
				}
				s.logger.Infof("Command %s changed description", cmdWithPrefix)
				s.endpoints[cmdWithPrefix] = command.Description
			}
		})

		// delete unused or invalid commands
		if len(commandsToDelete) > 0 {
			err := s.bot.DeleteCommands(commandsToDelete...)
			if err != nil {
				return "", err
			}
		}
	}

	// add new and changed commands
	err := s.bot.SetCommands(lo.MapToSlice(s.endpoints, func(text, description string) interface{} {
		text = strings.TrimPrefix(text, "/")
		if ok, err := regexp.MatchString("^[a-z0-9_]+$", text); !ok || err != nil {
			s.logger.Panicf("Invalid command format - matched pattern: %v; err: %v", ok, err)
		}
		return tele.Command{
			Text:        text,
			Description: description,
		}
	})...)
	if err != nil {
		return "", err
	}
	return genCommands(&s.endpoints), nil
}
