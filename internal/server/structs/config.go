package structs

import (
	"errors"
	"reflect"
)

type Config struct {
	Port               string
	SalesforceTargetID string
	SlackWebhookURL    string
}

func (c *Config) validate() error {
	v := reflect.ValueOf(c)
	for i := 0; i < v.Elem().NumField(); i++ {
		field := v.Elem().Field(i)
		if field.Kind() == reflect.String {
			if field.String() == "" {
				return errors.New("config field is empty")
			}
		}

	}
	return nil
}
func NewConfig(args ...string) (*Config, error) {
	// map the args to the config struct
	c := &Config{}
	if len(args) > 0 {
		for i, arg := range args {
			switch i {
			case 0:
				c.Port = arg
			case 1:
				c.SalesforceTargetID = arg
			case 2:
				c.SlackWebhookURL = arg
			default:
			}
		}
		err := c.validate()
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
