package core

import "io/ioutil"

type Challenge struct {
}

func (c *Challenge) GetData(path string) string {
	content, _ := ioutil.ReadFile(path)
	str := string(content)

	return str
}