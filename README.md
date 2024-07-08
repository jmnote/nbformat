# Notebook Format (nbformat) for Go 
This package provides Notebook Format structs for Go developers.
It currently supports nbformat versions from v4.0 to v4.5.
- https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.0.schema.json
- https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.5.schema.json

## Goal
To enable lossless bidirectional conversion between `nbformat` and Go struct `Notebook`.

## Testing
Tested using ipynb JSON files generated with nbconvert, as well as all the test ipynb files available in the [nbconvert](https://github.com/jupyter/nbconvert) repository.
For more details, see the `tests` directory.
