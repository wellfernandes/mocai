# Mocai

![mocai](../../.././img/mocai.svg)

#### [README (English)](/README.md)

Uma biblioteca Go para geração de dados de teste, permitindo criar mocks de entidades de forma simples e eficiente.

## 📖 Descrição
**Mocaí** é uma biblioteca open-source em Go projetada para simplificar a geração de dados fictícios realistas para entidades como Pessoa, Endereço, Telefone, Empresa, CPF, CNPJ, Certidões, RG e Título de Eleitor. O objetivo é tornar o desenvolvimento e os testes mais eficientes, fornecendo dados aleatórios, porém consistentes, que simulam cenários reais de forma prática e confiável.

## 🌟 Curiosidade sobre o Nome
O nome Mocai é uma homenagem à iniciativa brasileira por trás da biblioteca. Ele surgiu da combinação de "mock" (simulação ou dado fictício) com "açaí", uma fruta típica da Amazônia brasileira, conhecida por sua energia e versatilidade. Assim como o açaí é essencial para muitos brasileiros, o Mocai busca ser uma ferramenta essencial para desenvolvedores que precisam de dados de teste eficientes e de qualidade. 🇧🇷

## 🛠️ Principais Recursos
- **Geração de Dados Aleatórios:** Gere mocks para entidades com dados variados e realistas (Pessoa, Endereço, Telefone, Empresa, CPF, CNPJ, Certidões, RG, Título de Eleitor e mais).
- **Consistência:** Garante que os dados gerados sejam consistentes e adequados para testes.
- **Facilidade de Uso:** API simples e intuitiva para integração rápida.
- **Extensibilidade:** Adicione facilmente novas entidades ou idiomas. A arquitetura é modular e pronta para expansão.
- **Open Source:** Colabore, sugira melhorias e contribua para o crescimento da biblioteca.

## 📦 Entidades Suportadas
- Pessoa (com gênero, idade, CPF)
- Gênero (geração independente)
- Endereço (rua, número, cidade, estado, UF, CEP)
- Telefone (DDD, número)
- Empresa (nome, CNPJ)
- CPF (Cadastro de Pessoa Física)
- Certidões (Nascimento, Casamento, Óbito)
- RG (Identidade)
- Título de Eleitor

## 🌐 Suporte a Idiomas
- **ptbr** (Português do Brasil) atualmente suportado. A estrutura permite fácil adição de novos idiomas no futuro.

## 🚀 Por que usar o Mocai?
- **Produtividade:** Reduza o tempo gasto na criação de dados de teste.
- **Qualidade:** Melhore a cobertura e a eficácia dos seus testes com dados realistas.
- **Flexibilidade:** Adapte os mocks às necessidades específicas do seu projeto.
- **Comunidade:** Faça parte de uma comunidade open-source que valoriza a colaboração e a inovação.

## ⚙️ Requisitos
- Go 1.23.4 ou superior

## 🚀 Como Começar
### Instalação
Para começar a usar o Mocai, instale a biblioteca com:

```sh
go get github.com/brazzcore/mocai
```

### Uso Básico
Importe a biblioteca e gere mocks usando os métodos da struct `Mocker`:

```go
package main

import (
    "fmt"
    "log"
    "github.com/brazzcore/mocai/pkg/mocai"
)

func main() {
    // Cria uma instância do Mocker para português do Brasil
    mocker := mocai.NewMocker(
        mocai.WithLanguage("ptbr"),
        mocai.WithFormatted(true),
    )

    // Gera um endereço fictício
    address, err := mocker.NewAddress()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Endereço: %s, %d - %s, %s (%s) - %s\n", address.Street, address.Number, address.City, address.State, address.UF, address.ZIP)

    // Gera uma pessoa fictícia
    person, err := mocker.NewPerson()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Pessoa: %s %s, Gênero: %s, Idade: %d, CPF: %s\n", person.FirstNameMale, person.LastName, person.Gender.Identity, person.Age, person.CPF.Number)

    // Gera uma empresa fictícia
    company, err := mocker.NewCompany()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Empresa: %s, CNPJ: %s\n", company.BrazilianCompany.Name, company.BrazilianCompany.CNPJ)
}
```

### Mensagens de Erro e Localização

As mensagens de erro são localizadas quando há chave de tradução disponível. Quando ocorre um erro (ex: dados inválidos, idioma não suportado ou falha na geração), a mensagem será apresentada no idioma configurado para a instância do `Mocker`, se houver tradução. Em alguns casos, mensagens fixas podem aparecer caso a cobertura de traduções não seja completa.

Não é necessário realizar nenhuma etapa extra para a localização das mensagens de erro — o Mocai faz isso automaticamente para todos os idiomas suportados onde há chave de tradução.

> **Nota:** Cada chamada de método como `NewPerson()` ou `NewAddress()` gera um novo mock com dados aleatórios. A instância do `Mocker` é imutável quanto à configuração (idioma, formatação, fonte de aleatoriedade).

#### Sobre Idiomas
Atualmente, apenas "ptbr" está implementado. Para suportar outros idiomas, contribua com novos arquivos de tradução e mocks.

#### Sobre Formatação
A opção `WithFormatted` controla se documentos como CPF/CNPJ são retornados formatados (ex: `123.456.789-00`) ou apenas números (`12345678900`).

### Uso Avançado

#### Interface MockGenerator
A interface `MockGenerator` define o contrato para geração de dados fictícios. Use-a para injeção de dependência nos seus testes:

```go
func CreateUser(generator mocai.MockGenerator) (*User, error) {
    person, err := generator.NewPerson()
    if err != nil {
        return nil, err
    }
    return &User{Name: person.FirstNameMale + " " + person.LastName}, nil
}
```

#### Providers Customizados
Você pode injetar providers customizados para Address, Person e Company, integrando com APIs externas ou bancos de dados:

```go
mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithAddressProvider(meuProviderDeEndereco),
    mocai.WithPersonProvider(meuProviderDePessoa),
    mocai.WithCompanyProvider(meuProviderDeEmpresa),
)
```

#### Geração Determinística
Use `WithRandSource` para fornecer uma fonte de aleatoriedade customizada para dados reproduzíveis em testes:

```go
rnd := translations.NewSafeRandSource(meuRandFixo)
mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithRandSource(rnd),
)
```

### Exemplos
O diretório ***examples*** contém exemplos de uso:

- `mocker/`: Exemplo usando o ponto de entrada principal `mocai.NewMocker(opts ...Option)` com functional options para gerar mocks de maneira fluida e simplificada.

Para executar um exemplo:
```sh
cd examples/mocker
go run main.go
```

## 🧪 Rodando os Testes
Para rodar todos os testes:
```sh
go test ./...
```

## 🤝 Contribua
O Mocai é um projeto open-source, e sua contribuição é muito bem-vinda! Seja reportando bugs, sugerindo novas funcionalidades ou enviando pull requests, sua participação ajuda a melhorar a biblioteca para todos.

### Como Contribuir
1. **Reporte Problemas:** Encontrou um bug ou tem uma sugestão? Abra uma issue.
2. **Envie Pull Requests:** Siga as diretrizes de contribuição e envie suas melhorias. Sempre rode os testes antes de enviar.
3. **Discuta Ideias:** Participe das discussões e compartilhe suas ideias para o projeto.

### Diretrizes de Contribuição
1. Siga o padrão de código do projeto.
2. Adicione testes para novas funcionalidades.
3. Caso necessário, documente suas alterações no **README.md**.

## 📄 Licença
O Mocai é distribuído sob a **MIT License**.

### 🌟 Mocai: Gerando mocks, simplificando testes, acelerando desenvolvimento.
Seja bem-vindo ao projeto e sinta-se à vontade para explorar, usar e contribuir!

## 🌐 Servidor no Discord
Participe do servidor **Mocai** no Discord:
[Clique aqui e junte-se a nós!](https://discord.gg/TFRnQBkAMt)
