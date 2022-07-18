package carapace

import (
	"github.com/rsteube/carapace/internal/config"
)

type carapaceConfig struct {
	DescriptionLength int
}

func (c *carapaceConfig) Completion() ActionMap {
	return ActionMap{
		"DescriptionLength": ActionValues("40", "30"),
	}
}

var conf = carapaceConfig{
	DescriptionLength: 40,
}

func init() {
	config.RegisterConfig("carapace", &conf)
}
