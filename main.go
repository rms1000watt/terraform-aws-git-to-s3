package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

var (
	gitUser = ""
	gitPass = ""
	gitURL  = ""
	svc     *s3.S3
)

func maisn() {
	createDir()

}

func createDir() {
	os.RemoveAll("./git")
	os.Mkdir("./git", os.ModePerm)
}

// Handler processes the API Gateway proxy request
func Handler(request events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	clone()

	if err = putObject(); err != nil {
		fmt.Println("Failed putting object:", err)
		return
	}

	out, err := getObject()
	if err != nil {
		fmt.Println("Failed getting object:", err)
		return
	}

	return events.APIGatewayProxyResponse{Body: out, StatusCode: 200}, nil
}

func clone() (dir string, err error) {
	auth := http.BasicAuth{
		Username: gitUser,
		Password: gitPass,
	}

	dir, err = ioutil.TempDir("", "git")
	if err != nil {
		fmt.Println("Failed creating tmp dir:", err)
		return
	}

	opts := git.CloneOptions{
		URL:  gitURL,
		Auth: &auth,
	}

	if _, err = git.PlainClone(dir, false, &opts); err != nil {
		fmt.Println("Failed cloning repo:", err)
		return
	}

	return
}

func setEnv() (err error) {
	gitUser = os.Getenv("GIT_USER")
	if strings.TrimSpace(gitUser) == "" {
		err = errors.New("GIT_USER env var not defined")
		return
	}

	gitPass = os.Getenv("GIT_PASS")
	if strings.TrimSpace(gitPass) == "" {
		err = errors.New("GIT_PASS env var not defined")
		return
	}

	gitURL = os.Getenv("GIT_URL")
	if strings.TrimSpace(gitURL) == "" {
		err = errors.New("GIT_URL env var not defined")
		return
	}

	return
}

func main() {
	if err := setEnv(); err != nil {
		fmt.Println("Failed getting env:", err)
		return
	}

	svc = s3.New(session.New())
	lambda.Start(Handler)
}
