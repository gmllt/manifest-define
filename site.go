package main

import (
	"regexp"
	"strings"
)

type Site struct {
	Trigram string
	Short string
	Long string
}

type Replacement struct {
	Search string
	Replace string
}

func GetSites() []Site  {
	var Sites []Site

	Sites = append(Sites, Site{"sph", "s0", "sophia"})
	Sites = append(Sites, Site{"bgl", "b0", "bagnolet"})
	Sites = append(Sites, Site{"mts", "m0", "montsouris"})

	return Sites
}

func GetReplacements(identifier string) []Replacement {
	var Replacements []Replacement
	for _, Site := range GetSites() {
		if Site.Trigram == identifier {
			var TrigramRegex []string
			var ShortRegex []string
			var LongRegex []string
			for _, SiteSearch := range GetSites() {
				if SiteSearch.Trigram != identifier {
					TrigramRegex = append(TrigramRegex, SiteSearch.Trigram)
					LongRegex = append(LongRegex, SiteSearch.Long)
					ShortRegex = append(ShortRegex, SiteSearch.Short)
				}
			}
			Replacements = append(Replacements, Replacement{"("+strings.Join(TrigramRegex, "|")+")", Site.Trigram})
			Replacements = append(Replacements, Replacement{"("+strings.Join(LongRegex, "|")+")", Site.Long})
			Replacements = append(Replacements, Replacement{"("+strings.Join(ShortRegex, "|")+")", Site.Short})
		}
	}
	return Replacements
}

func (s Site) Replace(Var string) string {
	for _, Replace := range GetReplacements(s.Trigram) {
		var re = regexp.MustCompile(Replace.Search)
		Var = re.ReplaceAllString(Var, Replace.Replace)
	}
	return Var
}