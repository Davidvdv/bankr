package model

type Classification string

// Income
const (
	Salary Classification = "Salary"
	Board  Classification = "Board"
)

// House & Utilities
const (
	HomeLoan     Classification = "HomeLoan"
	CouncilRates Classification = "CouncilRates"
	Internet     Classification = "Internet"
	Electricity  Classification = "Electricity"
	Water        Classification = "Water"
	Gas          Classification = "Gas"
)

// Subscriptions & Plans
const (
	Spotify    Classification = "Spotify"
	Netflix    Classification = "Netflix"
	TwoDegrees Classification = "TwoDegrees"
	GoogleOne  Classification = "GoogleOne"
)

// Insurance
const (
	HouseAndContents Classification = "HouseAndContents"
	AAInsurance      Classification = "AAInsurance"
)

// Personal & medical
const (
	Haircuts           Classification = "Haircuts"
	Dental             Classification = "Dental"
	DoctorAndMedicaton Classification = "DoctorAndMedicaton"
	Clothing           Classification = "Clothing"
	Eyewear            Classification = "Eyewear"
)

const (
	Donations Classification = "Donations"
)

const (
	Other Classification = "Other"
)

var AllClassifications = []Classification{
	Salary,
	Board,
	HomeLoan,
	CouncilRates,
	Internet,
	Electricity,
	Water,
	Gas,
	Spotify,
	Netflix,
	TwoDegrees,
	GoogleOne,
	HouseAndContents,
	AAInsurance,
	Haircuts,
	Dental,
	DoctorAndMedicaton,
	Clothing,
	Eyewear,
	Donations,
	Other,
}

func (c Classification) String() string {
	return string(c)
}
