package main

import (
	"flag"
	"fmt"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Config struct contains the settings from the config file.
type Config struct {
	Containers []string
	Enable     bool
}

// ContainerStatus struct contains the elements of docker ps
type ContainerStatus struct {
	Id      string
	Image   string
	Command string
	Created string
	Status  string
	Port    string
	Names   string
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

/*func BuildContainerStatus(output string) []ContainerStatus {
	lines := strings.Split(output, "\n")

	// no running containers where found
	if len(lines) <= 2 {
		return []ContainerStatus{}
	}

	for i := 1; i < len(lines) - 1; i++ {
		ParseStatusString(lines[i])
	}

	return []ContainerStatus{}
}

func ParseStatusString(line string) ContainerStatus {
	containerStatus := ContainerStatus{}
	inString := false

	attributeIndex := 0
	attribute := ""
	fmt.Println(line)
	for pos, char := range line {
		if pos == len(line) -1 {
			containerStatus.Names = attribute
		} else if char == '"' {
			inString = !inString
		} else if char == ' ' {
			if inString {
				attribute += string(char)
			} else if len(attribute) != 0 {
				switch attributeIndex {
					case 0:
						containerStatus.Id = attribute
						break
					case 1:
						containerStatus.Image = attribute
						break
					case 2:
						containerStatus.Command = attribute
						break
					case 3:
						containerStatus.Created = attribute
						break
					case 4:
						containerStatus.Status = attribute
						break
					case 5:
						containerStatus.Port = attribute
						break
					default:
						fmt.Println("Illegal attributeIndex reached")
						os.Exit(5)
				}

				print("old attribute "+attribute+"\n")
				attribute = ""
				attributeIndex++
			}
		} else {
			attribute += string(char)
		}
	}

	fmt.Println(containerStatus)
	return containerStatus
}*/

func main() {
	c := LoadConfig()
	if c.Enable == false {
		os.Exit(0)
	}

	client, err := client.NewEnvClient()

	/*
		outputBytes, err := exec.Command("docker", "ps").Output()
		if err != nil {
			fmt.Println("Cloud not get running docker containers. Reason: "+err.Error())
			os.Exit(4)
		}
		output := string(outputBytes)*/
}
