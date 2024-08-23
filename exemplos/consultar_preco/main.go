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

	config, err := corvo.NewConfig(
		"seu cartão postagem",
		"encoded user & senha", //https://cws.correios.com.br/ajudas
		urls,
		corvo.ConfigWithDefaultDeclaredValue(200),
		corvo.ConfigWithProductSpecification(500, 20, 20, 20),
		corvo.ConfigWithDeliveryType(2),                                      // Tipo do objeto da postagem: 1 - Envelope, 2 - Pacote; 3 - Rolo.
		corvo.ConfigWithCheckPriceAdditionalServices([]string{"001", "019"}), // "001" --> AVISO DE RECEBIMENTO | "019" --> VALOR DECLARADO NACIONAL
	)

	if err != nil {
		panic(err)
	}

	ws := corvo.NewCorreiosWebServices(config)
	codigoProduto := "03310" //PAC CONTRATO PGTO ENTREGA
	cepDestino := "22222222"

	data, er := ws.CheckDeliveryProductPrice(codigoProduto, cepDestino)
	if er != nil {
		// trate o erro
		log.Fatal(er)
	}

	log.Println(data.StrPrice) // string

	//se voê usar a corvo.ConfigWithFloatPriceEnabled(), é possível acessar desta forma
	log.Println(data.FloatPrice) // float64
}
