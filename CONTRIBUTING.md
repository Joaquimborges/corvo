# Contribuindo com o Corvo

Obrigado por considerar contribuir com o Corvo! Para mantermos um ambiente de desenvolvimento saudável e organizado, pedimos que siga as diretrizes abaixo.

## Como Contribuir

1. Faça um Fork do Repositório

   - Primeiro, faça um fork do repositório para a sua conta do GitHub.

   ```bash
   git clone https://github.com/Joaquimborges/corvo.git
   cd corvo
   ```

2. Crie uma Branch

   - Crie uma branch a partir da main para trabalhar na sua contribuição. Nomeie a branch de forma descritiva, por exemplo:

   ```bash
   git checkout -b feature/nome-da-sua-feature
   ```

3. Adicione seus Código e Testes

   - Ao adicionar novas funcionalidades ou corrigir bugs, certifique-se de incluir testes unitários para o código adicionado ou modificado.

   - Utilize uma cobertura de testes adequada para garantir que sua alteração não quebre outras partes do código.

Exemplo de criação de teste:

```go
func TestMinhaNovaFuncao(t *testing.T) {
    resultado := MinhaNovaFuncao()
    esperado := "resultado esperado"
    if resultado != esperado {
        t.Errorf("Resultado foi %s; esperado %s", resultado, esperado)
    }
}
```

4. Certifique-se de que Todos os Testes Passam

   - Antes de enviar sua contribuição, execute todos os testes para garantir que não há regressões

   ```bash
   make test
   ```

5. Envie um Pull Request (PR)

   - Depois de concluir suas alterações, faça o commit e envie o push para o seu fork. Em seguida, abra um Pull Request (PR) apontando para a branch main do repositório original.

   ```bash
   git push origin feature/nome-da-sua-feature
   ```

   - No PR, inclua uma descrição clara e objetiva do que foi alterado, o motivo da alteração e como testar as mudanças. Se necessário, referencie issues relacionadas.

6. Aguarde a Revisão

   - Um mantenedor revisará seu PR. Pode ser solicitado que você faça ajustes ou melhorias. Responda às solicitações e faça os commits adicionais conforme necessário.

## Regras de Estilo de Código

- Mantenha o código limpo e bem organizado.
- Siga as convenções de nomenclatura do Go.
- Documente as funções, especialmente as públicas, com comentários claros.
