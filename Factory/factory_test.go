package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want RuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: JsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: YamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRuleConfigParserFactory(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want RuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{str: "json"},
			want: JsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{str: "yaml"},
			want: YamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleConfigParserFactory(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestJsonConfigParserFactoryCreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		want SystemConfigParser
	}{
		{
			name: "json",
			want: JsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonConfigParserFactoryCreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want RuleConfigParser
	}{
		{
			name: "json",
			want: JsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
