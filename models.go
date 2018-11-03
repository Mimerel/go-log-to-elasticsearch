package main

type configuration struct {
	ElasticsearchURL string `yaml:"elasticSearchUrl,omitempty"`
	Port string `yaml:"port,omitempty"`
}