package models

type TicketsResp struct {
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
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	Score      float32 `json:"score"`
	Popularity float32 `json:"popularity"`
	Stats      Stat    `json:"stats"`
}

type Stat struct {
	Count                           int    `json:"listing_count"`
	Lowest_price_good_deals         string `json:"lowest_price_good_deals"`
	LowestPrice                     int    `json:"lowest_price"`
	AveragePrice                    int    `json:"average_price"`
	HighestPrice                    int    `json:"highest_price"`
	Visible_listing_count           int    `json:"visible_listing_count"`
	Median_price                    int    `json:"median_price"`
	Lowest_sg_base_price            int    `json:"lowest_sg_base_price"`
	Lowest_sg_base_price_good_deals string `json:"lowest_sg_base_price_good_deals"`
}
