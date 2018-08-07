[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/micnncim/medicine)](https://goreportcard.com/report/github.com/micnncim/medicine)

<div align="center">
    <img src="https://user-images.githubusercontent.com/21333876/43464345-ddb4b5be-9515-11e8-9de8-0ef5741aa658.png">
</div>
<br>

Come on, [Medium](medium.com) code blocks do not support syntax highlighting... Although we can use syntax highlight by [GitHubGist](gist.github.com), it's painful when we have many code blocks.

`medicine` automatically converts code blocks with markdown in [Medium](medium.com) into [GitHubGist](gist.github.com) embedded links.

## Table of Contents

<!-- TOC -->

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Author](#author)
- [License](#license)

<!-- /TOC -->

## Description

Look at an example of `medicine`.

Convert the code block in your Medium article

```sample.go
func main() {
    fmt.Println("Hello, World!")
}
```

into the embedded Gist!

<img width="700" alt="screenshot 2018-07-31 11 22 44" src="https://user-images.githubusercontent.com/21333876/43433844-1340e89a-94b4-11e8-9486-5630ff1c5702.png">


## Installation

```
$ go get github.com/micnncim/medicine
```

## Usage

1. Get [GitHub Access Token](https://github.com/settings/tokens) and [Medium Access Token](https://medium.com/me/settings).
2. Prepare a Markdown file. See the following rules.
  - The first line should be Medium article title.
  - Code blocks require a filename. See the [example](https://raw.githubusercontent.com/micnncim/medicine/master/example.md).
3. Run `medicine <MARKDOWN_FILE>` . If you run it for the first time, it requires your access tokens and saves the tokens in `~/.config/medicine/config.toml` .
4. A few seconds later, your browser opens your Medium article from the markdown file as a draft. Edit and publish it!

## Example

Clone this repository and `cd` the path. Then, just run it!

```
$ medicine example.md
```

## Author

[@micnncim](https://twitter.com/micnncim)

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmicnncim%2Fmedicine.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmicnncim%2Fmedicine?ref=badge_large)
