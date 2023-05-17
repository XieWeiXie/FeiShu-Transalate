package command

type Commander interface {
	Command() error
}

const (
	CMDHelp      = "/help"
	CMDWord      = "/word"
	CMDSentence  = "/sentence"
	CMDTranslate = "/translate"
)

type HelpCommander struct {
	des string
}

type WordCommander struct {
	des string
}

type SentenceCommander struct {
	des string
}

type TranslateCommander struct {
	des string
}
