package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Definition struct {
	Envs map[string]Manifest
}

type Docker struct {
	Image    string `yaml:"image,omitempty"`
	Username string `yaml:"username,omitempty"`
}

type Route struct {
	Route string `yaml:"route"`
}

type Application struct {
	Name                    string                 `yaml:"name,omitempty"`
	Memory                  string                 `yaml:"memory,omitempty"`
	DiskQuota               string                 `yaml:"disk_quota,omitempty"`
	Instances               int                    `yaml:"instances,omitempty"`
	Command                 string                 `yaml:"command,omitempty"`
	Docker                  Docker                 `yaml:"docker,omitempty"`
	HealthCheckHttpEndpoint string                 `yaml:"health-check-http-endpoint,omitempty"`
	HealthCheckType         string                 `yaml:"health-check-type,omitempty"`
	NoRoute                 bool                   `yaml:"no-route,omitempty"`
	RandomRoute             bool                   `yaml:"random-route,omitempty"`
	Stack                   string                 `yaml:"stack,omitempty"`
	Timeout                 int                    `yaml:"timeout,omitempty"`
	Domain                  string                 `yaml:"domain,omitempty"`
	Domains                 []string               `yaml:"domains,omitempty"`
	Host                    string                 `yaml:"host,omitempty"`
	Hosts                   []string               `yaml:"hosts,omitempty"`
	NoHostname              bool                   `yaml:"no-hostname,omitempty"`
	Routes                  []Route                `yaml:"routes,omitempty"`
	Services                []string               `yaml:"services,omitempty"`
	Path                    string                 `yaml:"path,omitempty"`
	BuildPack               string                 `yaml:"buildpack,omitempty"`
	BuildPacks              []string               `yaml:"buildpacks,omitempty"`
	Env                     map[string]interface{} `yaml:"env,omitempty"`
}

type Manifest struct {
	Applications []Application `yaml:"applications,omitempty"`
}

func (manifest *Manifest) parseByte(Byte []byte) error {
	var err = yaml.Unmarshal(Byte, manifest)
	if err != nil {
		return err
	}
	return nil
}

func (manifest *Manifest) parseFile(manifestFile string) error {
	yamlFile, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return err
	}
	return manifest.parseByte(yamlFile)
}

func (def *Definition) parseByte(Byte []byte) error {
	var err = yaml.Unmarshal(Byte, def)
	if err != nil {
		return err
	}
	return nil
}

func (def *Definition) parseFile(manifestFile string) error {
	yamlFile, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return err
	}
	return def.parseByte(yamlFile)
}

func (manifest *Manifest) GenerateFile(FileName string) error {
	Byte, err := yaml.Marshal(manifest)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(FileName, Byte, 0644)
}
