package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRoomValidate(t *testing.T) {
	tests := []struct {
		name    string
		room    Room
		wantErr bool
	}{
		{
			name: "Valid Room",
			room: Room{
				ID:            primitive.NewObjectID(),
				Name:          "Test Room",
				Details:       "Test Details",
				Neighborhood:  "Test Neighborhood",
				ServeStyle:    ServeStyleBar,
				Photo:         "https://example.com/photo.jpg",
				Price:         100,
				MichelinStars: 2,
				Rating:        4.5,
				ReviewCount:   10,
				Coordinates:   Coordinates{Latitude: 40.7128, Longitude: -74.0060},
			},
			wantErr: false,
		},
		{
			name: "Invalid Name (Empty)",
			room: Room{
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "Invalid Name (Too Long)",
			room: Room{
				Name: "This name is way too long and exceeds the maximum length",
			},
			wantErr: true,
		},
		{
			name: "Invalid Details (Too Long)",
			room: Room{
				Name:    "Test Room",
				Details: string(make([]byte, MaxDetailsLength+1)),
			},
			wantErr: true,
		},
		{
			name: "Invalid ServeStyle",
			room: Room{
				Name:       "Test Room",
				ServeStyle: "invalid",
			},
			wantErr: true,
		},
		{
			name: "Invalid Photo URL",
			room: Room{
				Name:  "Test Room",
				Photo: "not-a-url",
			},
			wantErr: true,
		},
		{
			name: "Invalid Price (Negative)",
			room: Room{
				Name:  "Test Room",
				Price: -10,
			},
			wantErr: true,
		},
		{
			name: "Invalid MichelinStars (Out of Range)",
			room: Room{
				Name:          "Test Room",
				MichelinStars: 5,
			},
			wantErr: true,
		},
		{
			name: "Invalid Rating (Out of Range)",
			room: Room{
				Name:   "Test Room",
				Rating: 6.0,
			},
			wantErr: true,
		},
		{
			name: "Invalid Coordinates (Latitude Out of Range)",
			room: Room{
				Name:        "Test Room",
				Coordinates: Coordinates{Latitude: 91, Longitude: 0},
			},
			wantErr: true,
		},
		{
			name: "Invalid Coordinates (Longitude Out of Range)",
			room: Room{
				Name:        "Test Room",
				Coordinates: Coordinates{Latitude: 0, Longitude: 181},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.room.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Room.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServeStyle(t *testing.T) {
	tests := []struct {
		name  string
		style ServeStyle
		want  string
	}{
		{"Valid Bar", ServeStyleBar, ""},
		{"Valid Table", ServeStyleTable, ""},
		{"Invalid Style", "invalid", "ServeStyle must be either \"bar\" or \"table\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateServeStyle(tt.style); got != tt.want {
				t.Errorf("validateServeStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		name  string
		coord Coordinates
		want  string
	}{
		{"Valid Coordinates", Coordinates{Latitude: 40.7128, Longitude: -74.0060}, ""},
		{"Invalid Latitude (Too High)", Coordinates{Latitude: 91, Longitude: 0}, "Latitude must be between -90 and 90"},
		{"Invalid Latitude (Too Low)", Coordinates{Latitude: -91, Longitude: 0}, "Latitude must be between -90 and 90"},
		{"Invalid Longitude (Too High)", Coordinates{Latitude: 0, Longitude: 181}, "Longitude must be between -180 and 180"},
		{"Invalid Longitude (Too Low)", Coordinates{Latitude: 0, Longitude: -181}, "Longitude must be between -180 and 180"},
		{"Zero Coordinates", Coordinates{Latitude: 0, Longitude: 0}, "Latitude is required; Longitude is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCoordinates(tt.coord); got != tt.want {
				t.Errorf("validateCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
