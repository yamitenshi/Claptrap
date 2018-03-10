package claptrap

import (
	"github.com/DSchalla/Claptrap/rules"
	"log"
)

type BotServer struct {
	config       Config
	mattermostHandler *MattermostHandler
	eventHandler *EventHandler
	ruleEngine   *rules.Engine
}

func NewBotServer(config Config) *BotServer {
	b := BotServer{}
	b.config = config
	b.mattermostHandler = NewMattermostHandler(config.ApiUrl, config.Username, config.Password, config.Team)
	b.ruleEngine = rules.NewEngine(b.config.CaseDir)
	return &b
}

func (b *BotServer) Start() {
	log.Println("[+] Claptrap BotServer starting")
	b.mattermostHandler.StartWS()

	if b.config.AutoJoinAllChannel {
		b.mattermostHandler.AutoJoinAllChannel()
	}
	b.eventHandler = NewEventHandler(b.mattermostHandler.GetMessages(), b.ruleEngine)
	respHandler := NewMattermostResponseHandler(b.mattermostHandler.Client, b.mattermostHandler.BotUser)
	b.ruleEngine.SetResponseHandler(respHandler)
	b.ruleEngine.Start()
	b.eventHandler.Start()
}

func (b *BotServer) AddCase(caseType string, newCase rules.Case) {
	b.ruleEngine.AddCase(caseType, newCase)
	log.Printf("[+] Dynamic Case '%s' with type '%s' loaded\n", newCase.Name, caseType)
}