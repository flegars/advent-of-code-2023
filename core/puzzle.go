package core

import "io/ioutil"

type Puzzle struct {
}

func (p *Puzzle) GetData(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic("something bad append.")
	}

	return string(content)
}