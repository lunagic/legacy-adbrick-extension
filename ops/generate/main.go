package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Thing struct {
	ID        int            `json:"id"`
	Priority  int            `json:"priority"`
	Action    ThingAction    `json:"action"`
	Condition ThingCondition `json:"condition"`
}

type ThingAction struct {
	Type string `json:"type"`
}

type ThingCondition struct {
	URLFilter string `json:"urlFilter"`
}

func main() {
	// https://firebog.net/
	files := []string{
		// Suspicious Lists
		"https://raw.githubusercontent.com/PolishFiltersTeam/KADhosts/master/KADhosts.txt",
		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Spam/hosts",
		"https://v.firebog.net/hosts/static/w3kbl.txt",
		// Advertising Lists
		"https://adaway.org/hosts.txt",
		"https://v.firebog.net/hosts/AdguardDNS.txt",
		"https://v.firebog.net/hosts/Admiral.txt",
		"https://raw.githubusercontent.com/anudeepND/blacklist/master/adservers.txt",
		"https://v.firebog.net/hosts/Easylist.txt",
		"https://pgl.yoyo.org/adservers/serverlist.php?hostformat=hosts&showintro=0&mimetype=plaintext",
		"https://raw.githubusercontent.com/FadeMind/hosts.extras/master/UncheckyAds/hosts",
		"https://raw.githubusercontent.com/bigdargon/hostsVN/master/hosts",
		// Tracking & Telemetry Lists
		// "https://v.firebog.net/hosts/Easyprivacy.txt",
		// "https://v.firebog.net/hosts/Prigent-Ads.txt",
		// "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.2o7Net/hosts",
		// "https://raw.githubusercontent.com/crazy-max/WindowsSpyBlocker/master/data/hosts/spy.txt",
		// "https://hostfiles.frogeye.fr/firstparty-trackers-hosts.txt",
		// Malicious Lists
		// "https://raw.githubusercontent.com/DandelionSprout/adfilt/master/Alternate%20versions%20Anti-Malware%20List/AntiMalwareHosts.txt",
		// "https://osint.digitalside.it/Threat-Intel/lists/latestdomains.txt",
		// "https://v.firebog.net/hosts/Prigent-Crypto.txt",
		// "https://raw.githubusercontent.com/FadeMind/hosts.extras/master/add.Risk/hosts",
		// "https://bitbucket.org/ethanr/dns-blacklists/raw/8575c9f96e5b4a1308f2f12394abd86d0927a4a0/bad_lists/Mandiant_APT1_Report_Appendix_D.txt",
		// "https://phishing.army/download/phishing_army_blocklist_extended.txt",
		// "https://gitlab.com/quidsup/notrack-blocklists/raw/master/notrack-malware.txt",
		// "https://v.firebog.net/hosts/RPiList-Malware.txt",
		// "https://v.firebog.net/hosts/RPiList-Phishing.txt",
		// "https://raw.githubusercontent.com/Spam404/lists/master/main-blacklist.txt",
		// "https://raw.githubusercontent.com/AssoEchap/stalkerware-indicators/master/generated/hosts",
		// "https://urlhaus.abuse.ch/downloads/hostfile/",
		// Other
		// "https://zerodot1.gitlab.io/CoinBlockerLists/hosts_browser",
	}

	things := []Thing{}

	m := map[string]bool{}
	for _, filePath := range files {
		for _, domain := range getDomainsFromURL(filePath) {
			if alreadyExists := m[domain]; alreadyExists {
				continue
			}
			m[domain] = true

			things = append(things, Thing{
				ID:       len(things) + 1,
				Priority: 1,
				Action: ThingAction{
					Type: "block",
				},
				Condition: ThingCondition{
					URLFilter: fmt.Sprintf("*://%s/*", domain),
				},
			})
		}
	}

	file, err := os.Create("src/assets/rules.json")
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(things, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	file.Write(b)
}

func getDomainsFromURL(url string) []string {
	log.Printf("Downloading: %s", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	domains := []string{}

	re := regexp.MustCompile(`(?m)\s([\w.-]+\.[\w-]+)$`)
	matches := re.FindAllStringSubmatch(string(b), -1)
	for _, match := range matches {
		domains = append(domains, match[1])
	}

	return domains
}
