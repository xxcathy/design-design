package factory

type IParser interface {
	Parse(data []byte)
}

type jsonParser struct{}

func (p *jsonParser) Parse(data []byte) {}

type tomlParser struct{}

func (p *tomlParser) Parse(data []byte) {}

/*************************** 简单工厂 *****************************/
func NewParser(name string) IParser {
	switch name {
	case "json":
		return &jsonParser{}
	case "toml":
		return &tomlParser{}
	}
	return nil
}

/*************************** 工厂方法 *****************************/
var factories = map[string]IParserFactory{ // map
	"json": &jsonParserFactory{},
	"toml": &tomlParserFactory{},
}

func NewParserFactory(name string) IParserFactory {
	if p, ok := factories[name]; ok {
		return p
	}
	return nil
}

type IParserFactory interface {
	CreateParser() IParser
}

type jsonParserFactory struct{}

func (f *jsonParserFactory) CreateParser() IParser { return &jsonParser{} }

type tomlParserFactory struct{}

func (f *tomlParserFactory) CreateParser() IParser { return &tomlParser{} }


/*************************** 抽象工厂 *****************************/
var configParserFactories = map[string]IConfigParserFactory{
	"json": &jsonConfigParserFactory{},
	"toml": &tomlConfigParserFactory{},
}

func NewRule1ConfigParserFactory(name string, ruleType int) IRule1Parser {
	if p, ok := configParserFactories[name]; ok {
		if ruleType == 1 {
			return p.CreateRule1Parser()
		}
	}
	return nil
}

func NewRule2ConfigParserFactory(name string, ruleType int) IRule2Parser {
	if p, ok := configParserFactories[name]; ok {
		if ruleType == 1 {
			return p.CreateRule2Parser()
		}
	}
	return nil
}

type IRule1Parser interface {
	Parse(data []byte)
}
type IRule2Parser interface {
	Parse(data []byte)
}
type IConfigParserFactory interface {
	CreateRule1Parser() IRule1Parser
	CreateRule2Parser() IRule2Parser
}

type jsonConfigParserFactory struct{}
func (f *jsonConfigParserFactory) CreateRule1Parser() IRule1Parser { return &jsonParser{} }
func (f *jsonConfigParserFactory) CreateRule2Parser() IRule2Parser { return &jsonParser{} }

type tomlConfigParserFactory struct{}
func (f *tomlConfigParserFactory) CreateRule1Parser() IRule1Parser { return &tomlParser{} }
func (f *tomlConfigParserFactory) CreateRule2Parser() IRule2Parser { return &tomlParser{} }

