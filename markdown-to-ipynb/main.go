package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/ipynb"
)

func main() {
	nb := ipynb.New()
	c := nb.AddCell(ipynb.Markdown)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if c.CellType == ipynb.Markdown && strings.HasPrefix(strings.TrimSpace(line), "```python") {
			c = nb.AddCell(ipynb.Code)
		} else if c.CellType == ipynb.Code && strings.HasPrefix(strings.TrimSpace(line), "```") {
			c = nb.AddCell(ipynb.Markdown)
		} else {
			c.AddLine(line + "\n")
		}
	}
	candy.Must(scanner.Err())

	b, e := json.MarshalIndent(nb, "", "  ")
	candy.Must(e)
	fmt.Printf("%s\n", b)
}
