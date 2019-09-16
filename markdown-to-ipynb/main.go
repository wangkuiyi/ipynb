package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/topicai/candy"
	"github.com/wangkuiyi/ipynb"
)

func isCodeBlock(line string, codeBlockType string) bool {
	if strings.HasPrefix(strings.TrimSpace(line), fmt.Sprintf("```%s", codeBlockType)) {
		return true
	}
	return false
}

func main() {
	codeBlockType := flag.String("code-block-type", "python", "the code block type which used to generate code block in ipynb.")
	flag.Parse()

	nb := ipynb.New()
	c := nb.AddCell(ipynb.Markdown)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if c.CellType == ipynb.Markdown && isCodeBlock(line, *codeBlockType) {
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
