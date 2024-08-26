package elsaticsearch

import "pengyou/model/entity"

// PostHit Define the structs to match the JSON structure
type PostHit struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	ID     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Source entity.Post `json:"_source"`
}

type Response struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total    Total     `json:"total"`
	MaxScore float64   `json:"max_score"`
	Hits     []PostHit `json:"hits"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}
