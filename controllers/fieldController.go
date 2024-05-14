package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/Qwerci/eos-api2/config"
	"github.com/Qwerci/eos-api2/models"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Field Field `json:"field"`
}

type Field struct {
	Type       string      `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
  }
  
  type Properties struct {
	Name      string      `json:"name"`
	Group     string      `json:"group"`
	YearsData []YearsData `json:"years_data"`
  }
  
  type YearsData struct {
	CropType string `json:"crop_type"`
	Year     int    `json:"year"`
	SowingDate string `json:"sowing_date"`
  }
  
  type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][]float64   `json:"coordinates"`
  }



  func  sendApiRequestwithparam(url string, requestBody RequestBody, apiKey string) ([]byte, error) {
	client := &http.Client{}

	// Create a JSON payload from the request data
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Build the URL with the API key
	urlWithApiKey := fmt.Sprintf("%s?api_key=%s", url, apiKey)
	fmt.Println(urlWithApiKey)

	req, err := http.NewRequest("POST", urlWithApiKey, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("can't create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't send request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %v", err)
	}
	fmt.Println(res)
	fmt.Println(string(body))

	return body, nil

}

func CreateField( c *gin.Context){
	configUrl, _ := config.LoadUrl()
	configApi, _ := config.LoadKey()
	// url =

	var requestBody RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	fmt.Println(configUrl.CreateFieldUrl)
	
	response, err := sendApiRequestwithparam(configUrl.CreateFieldUrl, requestBody, configApi.Eos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with third-party URL"})
		return
	}

	var responseData models.FieldResponse1
	if err := json.Unmarshal(response, &responseData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": responseData})
}


type SceneSearchParams struct {
	DateStart  string   `json:"date_start"`
	DateEnd    string   `json:"date_end"`
	DataSource []string `json:"data_source"`
}


 func SearchScenes(c *gin.Context) {
		fieldID := c.Param("fieldID")
		config, _ := config.LoadKey()

		// Read API key from environment variable (recommended)
		// apiKey := os.Getenv("X_API_KEY")
		// if apiKey == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-API-KEY"})
		// 	return
		// }

		// Read request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Parse request body
		var params SceneSearchParams
		err = json.Unmarshal(body, &params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Build the URL with field ID
		url := fmt.Sprintf("https://api-connect.eos.com/scene-search/for-field/%s", fieldID)

		// Prepare request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Set headers
		req.Header.Set("Content-Type", "text/plain")
		req.Header.Set("x-api-key", config.Eos)

		// Send request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Read response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Handle response (success or error)
		c.JSON(resp.StatusCode, responseBody)
	}