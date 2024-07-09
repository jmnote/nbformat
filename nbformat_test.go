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
	json1 string
	json2 string
	json3 string
	json4 string

	notebook1 Notebook
	notebook2 Notebook
	notebook3 Notebook
	notebook4 Notebook
)

func init() {
	json1 = readFile("tests/myfiles/python1.ipynb")
	json2 = readFile("tests/myfiles/python2.ipynb")
	json3 = readFile("tests/myfiles/r1.ipynb")
	json4 = readFile("tests/myfiles/r2.ipynb")
	setNotebooks()
}

func readFile(filePath string) string {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}

func setNotebooks() {
	notebook1 = Notebook{
		Cells: []Cell{
			{
				CellType: "code",
				Source:   []string{},
				Metadata: StringMap{},
				Outputs:  []Output{},
			},
		},
		Metadata: Metadata{
			Kernelspec: &Kernelspec{
				DisplayName: "Python 3",
				Language:    "Python",
				Name:        "python3",
			},
			LanguageInfo: StringMap{
				"codemirror_mode": StringMap{
					"name":    "ipython",
					"version": float64(3),
				},
				"file_extension":     ".py",
				"mimetype":           "text/x-python",
				"name":               "python",
				"nbconvert_exporter": "python",
				"pygments_lexer":     "ipython3",
				"version":            "3.10.1",
			},
		},
		Nbformat:      4,
		NbformatMinor: 4,
	}
	notebook2 = Notebook{
		Cells: []Cell{
			{
				CellType:       "code",
				ExecutionCount: 1,
				Metadata: StringMap{
					"execution": StringMap{
						"iopub.execute_input": "2024-07-06T15:48:16.194233Z",
						"iopub.status.busy":   "2024-07-06T15:48:16.190279Z",
						"iopub.status.idle":   "2024-07-06T15:48:16.197696Z",
						"shell.execute_reply": "2024-07-06T15:48:16.198094Z",
					},
				},
				Outputs: []Output{
					{
						Name:       "stdout",
						OutputType: "stream",
						Text:       []string{"hello\n"},
					},
				},
				Source: []string{"print('hello')"},
			}},
		Metadata: Metadata{
			Kernelspec: &Kernelspec{
				DisplayName: "Python 3",
				Language:    "Python",
				Name:        "python3",
			},
			LanguageInfo: StringMap{
				"codemirror_mode": StringMap{
					"name":    "ipython",
					"version": float64(3),
				},
				"file_extension":     ".py",
				"mimetype":           "text/x-python",
				"name":               "python",
				"nbconvert_exporter": "python",
				"pygments_lexer":     "ipython3",
				"version":            "3.10.1",
			},
		},
		Nbformat:      4,
		NbformatMinor: 4,
	}
	notebook3 = Notebook{
		Cells: []Cell{
			{
				CellType:       "code",
				ExecutionCount: 0,
				Metadata:       StringMap{},
				Outputs:        []Output{},
				Source:         []string{},
			},
		},
		Metadata: Metadata{
			Kernelspec: &Kernelspec{
				DisplayName: "R",
				Language:    "R",
				Name:        "ir",
			},
			LanguageInfo: StringMap{
				"codemirror_mode": "r",
				"file_extension":  ".r",
				"mimetype":        "text/x-r-source",
				"name":            "R",
				"pygments_lexer":  "r",
				"version":         "4.1.2",
			},
		},
		Nbformat:      4,
		NbformatMinor: 4,
	}
	notebook4 = Notebook{
		Cells: []Cell{
			{
				CellType:       "code",
				ExecutionCount: 1,
				Metadata: StringMap{
					"execution": StringMap{
						"iopub.execute_input": "2024-07-07T05:32:29.525611Z",
						"iopub.status.busy":   "2024-07-07T05:32:29.523594Z",
						"iopub.status.idle":   "2024-07-07T05:32:29.546156Z",
						"shell.execute_reply": "2024-07-07T05:32:29.545217Z",
					},
				},
				Outputs: []Output{
					{
						Name:       "stdout",
						OutputType: "stream",
						Text:       []string{"hello"},
					},
				},
				Source: []string{"cat('hello')"},
			},
		},
		Metadata: Metadata{
			Kernelspec: &Kernelspec{
				DisplayName: "R",
				Language:    "R",
				Name:        "ir",
			},
			LanguageInfo: StringMap{
				"codemirror_mode": "r",
				"file_extension":  ".r",
				"mimetype":        "text/x-r-source",
				"name":            "R",
				"pygments_lexer":  "r",
				"version":         "4.1.2",
			},
		},
		Nbformat:      4,
		NbformatMinor: 4,
	}
}

func TestToNotebook(t *testing.T) {
	type TestCase struct {
		name         string
		input        string
		wantNotebook Notebook
		wantError    string
	}
	testCases := []TestCase{
		{
			"empty string",
			"",
			Notebook{},
			"unexpected end of JSON input",
		},
		{
			"empty StringMapionary",
			"{}",
			Notebook{},
			"",
		},
		{
			"json1",
			json1,
			notebook1,
			"",
		},
		{
			"json2",
			json2,
			notebook2,
			"",
		},
		{
			"json3",
			json3,
			notebook3,
			"",
		},
		{
			"json4",
			json4,
			notebook4,
			"",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got Notebook
			err := json.Unmarshal([]byte(tc.input), &got)
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantError)
			}
			require.Equal(t, tc.wantNotebook, got)
		})
	}
}

func TestToJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    Notebook
		wantJSON string
	}{
		{
			"zero notebook",
			Notebook{},
			`{"cells":null,"metadata":{},"nbformat":0,"nbformat_minor":0}`,
		},
		{
			"notebook1",
			notebook1,
			json1,
		},
		{
			"notebook2",
			notebook2,
			json2,
		},
		{
			"notebook3",
			notebook3,
			json3,
		},
		{
			"notebook4",
			notebook4,
			json4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.input)
			require.NoError(t, err)
			require.JSONEq(t, tc.wantJSON, string(got))
		})
	}
}

func TestRemarshalFiles(t *testing.T) {
	for _, filePath := range getAllFiles() {
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

func getAllFiles() []string {
	var filePaths []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".ipynb" {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return filePaths
}
