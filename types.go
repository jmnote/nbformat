package notebookgo

type Dict = map[string]any

type Notebook struct {
	Cells         []Cell   `json:"cells"`
	Metadata      Metadata `json:"metadata"`
	Nbformat      int      `json:"nbformat"`
	NbformatMinor int      `json:"nbformat_minor"`
}

type Metadata struct {
	Kernelspec   *Kernelspec `json:"kernelspec,omitempty"`
	LanguageInfo *Dict       `json:"language_info,omitempty"`
}

type Kernelspec struct {
	DisplayName string `json:"display_name,omitempty"`
	Language    string `json:"language,omitempty"`
	Name        string `json:"name"`
}

type Cell struct {
	CellType       string   `json:"cell_type"`
	ExecutionCount *int     `json:"execution_count"`
	Metadata       Dict     `json:"metadata"`
	Outputs        []Output `json:"outputs"`
	Source         []string `json:"source"`
}

type Output struct {
	Name       string   `json:"name"`
	OutputType string   `json:"output_type"`
	Text       []string `json:"text"`
}
