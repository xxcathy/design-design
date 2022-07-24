package builder

import "errors"

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

/*************************** 通用实现 *****************************/
type ResourcePoolConfigBuilder struct {
	config *ResourcePoolConfig
}

func NewResourcePoolConfigBuilder() *ResourcePoolConfigBuilder {
	return &ResourcePoolConfigBuilder{config: &ResourcePoolConfig{}}
}

func (builder *ResourcePoolConfigBuilder) SetName(name string) *ResourcePoolConfigBuilder {
	builder.config.name = name
	return builder
}

func (builder *ResourcePoolConfigBuilder) SetMaxIdle(idle int) *ResourcePoolConfigBuilder {
	builder.config.maxIdle = idle
	return builder
}

func (builder *ResourcePoolConfigBuilder) SetMinIdle(idle int) *ResourcePoolConfigBuilder {
	builder.config.minIdle = idle
	return builder
}

func (builder *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if len(builder.config.name) == 0 {
		return nil, errors.New("invalid name")
	}
	if builder.config.minIdle < 0 || builder.config.maxIdle < 0 || builder.config.minIdle > builder.config.maxTotal {
		return nil, errors.New("invalid idle")
	}
	return builder.config, nil
}

//func main() {
//	config, err := NewResourcePoolConfigBuilder().
//		SetName("name").
//		SetMaxIdle(20).
//		Build()
//}

/*************************** golang 实现 *****************************/
type ResourcePoolConfigOptFunc func(conf *ResourcePoolConfig)

func NewResourcePoolConfig(options ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	conf := &ResourcePoolConfig{}
	for _, option := range options {
		option(conf)
	}
	if len(conf.name) == 0 {
		return nil, errors.New("invalid name")
	}
	if conf.minIdle < 0 || conf.maxIdle < 0 || conf.minIdle > conf.maxTotal {
		return nil, errors.New("invalid idle")
	}
	return conf, nil
}

func SetConfigName(name string) ResourcePoolConfigOptFunc {
	return func(conf *ResourcePoolConfig) {
		conf.name = name
	}
}

func SetConfigMaxIdle(idle int) ResourcePoolConfigOptFunc {
	return func(conf *ResourcePoolConfig) {
		conf.maxIdle = idle
	}
}

func SetConfigMinIdle(idle int) ResourcePoolConfigOptFunc {
	return func(conf *ResourcePoolConfig) {
		conf.minIdle = idle
	}
}

//func main() {
//	conf, err := NewResourcePoolConfig(
//		SetConfigName("name"),
//		SetConfigMaxIdle(20),
//	)
//}
