package routers

import (
	"cloudflare-pages-hook/pkg/notification"
	"cloudflare-pages-hook/response"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	CFScanWaitTime = 2 * time.Minute
	CFPerPage      = "5"
	CFPage         = "1"
)

type filterDeploymentsResponse struct {
	ProjectName     string
	Environment     string
	Branch          string
	URL             string
	LatestBuildTime time.Time
	CommitHash      string
	CommitMsg       string
	Status          string
}

type DeploymentsResponse struct {
	Result []struct {
		ID          string    `json:"id"`
		ShortID     string    `json:"short_id"`
		ProjectID   string    `json:"project_id"`
		ProjectName string    `json:"project_name"`
		Environment string    `json:"environment"`
		URL         string    `json:"url"`
		CreatedOn   time.Time `json:"created_on"`
		ModifiedOn  time.Time `json:"modified_on"`
		LatestStage struct {
			Name      string    `json:"name"`
			StartedOn time.Time `json:"started_on"`
			EndedOn   time.Time `json:"ended_on"`
			Status    string    `json:"status"`
		} `json:"latest_stage"`
		DeploymentTrigger struct {
			Type     string `json:"type"`
			Metadata struct {
				Branch        string `json:"branch"`
				CommitHash    string `json:"commit_hash"`
				CommitMessage string `json:"commit_message"`
				CommitDirty   bool   `json:"commit_dirty"`
			} `json:"metadata"`
		} `json:"deployment_trigger"`
	} `json:"result"`
}

func deploymentsHook(c *gin.Context) {
	project := c.DefaultQuery("project", "")
	commitHash := c.DefaultQuery("commitHash", "")
	branch := c.DefaultQuery("branch", "")
	if project == "" || commitHash == "" || branch == "" {
		response.Fail(c, gin.H{}, "should provide \"id\" and \"commitHash\" args")
	} else {
		//text := fmt.Sprintf("✈️A new build is on the fly✈️\n\n"+
		//	"Project: *%s*\n"+
		//	"CommitHash: *%s*\n"+
		//	"Branch: *%s*", project, commitHash, branch)
		//
		//err := deploymentSendTG(text)
		//if err != nil {
		//	log.Println(err.Error())
		//	response.Fail(c, gin.H{}, err.Error())
		//	return
		//}

		response.Success(c, gin.H{}, "api success")

		go func(project, commitHash, branch string) {
			time.Sleep(CFScanWaitTime)

			r, err := makeDeploymentsRequest(project)
			if err != nil {
				response.Fail(c, gin.H{}, err.Error())
			}

			f, err := filterDeploymentsByCommit(r, commitHash, branch)
			if err != nil {
				response.Fail(c, gin.H{}, err.Error())
			}
			text := fmt.Sprintf("👇CF Pages deploy result:👇\n\n"+
				"Project: *%s*\n"+
				"Env: *%s*\n"+
				"Branch: *%s*\n"+
				"Url: *%s*\n"+
				"BuildTime: *%s*\n"+
				"CommitHash: *%s*\n"+
				"CommitMsg: *%s*\n"+
				"Status: *%s*\n"+
				"",
				f.ProjectName,
				f.Environment,
				f.Branch,
				f.URL,
				f.LatestBuildTime,
				f.CommitHash,
				f.CommitMsg,
				f.Status,
			)
			deploymentSendTG(text)
			log.Printf("%v", f)
		}(project, commitHash, branch)
	}
}

func makeDeploymentsRequest(project string) (DeploymentsResponse, error) {
	_url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/pages/projects/%s/deployments", os.Getenv("CHEEMS_CF_ACCOUNT_ID"), project)
	params := url.Values{}
	Url, _ := url.Parse(_url)
	params.Set("sort_by", "created_on")
	params.Set("sort_order", "desc")
	params.Set("per_page", CFPerPage)
	params.Set("page", CFPage)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CHEEMS_CF_ACCOUNT_TOKEN")))
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return DeploymentsResponse{}, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var r DeploymentsResponse
	json.Unmarshal(body, &r)
	return r, nil
}

func filterDeploymentsByCommit(r DeploymentsResponse, commitHash string, branch string) (filterDeploymentsResponse, error) {
	var f filterDeploymentsResponse
	for _, v := range r.Result {
		if v.DeploymentTrigger.Metadata.CommitHash == commitHash && v.DeploymentTrigger.Metadata.Branch == branch {
			ShortCommitMsg := string([]byte(v.DeploymentTrigger.Metadata.CommitHash))[:7]
			f.ProjectName = v.ProjectName
			f.Environment = v.Environment
			f.Branch = branch
			f.URL = v.URL
			f.LatestBuildTime = v.LatestStage.StartedOn
			f.CommitHash = ShortCommitMsg
			f.CommitMsg = v.DeploymentTrigger.Metadata.CommitMessage
			if v.LatestStage.Status == "success" {
				f.Status = "✅success"
			} else if v.LatestStage.Status == "failure" {
				f.Status = "❌failure"
			} else {
				f.Status = v.LatestStage.Status
			}
			return f, nil
		}
	}
	return f, fmt.Errorf("no result")
}

func deploymentSendTG(text string) error {
	var n notification.Notifier
	n.SetNotification(notification.N)
	n.Send(text)
	return nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/deployment", deploymentsHook)
	return r
}
