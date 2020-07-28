package twtrlib

const baseStreamURL = "https://stream.twitter.com/1.1"

// Stream allows access to the Twitter Streaming API
type Stream struct {
	client TwitterClient
	C      chan interface{}
}

// StdStreamParams are the standard parameters to the streaming endpoints
type StdStreamParams struct {
	Count         string   `url:"count,omitempty"`
	Delimited     string   `url:"delimited,omitempty"`
	FilterLevel   string   `url:"filter_level,omitempty"`
	Follow        []string `url:"follow,omitempty"`
	Language      []string `url:"language,omitempty"`
	Locations     []string `url:"locations,omitempty"`
	StallWarnings bool     `url:"stall_warnings,omitempty"`
}

// Filter should return messages that mathes set conditions
func (strm *Stream) Filter(params *StdStreamParams) {}

