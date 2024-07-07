# Notebook Format (nbformat) for Go 
This package provides Notebook Format structs for Go developers.

## Goal
To enable lossless bidirectional conversion between `nbformat` and Go struct `Notebook`.

## Testing
Tested using ipynb JSON files generated with nbconvert, as well as all the test ipynb files available in the [nbconvert](https://github.com/jupyter/nbconvert) repository.
For more details, see the `tests` directory.

## Known Issue
Should we use `omitempty` for optionals or not? 
https://github.com/jmnote/notebook-go/issues/7
