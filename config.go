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

func ActionConfigs() Action {
	return ActionMultiParts("=", func(c Context) Action {
		switch len(c.Parts) {
		case 0:
			return ActionMultiParts(".", func(c Context) Action {
				switch len(c.Parts) {
				case 0:
					return ActionValues(config.GetConfigs()...).Invoke(c).Suffix(".").ToA()
				case 1:
					fields, err := config.GetConfigFields(c.Parts[0])
					if err != nil {
						return ActionMessage(err.Error())
					}
					return ActionStyledValuesDescribed(fields...).Invoke(c).Suffix("=").ToA()
				default:
					return ActionValues()
				}
			})
		case 1:
			return ActionValues()
		default:
			return ActionValues()
		}
	})
}
