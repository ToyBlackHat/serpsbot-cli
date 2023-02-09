# SerpsBot-CLI v1.0 by ToyBlackHat
[![Build](https://github.com/ToyBlackHat/serpsbot-cli/actions/workflows/build.yml/badge.svg)](https://github.com/ToyBlackHat/serpsbot-cli/actions/workflows/build.yml)

A command line (cli) tool for [SerpsBot](https://www.serpsbot.com) service, written in Golang

## Limitations
At this moment command line tool supports only fetching keyword suggestions (search suggestions) from Google.
If you need additional features please [contact me](mailto:toyblackhat@pm.me).

## Usage:
Most basic query, get a list of suggestions for presented keywords separated by ',' *(comma)*:
```
./serpsbot-cli --keywords=<KEYWORD>,<KEYWORD2>,<KEYWORD3> --apikey=<YOUR API KEY>
```
All available command line options:
```
./serpsbot-cli --help
```

## Installation
 Install golang on your platform by following [this instruction](https://go.dev/doc/install). Then build this project by 
```
go build
```
 or take a look at makefile and use `make` tool.

## Setup
You can write a Serpsbot's API code to a JSON config file by running the command:
```
serpsbot-cli --setup
```
The program will ask you for an API code and save it in your home directory in file `~/.serpsbot-cli.json`. From now, you'll no need to enter the API code in `--apikey` command line option.


## Useful scenarios:

1. Process keyword list from input.txt and receive keyword extensions (suggestions) in output.txt. Results and input keywords are merged in the output file.
    ```
    ./serpsbot-cli --apikey=<YOUR API KEY> --outputfile=output.txt --inputfile=input.txt --merge
    ```
2. Use --gl= and --hl== params to use desired language and location (Poland and polish language in this case). In the first step, use --keyword= to fetch google suggestions to your main keyword, and then in the second step, use the output file from the first step as an input file for the second step to get even more keyword suggestions.
    ```
    ./serpsbot-cli --apikey=<YOUR API KEY> --keywords="mama madzi" --gl=PL --hl=pl --outputfile=mama_madzi.txt  
    ./serpsbot-cli --apikey=<YOUR API KEY> --gl=PL --hl=pl --outputfile=mama_madzi2.txt --inputfile=mama_madzi.txt --merge
    ```
