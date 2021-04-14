package config

type _defaultType struct {
	Country, Name, Org, OrgUnit, Province, Locality, Street, Postcode, KeyPassword string
	RootCAPassword                                                                 string
}

var (
	Default = _defaultType{
		Country:        "CN",
		Name:           "ArkRootCA",
		Org:            "Ryan Ark Center",
		OrgUnit:        "Box",
		Province:       "Beijing",
		Locality:       "Beijing",
		Street:         "NoWhere Road 9+3/4 Site Corner",
		Postcode:       "061219",
		KeyPassword:    "wingca",
		RootCAPassword: "wingCA2021",
	}
)
