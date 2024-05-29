package utils

import (
	"net/url"
)

func ExtractPathFromURL(inputURL string) (string, error) {
	// Parse the input URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Extract the path component
	path := parsedURL.Path

	return path, nil
}

func ExtractBaseFromURL(inputURL string) (string, error) {
	// Parse the input URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Extract the path component
	path := parsedURL.Scheme + "://" + parsedURL.Host

	return path, nil
}

