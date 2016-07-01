package azuregraph

import "encoding/json"

// GroupGet gets a specified group. Specify the group by its object ID (GUID).
func (d *Dispatcher) GroupGet(objectID string) (*Group, error) {
	var group Group
	endpoint, err := d.getEndpoint("group", objectID)
	if err != nil {
		return nil, err
	}
	buf, err := d.dispatch("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

// GroupList gets groups. You can add query parameters to the request to filter,
// sort and page the response.
func (d *Dispatcher) GroupList(query *OdataQuery) (*[]Group, error) {
	var groups struct {
		Value []Group `json:"value"`
	}
	endpoint, err := d.getEndpoint("group")
	if err != nil {
		return nil, err
	}
	values := endpoint.Query()
	query.setQuery(&values)
	endpoint.RawQuery = values.Encode()
	buf, err := d.dispatch("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(buf, &groups); err != nil {
		return nil, err
	}
	return &groups.Value, nil
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
