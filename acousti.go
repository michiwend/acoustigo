package acoustigo

import (
	"github.com/michiwend/gomusicbrainz"
	"net/url"
)

type GenreTzanetakis struct {
	Value       string
	Probability float64
}

type HLResponse struct {
	MetaData
	HighLevel []HighLevelEntity
}

type MetaData struct {
	Version         map[string]interface{}
	AudioProperties struct {
		SampleRate         int
		Codec              string
		BitRate            int
		EqualLoudness      int
		AnalysisSampleRate int
		Length             float64
		MD5Encoded         string
		ReplayGain         float64
		DownMix            string
		Lossless           bool
	}
	Tags map[string]interface{}
}

type HighLevelEntity struct {
	Value       string
	Probability float64
	All         map[string]float64
}

type ABClient struct {
	BaseURL url.URL
}

func (a *ABClient) HighLevelFromGMBRecording(r *gomusicbrainz.Recording) (*HighLevel, error) {
	return nil, nil
}

func NewABClient(apiUrl string) *ABClient {

	c := ABClient{}
	var err error

	c.BaseURL, err = url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
