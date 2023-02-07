# SerpsBot-CLI v1.0 by ToyBlackHat
[![Go](https://github.com/ToyBlackHat/serpsbot-cli/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ToyBlackHat/serpsbot-cli/actions/workflows/go.yml)
A command line (cli) tool, for [SerpsBot](https://www.serpsbot.com) service, written in Golang

## Limitations
At this moment command line tool supports only fetching keyword suggestions (search suggestions) from Google.
If you need additional features please [contact me](mailto:toyblackhat@pm.me), and send me some money ;)

## Usage:
Most basic query, get a list of suggestions for presented keywords (separated by ,):
```
./serpsbot-cli --keywords=<KEYWORD>,<KEYWORD2>,<KEYWORD3> --apikep=<YOUR API KEY>
```
All available command line options:
```
./serpsbot-cli --help
```

### Useful scenarios:

1. Process keyword list from input.txt and receive keyword extensions (suggestions) in output.txt. Results and input keywords are merged in output file.
    ```
    ./serpsbot-cli --apikey=<YOUR API KEY> --outputfile=output.txt --inputfile=input.txt --merge
    ```
2. Use --gl= and --hl== params to use desired language and location (Poland and polish language in this case).At first step use --keyword= to fetch google suggestions to your main keyword, and than in second step, use output file from first step as a input file for second step, to get even more keywords suggestions.
    ```
    ./serpsbot-cli --apikey=<YOUR API KEY> --keywords="mama madzi" --gl=PL --hl=pl --outputfile=mama_madzi.txt  
    ./serpsbot-cli --apikey=<YOUR API KEY> --gl=PL --hl=pl --outputfile=mama_madzi2.txt --inputfile=mama_madzi.txt --merge
    ```
