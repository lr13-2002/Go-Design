package factory

type SystemConfigParser interface {
	ParseSystem(data []byte)
}

type JsonSystemConfigParser struct{}

func (J JsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type ConfigParserFactory interface {
	CreateRuleParser() RuleConfigParser
	CreateSystemParser() SystemConfigParser
}

type JsonConfigParserFactory struct{}

func (J JsonConfigParserFactory) CreateRuleParser() RuleConfigParser {
	return JsonRuleConfigParser{}
}
func (J JsonConfigParserFactory) CreateSystemParser() SystemConfigParser {
	return JsonSystemConfigParser{}
}
