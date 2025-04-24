package utiles

import (
	"encoding/json"
	"fmt"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strings"
)

type ReleaseData struct {
	Version   string
	Link      string
	Changelog string
}

func NewGetRelease() *ReleaseData {
	return new(ReleaseData)
}

func (r *ReleaseData) RefreshRelease(user, repo, proxy string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", user, repo)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logx.Error("Error creating request:", err)
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		logx.Error("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logx.Error("Error response from server:", resp.Status)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Error("Error reading response body:", err)
		return err
	}
	tmpdata := types.GhData{}
	err = json.Unmarshal(body, &tmpdata)
	if err != nil {
		logx.Error("Error parsing JSON data:", err)
		return err
	}
	r.Version = tmpdata.TagName
	r.Changelog = tmpdata.Body
	for _, assets := range tmpdata.Assets {
		if strings.Contains(assets.Name, "GUI-BilibiliDanmuRobot_Windows_amd64_") {
			r.Link = proxy + assets.BrowserDownloadUrl
			break
		}
	}
	return nil
}
func (r *ReleaseData) RefreshUpgraderRelease(user, repo, proxy string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", user, repo)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logx.Error("Error creating request:", err)
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		logx.Error("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logx.Error("Error response from server:", resp.Status)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Error("Error reading response body:", err)
		return err
	}
	tmpdata := types.GhData{}
	err = json.Unmarshal(body, &tmpdata)
	if err != nil {
		logx.Error("Error parsing JSON data:", err)
		return err
	}
	r.Version = tmpdata.TagName
	r.Changelog = tmpdata.Body
	for _, assets := range tmpdata.Assets {
		if assets.Name == "upgrader.exe" {
			r.Link = proxy + assets.BrowserDownloadUrl
			break
		}
	}
	return nil
}
