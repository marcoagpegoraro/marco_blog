package enum

type Language uint8

const (
	BrazilianPortuguese Language = iota + 1
	English
	Italian
)

func (l Language) String() string {
	switch l {
	case BrazilianPortuguese:
		return "BrazilianPortuguese"
	case English:
		return "English"
	case Italian:
		return "Italian"
	}

	return "unknown"
}

func LanguageEnumValues() map[uint8]string {

	var languages = map[uint8]string{
		uint8(BrazilianPortuguese): BrazilianPortuguese.String(),
		uint8(English):             English.String(),
		uint8(Italian):             Italian.String(),
	}
	return languages
}
