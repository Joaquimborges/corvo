package main

import (
	"log"

	"github.com/Joaquimborges/corvo"
)

func main() {
	urls := map[corvo.EndpointURL]string{
		corvo.GenerateAccessTokenURL:       "url referente ao tipo de autenticacao", //autentica, cartaopostagem ou contrato
		corvo.CheckDeliveryProductPriceURL: "url refente a api preço",
	}

	config, er := corvo.NewConfig(
		"seu cartão postagem",
		"encoded user & senha", //https://cws.correios.com.br/ajudas
		urls,
		corvo.ConfigWithOriginZipCode("11111111"),
	)

	if er != nil {
		panic(er)
	}

	ws := corvo.NewCorreiosWebServices(config)
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
