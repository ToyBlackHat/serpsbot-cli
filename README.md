# SerpsBot-CLI v1.0 by ToyBlackHat

A command line (cli) tool, for [SerpsBot](https://www.example.com) service, written in Golang

## Limitations
At this moment command line tool supports only fetching keyword suggestions (search suggestions) from Google.
If you need additional features please [contact me](mailto:toyblackhat@pm.me), and send me some money ;)

## Usage:
Most basic query, get a list of suggestions for presented keywords (separated by ,):
```
./serpsbot-cli --keywords=<KEYWORD>,<KEYWORD2>,<KEYWORD3> --apikep=<YOUR API KEY>
```

### Other useful scenarios:

1. Process keyword list from input.txt and receive keyword extensions (suggestions) in output.txt. Results and input keywords are merged in output file.
```
./serpsbot-cli --apikey=<YOUR API KEY> --outputfile=output.txt --inputfile=input.txt --merge
```