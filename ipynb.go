// Package ipynb implements the 3.x iPython Notebook format:
// http://ipython.org/ipython-doc/3/notebook/nbformat.html
package ipynb

// Old versions of notebook format might support more types of cells,
// but version 3.x supports only markdown and code:
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#cell-types,
const (
	Markdown = "markdown"
	Code     = "code"
)

// Cell is the basic units in a notebook.
//
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#markdown-cells
// and
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#code-cells
type Cell struct {
	CellType       string `json:"cell_type"`
	Metadata       `json:"metadata"`
	Source         []string  `json:"source"`
	Outputs        []*Output `json:"outputs,omitempty"`
	ExecutionCount int       `json:"execution_count,omitempty"`
}

// Metadata is a property of code cells.
//
// As shown in
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#markdown-cells,
// metadata of markdown cells is empty.
type Metadata struct {
	Collapsed bool `json:"collapsed,omitempty"`
	Deletable bool `json:"deletable,omitempty"`
	Editable  bool `json:"editable,omitempty"`
}

// Output exists only in code cells.
//
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#code-cell-outputs
type Output struct {
	Name       string   `json:"name,omitempty"`
	OutputType string   `json:"output_type,omitempty"`
	Text       []string `json:"text,omitempty"`
}

func newOutput() *Output {
	return &Output{
		Name:       "stdout",
		OutputType: "stream",
		Text:       []string{"\n"},
	}
}

// AddLine appends a line to the source field of a cell.
func (c *Cell) AddLine(line string) {
	c.Source = append(c.Source, line)
}

// Notebook represents a Jupyter notebook.
type Notebook struct {
	Cells            []*Cell `json:"cells"`
	NotebookMetadata `json:"metadata"`
	FormatMajor      int `json:"nbformat"`
	FormatMinor      int `json:"nbformat_minor"`
}

// NotebookMetadata is a property of Notebook.
//
// NotebookMetadata is defined in
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#metadata.
type NotebookMetadata struct {
	KernelSpec   `json:"kernelspec"`
	LanguageInfo `json:"language_info"`
}

// KernelSpec is a property of NotebookMetadata.
//
// KernelSpec is defined in
// http://ipython.org/ipython-doc/3/development/kernels.html#kernelspecs.
type KernelSpec struct {
	DisplayName string `json:"display_name"`
	Language    string `json:"language"`
	Name        string `json:"name"`
}

// LanguageInfo is a property of NotebookMetadata.
//
// LanguageInfo is described in
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#top-level-structure.
type LanguageInfo struct {
	CodeMirrorMode    `json:"codemirror_mode"`
	FileExtension     string `json:"file_extension"`
	MIMEType          string `json:"mimetype"`
	Name              string `json:"name"`
	NBConvertExporter string `json:"nbconvert_exporter"`
	PygmentsLexer     string `json:"pygments_lexer"`
	Version           string `json:"version"`
}

// CodeMirrorMode is described in
// http://ipython.org/ipython-doc/3/notebook/nbformat.html#top-level-structure.
type CodeMirrorMode struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// New creates a new Notebook struct and returns a pointer to it.
func New() *Notebook {
	return &Notebook{
		Cells: make([]*Cell, 0),

		NotebookMetadata: NotebookMetadata{
			KernelSpec: KernelSpec{
				DisplayName: "Python 3",
				Language:    "python",
				Name:        "python3",
			},
			LanguageInfo: LanguageInfo{
				CodeMirrorMode: CodeMirrorMode{
					Name:    "ipython",
					Version: 3,
				},
				FileExtension:     ".py",
				MIMEType:          "text/x-python",
				Name:              "python",
				NBConvertExporter: "python",
				PygmentsLexer:     "ipython3",
				Version:           "3.6.0",
			},
		},
		FormatMajor: 4,
		FormatMinor: 0,
	}
}

// AddCell add a new cell at the end of a notebook.
func (nb *Notebook) AddCell(cellType string) *Cell {
	cell := &Cell{
		CellType: cellType,
		Metadata: Metadata{
			Collapsed: false,
			Deletable: false,
			Editable:  cellType == Code,
		},
		Source: make([]string, 0),
	}
	if cellType == Code {
		cell.Outputs = []*Output{newOutput()}
		cell.ExecutionCount = 1
	}
	nb.Cells = append(nb.Cells, cell)
	return cell
}
