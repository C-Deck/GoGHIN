package ghin

type (
	FacilityStatus string

	Facility struct {
		FacilityId        int            `json:"FacilityId"`
		FacilityStatus    FacilityStatus `json:"FacilityStatus"`
		FacilityName      string         `json:"FacilityName"`
		FacilityNumber    *string        `json:"FacilityNumber"`
		GolfAssociationId int            `json:"GolfAssociationId"`
	}

	Season struct {
		SeasonName      string `json:"SeasonName"`
		SeasonStartDate string `json:"SeasonStartDate"`
		SeasonEndDate   string `json:"SeasonEndDate"`
		IsAllYear       bool   `json:"IsAllYear"`
	}

	HoleDetails struct {
		Number     int `json:"Number"`
		HoleId     int `json:"HoleId"`
		Length     int `json:"Length"`
		Par        int `json:"Par"`
		Allocation int `json:"Allocation"`
	}

	TeeSetRatingType string

	TeeSetGender string

	TeeSetRating struct {
		TeeSetRatingType string  `json:"RatingType"`
		CourseRating     float64 `json:"CourseRating"`
		SlopeRating      float64 `json:"SlopeRating"`
		BogeyRating      float64 `json:"BogeyRating"`
	}
	TeeSetDetails struct {
		Ratings          []TeeSetRating `json:"Ratings"`
		Holes            []HoleDetails  `json:"Holes"`
		TeeSetRatingId   int            `json:"TeeSetRatingId"`
		TeeSetRatingName string         `json:"TeeSetRatingName"`
		Gender           TeeSetGender   `json:"Gender"`
		HolesNumber      HolesPlayed    `json:"HolesNumber"`
		TotalYardage     int            `json:"TotalYardage"`
		TotalMeters      int            `json:"TotalMeters"`
		LegacyCRPTeeId   int            `json:"LegacyCRPTeeId"`
		StrokeAllocation bool           `json:"StrokeAllocation"`
		TotalPar         int            `json:"TotalPar"`
	}

	CourseStatus string

	CourseDetails struct {
		Facility     Facility        `json:"Facility"`
		Season       Season          `json:"Season"`
		TeeSets      []TeeSetDetails `json:"TeeSets"`
		CourseId     int             `json:"CourseId"`
		CourseName   string          `json:"CourseName"`
		CourseStatus string          `json:"CourseStatus"`
		CourseNumber *string         `json:"CourseNumber"`
		CourseCity   string          `json:"CourseCity"`
		CourseState  string          `json:"CourseState"`
	}

	CourseOverview struct {
		CourseID             int                    `json:"CourseID"`
		CourseStatus         string                 `json:"CourseStatus"`
		CourseName           string                 `json:"CourseName"`
		GeoLocationLatitude  int                    `json:"GeoLocationLatitude"`
		GeoLocationLongitude int                    `json:"GeoLocationLongitude"`
		FacilityID           int                    `json:"FacilityID"`
		FacilityStatus       string                 `json:"FacilityStatus"`
		FacilityName         string                 `json:"FacilityName"`
		FullName             string                 `json:"FullName"`
		Address1             string                 `json:"Address1"`
		Address2             *string                `json:"Address2"`
		City                 string                 `json:"City"`
		State                string                 `json:"State"`
		Zip                  string                 `json:"Zip"`
		Country              string                 `json:"Country"`
		EntCountryCode       int                    `json:"EntCountryCode"`
		EntStateCode         int                    `json:"EntStateCode"`
		LegacyCRPCourseId    int                    `json:"LegacyCRPCourseId"`
		Telephone            string                 `json:"Telephone"`
		Email                *string                `json:"Email"`
		UpdatedOn            string                 `json:"UpdatedOn"`
		Ratings              []CourseOverviewRating `json:"Ratings"`
	}
	CourseOverviewRating struct {
		TeeSetRatingId   int    `json:"TeeSetRatingId"`
		TeeSetRatingName string `json:"TeeSetRatingName"`
		TeeSetStatus     string `json:"TeeSetStatus"`
	}
)

const (
	TeeSetGenderMale   TeeSetGender = "Male"
	TeeSetGenderFemale TeeSetGender = "Female"

	FacilityStatusActive FacilityStatus = "Active"

	CourseStatusActive CourseStatus = "Active"

	TeeSetRatingTypeTotal TeeSetRatingType = "Total"
	TeeSetRatingTypeFront TeeSetRatingType = "Front"
	TeeSetRatingTypeBack  TeeSetRatingType = "Back"
)
