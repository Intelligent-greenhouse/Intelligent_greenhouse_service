package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	proto_api_directory = "./proto/api"
)

func ensureDir(dirName string) error {
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		return os.MkdirAll(dirName, 0755)
	}
	return err
}

func findProto(directory, t string) []string {
	var protoFiles []string
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, fmt.Sprintf("/%s/", t)) && strings.HasSuffix(path, ".proto") {
			relativePath, _ := filepath.Rel(directory, path)
			protoFiles = append(protoFiles, filepath.Join(directory, relativePath))
		}
		return nil
	})
	return protoFiles
}

func generateProtoWEB(protoFiles []string) {
	if err := ensureDir("./api"); err != nil {
		panic(err)
	}
	args := []string{
		"--proto_path=./proto/api",
		"--proto_path=./proto/extension",
		"--go_out=paths=source_relative:./api",
		"--go-http_out=paths=source_relative:./api",
		"--go-grpc_out=paths=source_relative:./api",
	}
	args = append(args, protoFiles...)
	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("protoc %s\n", strings.Join(args, " "))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func generateProtoAllWEB() {
	protoFiles := findProto(proto_api_directory, "web")
	generateProtoWEB(protoFiles)
}

func generateProtoMQ(protoFiles []string) {
	if err := ensureDir("./api"); err != nil {
		panic(err)
	}
	args := []string{
		"--proto_path=./proto/api",
		"--proto_path=./proto/extension",
		"--go_out=paths=source_relative:./api",
	}
	args = append(args, protoFiles...)
	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("protoc %s\n", strings.Join(args, " "))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func generateProtoAllMQ() {
	protoFiles := findProto(proto_api_directory, "mq")
	generateProtoMQ(protoFiles)
}

func generateProtoConfig() {
	args := []string{
		"--proto_path=./conf",
		"--go_out=paths=source_relative:./conf",
		"./conf/conf.proto",
	}
	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("protoc %s\n", strings.Join(args, " "))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func generateProtoOpenAPI() {
	protoFiles := findProto(proto_api_directory, "web")
	args := []string{
		"--proto_path=./proto/api",
		"--proto_path=./proto/extension",
		"--openapi_out=fq_schema_naming=true,default_response=false:.",
	}
	args = append(args, protoFiles...)
	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("protoc %s\n", strings.Join(args, " "))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	arguments := os.Args[1:]
	command := "all"
	if len(arguments) > 0 {
		command, arguments = arguments[0], arguments[1:]
	}
	_ = arguments

	switch command {
	case "api":
		generateProtoAllWEB()
		generateProtoAllMQ()
	case "conf":
		generateProtoConfig()
	case "openapi":
		generateProtoOpenAPI()
	case "all":
		generateProtoAllWEB()
		generateProtoAllMQ()
		generateProtoConfig()
		generateProtoOpenAPI()
	}
}
