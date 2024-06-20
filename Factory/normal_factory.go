package factory

func NewRuleConfigParserFactory(str string) RuleConfigParserFactory {
	switch str {
	case "json":
		return JsonRuleConfigParserFactory{}
	case "yaml":
		return YamlRuleConfigParserFactory{}
	}
	return nil
}
