package locationmap

import (
	"geodistance/distance"
)

// Define the PlacesCoordinates map
var PlacesCoordinates = map[string]distance.GeoCoordinate{
	"Mount Everest, Nepal": {
		Glat:  distance.Coordinate{D: 27, M: 59, S: 17, H: 'N'},
		Glong: distance.Coordinate{D: 86, M: 55, S: 31, H: 'E'},
	},
	"Pyramids of Giza, Egypt": {
		Glat:  distance.Coordinate{D: 29, M: 58, S: 45, H: 'N'},
		Glong: distance.Coordinate{D: 31, M: 8, S: 3, H: 'E'},
	},
	"Eiffel Tower, France": {
		Glat:  distance.Coordinate{D: 48, M: 51, S: 24, H: 'N'},
		Glong: distance.Coordinate{D: 2, M: 21, S: 8, H: 'E'},
	},
	"Stonehenge, UK": {
		Glat:  distance.Coordinate{D: 51, M: 10, S: 44, H: 'N'},
		Glong: distance.Coordinate{D: 1, M: 49, S: 34, H: 'W'},
	},
	"Tokyo Tower, Japan": {
		Glat:  distance.Coordinate{D: 35, M: 41, S: 22, H: 'N'},
		Glong: distance.Coordinate{D: 139, M: 41, S: 30, H: 'E'},
	},
	"Area 51, USA": {
		Glat:  distance.Coordinate{D: 37, M: 14, S: 36, H: 'N'},
		Glong: distance.Coordinate{D: 115, M: 48, S: 40, H: 'W'},
	},
	"Statue of Liberty, USA": {
		Glat:  distance.Coordinate{D: 40, M: 41, S: 21, H: 'N'},
		Glong: distance.Coordinate{D: 74, M: 2, S: 40, H: 'W'},
	},
	"Cape Town, South Africa": {
		Glat:  distance.Coordinate{D: 34, M: 3, S: 8, H: 'S'},
		Glong: distance.Coordinate{D: 18, M: 25, S: 24, H: 'E'},
	},
	"Easter Island, Chile": {
		Glat:  distance.Coordinate{D: 27, M: 10, S: 30, H: 'S'},
		Glong: distance.Coordinate{D: 109, M: 25, S: 10, H: 'W'},
	},
	"North Pole": {
		Glat:  distance.Coordinate{D: 90, M: 0, S: 0, H: 'N'},
		Glong: distance.Coordinate{D: 0, M: 0, S: 0, H: 'E'},
	},
	"South Pole": {
		Glat:  distance.Coordinate{D: 90, M: 0, S: 0, H: 'S'},
		Glong: distance.Coordinate{D: 0, M: 0, S: 0, H: 'W'},
	}, "Great Wall of China, China": {
		Glat:  distance.Coordinate{D: 40, M: 25, S: 0, H: 'N'},
		Glong: distance.Coordinate{D: 116, M: 34, S: 13, H: 'E'},
	},
	"Christ the Redeemer, Brazil": {
		Glat:  distance.Coordinate{D: 22, M: 57, S: 8, H: 'S'},
		Glong: distance.Coordinate{D: 43, M: 12, S: 44, H: 'W'},
	},
	"Grand Canyon, USA": {
		Glat:  distance.Coordinate{D: 36, M: 6, S: 3, H: 'N'},
		Glong: distance.Coordinate{D: 112, M: 5, S: 24, H: 'W'},
	},
	"Machu Picchu, Peru": {
		Glat:  distance.Coordinate{D: 13, M: 9, S: 48, H: 'S'},
		Glong: distance.Coordinate{D: 72, M: 32, S: 44, H: 'W'},
	},
	"Niagara Falls, Canada/USA": {
		Glat:  distance.Coordinate{D: 43, M: 4, S: 36, H: 'N'},
		Glong: distance.Coordinate{D: 79, M: 4, S: 33, H: 'W'},
	},
	"Petra, Jordan": {
		Glat:  distance.Coordinate{D: 30, M: 19, S: 43, H: 'N'},
		Glong: distance.Coordinate{D: 35, M: 26, S: 30, H: 'E'},
	},
	"Angel Falls, Venezuela": {
		Glat:  distance.Coordinate{D: 5, M: 58, S: 3, H: 'N'},
		Glong: distance.Coordinate{D: 62, M: 32, S: 30, H: 'W'},
	},
	"Uluru (Ayers Rock), Australia": {
		Glat:  distance.Coordinate{D: 25, M: 20, S: 41, H: 'S'},
		Glong: distance.Coordinate{D: 131, M: 2, S: 10, H: 'E'},
	},
	"Burj Khalifa, UAE": {
		Glat:  distance.Coordinate{D: 25, M: 11, S: 49, H: 'N'},
		Glong: distance.Coordinate{D: 55, M: 16, S: 26, H: 'E'},
	},
	"Amazon Rainforest, Brazil": {
		Glat:  distance.Coordinate{D: 3, M: 22, S: 24, H: 'S'},
		Glong: distance.Coordinate{D: 60, M: 1, S: 18, H: 'W'},
	},
	"Las Vegas, USA": {
		Glat:  distance.Coordinate{D: 36, M: 10, S: 30, H: 'N'},
		Glong: distance.Coordinate{D: 115, M: 8, S: 11, H: 'W'},
	},

}
