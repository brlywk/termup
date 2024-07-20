package config

import "github.com/BurntSushi/toml"

type Data struct {
	Usage string  `toml:"usage"`
	Git   DataGit `toml:"git"`
	Ui    DataUi  `toml:"ui"`
}

type DataUi struct {
	Theme string `toml:"theme"`
	Font  string `toml:"font"`
}

type DataGit struct {
	DefaultBranch string `toml:"default_branch"`
	Name          string `toml:"name"`
	Email         string `toml:"email"`
}

// Parse reads config Data from toml file in path.
func Parse(path string) (*Data, error) {
	var data Data
	_, err := toml.DecodeFile(path, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
