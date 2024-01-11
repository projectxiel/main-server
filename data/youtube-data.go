package data

// YouTubeResponse represents the response format from the YouTube API.
type YouTubeResponse struct {
	Etag       string        `json:"etag"`
	Items      []YouTubeItem `json:"items"`
	Kind       string        `json:"kind"`
	PageInfo   PageInfo      `json:"pageInfo"`
	RegionCode string        `json:"regionCode,omitempty"`
}

// YouTubeItem represents an individual item in the YouTube API response.
type YouTubeItem struct {
	Etag    string         `json:"etag"`
	ID      YouTubeID      `json:"id"`
	Kind    string         `json:"kind"`
	Snippet YouTubeSnippet `json:"snippet"`
}

// YouTubeID represents the ID section of an item in the YouTube API response.
type YouTubeID struct {
	ChannelID string `json:"channelId,omitempty"`
	Kind      string `json:"kind"`
	VideoID   string `json:"videoId,omitempty"`
}

// YouTubeSnippet represents the snippet section of an item in the YouTube API response.
type YouTubeSnippet struct {
	ChannelID            string            `json:"channelId"`
	ChannelTitle         string            `json:"channelTitle"`
	Description          string            `json:"description"`
	LiveBroadcastContent string            `json:"liveBroadcastContent"`
	PublishTime          string            `json:"publishTime"`
	PublishedAt          string            `json:"publishedAt"`
	Thumbnails           YouTubeThumbnails `json:"thumbnails"`
	Title                string            `json:"title"`
}

// YouTubeThumbnails represents the thumbnails section of a snippet in the YouTube API response.
type YouTubeThumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

// Thumbnail represents a single thumbnail in the YouTube API response.
type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// PageInfo represents the pageInfo section of the YouTube API response.
type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}
