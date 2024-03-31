package ghin

import (
	"time"
)

type (
	// ShotAccuracy represents the possible values for shot statistics.
	ShotAccuracy int

	// TeeSetSide represents the holes that were played.
	TeeSetSide string

	// ScoringType represents the options for when/where the round was played. (Away, Home, or Tournament)
	ScoringType string

	// PlayerGender is the gender of the player
	PlayerGender string

	HolesPlayed int

	ScoreStatus string
)

const (
	ShotAccuracyMissedLeft  ShotAccuracy = 0
	ShotAccuracyMissedRight ShotAccuracy = 1
	ShotAccuracyMissedLong  ShotAccuracy = 2
	ShotAccuracyMissedShort ShotAccuracy = 3

	TeeSetSide18    TeeSetSide = "ALL18"
	TeeSetSideFront TeeSetSide = "F9"
	TeeSetSideBack  TeeSetSide = "B9"

	ScoringTypeHome       = "H"
	ScoringTypeAway       = "A"
	ScoringTypeTournament = "T"

	PlayerGenderMale   PlayerGender = "M"
	PlayerGenderFemale PlayerGender = "F"

	EighteenHolesPlayed HolesPlayed = 18
	NineHolesPlayed     HolesPlayed = 9

	ScoreStatusValidated ScoreStatus = "Validated"
)

// LongString returns the non-abbreviated version of the gender.
func (g PlayerGender) LongString() string {
	if g == PlayerGenderMale {
		return "Male"
	}
	return "Female"
}

type ScoreSubmission struct {
	// GolferID is the ID of the golfer that the ScoreSubmission is for.
	GolferID string `json:"golfer_id"`
	// Gender is the gender of the golfer.
	Gender PlayerGender `json:"gender"`
	// CourseID is the ID of the course played.
	CourseID int `json:"course_id"`
	// TeeSetID is the ID of the tee box that was played (whites, blues, etc) for the course.
	TeeSetID int `json:"tee_set_id"`
	// TeeSetSide represents the which holes were played. Front/Back 9 or that 18 holes.
	TeeSetSide TeeSetSide `json:"tee_set_side"`
	// PlayedAt is the date the round was played.
	PlayedAt string `json:"played_at"`
	// HoleDetails are the scores and statistics for each hole.
	HoleDetails []HoleScore `json:"hole_details"`
	// NumberOfHoles represents whether 9 or 18 holes were played.
	NumberOfHoles HolesPlayed `json:"number_of_holes"`
	// ScoreType is the situation the round was played (Away, Home, or Tournament).
	ScoreType            ScoringType `json:"score_type"`
	OverrideConfirmation *bool       `json:"override_confirmation,omitempty"`
	IsManual             *bool       `json:"is_manual"`
	Source               *string     `json:"source,omitempty"`
}

type HoleScoreDetails struct {
	HoleScore
	AdjustedGrossScore int `json:"adjusted_gross_score"`
}

type HoleDetail struct {
	Id                 int `json:"id"`
	AdjustedGrossScore int `json:"adjusted_gross_score"`
	// HoleNumber is the hole number the score is for.
	HoleNumber int `json:"hole_number"`
	// RawScore is the score for the hole.
	RawScore int `json:"raw_score"`

	Par              int  `json:"par"`
	StrokeAllocation int  `json:"stroke_allocation"`
	XHole            bool `json:"x_hole"`
	MostLikelyScore  *int `json:"most_likely_score"`

	// Optional hole statistic values
	Putts *int `json:"putts"`
	// FairwayHit is whether the fairway was hit for the hole.
	FairwayHit *bool `json:"fairway_hit"`
	// GreenInRegulation is whether the player got to the green in regulation.
	GreenInRegulation *bool `json:"gir_flag"`
	// DriveAccuracy represents where the drive missed the fairway (should only be provided if FairwayHit is false).
	DriveAccuracy *ShotAccuracy `json:"drive_accuracy"`
	// ApproachShotAccuracy represents where the approach shot missed the green (should only be provided if GreenInRegulation is false).
	ApproachShotAccuracy *ShotAccuracy `json:"approach_shot_accuracy,omitempty"`
}

type HoleScore struct {
	// HoleNumber is the hole number the score is for.
	HoleNumber int `json:"hole_number"`
	// RawScore is the score for the hole.
	RawScore int `json:"raw_score"`
	Par      int `json:"par,omitempty"`
	// Optional hole statistic values
	Putts *int `json:"putts,omitempty"`
	// FairwayHit is whether the fairway was hit for the hole.
	FairwayHit *bool `json:"fairway_hit,omitempty"`
	// GreenInRegulation is whether the player got to the green in regulation.
	GreenInRegulation *bool `json:"gir_flag,omitempty"`
	// DriveAccuracy represents where the drive missed the fairway (should only be provided if FairwayHit is false).
	DriveAccuracy *ShotAccuracy `json:"drive_accuracy,omitempty"`
	// ApproachShotAccuracy represents where the approach shot missed the green (should only be provided if GreenInRegulation is false).
	ApproachShotAccuracy *ShotAccuracy `json:"approach_shot_accuracy,omitempty"`
}

type RoundStatistics struct {
	GirPercent                               int     `json:"gir_percent"`
	PuttsTotal                               int     `json:"putts_total"`
	OnePuttOrBetterPercent                   float64 `json:"one_putt_or_better_percent"`
	TwoPuttPercent                           float64 `json:"two_putt_percent"`
	ThreePuttOrWorsePercent                  float64 `json:"three_putt_or_worse_percent"`
	TwoPuttOrBetterPercent                   float64 `json:"two_putt_or_better_percent"`
	UpAndDownsTotal                          int     `json:"up_and_downs_total"`
	ParsPercent                              float64 `json:"pars_percent"`
	Par3SAverage                             float64 `json:"par3s_average"`
	Par4SAverage                             float64 `json:"par4s_average"`
	Par5SAverage                             float64 `json:"par5s_average"`
	BogeysPercent                            float64 `json:"bogeys_percent"`
	MissedLeftPercent                        int     `json:"missed_left_percent"`
	MissedLongPercent                        int     `json:"missed_long_percent"`
	FairwayHitsPercent                       int     `json:"fairway_hits_percent"`
	MissedRightPercent                       int     `json:"missed_right_percent"`
	MissedShortPercent                       int     `json:"missed_short_percent"`
	DoubleBogeysPercent                      float64 `json:"double_bogeys_percent"`
	BirdiesOrBetterPercent                   int     `json:"birdies_or_better_percent"`
	TripleBogeysOrWorsePercent               float64 `json:"triple_bogeys_or_worse_percent"`
	MissedLeftApproachShotAccuracyPercent    int     `json:"missed_left_approach_shot_accuracy_percent"`
	MissedRightApproachShotAccuracyPercent   int     `json:"missed_right_approach_shot_accuracy_percent"`
	MissedLongApproachShotAccuracyPercent    int     `json:"missed_long_approach_shot_accuracy_percent"`
	MissedShortApproachShotAccuracyPercent   int     `json:"missed_short_approach_shot_accuracy_percent"`
	MissedGeneralApproachShotAccuracyPercent int     `json:"missed_general_approach_shot_accuracy_percent"`
}

type GolferScores struct {
	Scores       []Score `json:"scores"`
	TotalCount   int     `json:"total_count"`
	HighestScore int     `json:"highest_score"`
	LowestScore  int     `json:"lowest_score"`
	Average      float64 `json:"average"`
}

type Score struct {
	Id                       int               `json:"id"`
	OrderNumber              int               `json:"order_number"`
	ScoreDayOrder            int               `json:"score_day_order"`
	Gender                   PlayerGender      `json:"gender"`
	Status                   ScoreStatus       `json:"status"`
	IsManual                 bool              `json:"is_manual"`
	NumberOfHoles            HolesPlayed       `json:"number_of_holes"`
	NumberOfPlayedHoles      HolesPlayed       `json:"number_of_played_holes"`
	GolferId                 string            `json:"golfer_id"`
	CourseId                 string            `json:"course_id"`
	CourseName               string            `json:"course_name"`
	FacilityName             *string           `json:"facility_name"`
	PlayedAt                 string            `json:"played_at"`
	AdjustedGrossScore       int               `json:"adjusted_gross_score"`
	PostedOnHomeCourse       *bool             `json:"posted_on_home_course"`
	Differential             float64           `json:"differential"`
	UnadjustedDifferential   float64           `json:"unadjusted_differential"`
	ScoreType                ScoringType       `json:"score_type"`
	Front9CourseName         *string           `json:"front9_course_name"`
	Back9CourseName          *string           `json:"back9_course_name"`
	Front9CourseRating       *string           `json:"front9_course_rating"`
	Back9CourseRating        *string           `json:"back9_course_rating"`
	TeeName                  *string           `json:"tee_name"`
	TeeSetId                 string            `json:"tee_set_id"`
	TeeSetSide               TeeSetSide        `json:"tee_set_side"`
	CourseRating             float64           `json:"course_rating"`
	SlopeRating              int               `json:"slope_rating"`
	Penalty                  *bool             `json:"penalty"`
	PenaltyType              *string           `json:"penalty_type"`
	PenaltyMethod            *string           `json:"penalty_method"`
	ParentId                 *int              `json:"parent_id"`
	ScoreTypeDisplayFull     string            `json:"score_type_display_full"`
	ScoreTypeDisplayShort    string            `json:"score_type_display_short"`
	Edited                   bool              `json:"edited"`
	PostedAt                 time.Time         `json:"posted_at"`
	SeasonStartDateAt        string            `json:"season_start_date_at"`
	SeasonEndDateAt          string            `json:"season_end_date_at"`
	CourseDisplayValue       string            `json:"course_display_value"`
	GhinCourseNameDisplay    string            `json:"ghin_course_name_display"`
	Used                     bool              `json:"used"`
	Revision                 bool              `json:"revision"`
	Pcc                      *int              `json:"pcc"`
	Adjustments              []ScoreAdjustment `json:"adjustments"`
	EstimatedHandicap        float64           `json:"estimated_handicap"`
	EstimatedHandicapDisplay string            `json:"estimated_handicap_display"`
	HoleDetails              []HoleScore       `json:"hole_details"`
	Statistics               *RoundStatistics  `json:"statistics,omitempty"`
	Exceptional              bool              `json:"exceptional"`
	IsRecent                 bool              `json:"is_recent"`
	ESR                      *int              `json:"ESR"`
	NetScoreDifferential     float64           `json:"net_score_differential"`
}

type ScoreAdjustment struct {
	Type    string  `json:"type"`
	Value   float64 `json:"value"`
	Display string  `json:"display"`
}

const (
	playedAtDateFormat = "2006-01-02"
)

func ToPlayedAtString(t time.Time) string {
	return t.Format(playedAtDateFormat)
}
