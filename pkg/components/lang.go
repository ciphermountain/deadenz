package components

const (
	DefaultFindPattern        = "you find {{item}}"
	DefaultSpawninPattern     = "you spawned in as a {{character}}"
	DefaultEarnedXPPattern    = "you earned {{amount}} xp"
	DefaultEarnedTokenPattern = "you earned {{amount}} tokens"
	DefaultEncounterPattern   = "you encounter {{encounter}}"
)

type LanguagePack struct {
	FindPattern        string
	SpawninPattern     string
	EarnedXPPattern    string
	EarnedTokenPattern string
	EncounterPattern   string
}

var defaultLanguagePack = LanguagePack{
	FindPattern:        DefaultFindPattern,
	SpawninPattern:     DefaultSpawninPattern,
	EarnedXPPattern:    DefaultEarnedXPPattern,
	EarnedTokenPattern: DefaultEarnedTokenPattern,
	EncounterPattern:   DefaultEncounterPattern,
}

var eventLangPatterns = map[string]LanguagePack{
	"en": defaultLanguagePack,
	"es": {
		FindPattern:        "encuentras {{item}}",
		SpawninPattern:     "naces como {{character}}",
		EarnedXPPattern:    "te ganaste {{amount}} xp",
		EarnedTokenPattern: "te ganaste {{amount}} fichas",
		EncounterPattern:   "encuentras con {{encounter}}",
	},
}

type language struct {
	lang LanguagePack
}

func (l *language) SetLanguage(lang string) {
	pack, ok := eventLangPatterns[lang]
	if !ok {
		l.lang = defaultLanguagePack
	}

	l.lang = pack
}
