package statistics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const githubAPI = "https://api.github.com/repos"

func getPulls(ownerRepo, token string) ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/pulls", githubAPI, ownerRepo), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+token)
	req.Header.Add("User-Agent", "GoLang Pulls App")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pulls []map[string]interface{} // dictionary of objects
	err = json.Unmarshal(body, &pulls)
	if err != nil {
		return nil, err
	}
	return pulls, nil
}

func Cli_options() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go [owner/repo] [Your-GitHub-Token]")
		return
	}
	repo := os.Args[1]
	token := os.Args[2]

	pulls, err := getPulls(repo, token)
	if err != nil {
		fmt.Println("Error fetching pull requests:", err)
		return
	}

	if len(pulls) == 0 {
		fmt.Println("No pull requests found for", repo)
		return
	}

	fmt.Println("Pull Requests for", repo)
	for _, pull := range pulls {
		fmt.Printf("#%v: %v by %v\n",
			pull["number"].(float64),
			pull["title"].(string),
			pull["user"].(map[string]interface{})["login"].(string),
		)
	}
}
