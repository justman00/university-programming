package trustpilot

// implement an http client that can be used to make requests to the trustpilot api
// it should use an API key for that
// should also go through all the results until there are no more pages
// should also be able to handle errors
// should also be able to handle rate limiting

// implement a function that can be used to import all the reviews from trustpilot
// it should use the http client to make requests to the trustpilot api
// it should use the review repository to store the reviews

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type TrustpilotReview struct {
	ID            string    `json:"id"`
	Stars         int       `json:"stars"`
	Title         string    `json:"title"`
	Text          string    `json:"text"`
	Language      string    `json:"language"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	ExperiencedAt time.Time `json:"experiencedAt"`
	IsVerified    bool      `json:"isVerified"`
	CompanyReply  struct {
		Text      string `json:"text"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	} `json:"companyReply"`
	Consumer struct {
		DisplayLocation string `json:"displayLocation"`
		DisplayName     string `json:"displayName"`
		ID              string `json:"id"`
	} `json:"consumer"`
	Location struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
}

func (r *TrustpilotReview) URL() string {
	return fmt.Sprintf("https://www.trustpilot.com/reviews/%s", r.ID)
}

type TrustpilotAPIResponse struct {
	Reviews       []TrustpilotReview `json:"reviews"`
	NextPageToken string             `json:"nextPageToken"`
}

type TrustpilotClient struct {
	apiKey string
	client *http.Client
}

func NewTrustpilotClient(apiKey string) *TrustpilotClient {
	return &TrustpilotClient{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

const trustpilotAPISchema = "https://api.trustpilot.com/v1/business-units/%s/all-reviews"

func (c *TrustpilotClient) GetReviews(ctx context.Context, businessUnitID string) ([]TrustpilotReview, error) {
	var reviews []TrustpilotReview
	var nextPageToken string

	for {
		logrus.Infof("fetching reviews from trustpilot api, next page: %s, current reviews: %d", nextPageToken, len(reviews))
		// build the request URL
		u, err := url.Parse(fmt.Sprintf(trustpilotAPISchema, businessUnitID))
		if err != nil {
			return nil, fmt.Errorf("parse trustpilot api url: %w", err)
		}

		q := u.Query()
		q.Set("apikey", c.apiKey)
		if nextPageToken != "" {
			q.Set("pageToken", nextPageToken)
		}
		u.RawQuery = q.Encode()

		// make the request
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, fmt.Errorf("create trustpilot api request: %w", err)
		}

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("make trustpilot api request: %w", err)
		}
		defer resp.Body.Close()

		// check for rate limiting
		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter, err := time.ParseDuration(resp.Header.Get("Retry-After"))
			if err != nil {
				return nil, fmt.Errorf("parse retry after header: %w", err)
			}
			time.Sleep(retryAfter)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("trustpilot api returned status code %d", resp.StatusCode)
		}

		// decode the response
		var response TrustpilotAPIResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, fmt.Errorf("decode trustpilot api response: %w", err)
		}
		reviews = append(reviews, response.Reviews...)
		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return reviews, nil
}
