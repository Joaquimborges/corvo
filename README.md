# corvo

Corvo é uma biblioteca escrita em Go para acessar as APIs dos Correios Brasil.

> [!NOTE]
> A biblioteca atualmente está em fase de desenvolvimento.

> [!WARNING]
> A versão principal atual é zero (v0.x.x) para acomodar o desenvolvimento rápido e a iteração rápida enquanto obtém feedback antecipado dos usuários (agradecemos comentários sobre APIs!). A API pública pode mudar sem uma atualização de versão principal antes do lançamento da v1.0.0.

## Instalação

```bash
go get github.com/Joaquimborges/corvo
```

## Funcionalidades

- [Consulta prazo de entrega](#consulta-prazo-de-entrega)
- Consultar custo total da entrega

## [Contribuições](CONTRIBUTING.md)

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests. Por favor, certifique-se de seguir as boas práticas de codificação e inclua testes ao enviar uma nova funcionalidade.

> [!IMPORTANT]
> Serão apenas aceites pull requests provenientes de forks.

## Exemplos

#### Consulta prazo de entrega

```go
package main

import "github.com/Joaquimborges/corvo"

func main() {
	configs := &corvo.Config{
		PostCard:          "seu cartão postagem",
		AuthorizationCode: "encoded user & senha", //https://cws.correios.com.br/ajudas
		UrlMapper: map[urlKey]string{
			corvo.GenerateAccessTokenUrlKey:  "url referente ao tipo de autenticacao", //autentica, cartaopostagem ou contrato
			corvo.CheckDeliveryDueDateUrlKey: "url refente a api prazo",
		},
	}

	ws := corvo.NewCorreiosWebServices(configs)
	codigoProduto := "03310" //PAC CONTRATO PGTO ENTREGA

	response, err := ws.CheckDeliveryDueDate(codigoProduto, "11111111", "22222222")
	if err != nil {
		//...
	}

	response.MaxDueDate   //data maxima da entrega
	response.DeliveryTime // prazo da entrega (qtd dias)
}

```

### Contato

Se você tiver alguma dúvida ou sugestão, entre em contato através do e-mail joaquim.borges@alabuta.com
