package profile

import (
	"fmt"
	"regexp"
)

//Profile simple structure to keep profile settings
type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//SetName name setter
func (p *Profile) SetName(name string) error {
	p.Name = name
	return nil
}

//SetEmail setter with an email validation
func (p *Profile) SetEmail(email string) error {
	// email validation regexp
	validationRp, err := regexp.Compile("^\\w+@[a-zA-Z_]+?\\.[a-zA-Z]{2,3}$")
	if err != nil {
		return err
	}

	if validationRp.Match([]byte(email)) {
		p.Email = email
		return nil
	}

	return fmt.Errorf("email validation failed. email<%s>", email)
}
