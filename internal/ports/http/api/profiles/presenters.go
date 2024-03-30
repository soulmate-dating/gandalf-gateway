package profiles

import "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profile"

type Profile struct {
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name"`
	BirthDate        string `json:"birth_date"`
	Sex              string `json:"sex,omitempty"`
	PreferredPartner string `json:"preferred_partner,omitempty"`
	Intention        string `json:"intention,omitempty"`
	Height           uint32 `json:"height,omitempty"`
	HasChildren      bool   `json:"has_children,omitempty"`
	FamilyPlans      string `json:"family_plans,omitempty"`
	Location         string `json:"location,omitempty"`
	DrinksAlcohol    string `json:"drinks_alcohol,omitempty"`
	Smokes           string `json:"smokes,omitempty"`
}

func NewProfile(p *profile.ProfileResponse) *Profile {
	info := p.PersonalInfo
	return &Profile{
		FirstName:        info.FirstName,
		LastName:         info.LastName,
		BirthDate:        info.BirthDate,
		Sex:              info.Sex,
		PreferredPartner: info.PreferredPartner,
		Intention:        info.Intention,
		Height:           info.Height,
		HasChildren:      info.HasChildren,
		FamilyPlans:      info.FamilyPlans,
		Location:         info.Location,
		DrinksAlcohol:    info.DrinksAlcohol,
		Smokes:           info.Smokes,
	}
}
