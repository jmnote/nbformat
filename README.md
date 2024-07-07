# Notebook Format (nbformat) For Go 
This package provides Notebook Format structs for Go developers.

## Goal
Enable lossless bidirectional conversion between `nbformat` and Go struct `Notebook`.

## Testing
Tested using ipynb json files generated with nbconvert and all the test ipynb files available in the [nbconvert](https://github.com/jupyter/nbconvert) repository.
See the `tests` directory for more details.

## Known issue
For optionals, should we use omitempty or not?
https://github.com/jmnote/notebook-go/issues/7
