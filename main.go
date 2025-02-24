package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	gitlabURL = "<URL DO GITLAB>/api/v4"
	groupID   = "<NOME DO GRUPO>"
)

type Project struct {
	ID  int    `json:"id"`
	URL string `json:"web_url"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		fmt.Println("Token de acesso não fornecido. Defina GITLAB_TOKEN no arquivo .env.")
		return
	}

	projects := getProjects(token)
	for _, project := range projects {
		if !checkLivenessProbe(token, project.ID) {
			fmt.Println(project.URL)
		}
	}
}

func getProjects(token string) []Project {
	url := fmt.Sprintf("%s/groups/%s/projects?per_page=100&include_subgroups=true", gitlabURL, groupID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao buscar projetos:", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var projects []Project
	json.Unmarshal(body, &projects)

	return projects
}

func checkLivenessProbe(token string, projectID int) bool {
	url := fmt.Sprintf("%s/projects/%d/repository/files/%s/raw?ref=main", gitlabURL, projectID, "values.yaml")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	return strings.Contains(string(content), "livenessProbe:") && strings.Contains(string(content), "enabled: true")
}
