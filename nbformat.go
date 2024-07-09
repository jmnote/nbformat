package nbformat

import "encoding/json"

type (
	CellType   string
	OutputType string
	StringMap  = map[string]any
)

const (
	CellTypeRaw      CellType = "raw"
	CellTypeMarkdown CellType = "markdown"
	CellTypeCode     CellType = "code"

	OutputTypeDisplayData   OutputType = "display_data"
	OutputTypeExecuteResult OutputType = "execute_result"
	OutputTypeStream        OutputType = "stream"
	OutputTypeError         OutputType = "error"
)

// "required": ["metadata", "nbformat_minor", "nbformat", "cells"]
type Notebook struct {
	Metadata      Metadata `json:"metadata"`
	NbformatMinor int      `json:"nbformat_minor"`
	Nbformat      int      `json:"nbformat"`
	Cells         []Cell   `json:"cells"`
}

// no required
type Metadata struct {
	Authors      []StringMap `json:"authors,omitempty"`
	CellToolbar  string      `json:"celltoolbar,omitempty"`
	Kernelspec   *Kernelspec `json:"kernelspec,omitempty"`
	LanguageInfo StringMap   `json:"language_info,omitempty"`
	RecordTiming *bool       `json:"record_timing,omitempty"`
	Signature    string      `json:"signature,omitempty"`
	Title        string      `json:"title,omitempty"`
	Widgets      StringMap   `json:"widgets,omitempty"`
}

// "required": ["name", "display_name"]
type Kernelspec struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Language    string `json:"language,omitempty"`
}

// For nbformat versions v4.0 to v4.4:
// A raw      cell must include the fields: ["cell_type", "metadata", "source"]
// A markdown cell must include the fields: ["cell_type", "metadata", "source"]
// A code     cell must include the fields: ["cell_type", "metadata", "source", "outputs", "execution_count"]
//
// Starting from nbformat version v4.5:
// A raw      cell must include the fields: ["id", "cell_type", "metadata", "source"]
// A markdown cell must include the fields: ["id", "cell_type", "metadata", "source"]
// A code     cell must include the fields: ["id", "cell_type", "metadata", "source", "outputs", "execution_count"]
//
// This means that in version v4.5, the "id" field is added as a requirement for all types of cells, in addition to the previously required fields.
// However, the "id" field will be set to "omitempty" to maintain backward compatibility.
type Cell struct {
	ID             string    `json:"id,omitempty"`
	CellType       CellType  `json:"cell_type"`
	Metadata       StringMap `json:"metadata"`
	Source         []string  `json:"source"`
	Outputs        []Output  `json:"outputs,omitempty"`
	ExecutionCount int       `json:"execution_count,omitempty"`
	Attachments    StringMap `json:"attachments,omitempty"`
}

func (c Cell) MarshalJSON() ([]byte, error) {
	type Alias Cell
	aux := struct {
		*Alias
		Outputs        *[]Output `json:"outputs,omitempty"`
		ExecutionCount *int      `json:"execution_count,omitempty"`
	}{
		Alias: (*Alias)(&c),
	}
	if c.CellType == CellTypeCode {
		aux.Outputs = &c.Outputs
		aux.ExecutionCount = &c.ExecutionCount
	}
	return json.Marshal(aux)
}

// For nbformat versions v4.0 to v4.5:
// A execute_result output must include the fields: ["output_type", "data", "metadata", "execution_count"]
// A display_data   output must include the fields: ["output_type", "data", "metadata"]
// A stream         output must include the fields: ["output_type", "name", "text"]
// A error          output must include the fields: ["output_type", "ename", "evalue", "traceback"]
// The output_type field is the only required field across different OutputTypes,
// and all other fields are tagged with omitempty.
type Output struct {
	OutputType     OutputType `json:"output_type"`
	Data           StringMap  `json:"data,omitempty"`
	Metadata       *StringMap `json:"metadata,omitempty"`
	ExecutionCount *int       `json:"execution_count,omitempty"`
	Name           string     `json:"name,omitempty"`
	Text           []string   `json:"text,omitempty"`
	Ename          string     `json:"ename,omitempty"`
	Evalue         string     `json:"evalue,omitempty"`
	Traceback      []string   `json:"traceback,omitempty"`
}
