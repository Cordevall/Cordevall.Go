package main

import (
	"github.com/bwmarrin/discordgo"
    "github.com/MikeModder/anpan"
)

type Command interface {
	Name() string
	Execute(session *discordgo.Session, msg *discordgo.MessageCreate)
}

var commandMap = make(map[string]Command)

func main() {
	dg, _ := discordgo.New("Bot " + "your-token-goes-here")
	dg.AddHandler(messageCreate)
	dg.Open()
	defer dg.Close()

	loadCommands()

	// Keep the program running until interrupted
	select {}
}

func loadCommands(commands) {
	log.Println("This is a simple log message.")
	files, _ := ioutil.ReadDir("./commands")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") {
			// Assuming each file only contains one command
			commandName := strings.TrimSuffix(f.Name(), ".go")
			switch commandName {
			case "ping":
				commandMap[commandName] = &commands.PingCommand{}
			// Add cases for other commands here
			}
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if command, ok := commandMap[m.Content]; ok {
		command.Execute(s, m)
	}
}