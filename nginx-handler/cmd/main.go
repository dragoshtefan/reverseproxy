package main

import (
	"context"
	"fmt"
	"log"
	"nginx-handler/internal/config"
	"nginx-handler/pkg/dockerclient"
	"nginx-handler/pkg/filewriter"
	"nginx-handler/pkg/filewriter/generator"
	"os"
	"path/filepath"
)

func main() {

	inPath := "./configs/config.yml"
	ngingxConf := "./out/nginx.conf"
	dockerYaml := "./out/docker-compose.yml"

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	inputConfig, err := config.ParseYAML(inPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var forwardConfigs []generator.FWDProxyDescriptor = make([]generator.FWDProxyDescriptor, len(inputConfig.Services))

	for i, service := range inputConfig.Services {
		port, err := dockerclient.GetExternalPort(ctx, service.ContainerName, service.ContainerPort)
		if err != nil {
			log.Printf("Error getting port for service %s: %v", service.Name, err)
			os.Exit(1)
		}
		log.Println("got port for service ", service.Name, port)
		forwardConfigs[i] = generator.FWDProxyDescriptor{
			ListenPort:   service.ListenPort,
			RedirectPort: port,
			Name:         service.Name,
		}
	}

	fileContent := generator.GenerateNginxConf(&forwardConfigs)
	written, err := filewriter.WriteToFile(ngingxConf, []byte(fileContent))
	if err != nil {
		log.Fatal("failed writing to nginx conf file", err)
	}
	log.Printf("written %d bytes in %v", written, ngingxConf)

	ymlContent := generator.GenerateDockerYmlFile(dockerYaml, &forwardConfigs)
	written, err = filewriter.WriteToFile(dockerYaml, []byte(ymlContent))
	if err != nil {
		log.Fatal("failed writing to docker-compose file", err)
	}
	log.Printf("written %d bytes in %v", written, dockerYaml)
}
