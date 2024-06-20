package factory

type RuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct{}

func (J JsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type YamlRuleConfigParser struct{}

func (Y YamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type RuleConfigParserFactory interface {
	CreateParser() RuleConfigParser
}

type YamlRuleConfigParserFactory struct{}

func (Y YamlRuleConfigParserFactory) CreateParser() RuleConfigParser {
	return YamlRuleConfigParser{}
}

type JsonRuleConfigParserFactory struct{}

func (J JsonRuleConfigParserFactory) CreateParser() RuleConfigParser {
	return JsonRuleConfigParser{}
}
