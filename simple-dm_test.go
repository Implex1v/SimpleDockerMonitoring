package main

import (
	"github.com/docker/docker/api/types"
	"os"
	"os/exec"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	cft := gopath+"/src/de.implex1v/simple-dm/config.yml"
	LoadConfig(&cft)
	return
}

func TestLoadConfig_F(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	cfg := gopath+"/src/de.implex1v/simple-dm/confiq.yml"

	if os.Getenv("BE_CRASHER") == "1" {
		LoadConfig(&cfg)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoadConfig_F")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestLoadConfig_I(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	cfg := gopath+"/src/de.implex1v/simple-dm/tests/invalid.yml"

	if os.Getenv("BE_CRASHER") == "1" {
		LoadConfig(&cfg)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLoadConfig_F")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestCheckContainers(t *testing.T) {
	containerNames := []string{"Foo","LoremIpsum", "Bar", "XY"}
	config := Config{}
	config.Containers = containerNames

	container1 := types.Container{}
	container1.Names = []string{"Foo", "XY"}

	container2 := types.Container{}
	container2.Names = []string{"LoremIpsum"}
	containers := []types.Container{container1, container2}

	notFound := CheckContainers(config, containers)

	if len(notFound) != 1 {
		t.Fatal("")
	}
}

func TestCheckContainers_All(t *testing.T) {
	containerNames := []string{"Foo","LoremIpsum", "Bar", "XY"}
	config := Config{}
	config.Containers = containerNames

	container1 := types.Container{}
	container1.Names = []string{"Foo", "XY"}

	container2 := types.Container{}
	container2.Names = []string{"LoremIpsum", "y"}

	container3 := types.Container{}
	container3.Names = []string{"Bar", "Z"}

	containers := []types.Container{container1, container2, container3}

	notFound := CheckContainers(config, containers)

	if len(notFound) != 0 {
		t.Fatal("")
	}
}

func TestRemove_Simple(t *testing.T) {
	hay := []string{"A","B","C","D"}
	needle := "A"

	result := remove(hay, needle)
	if len(result) != 3 {
		t.Fatal("Hay has wrong length")
	}
}

func TestRemove_Slash(t *testing.T) {
	hay := []string{"A","B","C","D"}
	needle := "/A"

	result := remove(hay, needle)
	if len(result) != 3 {
		t.Fatal("Hay has wrong length")
	}
}

func TestRemove_Double(t *testing.T) {
	hay := []string{"A","B","C","A"}
	needle := "/A"

	result := remove(hay, needle)
	if len(result) != 3 {
		t.Fatal("Hay has wrong length")
	}
}