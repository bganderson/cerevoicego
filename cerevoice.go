// CereVoice Cloud API Library for Go
// https://www.cereproc.com/files/CereVoiceCloudGuide.pdf
// This is a pre-release version and is subject to change

// Copyright 2018 Bryan Anderson (https://www.bganderson.com)
// Relesed under a BSD-style license which can be found in the LICENSE file

package cerevoicego

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	// VERSION is the global package version
	VERSION = "0.3.0"
	// DefaultRESTAPIURL is the default CereVoice Cloud REST API endpoint
	DefaultRESTAPIURL = "https://cerevoice.com/rest/rest_1_1.php"
)

// Client API connection settings
type Client struct {
	AccountID       string // CereVoice Cloud API AccountID
	Password        string // CereVoice Cloud API Password
	CereVoiceAPIURL string // CereVoice Cloud API URL
}

// Request to CereVoice Cloud API
type Request struct {
	XMLName          xml.Name
	AccountID        string `xml:"accountID"`
	Password         string `xml:"password"`
	Voice            string `xml:"voice,omitempty"`
	Text             string `xml:"text,omitempty"`
	AudioFormat      string `xml:"audioFormat,omitempty"`
	SampleRate       string `xml:"sampleRate,omitempty"`
	Audio3D          bool   `xml:"audio3D,omitempty"`
	Metadata         bool   `xml:"metadata,omitempty"`
	LexiconFile      string `xml:"lexiconFile,omitempty"`
	AbbreviationFile string `xml:"abbreviationFile,omitempty"`
	Language         string `xml:"language,omitempty"`
	Accent           string `xml:"accent,omitempty"`
}

// Response from CereVoice Cloud API
type Response struct {
	Raw   []byte
	Error error
}

// SpeakSimpleInput contains speakSimple parameters
type SpeakSimpleInput struct {
	Voice string
	Text  string
}

// SpeakExtendedInput contains speakExtended parameters
type SpeakExtendedInput struct {
	Voice       string
	Text        string
	AudioFormat string
	SampleRate  string
	Audio3D     bool
	Metadata    bool
}

// UploadLexiconInput contains uploadLexicon paramters
type UploadLexiconInput struct {
	LexiconFile string
	Language    string
	Accent      string
}

// UploadAbbreviationsInput contains uploadAbbreviations parameters
type UploadAbbreviationsInput struct {
	AbbreviationFile string
	Language         string
}

// SpeakSimpleResponse contains response from speakSimple
type SpeakSimpleResponse struct {
	FileURL           string `xml:"fileUrl"`
	CharCount         string `xml:"charCount"`
	ResultCode        string `xml:"resultCode"`
	ResultDescription string `xml:"resultDescription"`
	Error             error
}

// SpeakExtendedResponse contains response from speakExtended
type SpeakExtendedResponse struct {
	FileURL           string `xml:"fileUrl"`
	CharCount         string `xml:"charCount"`
	ResultCode        string `xml:"resultCode"`
	ResultDescription string `xml:"resultDescription"`
	Metadata          string `xml:"metadataUrl"`
	Error             error
}

// ListVoicesResponse contains response from listVoices
type ListVoicesResponse struct {
	VoiceList []Voice `xml:"voicesList>voice"`
	Error     error
}

// UploadLexiconResponse contains response from uploadLexicon
type UploadLexiconResponse struct {
	ResultCode        int    `xml:"resultCode"`
	ResultDescription string `xml:"resultDescription"`
	Error             error
}

// ListLexiconsResponse contains response from listLexicons
type ListLexiconsResponse struct {
	LexiconList []Lexicon `xml:"lexiconList>lexiconFile"`
	Error       error
}

// UploadAbbreviationsResponse contains response from uploadAbbreviations
type UploadAbbreviationsResponse struct {
	ResultCode        int    `xml:"resultCode"`
	ResultDescription string `xml:"resultDescription"`
	Error             error
}

// ListAbbreviationsResponse contains response from listAbbreviations
type ListAbbreviationsResponse struct {
	AbbreviationList []Abbreviation `xml:"abbreviationList>abbreviationFile"`
	Error            error
}

// ListAudioFormatsResponse contains response from listAudioFormats
type ListAudioFormatsResponse struct {
	AudioFormats []string `xml:"formatList>format"`
	Error        error
}

// GetCreditResponse contains response from getCredit
type GetCreditResponse struct {
	Credit Credit `xml:"credit"`
	Error  error
}

// Voice contains details about a voice
type Voice struct {
	SampleRate            string `xml:"sampleRate"`
	VoiceName             string `xml:"voiceName"`
	LanguageCodeISO       string `xml:"languageCodeISO"`
	CountryCodeISO        string `xml:"countryCodeISO"`
	AccentCode            string `xml:"accentCode"`
	Sex                   string `xml:"sex"`
	LanguageCodeMicrosoft string `xml:"languageCodeMicrosoft"`
	Country               string `xml:"country"`
	Region                string `xml:"region"`
	Accent                string `xml:"accent"`
}

// Lexicon contains details about a lexicon
type Lexicon struct {
	URL          string `xml:"url"`
	Language     string `xml:"language"`
	Accent       string `xml:"accent"`
	LastModified string `xml:"lastModified"`
	Size         string `xml:"size"`
}

// Abbreviation contains details about an abbreviation
type Abbreviation struct {
	URL          string `xml:"url"`
	Language     string `xml:"language"`
	LastModified string `xml:"lastModified"`
	Size         string `xml:"size"`
}

// Credit contains details about CereVoice Cloud credits
type Credit struct {
	FreeCredit     string `xml:"freeCredit"`
	PaidCredit     string `xml:"paidCredit"`
	CharsAvailable string `xml:"charsAvailable"`
}

// SpeakSimple synthesises input text with the selected voice
func (c *Client) SpeakSimple(input *SpeakSimpleInput) (r *SpeakSimpleResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "speakSimple"},
		AccountID: c.AccountID,
		Password:  c.Password,
		Voice:     input.Voice,
		Text:      input.Text,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// SpeakExtended allows for more control over the audio output
func (c *Client) SpeakExtended(input *SpeakExtendedInput) (r *SpeakExtendedResponse) {
	resp := c.queryAPI(&Request{
		XMLName:     xml.Name{Local: "speakExtended"},
		AccountID:   c.AccountID,
		Password:    c.Password,
		Voice:       input.Voice,
		Text:        input.Text,
		AudioFormat: input.AudioFormat,
		SampleRate:  input.SampleRate,
		Audio3D:     input.Audio3D,
		Metadata:    input.Metadata,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// ListVoices outputs information about the available voices
func (c *Client) ListVoices() (r *ListVoicesResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "listVoices"},
		AccountID: c.AccountID,
		Password:  c.Password,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// UploadLexicon uploads and stores a custom lexicon file
func (c *Client) UploadLexicon(input *UploadLexiconInput) (r *UploadLexiconResponse) {
	resp := c.queryAPI(&Request{
		XMLName:     xml.Name{Local: "uploadLexicon"},
		AccountID:   c.AccountID,
		Password:    c.Password,
		LexiconFile: input.LexiconFile,
		Language:    input.Language,
		Accent:      input.Accent,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return

}

// ListLexicons lists custom lexicon file(s)
func (c *Client) ListLexicons() (r *ListLexiconsResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "listLexicons"},
		AccountID: c.AccountID,
		Password:  c.Password,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// UploadAbbreviations uploads and stores a custom abbreviation file
func (c *Client) UploadAbbreviations(input *UploadAbbreviationsInput) (r *UploadAbbreviationsResponse) {
	resp := c.queryAPI(&Request{
		XMLName:     xml.Name{Local: "uploadAbbreviations"},
		AccountID:   c.AccountID,
		Password:    c.Password,
		LexiconFile: input.AbbreviationFile,
		Language:    input.Language,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// ListAbbreviations lists custom abbreviation file(s)
func (c *Client) ListAbbreviations() (r *ListAbbreviationsResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "listAbbreviations"},
		AccountID: c.AccountID,
		Password:  c.Password,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	err := xml.Unmarshal(resp.Raw, &r)
	if err != nil {
		r.Error = err
	}

	return
}

// ListAudioFormats lists the available audio encoding formats
func (c *Client) ListAudioFormats() (r *ListAudioFormatsResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "listAudioFormats"},
		AccountID: c.AccountID,
		Password:  c.Password,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// GetCredit retrieves the credit information for the given account
func (c *Client) GetCredit() (r *GetCreditResponse) {
	resp := c.queryAPI(&Request{
		XMLName:   xml.Name{Local: "getCredit"},
		AccountID: c.AccountID,
		Password:  c.Password,
	})
	if resp.Error != nil {
		r.Error = resp.Error
		return
	}

	if err := xml.Unmarshal(resp.Raw, &r); err != nil {
		r.Error = err
	}

	return
}

// Query CereVoice Cloud API
func (c *Client) queryAPI(req *Request) (r *Response) {
	output, err := xml.MarshalIndent(req, "", "    ")
	if err != nil {
		r.Error = err
		return
	}

	request := bytes.NewReader(append([]byte(xml.Header), output...))
	resp, err := http.Post(c.CereVoiceAPIURL, "text/xml", request)
	if err != nil {
		r.Error = err
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Error = err
		return
	}

	return &Response{Raw: body}
}
