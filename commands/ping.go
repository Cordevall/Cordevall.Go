package commands

import (
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct{}

func (pc *PingCommand) Name() string {
	return "ping"
}

func (pc *PingCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}