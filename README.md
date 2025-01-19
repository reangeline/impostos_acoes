## Decisões Técnicas e Arquiteturais

### Clean Architecture

A solução foi desenvolvida seguindo os princípios da Clean Architecture para garantir separação de responsabilidades, modularidade e testabilidade. As camadas estão organizadas da seguinte forma:

- **internal/domain/entity**: Contém as entidades principais (Wallet, Book) e as regras de negócio relacionadas, como cálculos de impostos e preço médio ponderado.

- **internal/domain/usecase**: Implementa os casos de uso, como o processamento das operações e cálculos de impostos.

- **internal/domain/presentation**: Gerencia a interação com o mundo externo (neste caso, entrada via stdin e saída formatada em JSON).

- **cmd**: Inclui o arquivo principal (main.go) para inicialização.

### Arredondamento

- **internal/util** Foi implementado um tipo personalizado chamado Decimal para garantir que valores decimais sejam sempre exibidos com duas casas decimais no JSON.

### Estado em Memória

O estado interno da aplicação é gerenciado por uma estrutura chamada AppState, que é reinicializada a cada execução do programa, garantindo que o estado seja limpo no início de cada execução, conforme solicitado no desafio.

### Justificativa para Uso de Frameworks ou Bibliotecas

- **`encoding/json`**: Usado para serializar e desserializar dados no formato JSON, que é o formato padrão de entrada e saída do programa. Esta é a biblioteca padrão do Go para manipulação de JSON e é eficiente e confiável.

- **`math`**: Utilizado para cálculos matemáticos, como arredondamento de valores decimais. Faz parte da biblioteca nativa do Go.

- **`bufio` e `os`**: Para manipulação da entrada padrão (stdin) e saída padrão (stdout).


_Nenhum framework externo adicional foi utilizado._

## Instruções para Executar o Projeto

### Pré-requisitos

Go 1.20 ou superior instalado.

### Passos

#### _Para copilar o projeto:_

1. **No terminal acesse a pasta correspondente ao projeto**

2. **Compile o projeto:**
    ```bash
    go build -o impostos-acoes ./cmd/main.go
    ```

#### _Para executar o projeto utilizando JSON direto no terminal:_

1. **Execute o binário gerado:**
    ```bash
    ./impostos-acoes
    ```

2. **Insira o JSON direto no terminal**
Exemplo com entrada direta:
    
    ```bash
    [{"operation":"buy", "unit-cost":10.00, "quantity":10000}, {"operation":"sell", "unit-cost":20.00, "quantity":5000}]
    ```

#### _Para executar o projeto utilizando um arquivo JSON:_
1. **Insira o JSON seguido do arquivo copilado**
Exemplo com entrada de um arquivo:

    ```bash
    cat input.json | ./impostos-acoes
    ```

## Instruções para Executar os Testes

### Certifique-se de estar no diretório raiz do projeto.

1. **Execute os testes com o comando:**
    ```bash
    go test ./test/...
    ```

2. **Verifique os resultados no terminal.**



## Notas Adicionais

A aplicação foi projetada utilizando clean code e clean architeture.

Para os teste foi utilizado AAA Pattern (Arrange, Act, Assert).

O github do projeto é https://github.com/reangeline/impostos_acoes
