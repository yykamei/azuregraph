package azuregraph

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const apiVersion = "1.6"

// Dispatcher dispach request.
type Dispatcher struct {
	tenantID     string
	clientID     string
	clientSecret string
	accessToken  string
	tokenType    string
}

// NewDispatcher constructs Dispatcher object
func NewDispatcher(tenantID string, clientID string, clientSecret string) (*Dispatcher, error) {
	tokenInfo, err := getToken(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	dispatcher := &Dispatcher{
		tenantID:     tenantID,
		clientID:     clientID,
		clientSecret: clientSecret,
		accessToken:  tokenInfo.AccessToken,
		tokenType:    tokenInfo.TokenType,
	}
	return dispatcher, nil
}

func (d *Dispatcher) getEndpoint(resourceType string, paths ...string) (*url.URL, error) {
	switch resourceType {
	case "user":
		resourceType = "users"
	case "group":
		resourceType = "groups"
	default:
		return nil, errors.New("Undefined resource type")
	}
	query := url.Values{
		"api-version": []string{apiVersion},
	}
	path := fmt.Sprintf("/%s/%s", d.tenantID, resourceType)
	if len(paths) > 0 {
		path = fmt.Sprintf("%s/%s", path, strings.Join(paths, "/"))
	}
	u := &url.URL{
		Scheme:   "https",
		Host:     "graph.windows.net",
		RawQuery: query.Encode(),
		Path:     path,
	}
	return u, nil
}

func (d *Dispatcher) dispatch(method string, requestURL *url.URL, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, requestURL.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", d.tokenType, d.accessToken))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// TokenInfo is Token data sent by Azure AD.
// This includes OAuth2 AccessToken.
type TokenInfo struct {
	NotBefore    string `json:"not_before"`
	Scope        string `json:"scope"`
	Resource     string `json:"resource"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	ExtExpiresIn string `json:"ext_expires_in"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
}

// getToken do authenticate by OAuth2 FIXME: expired?
func getToken(tenantID string, clientID string, clientSecret string) (*TokenInfo, error) {
	var tokenInfo TokenInfo
	req, err := http.NewRequest(
		"POST",
		getTokenEndpoint(tenantID),
		bytes.NewBufferString(url.Values{
			"grant_type":    []string{"client_credentials"},
			"client_id":     []string{clientID},
			"client_secret": []string{clientSecret},
		}.Encode()),
	)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf, &tokenInfo); err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

func getTokenEndpoint(tenantID string) string {
	query := url.Values{
		"api-version": []string{apiVersion},
	}
	u := url.URL{
		Scheme:   "https",
		Host:     "login.windows.net",
		RawQuery: query.Encode(),
		Path:     fmt.Sprintf("/%s/oauth2/token", tenantID),
	}
	return u.String()
}

// AssignedLicense FIXME
type AssignedLicense struct {
	SkuID         string   `json:"skuId"`
	DisabledPlans []string `json:"disabledPlans"`
}

// AssignedPlan FIXME
type AssignedPlan struct {
	ServicePlanID     string `json:"servicePlanId"`
	AssignedTimestamp string `json:"assignedTimestamp"`
	CapabilityStatus  string `json:"capabilityStatus"`
	Service           string `json:"service"`
}

// PasswordProfile FIXME
type PasswordProfile struct {
	Password                     string `json:"password"`
	ForceChangePasswordNextLogin bool   `json:"forceChangePasswordNextLogin"`
}

// ProvisionedPlan FIXME
type ProvisionedPlan struct {
	CapabilityStatus   string `json:"capabilityStatus"`
	ProvisioningStatus string `json:"provisioningStatus"`
	Service            string `json:"service"`
}

// ProvisioningError FIXME
type ProvisioningError struct {
	ErrorDetail     string `json:"errorDetail"`
	Resolved        bool   `json:"resolved"`
	ServiceInstance string `json:"serviceInstance"`
	Timestamp       string `json:"timestamp"`
}

// SignInName FIXME
type SignInName struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
