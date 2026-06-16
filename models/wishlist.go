package models

import "time"

// WishlistItem represents a user's saved destination in the wishlist.
type WishlistItem struct {
	ID          string    `json:"id"`
	CountryName string    `json:"country_name"`
	Note        string    `json:"note"`
	Status      string    `json:"status"` // "Planned" or "Visited"
	CreatedAt   time.Time `json:"created_at"`
}
