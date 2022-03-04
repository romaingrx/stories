package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/andremueller/md-publisher/config"
	"github.com/andremueller/md-publisher/publisher"
)

func VerifyDomain(value string, values []string) error {
	for _, v := range values {
		if value == v {
			return nil
		}
	}
	return errors.New("Value " + value + " not in the restricted domain " + strings.Join(values, ", "))
}

func main() {
	var filename, token string
	flag.StringVar(&filename, "f", "", "The filename of the story. Accepted extensions: html, md. ")
	flag.StringVar(&token, "token", "", "The Medium access token.")
	flag.Parse()

	// Verify the extension of the filename
	var extension string
	extension = filepath.Ext(filename)
	if err := VerifyDomain(extension, []string{".md", ".html"}); err != nil {
		log.Fatalf(err.Error())
	}

	// Verify the Medium access token
	if token == "" {
		token = os.Getenv("MediumAccessToken")
	}
	if token == "" {
		log.Fatal("Couldn't retrieve the Medium access token. Either pass it as an argument or set the environment variable `MediumAccessToken`")
	}

	_, err := publisher.PublishMedium(filename, config.Config{
		NoImages:          false,
		MediumAccessToken: token,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

}
