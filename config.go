package corvo

type Config struct {
	// cartão postagem: https://www.correios.com.br/correios-facil
	PostCard string

	// código de acesso API: https://cws.correios.com.br/acesso-componentes
	AuthorizationCode string

	// CEP de origem
	// usado para calcular o frete e prazo de entrega
	OriginZipCode string

	// valor declarado do produto
	// 	será usado na base para calcular o preço do frete.
	DefaultDeclaredValue int

	// Peso do objeto em gramas.
	// 	será usado na base para calcular o preço do frete.
	ObjectBaseWeight int

	// cumprimento base
	BaseFulfillment int
	// largura base
	BaseHeight int
	// altura base
	BaseWidth int

	// Tipo do objeto da postagem: 1 - Envelope, 2 - Pacote; 3 - Rolo.
	// 	será usado na base para calcular o preço do frete.
	DeliveryType int

	// serviços adicionais para a consulta de preços
	/*
		 	exemplo: []string{"001", "019"}

			"001" --> AVISO DE RECEBIMENTO
			"019" --> VALOR DECLARADO NACIONAL
	*/
	// 	será usado na base para calcular o preço do frete.
	AdditionalServices []string
	UrlMapper          map[urlKey]string
}
