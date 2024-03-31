package ghin

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

const (
	loginPath         = "golfer_login.json"
	logoutPath        = "users/logout.json"
	baseURL           = "https://api2.ghin.com/api/v1/"
	maxScoresPath     = "maximum_hole_scores.json"
	lookupPath        = "golfers.json"
	searchGolfersPath = "golfers/search.json"
	searchCoursePath  = "crsCourseMethods.asmx/SearchCourses.json"
	courseDetailsPath = "crsCourseMethods.asmx/GetCourseDetails.json"
	postScorePath     = "scores/hbh.json"
)

type (
	httpClient struct {
		baseURL string
		Client  interface {
			Do(req *http.Request) (*http.Response, error)
		}
		authToken string
	}

	Golfer struct {
		GhinNumber                 string  `json:"ghin_number"`
		Suffix                     string  `json:"suffix"`
		FirstName                  string  `json:"first_name"`
		MiddleName                 string  `json:"middle_name"`
		LastName                   string  `json:"last_name"`
		Prefix                     string  `json:"prefix"`
		PlayerName                 string  `json:"player_name"`
		Gender                     string  `json:"gender"`
		ClubName                   string  `json:"club_name"`
		ClubId                     string  `json:"club_id"`
		GolfAssociationName        string  `json:"golf_association_name"`
		GolfAssociationId          string  `json:"golf_association_id"`
		Display                    string  `json:"display"`
		DateOfBirth                string  `json:"date_of_birth"`
		LowHiDisplay               string  `json:"low_hi_display"`
		Email                      string  `json:"email"`
		PrimaryClubCountry         string  `json:"primary_club_country"`
		PrimaryClubState           string  `json:"primary_club_state"`
		PrimaryClubName            string  `json:"primary_club_name"`
		PrimaryClubId              int     `json:"primary_club_id"`
		PrimaryGolfAssociationId   int     `json:"primary_golf_association_id"`
		PrimaryGolfAssociationName string  `json:"primary_golf_association_name"`
		RevDate                    string  `json:"rev_date"`
		Status                     string  `json:"status"`
		TechnologyProvider         string  `json:"technology_provider"`
		SoftCap                    string  `json:"soft_cap"`
		HardCap                    string  `json:"hard_cap"`
		MessageClubAuthorized      *string `json:"message_club_authorized"`
	}

	GolferSubscription struct {
		Active                         bool   `json:"active"`
		SubscriptionAppType            string `json:"subscription_app_type"`
		SubscriptionType               string `json:"subscription_type"`
		InitialSubscriptionDate        string `json:"initial_subscription_date"`
		CurrentSubscriptionStartDate   string `json:"current_subscription_start_date"`
		CurrentSubscriptionEndDate     string `json:"current_subscription_end_date"`
		CurrentSubscriptionRenewalType string `json:"current_subscription_renewal_type"`
	}

	User struct {
		GolferUserToken         string              `json:"golfer_user_token"`
		GolferId                int                 `json:"golfer_id"`
		GuardianId              *int                `json:"guardian_id"`
		GolferUserAcceptedTerms bool                `json:"golfer_user_accepted_terms"`
		GolferCreationDate      time.Time           `json:"golfer_creation_date"`
		Golfers                 []Golfer            `json:"golfers"`
		MinorAccounts           []any               `json:"minor_accounts"`
		Subscription            *GolferSubscription `json:"subscription"`
	}
)

func newDefaultClient() *httpClient {
	return &httpClient{Client: http.DefaultClient, baseURL: baseURL}
}

func (c *httpClient) Login(ctx context.Context, email, password string) (*User, error) {
	reqURL := filepath.Join(c.baseURL, loginPath)
	login := struct {
		User struct {
			Email      string `json:"email"`
			Password   string `json:"password"`
			RememberMe bool   `json:"remember_me"`
		} `json:"user"`
		Token     string `json:"token"`
		UserToken string `json:"user_token"`
	}{
		User: struct {
			Email      string `json:"email"`
			Password   string `json:"password"`
			RememberMe bool   `json:"remember_me"`
		}{
			Email:      email,
			Password:   password,
			RememberMe: true,
		},
		Token:     "someValue",
		UserToken: "someValue",
	}
	body, err := json.Marshal(login)
	if err != nil {
		return nil, errors.Wrap(err, "problem encoding login request")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "problem creating request to send to server")
	}

	result, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var out struct {
		User *User `json:"golfer_user"`
	}
	err = json.Unmarshal(result, &out)
	if err != nil {
		return nil, errors.Wrap(err, "problem unmarshalling result")
	}
	c.authToken = out.User.GolferUserToken

	return out.User, nil
}

func (c *httpClient) Post(ctx context.Context, path string, body []byte) ([]byte, error) {
	requestURL, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return nil, errors.Wrap(err, "problem building request url")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "problem creating request to send to server")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return c.sendRequest(req)
}

func (c *httpClient) Get(ctx context.Context, path string, params url.Values) ([]byte, error) {
	requestURL, err := url.JoinPath(c.baseURL, path)
	if err != nil {
		return nil, errors.Wrap(err, "problem building request url")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "problem creating request to send to server")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return c.sendRequest(req)
}

// sendRequest sends a single http.Request and verifies that the response code is valid.
func (c *httpClient) sendRequest(r *http.Request) ([]byte, error) {
	// Set the auth token of the request
	if c.authToken != "" {
		r.Header.Set("Authorization", c.authToken)
	}

	// Perform Operation
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}

	// Evaluate response
	_, ok := successfulResponseCodes[resp.StatusCode]
	if !ok {
		return nil, errors.Errorf("non-success status returned from request %s: %d", r.URL.Path, resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO Log
		}
	}(resp.Body)
	return respBody, nil
}

var successfulResponseCodes = map[int]bool{
	http.StatusOK:      true,
	http.StatusCreated: true,
}
