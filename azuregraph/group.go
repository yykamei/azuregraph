package azuregraph

import (
	"encoding/json"
	"net/url"
)

// GroupGet gets a specified group. Specify the group by its object ID (GUID).
func (d *Dispatcher) GroupGet(objectID string) (*Group, error) {
	var group Group
	if err := d.get(objectID, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

// GroupList gets groups. You can add query parameters to the request to filter,
// sort and page the response.
func (d *Dispatcher) GroupList(query *OdataQuery) (*[]Group, *string, error) {
	var groups struct {
		NextLink string  `json:"odata.nextLink"`
		Value    []Group `json:"value"`
	}
	endpoint, err := d.getEndpoint("group")
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
	if err := json.Unmarshal(buf, &groups); err != nil {
		return nil, nil, err
	}
	if groups.NextLink != "" {
		nextLinkURL, err := url.Parse(groups.NextLink)
		if err != nil {
			return nil, nil, err
		}
		skiptoken := nextLinkURL.Query().Get("$skiptoken")
		return &groups.Value, &skiptoken, nil
	}
	return &groups.Value, nil, nil
}

// Group FIXME
type Group struct {
	DeletionTimestamp            string              `json:"deletionTimestamp"`
	Description                  string              `json:"description"`
	DirSyncEnabled               bool                `json:"dirSyncEnabled"`
	DisplayName                  string              `json:"displayName"`
	LastDirSyncTime              string              `json:"lastDirSyncTime"`
	Mail                         string              `json:"mail"`
	MailEnabled                  bool                `json:"mailEnabled"`
	MailNickname                 string              `json:"mailNickname"`
	ObjectID                     string              `json:"objectId"`
	ObjectType                   string              `json:"objectType"`
	OnPremisesSecurityIdentifier string              `json:"onPremisesSecurityIdentifier"`
	ProvisioningErrors           []ProvisioningError `json:"provisioningErrors"`
	ProxyAddresses               []string            `json:"proxyAddresses"`
	SecurityEnabled              bool                `json:"securityEnabled"`
}

func (group *Group) resourceName() string {
	return "group"
}
