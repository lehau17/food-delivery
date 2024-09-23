package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension" gorm:"-"`
}

// Scan implements the sql.Scanner interface.
// This function converts a database value (stored as JSON string) into an Image struct.
func (i *Image) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	// Convert the value to a string, check if it's valid
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid image data ")
	}
	var img Image
	// Decode the JSON string into the Image struct
	err := json.Unmarshal(s, &img)
	if err != nil {
		return errors.New("failed to decode image data")
	}
	*i = img

	return nil
}

// Value implements the driver.Valuer interface.
// This function converts the Image struct to a JSON string for storing in the database.
func (i *Image) Value() (driver.Value, error) {
	// Convert the Image struct to a JSON string
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

type Images []Image

func (is *Images) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	//
	// check covert to byte[]
	s, Ok := value.([]byte)
	if !Ok {
		return errors.New("invalid image")
	}
	// UmMarsol
	var images Images
	err := json.Unmarshal(s, &images)
	if err != nil {
		return errors.New("invalid image")
	}
	*is = images
	return nil
}

func (is *Images) Value() (driver.Value, error) {
	if is == nil {
		return nil, nil
	}
	return json.Marshal(is)
}
