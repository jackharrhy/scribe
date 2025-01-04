package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/nint8835/scribe/pkg/database"
)

type AddArgs struct {
	Text   string         `description:"Text for the quote to add. To insert a new line, insert \\n."`
	Author discordgo.User `description:"Author of the quote."`
	Source *string        `description:"Link to a source for the quote, if available (such as a Discord message, screenshot, etc.)"`
}

func (b *Bot) AddQuoteCommand(_ *discordgo.Session, interaction *discordgo.InteractionCreate, args AddArgs) {
	quote := database.Quote{
		Text:    strings.Replace(args.Text, "\\n", "\n", -1),
		Authors: []*database.Author{{ID: args.Author.ID}},
		Source:  args.Source,
	}

	b.addQuote(quote, interaction)
}
