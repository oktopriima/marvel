package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const accessKey = "PJjZ_Uy_bsJgq4NpuIcd02OLcbFRhSDxIu30PhoeJx4" // Replace with your Unsplash API access key

type UnsplashResponse struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		Id               string `json:"id"`
		Slug             string `json:"slug"`
		AlternativeSlugs struct {
			En string `json:"en"`
			Es string `json:"es"`
			Ja string `json:"ja"`
			Fr string `json:"fr"`
			It string `json:"it"`
			Ko string `json:"ko"`
			De string `json:"de"`
			Pt string `json:"pt"`
		} `json:"alternative_slugs"`
		CreatedAt      time.Time     `json:"created_at"`
		UpdatedAt      time.Time     `json:"updated_at"`
		PromotedAt     interface{}   `json:"promoted_at"`
		Width          int           `json:"width"`
		Height         int           `json:"height"`
		Color          string        `json:"color"`
		BlurHash       string        `json:"blur_hash"`
		Description    *string       `json:"description"`
		AltDescription string        `json:"alt_description"`
		Breadcrumbs    []interface{} `json:"breadcrumbs"`
		Urls           struct {
			Raw     string `json:"raw"`
			Full    string `json:"full"`
			Regular string `json:"regular"`
			Small   string `json:"small"`
			Thumb   string `json:"thumb"`
			SmallS3 string `json:"small_s3"`
		} `json:"urls"`
		Links struct {
			Self             string `json:"self"`
			Html             string `json:"html"`
			Download         string `json:"download"`
			DownloadLocation string `json:"download_location"`
		} `json:"links"`
		Likes                  int           `json:"likes"`
		LikedByUser            bool          `json:"liked_by_user"`
		CurrentUserCollections []interface{} `json:"current_user_collections"`
		Sponsorship            interface{}   `json:"sponsorship"`
		TopicSubmissions       struct {
		} `json:"topic_submissions"`
		AssetType string `json:"asset_type"`
		User      struct {
			Id              string      `json:"id"`
			UpdatedAt       time.Time   `json:"updated_at"`
			Username        string      `json:"username"`
			Name            string      `json:"name"`
			FirstName       string      `json:"first_name"`
			LastName        string      `json:"last_name"`
			TwitterUsername *string     `json:"twitter_username"`
			PortfolioUrl    interface{} `json:"portfolio_url"`
			Bio             interface{} `json:"bio"`
			Location        *string     `json:"location"`
			Links           struct {
				Self      string `json:"self"`
				Html      string `json:"html"`
				Photos    string `json:"photos"`
				Likes     string `json:"likes"`
				Portfolio string `json:"portfolio"`
				Following string `json:"following"`
				Followers string `json:"followers"`
			} `json:"links"`
			ProfileImage struct {
				Small  string `json:"small"`
				Medium string `json:"medium"`
				Large  string `json:"large"`
			} `json:"profile_image"`
			InstagramUsername          string `json:"instagram_username"`
			TotalCollections           int    `json:"total_collections"`
			TotalLikes                 int    `json:"total_likes"`
			TotalPhotos                int    `json:"total_photos"`
			TotalPromotedPhotos        int    `json:"total_promoted_photos"`
			TotalIllustrations         int    `json:"total_illustrations"`
			TotalPromotedIllustrations int    `json:"total_promoted_illustrations"`
			AcceptedTos                bool   `json:"accepted_tos"`
			ForHire                    bool   `json:"for_hire"`
			Social                     struct {
				InstagramUsername string      `json:"instagram_username"`
				PortfolioUrl      interface{} `json:"portfolio_url"`
				TwitterUsername   *string     `json:"twitter_username"`
				PaypalEmail       interface{} `json:"paypal_email"`
			} `json:"social"`
		} `json:"user"`
	} `json:"results"`
}

func main() {
	places := []string{
		"Tugu Yogyakarta",
		"Malioboro",
		"Pantai Parangtritis",
		"Dufan",
	}

	for i := 0; i < len(places); i++ {
		url := fmt.Sprintf("https://api.unsplash.com/search/photos?query=%s&client_id=%s", strings.ReplaceAll(places[i], " ", "%20"), accessKey)
		get, err := http.Get(url)
		if err != nil {
			return
		}

		all, err := io.ReadAll(get.Body)
		if err != nil {
			fmt.Println(get.StatusCode)
			continue
		}

		var res UnsplashResponse
		err = json.Unmarshal(all, &res)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Convert the struct to JSON
		jsonData, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling to JSON:", err)
			return
		}

		// Write JSON data to a file
		fileName := fmt.Sprintf("images-%s.json", strings.ReplaceAll(places[i], " ", "-"))
		err = os.WriteFile(strings.ToLower(fileName), jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Printf("Data written to %s", fileName)
	}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	fmt.Println(runtime.GOOS)

	// Determine the OS and set the command accordingly
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url) // For Linux
	case "darwin":
		cmd = exec.Command("open", url) // For macOS
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url) // For Windows
	default:
		return nil // Unsupported OS
	}

	err := cmd.Start()
	time.Sleep(time.Second * 2) // Start the command

	return err
}
