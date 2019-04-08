package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/stacktitan/smb/smb"
)

type Configuration struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	User        string `json:"user"`
	Domain      string `json:"domain"`
	Workstation string `json:"workstation"`
	Password    string `json:"password"`
}

func main() {

	cfg := LoadConfiguration("config")
	options := smb.Options{
		Host:        cfg.Host,
		Port:        cfg.Port,
		User:        cfg.User,
		Domain:      cfg.Domain,
		Workstation: cfg.Workstation,
		Password:    cfg.Password,
	}
	//fmt.Println(cfg)
	debug := true

	session, err := smb.NewSession(options, debug)
	if err != nil {
		log.Fatalln("[!]", err)
	}
	defer session.Close()

	if session.IsSigningRequired {
		log.Println("[-] Signing is required")
	} else {
		log.Println("[+] Signing is NOT required")
	}

	if session.IsAuthenticated {
		log.Println("[+] Login successful")
	} else {
		log.Println("[-] Login failed")
	}

	if err != nil {
		log.Fatalln("[!]", err)
	}
}

func LoadConfiguration(file string) Configuration {
	var config Configuration
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
