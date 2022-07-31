package configs

import "embed"

//go:embed conf.toml
var EmbedConf embed.FS
