package config


import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)



func InitCloudinary() *cloudinary.Cloudinary {

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env: %v", err)
	}

	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	fmt.Printf("CLOUDINARY_NAME: '%s'\n", cloudName)
	fmt.Printf("API_KEY: '%s'\n", apiKey)
	fmt.Printf("API_SECRET: '%s'\n", apiSecret)

	if cloudName == "" {
		log.Fatal("ERROR: CLOUDINARY_NAME empty")
	}
	if apiKey == "" {
		log.Fatal("ERROR: API_KEY empty")
	}
	if apiSecret == "" {
		log.Fatal("ERROR: API_SECRET empty")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("ERROR creating instance: %v", err)
	}


	return cld
}