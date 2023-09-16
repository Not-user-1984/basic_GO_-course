package main


import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	Responce struct {
        Header ResponseHeader `json:"header"`
        Data   ResponseData   `json:"data,omitempty"`
	}
    ResponseHeader struct {
        Code    int    `json:"code"`
        Message string `json:"message,omitempty"`
    }	

    ResponseData []ResponseDataItem

    ResponseDataItem struct {
        Type       string                `json:"type"`
        Id         int                   `json:"id"`
        Attributes ResponseDataItemAttrs `json:"attributes"`
    }

    ResponseDataItemAttrs struct {
        Email      string `json:"email"`
        ArticleIds []int  `json:"article_ids"`
    }
)



func ReadResponse(filePath string) (Responce, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Responce{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	resp := Responce{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&resp); err != nil {
		return Responce{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

func validateResponseStructure(resp Responce)  error {
	if resp.Header.Code != 0 {
		return   fmt.Errorf("invalid JSON structure: header code is not 0")
	}

	if resp.Header.Message != "" {
		return   fmt.Errorf("invalid JSON structure: header message is not empty")
	}

	if len(resp.Data) == 0 {
		return   fmt.Errorf("invalid JSON structure: data is empty")
	}

	for _, item := range resp.Data {
		if item.Type == "" || item.Id == 0 || item.Attributes.Email == "" {
			return   fmt.Errorf("invalid JSON structure: missing required fields")
		}
	}

	return  nil
}

func main() {
	filePath := "test/json/list.json" // Путь к файлу JSON
	resp, err := ReadResponse(filePath)
	if err != nil {
		panic(err)
	}
	
    validateResponseStructure(resp)
	jsonData, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	// Выводим JSON-строку в терминале
	fmt.Printf("%s\n", jsonData)
}