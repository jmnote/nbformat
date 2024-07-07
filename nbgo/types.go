package nbgo

type Dict = map[string]any

type Notebook struct {
	Cells         []Cell   `json:"cells"`
	Metadata      Metadata `json:"metadata"`
	Nbformat      int      `json:"nbformat"`
	NbformatMinor int      `json:"nbformat_minor"`
}

type Metadata struct {
	CellToolbar  string      `json:"celltoolbar,omitempty"`
	Kernelspec   *Kernelspec `json:"kernelspec,omitempty"`
	LanguageInfo Dict        `json:"language_info,omitempty"`
	Signature    string      `json:"signature,omitempty"`
	Title        string      `json:"title,omitempty"`
	Widgets      Dict        `json:"widgets,omitempty"`
}

// omitempty issue
// omitable: language
// https://raw.githubusercontent.com/jupyter/nbconvert/v7.16.4/tests/exporters/files/pngmetadata.ipynb
type Kernelspec struct {
	DisplayName string `json:"display_name"`
	Language    string `json:"language,omitempty"`
	Name        string `json:"name"`
}

// omitempty issue
// omitable: execution_count, outputs
// However, sometimes we also want the outputs to be empty slices(`outputs: {}`). So, it will not be omitted.
// https://nbformat.readthedocs.io/en/latest/format_description.html#cell-types
// TODO: we need omitnil https://github.com/golang/go/issues/22480
type Cell struct {
	Attachments    Dict     `json:"attachments,omitempty"`
	CellType       string   `json:"cell_type"`
	ExecutionCount *int     `json:"execution_count,omitempty"`
	ID             string   `json:"id,omitempty"`
	Metadata       Dict     `json:"metadata"`
	Outputs        []Output `json:"outputs"`
	Source         []string `json:"source"`
}

type Output struct {
	Data           Dict     `json:"data,omitempty"`
	ExecutionCount *int     `json:"execution_count,omitempty"`
	Metadata       *Dict    `json:"metadata,omitempty"`
	Name           string   `json:"name,omitempty"`
	OutputType     string   `json:"output_type,omitempty"`
	Text           []string `json:"text,omitempty"`
}
