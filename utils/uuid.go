package utils

import (
	"fmt"
	"strings"

	"github.com/houndgo/suuid"
	"github.com/satori/go.uuid"
)

// CreateUUID is creation uuid
func CreateUUID() (string, error) {
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return "", err
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	// u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// if err != nil {
	// 	fmt.Printf("Something went wrong: %s", err)
	// }
	fmt.Printf("Successfully parsed: %s", u2)

	id := suuid.New()
	fmt.Println("\n", id)
	v4 := strings.Replace(u2.String(), "-", "", -1)
	return v4, nil
}
