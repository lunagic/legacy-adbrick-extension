package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no arg provided")
	}
	newVersionNumber := os.Args[1]

	if err := ChangeFile(
		"apple/AdBrick/AdBrick.xcodeproj/project.pbxproj",
		func(data []byte) ([]byte, error) {
			re := regexp.MustCompile(`(?m)MARKETING_VERSION = [\d.]+;$`)
			newData := re.ReplaceAll(data, []byte(fmt.Sprintf("MARKETING_VERSION = %s;", newVersionNumber)))
			return newData, nil

		},
	); err != nil {
		log.Fatal(err)
	}

	manifestVersionChanger := func(data []byte) ([]byte, error) {
		re := regexp.MustCompile(`(?m)^\t"version": "[\d.]+",$`)
		newData := re.ReplaceAll(data, []byte(fmt.Sprintf("\t\"version\": \"%s\",", newVersionNumber)))
		return newData, nil
	}

	if err := ChangeFile(
		"targets/apple/public/manifest.json",
		manifestVersionChanger,
	); err != nil {
		log.Fatal(err)
	}

	if err := ChangeFile(
		"targets/chromium/public/manifest.json",
		manifestVersionChanger,
	); err != nil {
		log.Fatal(err)
	}

	if err := ChangeFile(
		"targets/firefox/public/manifest.json",
		manifestVersionChanger,
	); err != nil {
		log.Fatal(err)
	}
}

func ChangeFile(name string, changer func(data []byte) ([]byte, error)) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	newBytes, err := changer(data)
	if err != nil {
		return err
	}

	return os.WriteFile(name, newBytes, 0644)
}
