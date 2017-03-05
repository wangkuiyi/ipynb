`ipynb` is a Go package for
reading/writing
[Jupyter Notebook files](http://ipython.org/ipython-doc/3/notebook/nbformat.html).
The document is [here](https://godoc.org/github.com/wangkuiyi/ipynb).

`markdown-to-ipynb` is a command line tool that converts Markdown
files (`.md`) into Jupyter Notebooks using `ipynb`.

To build and use `markdown-to-ipynb`:

1. Download and install Go: https://golang.org/doc/install
1. Make sure that you have Git installed.
1. Checkout and install `markdown-to-ipynb`:
   ```bash
   export GOPATH=...any where you like...
   go get -u github.com/wangkuiyi/ipynb/markdown-to-ipynb`
   ```
1. Convert an example Markdown file:
   ```bash
   GOPATH/bin/markdown-to-ipynb example.md > example.ipynb
   ```
1. (Optionally) view the Notebook file:
   ```bash
   jupyter notebook example.ipynb
   ```
