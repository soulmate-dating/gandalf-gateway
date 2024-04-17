package profiles

import "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"

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

type Prompt struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Question string `json:"question"`
	Content  string `json:"content"`
	Position int32  `json:"position"`
}

type FullProfile struct {
	Profile Profile  `json:"profile"`
	Prompts []Prompt `json:"prompts"`
}

func mapPrompts(prompts []Prompt) []*profiles.Prompt {
	res := make([]*profiles.Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = &profiles.Prompt{
			Id:       p.ID,
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
		}
	}
	return res
}

func NewProfile(p *profiles.ProfileResponse) *Profile {
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

func NewPrompt(p *profiles.Prompt) Prompt {
	return Prompt{
		ID:       p.Id,
		Type:     "text",
		Question: p.Question,
		Content:  p.Content,
		Position: p.Position,
	}
}

func NewFullProfile(response *profiles.FullProfileResponse) *FullProfile {
	info := response.PersonalInfo
	profile := Profile{
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
	prompts := make([]Prompt, len(response.Prompts))
	for i, p := range response.Prompts {
		prompts[i] = Prompt{
			ID:       p.Id,
			Type:     p.Type,
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
		}
	}
	return &FullProfile{
		Profile: profile,
		Prompts: prompts,
	}
}

func Prompts(response *profiles.PromptsResponse) []Prompt {
	prompts := response.Prompts
	res := make([]Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = Prompt{
			ID:       p.Id,
			Type:     p.Type,
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
		}
	}
	return res
}
