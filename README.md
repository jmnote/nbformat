# Notebook Format (nbformat) for Go

[![main](https://github.com/jmnote/nbformat/actions/workflows/main.yml/badge.svg)](https://github.com/jmnote/nbformat/actions/workflows/main.yml)
[![pull-request](https://github.com/jmnote/nbformat/actions/workflows/pull-request.yml/badge.svg)](https://github.com/jmnote/nbformat/actions/workflows/pull-request.yml)
[![Coverage Status](https://coveralls.io/repos/github/jmnote/nbformat/badge.svg?branch=main)](https://coveralls.io/github/jmnote/nbformat?branch=main)
[![GitHub license](https://img.shields.io/github/license/jmnote/nbformat.svg)](https://github.com/jmnote/nbformat/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jmnote/nbformat)](https://goreportcard.com/report/github.com/jmnote/nbformat)
[![GitHub stars](https://img.shields.io/github/stars/jmnote/nbformat.svg)](https://github.com/jmnote/nbformat/stargazers)

This package provides Notebook Format structs for Go developers.
It currently supports nbformat versions from [v4.0](https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.0.schema.json) to [v4.5](https://github.com/jupyter/nbformat/blob/v5.10.4/nbformat/v4/nbformat.v4.5.schema.json).

## Goal
To enable lossless bidirectional conversion between `nbformat` and Go struct `Notebook`.

## Testing
Tested using ipynb files generated by [nbconvert](https://github.com/jupyter/nbconvert), as well as all valid ipynb files from the [nbconvert](https://github.com/jupyter/nbconvert) and [nbformat](https://github.com/jupyter/nbformat) repositories (excluding v2 and v3 files). The ipynb files are located in the `tests` directory.
