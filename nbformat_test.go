package nbformat

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	files []string
)

func init() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".ipynb" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestRemarshalFiles(t *testing.T) {
	for _, filePath := range files {
		t.Run(filePath, func(t *testing.T) {
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			var nb Notebook
			err = json.Unmarshal(fileBytes, &nb)
			require.NoError(t, err)
			gotBytes, err := json.Marshal(nb)
			require.NoError(t, err)

			// WORKAROUND: optional and omitempty. However, it may exist in JSON (see types.go).
			// If there are keys in the original JSON and the values ​​for them are empty strings ("") or null,
			// the keys and values ​​will be missing in the result due to omitempty.
			want := string(fileBytes)
			want = strings.ReplaceAll(want, `"language": "",`, "")
			want = strings.ReplaceAll(want, `     "execution_count": null,`, ``)
			want = strings.ReplaceAll(want, `"execution_count": null,`, `"execution_count": 0,`)

			require.JSONEq(t, want, string(gotBytes))
		})
	}
}

func TestGetCellType(t *testing.T) {
	testCases := []struct {
		cell         Cell
		wantCellType CellType
	}{
		{&CodeCell{}, CellTypeCode},
		{&MarkdownCell{}, CellTypeMarkdown},
		{&RawCell{}, CellTypeRaw},
	}
	for _, tc := range testCases {
		t.Run(string(tc.wantCellType), func(t *testing.T) {
			got := tc.cell.GetCellType()
			require.Equal(t, tc.wantCellType, got)
		})
	}
}
