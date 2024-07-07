package nbgo

type StringMap = map[string]any

type Notebook struct {
	Cells         []Cell   `json:"cells"`
	Metadata      Metadata `json:"metadata"`
	Nbformat      int      `json:"nbformat"`
	NbformatMinor int      `json:"nbformat_minor"`
}

type Metadata struct {
	CellToolbar  string      `json:"celltoolbar,omitempty"`
	Kernelspec   *Kernelspec `json:"kernelspec,omitempty"`
	LanguageInfo StringMap   `json:"language_info,omitempty"`
	Signature    string      `json:"signature,omitempty"`
	Title        string      `json:"title,omitempty"`
	Widgets      StringMap   `json:"widgets,omitempty"`
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
	Attachments    StringMap `json:"attachments,omitempty"`
	CellType       CellType  `json:"cell_type"`
	ExecutionCount *int      `json:"execution_count,omitempty"`
	ID             string    `json:"id,omitempty"`
	Metadata       StringMap `json:"metadata"`
	Outputs        []Output  `json:"outputs"`
	Source         []string  `json:"source"`
}

type CellType string

const (
	CellTypeMarkdown CellType = "markdown"
	CellTypeCode     CellType = "code"
	CellTypeRaw      CellType = "raw"
)

type Output struct {
	Data           StringMap  `json:"data,omitempty"`
	ExecutionCount *int       `json:"execution_count,omitempty"`
	Metadata       *StringMap `json:"metadata,omitempty"`
	Name           string     `json:"name,omitempty"`
	OutputType     OutputType `json:"output_type,omitempty"`
	Text           []string   `json:"text,omitempty"`
}

type OutputType string

const (
	OutputTypeDisplayData   OutputType = "display_data"
	OutputTypeExecuteResult OutputType = "execute_result"
	OutputTypeStream        OutputType = "stream"
	OutputTypeError         OutputType = "error"
)
