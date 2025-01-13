# Mocai

![mocai](img/mocai-mascot.png)

Uma biblioteca Go para geração de dados de teste, permitindo criar mocks de entidades de forma simples e eficiente.
 
## 📖 Descrição

**Mocaí** é uma biblioteca open-source em Go projetada para simplificar a geração de mocks de entidades como Pessoa, Endereço, Telefone e muitas outras. Nosso objetivo é tornar o desenvolvimento e teste de aplicações mais eficiente, fornecendo dados aleatórios, porém consistentes, que simulam cenários reais de forma prática e confiável.

## 🌟 Curiosidade sobre o Nome
O nome Mocai é uma homenagem à iniciativa brasileira por trás da biblioteca. Ele surgiu da combinação de "mock" (termo em inglês para simulação ou dados fictícios) com "açaí", uma fruta típica da Amazônia brasileira, conhecida por sua energia e versatilidade. Assim como o açaí é essencial para muitos brasileiros, o Mocai busca ser uma ferramenta essencial para desenvolvedores que precisam de dados de teste eficientes e de qualidade. 🇧🇷

## 🛠️ Principais Recursos

- Geração de Dados Aleatórios: Crie mocks de entidades com dados variados e realistas.
- Consistência: Garanta que os dados gerados sejam consistentes e adequados para testes.
- Facilidade de Uso: Interface simples e intuitiva para integração rápida em seus projetos.
- Extensibilidade: Adicione novas entidades ou personalize as existentes conforme suas necessidades.
- Open Source: Colabore, sugira melhorias e contribua para o crescimento da biblioteca.

## 🚀 Por que usar o Mocai?
- Produtividade: Reduza o tempo gasto na criação de dados de teste.
- Qualidade: Melhore a cobertura e a eficácia dos seus testes com dados realistas.
- Flexibilidade: Adapte os mocks às necessidades específicas do seu projeto.
- Comunidade: Faça parte de uma comunidade open-source que valoriza a colaboração e a inovação.

## 📂 Estrutura do Projeto
Aqui está a estrutura de pastas do projeto:

````
mocai/

├── .github/

│ └── workflows/

│ └── workflow_dev.yaml

├── pkg/

│ ├── mocai/

│ │ ├── person.go

│ │ ├── address.go

│ │ ├── phone.go

│ │ ├── mocai.go

│ │ └── utils.go

│ └── locale/

│ ├── pt-br/

│ │ ├── constants/

│ │ │ ├── streets.go

│ │ │ ├── cities.go

│ │ │ ├── states.go

│ │ │ └── zipcodes.go

│ │ ├── person.go

│ │ ├── address.go

│ │ └── phone.go

│ └── en-us/

│ ├── constants/

│ │ ├── streets.go

│ │ ├── cities.go

│ │ ├── states.go

│ │ └── zipcodes.go

│ ├── person.go

│ ├── address.go

│ └── phone.go

├── examples/

│ ├── main.go

│ ├── mock_ptbr.go

│ └── mock_enus.go

├── go.mod

├── go.sum

└── README.md

````

---
