package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gudn/vkpredict"
	"github.com/gudn/vkpredict/pkg/match/preprocessed"
	"github.com/gudn/vkpredict/pkg/match/prev"
	"github.com/gudn/vkpredict/pkg/preprocessing/norm"
	"github.com/gudn/vkpredict/pkg/preprocessing/sequence"
	"github.com/gudn/vkpredict/pkg/revidx/revstore"
	"github.com/gudn/vkpredict/pkg/store/memory"
)

var prep = sequence.New(
	norm.Norm,
	strings.ToLower,
)
var matcher = preprocessed.New(
	prep,
	&prev.PRevMatcher{
		ReverseIndex: &revstore.RevStore{
			Store: memory.New(),
		},
		MinN: 3,
	},
)
var predictor = vkpredict.Predictor{
	Store:   memory.New(),
	Matcher: matcher,
}

func loadEntries(fname string) ([]string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func main() {
	limit := flag.Uint("limit", 5, "limit of results")
	entriesFile := flag.String("entries", "", "path to newline-separated entries")
	flag.Parse()
	entries, err := loadEntries(*entriesFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = predictor.Add(entries)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		query := scanner.Text()
		results, err := predictor.Predict(query, *limit)
		if err != nil {
			log.Println(err)
		} else {
			for i, r := range results {
				fmt.Printf("%v: %v\n", i+1, r)
			}
		}
		fmt.Print("> ")
	}
}
