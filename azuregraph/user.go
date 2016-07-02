package azuregraph

import (
	"encoding/json"
	"net/url"
)

// UserGet gets a specified user. You can use either the object ID (GUID) or
// the user principal name (UPN) to identify the target user
func (d *Dispatcher) UserGet(userID string) (*User, error) {
	var user User
	endpoint, err := d.getEndpoint("user", userID)
	if err != nil {
		return nil, err
	}
	buf, err := d.dispatch("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UserList gets users. You can add query parameters to the request to filter,
// sort and page the response.
func (d *Dispatcher) UserList(query *OdataQuery) (*[]User, *string, error) {
	var users struct {
		NextLink string `json:"odata.nextLink"`
		Value    []User `json:"value"`
	}
	endpoint, err := d.getEndpoint("user")
	if err != nil {
		return nil, nil, err
	}
	values := endpoint.Query()
	if query != nil {
		query.setQuery(&values)
		endpoint.RawQuery = values.Encode()
	}
	buf, err := d.dispatch("GET", endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(buf, &users); err != nil {
		return nil, nil, err
	}
	if users.NextLink != "" {
		nextLinkURL, err := url.Parse(users.NextLink)
		if err != nil {
			return nil, nil, err
		}
		skiptoken := nextLinkURL.Query().Get("$skiptoken")
		return &users.Value, &skiptoken, nil
	}
	return &users.Value, nil, nil
}

// User FIXME
type User struct {
	AccountEnabled               bool                `json:"accountEnabled"`
	AssignedLicenses             []AssignedLicense   `json:"assignedLicenses"`
	AssignedPlans                []AssignedPlan      `json:"assignedPlans"`
	City                         string              `json:"city"`
	Country                      string              `json:"country"`
	CreationType                 string              `json:"creationType"`
	DeletionTimestamp            string              `json:"deletionTimestamp"`
	Department                   string              `json:"department"`
	DirSyncEnabled               bool                `json:"dirSyncEnabled"`
	DisplayName                  string              `json:"displayName"`
	FacsimileTelephoneNumber     string              `json:"facsimileTelephoneNumber"`
	GivenName                    string              `json:"givenName"`
	ImmutableID                  string              `json:"immutableId"`
	JobTitle                     string              `json:"jobTitle"`
	LastDirSyncTime              string              `json:"lastDirSyncTime"`
	Mail                         string              `json:"mail"`
	MailNickname                 string              `json:"mailNickname"`
	Mobile                       string              `json:"mobile"`
	ObjectID                     string              `json:"objectId"`
	ObjectType                   string              `json:"objectType"`
	OnPremisesSecurityIdentifier string              `json:"onPremisesSecurityIdentifier"`
	OtherMails                   []string            `json:"otherMails"`
	PasswordPolicies             string              `json:"passwordPolicies"`
	PasswordProfile              PasswordProfile     `json:"passwordProfile"`
	PhysicalDeliveryOfficeName   string              `json:"physicalDeliveryOfficeName"`
	PostalCode                   string              `json:"postalCode"`
	PreferredLanguage            string              `json:"preferredLanguage"`
	ProvisionedPlans             []ProvisionedPlan   `json:"provisionedPlans"`
	ProvisioningErrors           []ProvisioningError `json:"provisioningErrors"`
	ProxyAddresses               []string            `json:"proxyAddresses"`
	SignInNames                  []SignInName        `json:"signInNames"`
	SipProxyAddress              string              `json:"sipProxyAddress"`
	State                        string              `json:"state"`
	StreetAddress                string              `json:"streetAddress"`
	Surname                      string              `json:"surname"`
	TelephoneNumber              string              `json:"telephoneNumber"`
	ThumbnailPhoto               string              `json:"thumbnailPhoto"`
	UsageLocation                string              `json:"usageLocation"`
	UserPrincipalName            string              `json:"userPrincipalName"`
	UserType                     string              `json:"userType"`
}
