package main

import (
	"flag"
	"fmt"
)

type Args struct {
	DefinitionFile string
	OutputFile     string
	Env            string
	Site           string
}

func (a *Args) ParseArgs() {
	flag.StringVar(&a.DefinitionFile, "definition", "definition.yml", "Definition file")
	flag.StringVar(&a.OutputFile, "output", "manifest.yml", "Output file")
	flag.StringVar(&a.Env, "environment", "perso", "Environment (perso|dev|integ|staging|pcache|prod)")
	flag.StringVar(&a.Site, "site", "sph", "Site trigram (sph|bgl|mts)")
	flag.Parse()

	fmt.Println("definition :", a.DefinitionFile)
	fmt.Println("output :", a.OutputFile)
	fmt.Println("environment :", a.Env)
	fmt.Println("site :", a.Site)

}
