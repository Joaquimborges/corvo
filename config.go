package corvo

import "errors"

type cfgOption func(*Config)

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
	DeliveryType                int
	shouldGenerateFloatPrice    bool
	productSpecificationsWasSet bool

	// serviços adicionais para a consulta de preços
	/*
		 	exemplo: []string{"001", "019"}

			"001" --> AVISO DE RECEBIMENTO
			"019" --> VALOR DECLARADO NACIONAL
	*/
	// 	será usado na base para calcular o preço do frete.
	AdditionalServices []string
	UrlMapper          map[EndpointURL]string
}

func NewConfig(postCard, authorizationCode string, urls map[EndpointURL]string, options ...cfgOption) (*Config, error) {
	if postCard == "" || authorizationCode == "" {
		return nil, errors.New("cartão postagem e código de autorização são obrigatórios")
	}

	if len(urls) == 0 {
		return nil, errors.New("o mapper de urls não pode estar vazio")
	}

	var config Config
	for _, option := range options {
		option(&config)
	}

	if _, ok := urls[CheckDeliveryProductPriceURL]; ok {
		if len(config.AdditionalServices) == 0 {
			return nil, errors.New("se você pretende usar a api de preço, serviços adicionais é um parametro obrigatório")
		}

		if !config.productSpecificationsWasSet {
			return nil, errors.New("adicione as espcificações do produto para usar a api de preço [peso, cumprimento, altura, largura]")
		}
	}
	return &config, nil
}

func ConfigWithFloatPriceEnabled() cfgOption {
	return func(c *Config) {
		c.shouldGenerateFloatPrice = true
	}
}

func ConfigWithCheckPriceAdditionalServices(additionalServices []string) cfgOption {
	return func(c *Config) {
		c.AdditionalServices = additionalServices
	}
}

func ConfigWithDeliveryType(deliveryType int) cfgOption {
	return func(c *Config) {
		c.DeliveryType = deliveryType
	}
}

func ConfigWithProductSpecification(weight, fulfillment, height, width int) cfgOption {
	return func(c *Config) {
		c.ObjectBaseWeight = weight
		c.BaseFulfillment = fulfillment
		c.BaseHeight = height
		c.BaseWidth = width
		c.productSpecificationsWasSet = true
	}
}

func ConfigWithOriginZipCode(zipCode string) cfgOption {
	return func(c *Config) {
		c.OriginZipCode = zipCode
	}
}

func ConfigWithDefaultDeclaredValue(value int) cfgOption {
	return func(c *Config) {
		c.DefaultDeclaredValue = value
	}
}
