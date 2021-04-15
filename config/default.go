package config

type _defaultType struct {
	Country, Name, Org, OrgUnit, Province, Locality, Street, Postcode, KeyPassword string
	RootCAPassword                                                                 string
	KeyLen                                                                         int
}

var (
	Default = _defaultType{
		Country:        "CN",
		Name:           "ArkRootCA",
		Org:            "Ark Ryan Center",
		OrgUnit:        "Box 12383",
		Province:       "Beijing",
		Locality:       "Beijing",
		Street:         "NoWhere Road 9+3/4 Site Corner",
		Postcode:       "061219",
		KeyPassword:    "wingca",
		RootCAPassword: "wingCA2021",
		KeyLen:         2048,
	}
)
