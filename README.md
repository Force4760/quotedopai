# QuoteDoPai

A Command Line Interface (CLI) version of @guilhermebaos's quotedopai, that gives you the wisdom of **Nunes "O Pai" da Silva** insite your terminal.

## Tech Stack
* Languages: Golang
* Libraries:
	* Cobra:
	* Vyper:
	* Logrus:
	* Flinch:
* Tools:
	* Git
	* JSON
	* Make

## CLI
### Output Methods

* `-q` or `--quote`:
	* Shows a random quote

* `-c` or `--copy`:
	* Copies a random quote to your clipboard

* `-i` or `--image`:
	* Creates an Instagram ready image containing a random quote

* `-p` or `--print`:
	* Prints a random quote to your terminal

### Filtering Methods

* `-n [word]` or `--not [word]`:
	* Filter for a quote not containing the provided word

* `-w [word]` or `--with [word]`:
	* Filter for a quote containing the provided word

* `-a [word1] [word2]` or `--and [word1] [word2]`:
	* Filter for a quote containing `word1 && word2`

* `-o [word1] [word2]` or `--or [word1] [word2]`:
	* Filter for a quote containing `word1 || word2`

* `-x [word1] [word2]` or `--xor [word1] [word2]`:
	* Filter for a quote containing `word1 ^ word2`

### Methods Table

| Name         | Short Flag | Long Flag | Arguments | Type   | Explanation                   |
|--------------|------------|-----------|-----------|--------|-------------------------------|
| Quote        | -q         | --quote   |           | Output | Show a quote                  |
| Copy         | -c         | --copy    |           | Output | Copy a quote                  |
| Image        | -i         | --image   | size      | Output | sizeXsize image with a quote  |
| Print        | -p         | --print   |           | Output | Print a quote                 |
| Not          | -n         | --not     | word      | Filter | Not containing word           |
| With         | -w         | --with    | word      | Filter | Containing word               |
| And          | -a         | --and     | w1, w2    | Filter | Containing w1 && w2           |
| Or           | -o         | --or      | w1, w2    | Filter | Containing w1 \|\| w2         |
| Exclusive Or | -x         | --xor     | w1, w2    | Filter | Containing w1 ^ w2            |

## Contributing
If you have had the pleasure of having "O Pai" as a professor, please message me with any quotes you have containing his ancient wisdom that will allow us to have good grades on his exams (learning experiences I mean...).

## Socials

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/quotedopai
```

### Testing

``make test``
