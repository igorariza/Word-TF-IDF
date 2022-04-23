package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/igorariza/tfidf/utils"
)

// Returns a map of words and their respective term frequency in the document
// where the document is specified by the index of the record in the recordArray
func termFrequency(recordArray [][]string, threshold float64) (m map[string]int, err error) {
	saveMap := make(map[string]map[string]int)
	for _, record := range recordArray {
		url := record[0]

		if _, ok := saveMap[url]; ok {
			continue
		}
		words := utils.LowercaseWords(strings.Fields(record[0]))

		for i := range words {
			w, err := utils.RemoveNonAlphaNumeric(words[i])
			if err != nil {
				continue
			} else {
				words[i] = w
			}
		}

		words, err = utils.RemoveStopwords(words)
		if err != nil {
			return nil, err
		}

		saveMap[url] = utils.WordFrequency(words)
	}

	documentFrequencyMap := make(map[string]int)

	for _, wordCountMap := range saveMap {
		for word := range wordCountMap {
			if _, ok := documentFrequencyMap[word]; ok {
				documentFrequencyMap[word]++
			} else {
				documentFrequencyMap[word] = 1
			}
		}
	}

	if threshold != 0.0 {
		for word, value := range documentFrequencyMap {
			if float64(value)/float64(len(saveMap)) < threshold {
				delete(documentFrequencyMap, word)
			}
		}
	}
	return documentFrequencyMap, nil
}

// Inverse Document Frequency
// returns a map of words and their respective inverse document frequency
func inverseDocumentFrequency(recordArray [][]string) (m map[string]float64, err error) {
	d := float64(len(recordArray))

	wordCountMap := make(map[string]int)
	for _, record := range recordArray {
		words := utils.LowercaseWords(strings.Fields(record[0]))

		for i := range words {
			w, err := utils.RemoveNonAlphaNumeric(words[i])
			if err != nil {
				continue
			} else {
				words[i] = w
			}
		}

		words, err = utils.RemoveStopwords(words)
		if err != nil {
			return nil, err
		}

		words = utils.RemoveDuplicates(words)

		for _, word := range words {
			if _, ok := wordCountMap[word]; ok {
				wordCountMap[word]++
			} else {
				wordCountMap[word] = 1
			}
		}
	}

	idfMap := make(map[string]float64)
	for word, value := range wordCountMap {
		idfMap[word] = math.Log(d / float64(value))
	}
	return idfMap, nil
}

// Term Frequency-Inverse Document Frequency (TF-IDF) is a statistical measure of the importance of a word in a document.
// how important a word is to a document in a collection or corpus. It is
func termFrequencyInverseDocumentFrequency(fileName string) (m map[string]float64, err error) {
	recordArray, err := utils.ReadRecords(fileName)
	if err != nil {
		return nil, err
	}

	tfidfMap := make(map[string]float64)

	tf, err := termFrequency(recordArray, 0.0)
	if err != nil {
		return nil, err
	}

	idf, err := inverseDocumentFrequency(recordArray)
	if err != nil {
		return nil, err
	}

	for word, docFreq := range idf {
		tfidfMap[word] = float64(tf[word]) * docFreq
	}
	return tfidfMap, nil
}

func main() {
	word := flag.String("word", "", "Word to search")
	flag.Parse()
	tfidfMap, err := termFrequencyInverseDocumentFrequency("data/document_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		ok := false
		if word == nil || *word == "" {
			for k, v := range tfidfMap {
				ok = true
				fmt.Println("the TF-IDF weight of the word ", k, " in the document is ", "->", v)
			}

		} else {
			for k, v := range tfidfMap {
				fq := fmt.Sprintf(*word)
				if strings.ToLower(fq) == k {
					ok = true
					fmt.Println("the TF-IDF weight of the word ", k, " in the document is ", "->", v)
				}
			}
		}
		//If the word is not found in the document
		if !ok {
			fmt.Println("the word it is not in the document")
		}
	}
}
