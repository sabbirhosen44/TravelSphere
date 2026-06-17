package utils

import (
	"errors"
	"strings"
)

// ValidateWishlistInput validates wishlist payload inputs.
func ValidateWishlistInput(countryName string, status string, note string) error {
	country := strings.TrimSpace(countryName)
	if country == "" {
		return errors.New("country name cannot be empty")
	}
	if len(country) < 2 {
		return errors.New("country name must be at least 2 characters")
	}

	st := strings.TrimSpace(status)
	if st != "Planned" && st != "Visited" {
		return errors.New("status must be either 'Planned' or 'Visited'")
	}

	if len(note) > 500 {
		return errors.New("note must not exceed 500 characters")
	}

	return nil
}

// ValidateSlug validates that a country slug is non-empty and has alphanumeric/hyphen characters.
func ValidateSlug(slug string) error {
	s := strings.TrimSpace(slug)
	if s == "" {
		return errors.New("slug cannot be empty")
	}
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_') {
			return errors.New("slug contains invalid characters")
		}
	}
	return nil
}
