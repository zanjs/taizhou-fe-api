package app

import "github.com/iris-contrib/middleware/cors"

var crs = cors.New(cors.Options{
	AllowedOrigins:   []string{"*"},
	AllowedMethods:   []string{"GET", "PATCH", "PUT", "POST", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "Accept"},
	AllowCredentials: true, // allows everything, use that to change the hosts.
})
