# Notebook Format (nbformat) For Go 

## Goal
Enable lossless bidirectional conversion between `nbformat` and Go struct `Notebook`.

## Tested
Tested with `[nbconvert](https://github.com/jupyter/nbconvert)` ipynb files. See `tests` directory.

## Known issue
For optionals, should we use omitempty or not?
https://github.com/jmnote/notebook-go/issues/7
