package main

import (
	"encoding/json"
	"fmt"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/ipynb"
)

func main() {
	nb := ipynb.New()
	c := nb.AddCell(ipynb.Markdown)
	c.AddLine("# Hello")
	c = nb.AddCell(ipynb.Code)
	c.AddLine("print(\"Hello Yi!\")")
	b, e := json.MarshalIndent(nb, "", "  ")
	candy.Must(e)
	fmt.Printf("%s\n", b)
}
