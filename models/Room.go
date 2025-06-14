package models

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MaxNameLength         = 35
	MaxDetailsLength      = 150
	MaxNeighborhoodLength = 35
	MaxMichelinStars      = 3
	MaxRating             = 5
)

type Coordinates struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type ServeStyle string

const (
	ServeStyleBar   ServeStyle = "bar"
	ServeStyleTable ServeStyle = "table"
)

type Room struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Coordinates   Coordinates        `json:"coordinates" bson:"coordinates"`
	Details       string             `json:"details" bson:"details"`
	MichelinStars int                `json:"michelin_stars" bson:"michelin_stars"`
	Name          string             `json:"name" bson:"name"`
	Neighborhood  string             `json:"neighborhood" bson:"neighborhood"`
	Photo         string             `json:"photo" bson:"photo"`
	Price         int                `json:"price" bson:"price"`
	Rating        float64            `json:"rating" bson:"rating"`
	ReviewCount   int                `json:"review_count" bson:"review_count"`
	ServeStyle    ServeStyle         `json:"serve_style" bson:"serve_style"`
}

func (r *Room) Validate() error {
	var errors []string

	errors = append(errors,
		validateRequired("Name", r.Name),
		validateMaxLength("Name", r.Name, MaxNameLength),
		validateRequired("Details", r.Details),
		validateMaxLength("Details", r.Details, MaxDetailsLength),
		validateRequired("Neighborhood", r.Neighborhood),
		validateMaxLength("Neighborhood", r.Neighborhood, MaxNeighborhoodLength),
		validateRequired("ServeStyle", string(r.ServeStyle)),
		validateServeStyle(r.ServeStyle),
		validateRequired("Photo", r.Photo),
		validateURL("Photo", r.Photo),
		validatePositive("Price", r.Price),
		validateRange("MichelinStars", r.MichelinStars, 0, MaxMichelinStars),
		validateRange("Rating", r.Rating, 0, MaxRating),
		validateMinimum("ReviewCount", r.ReviewCount, 0),
		validateCoordinates(r.Coordinates),
	)

	errors = removeEmptyStrings(errors)
	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, "; "))
	}
	return nil
}

func validateRequired(field, value string) string {
	if value == "" {
		return fmt.Sprintf("%s is required", field)
	}
	return ""
}

func validateMaxLength(field, value string, maxLength int) string {
	if len(value) > maxLength {
		return fmt.Sprintf("%s cannot exceed %d characters", field, maxLength)
	}
	return ""
}

func validatePositive(field string, value int) string {
	if value <= 0 {
		return fmt.Sprintf("%s must be positive", field)
	}
	return ""
}

func validateRange(field string, value interface{}, min, max float64) string {
	switch v := value.(type) {
	case int:
		if float64(v) < min || float64(v) > max {
			return fmt.Sprintf("%s must be between %.0f and %.0f", field, min, max)
		}
	case float64:
		if v < min || v > max {
			return fmt.Sprintf("%s must be between %.1f and %.1f", field, min, max)
		}
	}
	return ""
}

func validateMinimum(field string, value, min int) string {
	if value < min {
		return fmt.Sprintf("%s cannot be less than %d", field, min)
	}
	return ""
}

func validateURL(field, value string) string {
	// Basic check for scheme and host
	if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		return fmt.Sprintf("%s must start with http:// or https://", field)
	}

	// Remove the scheme for further parsing
	withoutScheme := strings.TrimPrefix(strings.TrimPrefix(value, "http://"), "https://")

	// Check if there's at least one dot in the remaining string (simple host check)
	if !strings.Contains(withoutScheme, ".") {
		return fmt.Sprintf("%s must contain a valid host", field)
	}

	return ""
}

func validateCoordinates(coords Coordinates) string {
	var errors []string

	if coords.Latitude == 0 && coords.Longitude == 0 {
		return "Latitude is required; Longitude is required"
	}

	if coords.Latitude < -90 || coords.Latitude > 90 {
		errors = append(errors, "Latitude must be between -90 and 90")
	}

	if coords.Longitude < -180 || coords.Longitude > 180 {
		errors = append(errors, "Longitude must be between -180 and 180")
	}

	return strings.Join(errors, "; ")
}

func validateServeStyle(style ServeStyle) string {
	switch style {
	case ServeStyleBar, ServeStyleTable:
		return ""
	default:
		return fmt.Sprintf("ServeStyle must be either %q or %q", ServeStyleBar, ServeStyleTable)
	}
}

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
