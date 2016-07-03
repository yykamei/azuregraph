package azuregraph

// SubscribedSkuGet gets a specified subscribedSku. Specify the group by its object ID (GUID).
func (d *Dispatcher) SubscribedSkuGet(objectID string) (*SubscribedSku, error) {
	var subscribedSku SubscribedSku
	if err := d.get(objectID, &subscribedSku); err != nil {
		return nil, err
	}
	return &subscribedSku, nil
}

// SubscribedSkuList gets subscribedSkus. You can add query parameters to the request to filter,
// sort and page the response.
func (d *Dispatcher) SubscribedSkuList(query *OdataQuery) (*[]SubscribedSku, *string, error) {
	var subscribedSkus SubscribedSkus
	if skiptoken, err := d.list(query, &subscribedSkus); err != nil {
		return nil, nil, err
	} else if skiptoken != nil {
		return &subscribedSkus.Value, skiptoken, nil
	}
	return &subscribedSkus.Value, nil, nil
}

// LicenseUnitsDetail FIXME
type LicenseUnitsDetail struct {
	Enabled   int `json:"enabled"`
	Suspended int `json:"suspended"`
	Warning   int `json:"warning"`
}

// ServicePlanInfo FIXME
type ServicePlanInfo struct {
	AppliesTo          string `json:"appliesTo"`
	ProvisioningStatus string `json:"provisioningStatus"`
	ServicePlanID      string `json:"servicePlanId"`
	ServicePlanName    string `json:"servicePlanName"`
}

// SubscribedSku FIXME
type SubscribedSku struct {
	AppliesTo        string             `json:"appliesTo"`
	CapabilityStatus string             `json:"capabilityStatus"`
	ConsumedUnits    int                `json:"consumedUnits"`
	ObjectID         string             `json:"objectId"`
	PrepaidUnits     LicenseUnitsDetail `json:"prepaidUnits"`
	ServicePlans     []ServicePlanInfo  `json:"servicePlans"`
	SkuID            string             `json:"skuId"`
	SkuPartNumber    string             `json:"skuPartNumber"`
}

func (subscribedSku *SubscribedSku) resourceName() string {
	return "subscribedSku"
}

// SubscribedSkus represents collection of subscribedSkus.
type SubscribedSkus struct {
	NextLink string          `json:"odata.nextLink"`
	Value    []SubscribedSku `json:"value"`
}

func (subscribedSkus *SubscribedSkus) resourceName() string {
	return "subscribedSku"
}

func (subscribedSkus *SubscribedSkus) nextLink() string {
	return subscribedSkus.NextLink
}
