/*
Copyright 2015 Aylien, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package textapi

import (
	"errors"
	"net/url"
)

// ClassifyParams is the set of parameters that defines a document whose classification needs to be calculated.
type ClassifyParams struct {
	// Either URL or Text is required.
	URL  string
	Text string

	// Valid languages are en, de, fr, es, it, pt and auto.
	// Default is en.
	Language string
}

// A Category is the JSON description of a classification category.
type Category struct {
	// IPTC SubjectCode
	Code       string  `json:"code"`
	Label      string  `json:"label"`
	Confidence float32 `json:"confidence"`
}

// A ClassifyResponse is the JSON description of classify response.
type ClassifyResponse struct {
	Text       string     `json:"text"`
	Language   string     `json:"language"`
	Categories []Category `json:"categories"`
}

// Classify classifies the document defined by the given params information.
func (c *Client) Classify(params *ClassifyParams) (*ClassifyResponse, error) {
	body := &url.Values{}

	if len(params.Text) > 0 {
		body.Add("text", params.Text)
	} else if len(params.URL) > 0 {
		body.Add("url", params.URL)
	} else {
		return nil, errors.New("You must either provide url or text")
	}

	if len(params.Language) > 0 {
		body.Add("language", params.Language)
	}

	classification := &ClassifyResponse{}
	err := c.call("/classify", body, classification)
	if err != nil {
		return nil, err
	}

	return classification, err
}
