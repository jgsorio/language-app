package services

import (
	"bytes"
	"encoding/json"
	"io"
	"language-app-api/models"
	"net/http"
	"os"
)

type IaService struct {}

func (s IaService) SearchWords() []models.Word {
	openaiKey := os.Getenv("OPENAI_KEY")
	
	if openaiKey == "" {
		panic("OPENAI_KEY is not set")
	}

	message := map[string]interface{}{
        "model": "gpt-3.5-turbo",
        "messages": []map[string]string{
            {
				"role": "system",
				"content": "A sua missão será retornar um conjunto de palavras em inglês e suas traduções em portugues, contendo 4 alternativas sendo uma delas a correta e precisa estar na ordem aleatória. O formato retornado deverá ser assim: [{\"english_word\": \"hello\", \"translate\": \"olá\", \"options\": [\"olá\", \"tchau\", \"bem vindo\", \"surpresa\"]}].",
			},
			{
				"role": "user",
				"content": "Traga para mim 20 palavras em inglês e suas traduções em portugues.",
			},
        },
    }

	jsonData, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	response, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	response.Header.Set("Content-Type", "application/json")
	response.Header.Set("Authorization", "Bearer " + openaiKey)

	client := &http.Client{}

	resp, err := client.Do(response)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	content := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	words := []models.Word{}
	json.Unmarshal([]byte(content), &words)

	return words

}
