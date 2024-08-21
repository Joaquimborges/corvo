package main

import (
	"log"

	"github.com/Joaquimborges/corvo"
)

func main() {
	configs := corvo.Config{
		PostCard:             "seu cartão postagem",
		AuthorizationCode:    "encoded user & senha", //https://cws.correios.com.br/ajudas
		OriginZipCode:        "11111111",
		DefaultDeclaredValue: 200, // valor declarado do produto
		ObjectBaseWeight:     500, // Peso do objeto em gramas.
		BaseFulfillment:      20,
		BaseHeight:           20,
		BaseWidth:            20,
		DeliveryType:         2,                      // Tipo do objeto da postagem: 1 - Envelope, 2 - Pacote; 3 - Rolo.
		AdditionalServices:   []string{"001", "019"}, // "001" --> AVISO DE RECEBIMENTO | "019" --> VALOR DECLARADO NACIONAL
		UrlMapper: map[corvo.EndpointURL]string{
			corvo.GenerateAccessTokenURL:       "url referente ao tipo de autenticacao", //autentica, cartaopostagem ou contrato
			corvo.CheckDeliveryProductPriceURL: "url refente a api preço",
		},
	}

	ws := corvo.NewCorreiosWebServices(&configs)
	codigoProduto := "03310" //PAC CONTRATO PGTO ENTREGA
	cepDestino := "22222222"

	data, err := ws.CheckDeliveryProductPrice(codigoProduto, cepDestino)
	if err != nil {
		// trate o erro
		log.Fatal(err)
	}

	log.Println(data.Price)
}
