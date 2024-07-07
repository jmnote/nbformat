package nbconvertfiles

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jmnote/notebook-go/nbgo"
	"github.com/stretchr/testify/require"
)

type (
	Cell       = nbgo.Cell
	Dict       = nbgo.Dict
	Kernelspec = nbgo.Kernelspec
	Metadata   = nbgo.Metadata
	Notebook   = nbgo.Notebook
	Output     = nbgo.Output
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

func TestLossless(t *testing.T) {
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
			want = strings.ReplaceAll(want, `"execution_count": null,`, "")

			// WORKAROUND: optional but not omitempty for outputs (see types.go)
			// Even if the original JSON does not have the `outputs` key or the value of `outputs` is null,
			// the result will always have `"outputs": {}`.
			got := string(gotBytes)
			got = strings.ReplaceAll(got, `"outputs":null,`, "")

			require.JSONEq(t, want, got)
		})
	}
}
