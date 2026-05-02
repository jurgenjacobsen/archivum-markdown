package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/blang/semver"
)

const (
	owner   = "jurgenjacobsen"
	repo    = "archivum-markdown"
	version = "1.0.2"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type UpdateInfo struct {
	Available   bool   `json:"available"`
	LatestVersion string `json:"latestVersion"`
	DownloadURL  string `json:"downloadUrl"`
}

func (a *App) CheckForUpdates() (UpdateInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	
	resp, err := http.Get(url)
	if err != nil {
		return UpdateInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return UpdateInfo{}, fmt.Errorf("failed to check for updates: %s", resp.Status)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return UpdateInfo{}, err
	}

	currentV, err := semver.Parse(strings.TrimPrefix(version, "v"))
	if err != nil {
		return UpdateInfo{}, err
	}

	latestV, err := semver.Parse(strings.TrimPrefix(release.TagName, "v"))
	if err != nil {
		return UpdateInfo{}, err
	}

	if latestV.GT(currentV) {
		downloadURL := ""
		suffix := ""
		
		switch runtime.GOOS {
		case "windows":
			suffix = ".exe"
		case "darwin":
			suffix = ".dmg"
		case "linux":
			suffix = ".deb" // Or .AppImage
		}

		for _, asset := range release.Assets {
			if strings.HasSuffix(asset.Name, suffix) {
				downloadURL = asset.BrowserDownloadURL
				break
			}
		}

		// Fallback to first asset if no exact match
		if downloadURL == "" && len(release.Assets) > 0 {
			downloadURL = release.Assets[0].BrowserDownloadURL
		}

		return UpdateInfo{
			Available:     true,
			LatestVersion: release.TagName,
			DownloadURL:   downloadURL,
		}, nil
	}

	return UpdateInfo{Available: false}, nil
}
