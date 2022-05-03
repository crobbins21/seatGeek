package models

type EventsReturn struct {
	Events []Event  `json:"events"`
	Meta   MetaData `json:"meta"`
}

type MetaData struct {
	Total       int    `json:"total"`
	Took        int    `json:"took"`
	Page        int    `json:"page"`
	Per_page    int    `json:"per_page"`
	Geolocation string `json:"geolocation"`
}

type Event struct {
	ID         int     `json:"id" bson:"eventId"`
	Title      string  `json:"title" bson:"title"`
	Type       string  `json:"type" bson:"type"`
	Score      float32 `json:"score" bson:"score"`
	Popularity float32 `json:"popularity" bson:"popularity"`
	URL        string  `json:"url" bson:"url"`
	Stats      Stat    `json:"stats" bson:"stats"`
}

type Stat struct {
	Count                           int    `json:"listing_count" bson:"listing_count"`
	Lowest_price_good_deals         string `json:"lowest_price_good_deals" bson:"lowest_price_good_deals"`
	LowestPrice                     int    `json:"lowest_price" bson:"lowest_price"`
	AveragePrice                    int    `json:"average_price" bson:"average_price"`
	HighestPrice                    int    `json:"highest_price" bson:"highest_price"`
	Visible_listing_count           int    `json:"visible_listing_count" bson:"visible_listing_count"`
	Median_price                    int    `json:"median_price" bson:"median_price"`
	Lowest_sg_base_price            int    `json:"lowest_sg_base_price" bson:"lowest_sg_base_price"`
	Lowest_sg_base_price_good_deals string `json:"lowest_sg_base_price_good_deals" bson:"lowest_sg_base_price_good_deals"`
}
