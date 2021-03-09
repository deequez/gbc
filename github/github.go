package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiBase = "https://api.github.com/users"

// GithubUserInfo returns a user's details from Github.
// type GithubUserInfo struct {
// 	Name     string `json:"name"`
// 	NumRepos int    `json:"public_repos"`
// }

// githubUserInfo returns (for a github login) the username and number of
// public repositories.
func githubUserInfo(login string) (string, int, error) {
	url := fmt.Sprintf("%s/%s", apiBase, login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("Bad status: %d %s", resp.StatusCode, resp.Status)
	}

	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	// fmt.Println("Body:", resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// var reply map[string]interface{}
	// var userInfo GithubUserInfo
	var userInfo struct {
		Name     string `json:"name"`
		NumRepos int    `json:"public_repos"`
	}
	if err := json.Unmarshal(data, &userInfo); err != nil {
		log.Fatalf("error: %s:", err)
	}

	// fmt.Println("reply: ", reply)
	// name := reply["name"]
	// publicRepos := reply["public_repos"]
	// fmt.Println("name:", name)
	// fmt.Println("public repos:", publicRepos)

	return userInfo.Name, userInfo.NumRepos, nil
}

func main() {
	name, repos, err := githubUserInfo("tebeka")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s, num repos: %d\n", name, repos)
	// name: Miki Tebeka, num repos: 76
}
