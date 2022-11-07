# Emdien

Search and read [MDN](https://developer.mozilla.org/en-US/) locally in the terminal.
Similar to [`mdn-cli`](https://www.npmjs.com/package/mdn-cli), but works without an internet connection.

## Installation

You can download a precompiled binary from the [releases page](https://github.com/Hasnep/emdien/releases), rename the file to `mdn` and place it in a location in your `PATH`, e.g. `$HOME/.local/bin/mdn`.

## Usage

Download and index the MDN data files:

```shell
mdn --update
```

Then, you can query MDN by running the `mdn` command, e.g. to search for "html":

```shell
mdn html
```

Search terms can be separated by spaces:

```shell
mdn css grid layout
```

You'll need to use quotation marks around terms containing special characters:

```shell
mdn '<br />' tag
```

To get raw markdown output instead of pretty rendered markdown, use the `--raw-output` or `-r` flag:

```shell
mdn --raw-output html
```

## Development

### Building

Install [Go 1.19](https://go.dev/dl/) or later.

Install the Go dependencies:

```shell
go get .
```

Build the binary:

```shell
go build -o build/mdn
```

### Pre-commit

Install [pre-commit](https://pre-commit.com/#install), then enable the pre-commit hooks in the repo:

```shell
pre-commit install
```
