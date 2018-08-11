package eager

import (
	"github.com/BurntSushi/toml"
)

type Parser interface {
	ParseConfig(path string, config interface{}) error
}

type TOMLParser struct {
	MetaData toml.MetaData
}

// parseTOML
func (parser TOMLParser) ParseConfig(path string, config interface{}) error {
	var err error
	parser.MetaData, err = toml.DecodeFile(path, config)
	if err != nil {
		return newError("decode file error", err)
	}
	return nil
}
