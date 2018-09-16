package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

// Config struct contains the settings from the config file.
type Config struct {
	Containers []string
	Enable     bool
	Email      Email
}

// Email configuration structure
type Email struct {
	Enable    bool
	Password  string
	Username  string
	Url       string
	Sender    string
	Recipient string
	Hostname  string
}

// Load the config file into a Config struct object. If the config flag "-config" is not specified a default location
// will be used which is equivalent "$GOPATH/src/de.implex1v/simple-dm/config.yml". This function will write a message
// on error and exist with a code > 0.
func LoadConfig() Config {
	gopath := os.Getenv("GOPATH")
	configPtr := flag.String("config", gopath+"/src/de.implex1v/simple-dm/config.yml", "The path to the config file")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*configPtr)
	if err != nil {
		fmt.Println("Specified config file or default config file is not existing or not readable")
		os.Exit(2)
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println("Config file is not valid")
		os.Exit(3)
	}

	return config
}

// Loads the container information of all currently running containers.
func LoadRunningContainers() []types.Container {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		fmt.Println("Cloud not get docker cli. " + err.Error())
		os.Exit(4)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Println("Cloud not get all containers. " + err.Error())
		os.Exit(5)
	}

	return containers
}

// Checks if the containers of the configuration are running.
func CheckContainers(config Config, containers []types.Container) []string {
	notFoundContainers := make([]string, len(config.Containers))
	copy(notFoundContainers, config.Containers)

	for _, container := range containers {
		for _, name := range container.Names {
			notFoundContainers = remove(notFoundContainers, name)
		}
	}

	return notFoundContainers
}

// Removes needle from haystack if needle is in haystack
func remove(haystack []string, needle string) []string {
	if needle[0] == '/' {
		needle = needle[1:]
	}

	for pos, hay := range haystack {
		if hay == needle {
			return append(haystack[:pos], haystack[pos+1:]...)
		}
	}

	return haystack
}

// Sends a mail with the not running containers. The email information are taken from config.
func SendMail(config Config, missingContainers []string) {
	text := "Some of your Docker containers might not be running:\n" + strings.Join(missingContainers, "\n")
	msg := []byte("To: " + config.Email.Recipient + "\r\n" +
		"Subject: SimpleDockerMonitoring warning!\r\n" +
		"\r\n" + text)

	auth := smtp.PlainAuth("", config.Email.Username, config.Email.Password, config.Email.Hostname)
	err := smtp.SendMail(config.Email.Url, auth, config.Email.Sender, []string{config.Email.Recipient}, msg)
	if err != nil {
		fmt.Println("Cloud not send mail. " + err.Error())
		os.Exit(7)
	} else {
		os.Exit(0)
	}
}

func main() {
	c := LoadConfig()
	if c.Enable == false {
		os.Exit(0)
	}

	containers := LoadRunningContainers()
	missingContainers := CheckContainers(c, containers)
	if len(missingContainers) > 0 && c.Email.Enable {
		SendMail(c, missingContainers)
	}
}
