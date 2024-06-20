package factory

func NewRuleConfigParser(str string) RuleConfigParser {
	switch str {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	}
	return nil
}