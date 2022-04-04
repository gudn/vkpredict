package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gudn/vkpredict"
	"github.com/gudn/vkpredict/pkg/match"
	"github.com/gudn/vkpredict/pkg/match/builder"
	"github.com/gudn/vkpredict/pkg/match/compose"
	"github.com/gudn/vkpredict/pkg/match/lcs"
	"github.com/gudn/vkpredict/pkg/match/preprocessed"
	"github.com/gudn/vkpredict/pkg/match/prev"
	"github.com/gudn/vkpredict/pkg/preprocessing/norm"
	"github.com/gudn/vkpredict/pkg/preprocessing/sequence"
	"github.com/gudn/vkpredict/pkg/preprocessing/stopwords"
	"github.com/gudn/vkpredict/pkg/revidx/revstore"
	"github.com/gudn/vkpredict/pkg/store"
	"github.com/gudn/vkpredict/pkg/store/level"
	"github.com/gudn/vkpredict/pkg/store/memory"
	"github.com/gudn/vkpredict/pkg/store/unique"
	"github.com/syndtr/goleveldb/leveldb"
)

// Дефолтный набор препроцессинга
var prep = sequence.New(
	norm.Norm,
	strings.ToLower,
	stopwords.Stopwords,
	strings.TrimSpace,
)

// Собираем предиктор с указанным файлом базы данных. Если база данных пустая,
// будет использована `store/memory`
func makePredictor(dbname string) *vkpredict.Predictor {
	var db *leveldb.DB
	if dbname != "" {
		var err error
		db, err = leveldb.OpenFile(dbname, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	makeStore := func (prefix string) store.IterAnyStore {
		if db == nil {
			return memory.New()
		}
		return unique.New(level.New(prefix, db))
	}
	var matcher = preprocessed.New(
		prep,
		&compose.ComposeMatcher{
			Matchers: []match.Matcher{
				&prev.PRevMatcher{
					ReverseIndex: &revstore.RevStore{
						Store: makeStore("revstore"),
					},
					MinN: 3,
				},
				&builder.BuilderMatcher{
					Builder: lcs.BuildScorer,
					IterAnyStore: makeStore("lcs"),
				},
			},
			Coefs: []uint{3, 1},
		},
	)
	return &vkpredict.Predictor{
		Store: makeStore("predictor"),
		Matcher: matcher,
	}
}


func loadStrings(fname string) ([]string, error) {
	if fname == "" {
		return nil, nil
	}
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
	loadFile := flag.String("load", "", "path to newline-separated entries")
	qFile := flag.String("queries", "", "path to newline-separated queries")
	dbname := flag.String("db", "", "path to leveldb file")
	flag.Parse()
	entries, err := loadStrings(*loadFile)
	if err != nil {
		log.Fatalln(err)
	}
	predictor := makePredictor(*dbname)
	err = predictor.Add(entries)
	if err != nil {
		log.Fatalln(err)
	}
	qs, err := loadStrings(*qFile)
	if err != nil {
		log.Fatalln(err)
	}
	for _, query := range qs {
		fmt.Printf("Query %q:\n", query)
		results, err := predictor.Predict(query, *limit)
		if err != nil {
			log.Println(err)
		} else {
			for i, r := range results {
				fmt.Printf("%v: %v\n", i+1, r)
			}
		}
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
