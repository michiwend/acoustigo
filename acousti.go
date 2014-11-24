package acoustigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
)

type LLResponse struct {
}

type HLResponse struct {
	MetaData
	HighLevel map[string]HighLevelEntity
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
	BaseURL *url.URL
}

func (a *ABClient) getRequest(result interface{}, MBID, endpoint string) error {

	client := &http.Client{}

	reqUrl := a.BaseURL
	reqUrl.Path = path.Join(reqUrl.Path, MBID, endpoint)
	//reqUrl.RawQuery = params.Encode()

	fmt.Println(reqUrl.String())

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("User-Agent", a.userAgentHeader)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Bad response: " + resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(result); err != nil {
		return err
	}
	return nil
}

func (a *ABClient) HighLevel(recordingID string) (*HLResponse, error) {
	result := HLResponse{}
	err := a.getRequest(&result, recordingID, "high-level")

	return &result, err
}

func (a *ABClient) LowLevel(recordingID string) (*LLResponse, error) {
	return nil, nil
}

func NewABClient(apiUrl string) (*ABClient, error) {
	c := ABClient{}
	var err error

	c.BaseURL, err = url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
