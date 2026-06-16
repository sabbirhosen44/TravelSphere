package services

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"TravelSphere/models"
)

var (
	wishlistFile = "wishlist.json"
	mutex        sync.RWMutex
)

// SetWishlistFile changes the target wishlist storage file (primarily for testing).
func SetWishlistFile(filename string) {
	mutex.Lock()
	defer mutex.Unlock()
	wishlistFile = filename
}

// WishlistService handles operations on the travel wishlist.
type WishlistService struct{}

// NewWishlistService creates a new instance of WishlistService.
func NewWishlistService() *WishlistService {
	return &WishlistService{}
}

// GenerateID produces a pseudorandom UUID string.
func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// readFromFile reads wishlist items from the JSON file. Assumes lock is already held.
func readFromFile() ([]models.WishlistItem, error) {
	if _, err := os.Stat(wishlistFile); os.IsNotExist(err) {
		return []models.WishlistItem{}, nil
	}

	data, err := os.ReadFile(wishlistFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read wishlist file: %w", err)
	}

	if len(data) == 0 {
		return []models.WishlistItem{}, nil
	}

	var items []models.WishlistItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("failed to parse wishlist data: %w", err)
	}

	return items, nil
}

// writeToFile saves wishlist items to the JSON file. Assumes lock is already held.
func writeToFile(items []models.WishlistItem) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal wishlist: %w", err)
	}

	err = os.WriteFile(wishlistFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write wishlist to file: %w", err)
	}

	return nil
}

// List returns all wishlist entries.
func (s *WishlistService) List() ([]models.WishlistItem, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	return readFromFile()
}

// Create adds a new destination to the wishlist.
func (s *WishlistService) Create(countryName, note, status string) (models.WishlistItem, error) {
	mutex.Lock()
	defer mutex.Unlock()

	items, err := readFromFile()
	if err != nil {
		return models.WishlistItem{}, err
	}

	item := models.WishlistItem{
		ID:          generateID(),
		CountryName: countryName,
		Note:        note,
		Status:      status,
		CreatedAt:   time.Now(),
	}

	items = append(items, item)

	if err := writeToFile(items); err != nil {
		return models.WishlistItem{}, err
	}

	return item, nil
}

// Update modifies an existing wishlist entry's note or status.
func (s *WishlistService) Update(id, note, status string) (models.WishlistItem, error) {
	mutex.Lock()
	defer mutex.Unlock()

	items, err := readFromFile()
	if err != nil {
		return models.WishlistItem{}, err
	}

	foundIdx := -1
	for idx, item := range items {
		if item.ID == id {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return models.WishlistItem{}, errors.New("wishlist item not found")
	}

	items[foundIdx].Note = note
	items[foundIdx].Status = status

	if err := writeToFile(items); err != nil {
		return models.WishlistItem{}, err
	}

	return items[foundIdx], nil
}

// Delete removes a wishlist entry by its ID.
func (s *WishlistService) Delete(id string) error {
	mutex.Lock()
	defer mutex.Unlock()

	items, err := readFromFile()
	if err != nil {
		return err
	}

	foundIdx := -1
	for idx, item := range items {
		if item.ID == id {
			foundIdx = idx
			break
		}
	}

	if foundIdx == -1 {
		return errors.New("wishlist item not found")
	}

	items = append(items[:foundIdx], items[foundIdx+1:]...)

	return writeToFile(items)
}

// GetSummary returns counts for total, planned, and visited destinations.
func (s *WishlistService) GetSummary() (total int, planned int, visited int, err error) {
	mutex.RLock()
	defer mutex.RUnlock()

	items, err := readFromFile()
	if err != nil {
		return 0, 0, 0, err
	}

	total = len(items)
	for _, item := range items {
		if item.Status == "Planned" {
			planned++
		} else if item.Status == "Visited" {
			visited++
		}
	}

	return total, planned, visited, nil
}
