package ghin

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	client *httpClient
	user   *User
}

func (c *Client) Login(ctx context.Context, email, password string) error {
	user, err := c.client.Login(ctx, email, password)
	if err != nil {
		return err
	}
	c.user = user
	return nil
}

type SubmitScoreInput struct {
	// Gender is the gender of the golfer.
	Gender PlayerGender `json:"gender"`

	// CourseID is the ID of the course played.
	CourseID int `json:"course_id"`

	// TeeSetID is the ID of the tee box that was played (whites, blues, etc) for the course.
	TeeSetID int `json:"tee_set_id"`

	/* TeeSetSide represents the which holes were played. Front/Back 9 or that 18 holes.
	 * Default: All18
	 */
	TeeSetSide *TeeSetSide `json:"tee_set_side"`

	/* PlayedAt is the date the round was played.
	 * Default: Current date
	 */
	PlayedAt *time.Time `json:"played_at"`

	// HoleDetails are the scores and statistics for each hole.
	HoleDetails []HoleScore `json:"hole_details"`

	/* NumberOfHoles represents whether 9 or 18 holes were played.
	 * Default: 18
	 */
	NumberOfHoles *HolesPlayed `json:"number_of_holes"`

	/* ScoreType is the situation the round was played (Away, Home, or Tournament).
	 * Default: Away
	 */
	ScoreType *ScoringType `json:"score_type"`
}

// SubmitScore submits a score for a golfer.
func (c *Client) SubmitScore(ctx context.Context, input SubmitScoreInput) (*Score, error) {
	if c.user == nil {
		return nil, NewUserNotLoggedInError("cannot submit score without user login")
	}

	submission := ScoreSubmission{
		GolferID:      strconv.Itoa(c.user.GolferId),
		Gender:        input.Gender,
		CourseID:      input.CourseID,
		ScoreType:     ScoringTypeAway,
		PlayedAt:      ToPlayedAtString(time.Now()),
		HoleDetails:   input.HoleDetails,
		NumberOfHoles: EighteenHolesPlayed,
		TeeSetSide:    TeeSetSide18,
		TeeSetID:      input.TeeSetID,
	}
	if input.TeeSetSide != nil {
		submission.TeeSetSide = *input.TeeSetSide
	}
	if input.NumberOfHoles != nil {
		submission.NumberOfHoles = *input.NumberOfHoles
	}
	if input.ScoreType != nil {
		submission.ScoreType = *input.ScoreType
	}
	if input.PlayedAt != nil {
		submission.PlayedAt = ToPlayedAtString(*input.PlayedAt)
	}
	out, err := postAndDeserialize[Score, ScoreSubmission](c.client, ctx, postScorePath, submission)
	if err != nil {
		return nil, errors.Wrapf(err, "")
	}

	return out, nil
}

type GetUserInfoInput struct {
	Offset *int
	Limit  *int
	// Statuses are the status
	Statuses *string
}

func (c *Client) GetUserInfo(ctx context.Context, input GetUserInfoInput) (*GolferScores, error) {

}

type GetCourseDetailsInput struct {
	CourseID           string
	IncludeAlteredTees *bool
}

func (c *Client) GetCourseDetails(ctx context.Context, input GetCourseDetailsInput) (*CourseDetails, error) {
	params := url.Values{}
	var includeAlteredTees bool
	if input.IncludeAlteredTees != nil {
		includeAlteredTees = *input.IncludeAlteredTees
	}
	params.Set("courseId", input.CourseID)
	params.Set("include_altered_tees", strconv.FormatBool(includeAlteredTees))

	out, err := getAndDeserialize[CourseDetails](c.client, ctx, courseDetailsPath, params)
	if err != nil {
		return nil, errors.Wrapf(err, "problem retrieving course %q details", input.CourseID)
	}

	return out, nil
}

type SearchCoursesInput struct {
	CourseName     *string
	FacilityID     *int
	Country        *string
	State          *string
	CourseStatus   *CourseStatus
	FacilityStatus *FacilityStatus
	Offset         *int
	Limit          *int
	IncludeTeeSets *bool
}

type searchCoursesResponse struct {
	Courses []CourseOverview `json:"courses"`
}

func (c *Client) SearchCourses(ctx context.Context, input SearchCoursesInput) ([]CourseOverview, error) {
	params := url.Values{}
	if input.CourseName != nil {
		params.Set("name", *input.CourseName)
	}
	if input.FacilityID != nil {
		params.Set("facility_id", strconv.Itoa(*input.FacilityID))
	}
	if input.Country != nil {
		params.Set("country", *input.Country)
	}
	if input.State != nil {
		params.Set("state", *input.State)
	}
	if input.CourseStatus != nil {
		params.Set("course_status", string(*input.CourseStatus))
	}
	if input.FacilityStatus != nil {
		params.Set("facility_status", string(*input.FacilityStatus))
	}
	if input.IncludeTeeSets != nil {
		params.Set("include_tee_sets", strconv.FormatBool(*input.IncludeTeeSets))
	}
	if input.Limit != nil {
		params.Set("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		params.Set("offset", strconv.Itoa(*input.Offset))
	}
	out, err := getAndDeserialize[searchCoursesResponse](c.client, ctx, searchCoursePath, params)
	if err != nil {
		return nil, errors.Wrapf(err, "problem searching for course details")
	}

	return out.Courses, nil
}

func postAndDeserialize[T any, K any](client *httpClient, ctx context.Context, path string, input K) (*T, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrapf(err, "problem serializing data for post request")
	}

	data, err := client.Post(ctx, path, body)
	if err != nil {
		return nil, err
	}

	var out T
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, errors.Wrapf(err, "problem deserializing response from post request")
	}

	return &out, nil
}

func getAndDeserialize[T any](client *httpClient, ctx context.Context, path string, params url.Values) (*T, error) {
	data, err := client.Get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var out T
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, errors.Wrapf(err, "problem deserializing response from get request")
	}

	return &out, nil
}
