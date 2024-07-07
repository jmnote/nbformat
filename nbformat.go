package nbformat

import (
	"encoding/json"
	"fmt"
)

// covers schema v4.0 - v4.5
// https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.0.schema.json
// https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.5.schema.json

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
	CellToolbar  string      `json:"celltoolbar,omitempty"`
	Kernelspec   *Kernelspec `json:"kernelspec,omitempty"`
	LanguageInfo StringMap   `json:"language_info,omitempty"`
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

// v4.0 - v4.4
// raw_cell      "required": [      "cell_type", "metadata", "source"]
// markdown_cell "required": [      "cell_type", "metadata", "source"]
// code_cell     "required": [      "cell_type", "metadata", "source", "outputs", "execution_count"]
//
// v4.5
// raw_cell      "required": ["id", "cell_type", "metadata", "source"]
// markdown_cell "required": ["id", "cell_type", "metadata", "source"]
// code_cell     "required": ["id", "cell_type", "metadata", "source", "outputs", "execution_count"]
//
// "id" will not be set to "omitempty" to maintain backward compatibility.

type Cell interface {
	GetCellType() CellType
}

type RawCell struct {
	ID          string    `json:"id,omitempty"`
	CellType    CellType  `json:"cell_type"`
	Metadata    StringMap `json:"metadata"`
	Source      []string  `json:"source"`
	Attachments StringMap `json:"attachments,omitempty"`
}

type MarkdownCell struct {
	ID          string    `json:"id,omitempty"`
	CellType    string    `json:"cell_type"`
	Metadata    StringMap `json:"metadata"`
	Source      []string  `json:"source"`
	Attachments StringMap `json:"attachments,omitempty"`
}

type CodeCell struct {
	ID             string    `json:"id,omitempty"`
	CellType       CellType  `json:"cell_type"`
	Metadata       StringMap `json:"metadata"`
	Source         []string  `json:"source"`
	Outputs        []Output  `json:"outputs"`
	ExecutionCount int       `json:"execution_count"`
	Attachments    StringMap `json:"attachments,omitempty"`
}

func (c *RawCell) GetCellType() CellType {
	return CellTypeRaw
}

func (c *MarkdownCell) GetCellType() CellType {
	return CellTypeMarkdown
}

func (c *CodeCell) GetCellType() CellType {
	return CellTypeCode
}

// v4.0 - v4.5
// execute_result "required": ["output_type", "data", "metadata", "execution_count"]
// display_data   "required": ["output_type", "data", "metadata"]
// stream         "required": ["output_type", "name", "text"]
// error          "required": ["output_type", "ename", "evalue", "traceback"]
//
// The output_type field is the only required field across different OutputTypes,
// and all other fields are tagged with omitempty,
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

// UnmarshalJSON custom unmarshal function for Notebook
func (n *Notebook) UnmarshalJSON(data []byte) error {
	var aux struct {
		Metadata      Metadata          `json:"metadata"`
		NbformatMinor int               `json:"nbformat_minor"`
		Nbformat      int               `json:"nbformat"`
		Cells         []json.RawMessage `json:"cells"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	n.Metadata = aux.Metadata
	n.NbformatMinor = aux.NbformatMinor
	n.Nbformat = aux.Nbformat

	for _, rawCell := range aux.Cells {
		var cellType struct {
			CellType string `json:"cell_type"`
		}
		if err := json.Unmarshal(rawCell, &cellType); err != nil {
			return err
		}

		var cell Cell
		switch cellType.CellType {
		case "raw":
			var rawContent RawCell
			if err := json.Unmarshal(rawCell, &rawContent); err != nil {
				return err
			}
			cell = &rawContent
		case "markdown":
			var markdownCell MarkdownCell
			if err := json.Unmarshal(rawCell, &markdownCell); err != nil {
				return err
			}
			cell = &markdownCell
		case "code":
			var codeCell CodeCell
			if err := json.Unmarshal(rawCell, &codeCell); err != nil {
				return err
			}
			cell = &codeCell
		default:
			return fmt.Errorf("unknown cell type: %s", cellType.CellType)
		}
		n.Cells = append(n.Cells, cell)
	}
	return nil
}
