package options

import "golang.org/x/net/context"

type Person struct {
	name string
	age  int
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

type Config struct {
	name string
	age  int
}

type Option func(config *Config) error

func WithName(name string) Option {
	return func(config *Config) error {
		config.name = name
		return nil
	}
}

func WithAge(age int) Option {
	return func(config *Config) error {
		config.age = age
		return nil
	}
}

func (c *Config) ApplyOptions(options ...Option) error {
	for _, opt := range options {
		err := opt(c)
		if err != nil {
			return err
		}
	}
	return nil
}

var DefaultName Option = WithName("default")
var DefaultAge Option = WithAge(20)

var defaults = []struct {
	needDefault func(cfg *Config) bool
	opt         Option
}{
	{
		needDefault: func(cfg *Config) bool { return cfg.name == "" },
		opt:         DefaultName,
	},
	{
		needDefault: func(cfg *Config) bool { return cfg.age == 0 },
		opt:         DefaultAge,
	},
}

var FallBackDefault Option = func(config *Config) error {
	for _, def := range defaults {
		if !def.needDefault(config) {
			continue
		}
		if err := def.opt(config); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) NewPerson() *Person {
	person := &Person{
		name: c.name,
		age:  c.age,
	}
	return person
}

func NewWithDefault(ctx context.Context, options ...Option) (*Person, error) {
	var cfg Config
	for _, opt := range options {
		err := cfg.ApplyOptions(opt)
		if err != nil {
			return nil, err
		}
	}
	return cfg.NewPerson(), nil
}
