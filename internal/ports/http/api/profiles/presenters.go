package profiles

import "github.com/soulmate-dating/gandalf-gateway/internal/app/clients/profiles"

// Profile represents a user's profile.
type Profile struct {
	FirstName        string `json:"first_name,omitempty" binding:"required" example:"Elon"`
	LastName         string `json:"last_name" binding:"required" example:"Musk"`
	BirthDate        string `json:"birth_date" binding:"required" example:"1971-06-28" format:"date" pattern:"^\\d{4}-\\d{2}-\\d{2}$"`
	Sex              string `json:"sex" binding:"required" example:"man"`
	PreferredPartner string `json:"preferred_partner" binding:"required" example:"woman"`
	Intention        string `json:"intention" binding:"required" example:"long-term relationship"`
	Height           uint32 `json:"height" example:"180"`
	HasChildren      bool   `json:"has_children" binding:"required" example:"false"`
	FamilyPlans      string `json:"family_plans" binding:"required" example:"not sure yet""`
	Location         string `json:"location,omitempty"`
	DrinksAlcohol    string `json:"drinks_alcohol" binding:"required" example:"sometimes"`
	Smokes           string `json:"smokes" binding:"required" example:"no"`
}

// Prompt represents a user's prompt.
type Prompt struct {
	ID       string `json:"id" example:"75988450-f7c7-4022-b04b-6679e9294056"`
	Type     string `json:"type" example:"text"`
	Question string `json:"question,omitempty" example:"My most irrational fear is..."`
	Content  string `json:"content" binding:"required" example:"Spider Man"`
	Position int32  `json:"position"`
}

// FullProfile represents a user's full profile.
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
