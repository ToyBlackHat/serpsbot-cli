package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type SerpsbotSuggestions struct {
	Meta struct {
		Gl       string   `json:"gl"`
		Hl       string   `json:"hl"`
		Keywords []string `json:"keywords"`
	} `json:"meta"`
	Data []struct {
		Keyword     string   `json:"keyword"`
		Suggestions []string `json:"suggestions"`
	} `json:"data"`
}

type Serpsbot struct {
	Apikey string
}

func WriteStringToFile(filename string, line string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(line)
	if err != nil {
		return err
	}
	return nil
}

func SerpsGetSuggestions(keywords []string, apikey string, gl string, hl string) (SerpsbotSuggestions, error) {

	type SerpsbotQuery struct {
		Gl       string   `json:"gl"`
		Hl       string   `json:"hl"`
		Keywords []string `json:"keywords"`
	}

	type SerpsbotError struct {
		Detail []struct {
			Loc  []interface{} `json:"loc"`
			Msg  string        `json:"msg"`
			Type string        `json:"type"`
		} `json:"detail"`
	}

	api := SerpsbotQuery{
		Gl:       "us",
		Hl:       "en_US",
		Keywords: keywords,
	}

	var x SerpsbotSuggestions

	b, _ := json.Marshal(api)

	req, err := http.NewRequest("POST", "https://api.serpsbot.com/v2/google/search-suggestions", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return x, err
	}

	req.Header.Set("X-API-KEY", apikey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return x, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// probably 422 error, so decode the error
		var a SerpsbotError
		err = json.NewDecoder(res.Body).Decode(&a)
		if err != nil {
			log.Fatalln(err)
			return x, err
		}
		fmt.Println(a)
	} else {
		// status 200: decode the response and return it
		err = json.NewDecoder(res.Body).Decode(&x)
		if err != nil {
			log.Fatalln(err)
			return x, err
		}
		return x, nil
	}
	return x, errors.New("general error")
}

func main() {
	fmt.Println("SerpsBot v1.0 by ToyBlackHat")

	keywords_cli := flag.String("keywords", "", "Keywords to get suggestions from Google, separated by ',' (comma)")
	apikey_cli := flag.String("apikey", "", "Serpsbot API Key")
	outputfile_cli := flag.String("outputfile", "", "Output file to save results to")
	inputfile_cli := flag.String("inputfile", "", "Input file with keywords to get suggestions from Google")
	gl_cli := flag.String("gl", "US", "Geographic Location. The ISO code of the country for which you want to get the suggestions for")
	hl_cli := flag.String("hl", "en_US", "Home Language. The language to get the suggestions for.")
	merge_cli := flag.Bool("merge", false, "Add keywords from input to the output (result) file")

	setup_cli := flag.Bool("setup", false, "Setup Serpsbot API Key")

	flag.Parse()

	//open file ~/.serpsbot-cli.json and read apikey from there

	var sbot Serpsbot
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error obtaining $HOME directory:", err)
	}
	jsonFile, err := os.ReadFile(homedir + "/.serpsbot-cli.json")
	if err == nil {
		// open file ~/.serpsbot-cli.json and read apikey
		fmt.Println("Reading API key from ~/.serpsbot-cli.json")
		err = json.Unmarshal(jsonFile, &sbot)
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
		// fmt.Println(sbot.Apikey)
	}

	// if setup flag is set, ask for apikey and save it to ~/.serpsbot-cli.json
	if *setup_cli {
		fmt.Println("Please enter your API key:")
		_, _ = fmt.Scan(&sbot.Apikey)

		// save to ~/.serpsbot-cli.json
		b, err := json.Marshal(sbot)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
		err = os.WriteFile(homedir+"/.serpsbot-cli.json", b, 0o644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
		fmt.Println("API key saved to ~/.serpsbot-cli.json")
		os.Exit(0)
	}

	if *apikey_cli == "" {
		*apikey_cli = sbot.Apikey
	}

	if *keywords_cli == "" && *inputfile_cli == "" {
		fmt.Println("Please specify a keyword(s) for --keywords= or a file with keywords for --inputfile=")
		flag.Usage()
		os.Exit(1)
	}
	if *keywords_cli != "" && *inputfile_cli != "" {
		fmt.Println("You can only select one mode --keywords= or --inputfile=")
		flag.Usage()
		os.Exit(1)
	}
	if *apikey_cli == "" {
		fmt.Println("Please enter your Serpsbot apikey into --apikey=")
		flag.Usage()
		os.Exit(1)
	}

	var keywords []string

	if *keywords_cli != "" {
		// Process keywords from command line --keywords=
		keywords = strings.Split(*keywords_cli, ",")
	}

	if *inputfile_cli != "" {
		// Process keywords from file --inputfile=
		file, err := os.Open(*inputfile_cli)
		if err != nil {
			fmt.Println("error opening file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			keywords = append(keywords, strings.TrimSpace(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading file:", err)
			os.Exit(1)
		}
	}

	for i := 0; i < len(keywords); i += 50 {
		end := i + 50
		if end > len(keywords) {
			end = len(keywords)
		}
		keywords50 := keywords[i:end]

		fmt.Println("Processing", len(keywords50), "keywords, gl=", *gl_cli, "hl=", *hl_cli, "")
		result, err := SerpsGetSuggestions(keywords50, *apikey_cli, *gl_cli, *hl_cli)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, v := range result.Data {
			fmt.Println(v.Keyword)
			if *merge_cli {
				if *outputfile_cli != "" {
					err = WriteStringToFile(*outputfile_cli, v.Keyword+"\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
			for _, v2 := range v.Suggestions {
				fmt.Println("-", v2)
				if *outputfile_cli != "" {
					err = WriteStringToFile(*outputfile_cli, v2+"\n")
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
				}
			}
		}
	}

}
