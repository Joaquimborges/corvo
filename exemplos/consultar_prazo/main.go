package main

import (
	"log"

	"github.com/Joaquimborges/corvo"
)

func main() {
	configs := corvo.Config{
		PostCard:          "seu cartão postagem",
		AuthorizationCode: "encoded user & senha", //https://cws.correios.com.br/ajudas
		UrlMapper: map[corvo.EndpointURL]string{
			corvo.GenerateAccessTokenURL: "url referente ao tipo de autenticacao",
			//autentica, cartaopostagem ou contrato

			corvo.CheckDeliveryDueDateURL: "url refente a api prazo",
		},
		OriginZipCode: "11111111",
	}

	ws := corvo.NewCorreiosWebServices(&configs)
	codigoProduto := "03310" //PAC CONTRATO PGTO ENTREGA
	cepDestino := "22222222"

	response, err := ws.CheckDeliveryDueDate(codigoProduto, cepDestino)
	if err != nil {
		// trate o erro...
		log.Fatal(err)
	}

	log.Println(response.MaxDueDate)   //data maxima da entrega
	log.Println(response.DeliveryTime) // prazo da entrega (qtd dias)
}
