# QuoteDoPai

A Command Line Interface (CLI) version of @guilhermebaos's quotedopai, that gives you the wisdom of **Nunes "O Pai" da Silva** insite your terminal.

## Tech Stack
* Languages: Golang
* Libraries:
	* Cobra: Used to create the CLI
	* HasGo: Functional programming in Go inspired by Haskell
* Tools:
	* Git
	* Make
	* CookieCutter

## CLI

### Commands

* `not [word]`
	* Filter for a quote not containing the provided word

* `with [word]`:
	* Filter for a quote containing the provided word

* `and [word1] [word2]`:
	* Filter for a quote containing `word1 && word2`

* `or [word1] [word2]`:
	* Filter for a quote containing `word1 || word2`

* `xor [word1] [word2]`:
	* Filter for a quote containing `word1 ^ word2`

### Methods Table

| Name         | Command | Short Flag | Long Flag | Type   | Explanation                   |
|--------------|---------|------------|-----------|--------|-------------------------------|
| Quote        |   ---   | -q         | --quote   | Output | Show a quote    [Quote]       |
| Copy         |   ---   | -c         | --copy    | Output | Copy a quote                  |
| Image        |   ---   | -i         | --image   | Output | Get an insta ready image      |
|              |         |            |           |        |                               |
| Not          | not     |    ----    |    ---    | Filter | `Not` containing `word`       |
| With         | with    |    ----    |    ---    | Filter | Containing `word`             |
| And          | and     |    ----    |    ---    | Filter | Containing `w1 && w2`         |
| Or           | or      |    ----    |    ---    | Filter | Containing `w1 \|\| w2`       |
| Exclusive Or | xor     |    ----    |    ---    | Filter | Containing `w1 ^ w2`          |
|              |         |            |           |        |                               |
| Form         | form    |    ----    |    ---    | Other  | Open a quote suggestion form  |
| Version      | version |    ----    |    ---    | Other  | Version information           |
| Help         | help    |    ----    |    ---    | Other  | Show help and command info    |

## Socials

## Getting started

This project requires Go to be installed.

Running it then should be as simple as:

```console
$ make
$ ./bin/quotedopai
```

### Testing

``make test``
