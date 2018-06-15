package maxmind

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	loadEnv()
}

var (
	// URL of maxMind service, overwritten via ENV variable.
	URL = ""
)

// IP is an IPv4 address.
type IP string

// Email is an ISO email address.
type Email string

// RiskCheck queries the minFraud service from MaxMind for Credit Card
// fraud patterns matching.
func RiskCheck(firstSixDigits, lastFourDigits string, IPAddress IP, emailAddress Email) (score float64, err error) {
	// Checks
	if len(firstSixDigits) < 6 || len(lastFourDigits) < 4 {
		return 0.0, fmt.Errorf("error")
	}
	if IPAddress == "" {
		IPAddress = "1.1.1.1"
	}

	r := request{
		CreditCard: struct {
			CvvResult      string "json:\"cvv_result\""
			IssuerIDNumber string "json:\"issuer_id_number\""
			Last4Digits    string "json:\"last_4_digits\""
		}{
			CvvResult:      "Y",
			IssuerIDNumber: firstSixDigits,
			Last4Digits:    lastFourDigits,
		},
		Device: struct {
			IPAddress string "json:\"ip_address\""
		}{
			IPAddress: string(IPAddress),
		},
		Email: struct {
			Address string "json:\"address,omitempty\""
		}{
			Address: string(emailAddress),
		},
	}
	b, err := json.Marshal(r)
	if err != nil {
		return 0.0, fmt.Errorf("Unable to marshall: ", r)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(b))
	req.Header.Add("Accept", "application/vnd.maxmind.com-minfraud-score+json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", encodeAuth()))
	if err != nil {
		return 0.0, fmt.Errorf("NewRequest failed: %s", err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0.0, fmt.Errorf("No response from server: %v", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0.0, fmt.Errorf("Unable to read data: %v", err)
	}
	var responseData Response
	if err := json.Unmarshal(data, &responseData); err != nil {
		return 0.0, fmt.Errorf("Unable to read data: %v", err)
	}
	fmt.Println(responseData)
	return responseData.RiskScore, nil
}

func encodeAuth() string {
	user := os.Getenv("MAXMIND_USER")
	pass := os.Getenv("MAXMIND_PASSWORD")
	combined := fmt.Sprintf("%v:%v", user, pass)
	return base64.StdEncoding.EncodeToString([]byte(combined))
}

func BasicRiskCheck()    {}
func ExtendedRiskCheck() {}
