package maxmind

import "time"

type request struct {
	CreditCard struct {
		CvvResult      string `json:"cvv_result"`
		IssuerIDNumber string `json:"issuer_id_number"`
		Last4Digits    string `json:"last_4_digits"`
	} `json:"credit_card"`
	Device struct {
		IPAddress string `json:"ip_address"`
	} `json:"device"`
	Email struct {
		Address string `json:"address,omitempty"`
	} `json:"email"`
}

// Response stores all the data returned for MaxMind.
type Response struct {
	IPAddress struct {
		Risk    float64 `json:"risk"`
		Country struct {
			IsHighRisk bool   `json:"is_high_risk"`
			Confidence int    `json:"confidence"`
			IsoCode    string `json:"iso_code"`
			GeonameID  int    `json:"geoname_id"`
			Names      struct {
			} `json:"names,omitempty"`
		} `json:"country"`
		Location struct {
			LocalTime      time.Time `json:"local_time"`
			AccuracyRadius int       `json:"accuracy_radius"`
			Latitude       float64   `json:"latitude"`
			Longitude      float64   `json:"longitude"`
			TimeZone       string    `json:"time_zone"`
		} `json:"location"`
		City struct {
			Confidence int `json:"confidence"`
			GeonameID  int `json:"geoname_id"`
		} `json:"city"`
		Continent struct {
			Code      string `json:"code"`
			GeonameID int    `json:"geoname_id"`
		} `json:"continent"`
		Postal struct {
			Confidence int    `json:"confidence"`
			Code       string `json:"code"`
		} `json:"postal"`
		RegisteredCountry struct {
			IsoCode   string `json:"iso_code"`
			GeonameID int    `json:"geoname_id"`
		} `json:"registered_country"`
		Subdivisions []struct {
			Confidence int    `json:"confidence"`
			IsoCode    string `json:"iso_code"`
			GeonameID  int    `json:"geoname_id"`
		} `json:"subdivisions"`
		Traits struct {
			UserType                     string `json:"user_type"`
			AutonomousSystemNumber       int    `json:"autonomous_system_number"`
			AutonomousSystemOrganization string `json:"autonomous_system_organization"`
			Isp                          string `json:"isp"`
			Organization                 string `json:"organization"`
			IPAddress                    string `json:"ip_address"`
		} `json:"traits"`
	} `json:"ip_address"`
	CreditCard struct {
		Issuer struct {
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
		} `json:"issuer"`
		Brand     string `json:"brand"`
		Country   string `json:"country"`
		IsPrepaid bool   `json:"is_prepaid"`
		IsVirtual bool   `json:"is_virtual"`
		Type      string `json:"type"`
	} `json:"credit_card"`
	ID               string  `json:"id"`
	RiskScore        float64 `json:"risk_score"`
	FundsRemaining   float64 `json:"funds_remaining"`
	QueriesRemaining int     `json:"queries_remaining"`
}
