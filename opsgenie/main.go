package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func CalledOpsGenie(ctx context.Context, event *MyEvent) (*string, error) {
	url := "https://api.opsgenie.com/v2/alerts"

	jsonStr := []byte(`{
        "message": "Risk Analysis Anticipation Data Loader - Erro Execução Lambda ra_anticipation_process",
        "description": "Houve um erro na execução do lambda e a atualização da antecipação de recebíveis pode ser impactada. Consultar os respectivos logs para identificação da causa e notificar R&M para realizar o upload dos arquivos impactados novamente, quando resolvido o erro.",
        "responders": [
            {
                "name": "Fraud-Prevention",
                "type": "team"
            }
        ],
        "details": {
            "Team": "Fraud-Prevention",
			"Application": "Risk Analysis Anticipation Data Loader",
			"Priority":"P4"
        },
        "priority": "P4",
        "tags": ["Team:Fraud-Prevention", "Application:Risk Analysis Anticipation Data Loader", "Priority:P4", "P4"]
    }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", os.Getenv("OPSGENIE_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Status == "202 Accepted" {
		result := "Alerta criado com sucesso"
		return &result, nil
	} else {
		result := fmt.Sprintf("Falha ao criar o alerta. Status da resposta: %s", resp.Status)
		return &result, nil
	}
}

func main() {
	lambda.Start(CalledOpsGenie)
}
