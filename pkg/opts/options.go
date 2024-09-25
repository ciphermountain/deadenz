package opts

type LanguageSetter interface {
	SetLanguage(string)
}

type Option func(any)

func WithLanguage(lang string) Option {
	return func(os any) {
		if val, ok := os.(LanguageSetter); ok {
			val.SetLanguage(lang)
		}
	}
}
