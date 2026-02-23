package validator

import (
	"errors"
	"fmt"

	"github.com/omn1vor/resume-maker/internal/model"
)

func Validate(r *model.Resume) error {
	// main
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if len(r.Experience) == 0 {
		return errors.New("at least one experience required")
	}

	// experience
	for _, ex := range r.Experience {
		if ex.Company == "" {
			return fmt.Errorf("expeience %v: company name is required", ex)
		}
	}

	// education
	for _, ed := range r.Education {
		if ed.Institution == "" {
			return fmt.Errorf("education %v: insitution name is required", ed)
		}
	}

	// courses
	for _, c := range r.Courses {
		if c.Name == "" {
			return fmt.Errorf("course %v: name is required", c)
		}
	}

	// projects
	for _, p := range r.Projects {
		if p.Name == "" {
			return fmt.Errorf("project %v: name is required", p)
		}
	}

	return nil
}
