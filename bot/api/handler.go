package api

import (
	"github.com/nndergunov/tgBot/bot/pkg/domain"
)

func (b *ChatBot) Handle() {
	b.bot.Listen()
	b.pool.Start()

	for message := range b.bot.MessageChan {
		currMsg := message

		b.log.LogMessage(currMsg)

		handleTask := func() {
			b.Respond(currMsg)
		}

		b.pool.AddTask(handleTask)
	}
}

func (b *ChatBot) Respond(message domain.ReceiveMessage) {
	answer := b.router.Route(message)

	if err := b.bot.Send(answer); err != nil {
		b.log.Println(err)
	}
}
