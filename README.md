# fundamentos-go

Repositório de estudos dos fundamentos da linguagem Go, organizado por tópicos com exemplos práticos e comentados.

---

## Estrutura

```
fundamentos-go/
├── variaveis/
├── tipos-de-dados/
├── operadores/
├── if-else/
├── switch/
├── loops/
├── arrays_e_slices/
├── maps/
├── funcoes/
├── structs/
├── metodos/
├── ponteiros/
├── "heranca"/
├── interfaces/
│   ├── formas/
│   └── tipo-generico/
├── genericos/
├── funcoes_avancadas/
│   ├── anonimas/
│   ├── closure/
│   ├── defer/
│   ├── init/
│   ├── panic-recover/
│   ├── ponteiros/
│   ├── recursivas/
│   ├── retorno_nomeado/
│   └── variaticas/
└── pacotes/
    └── auxiliar/
```

---

## Conceitos

### Variáveis — `variaveis/`
- Declaração explícita com `var`
- Inferência de tipo com `:=`
- Declaração em bloco `var ( ... )`
- Múltiplas variáveis na mesma linha
- Constantes com `const`
- Troca de valores entre variáveis

### Tipos de Dados — `tipos-de-dados/`
- Inteiros: `int8`, `int16`, `int32`, `int64`
- Inteiros sem sinal: `uint`
- Aliases: `rune` (= `int32`), `byte` (= `uint8`)
- Reais: `float32`, `float64`
- Texto: `string` (aspas duplas) e `char` (aspas simples)
- Booleano: `bool`
- Erros: pacote `errors`

### Operadores — `operadores/`
- Aritméticos: `+`, `-`, `*`, `/`, `%`
- Atribuição: `=`, `:=`
- Relacionais: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Lógicos: `&&`, `||`, `!`
- Go não permite operações entre tipos diferentes (ex: `int16` + `int32`)

### Condicionais — `if-else/`
- `if / else`
- `if` com instrução de inicialização: `if variavel := valor; condicao { }`
- Variáveis declaradas no `if` não existem fora dele
- `else if` encadeado

### Switch — `switch/`
- `switch` com `case` e `default`
- Switch sem expressão (substitui if-else encadeado)
- Switch com `time.Now().Weekday()`
- Alternativa com `map` para tabelas de lookup

### Loops — `loops/`
- Go tem apenas o `for` — sem `while` ou `do-while`
- Forma 1: `for condicao { }` (equivale ao while)
- Forma 2: `for init; condicao; incremento { }`
- Forma 3: `for init; condicao; { }` (incremento no corpo)
- Forma 4: `for index, valor := range slice { }`
- Forma 5: `for chave, valor := range map { }`
- `_` (underline) para ignorar index ou chave
- Não é possível usar `range` em `struct`

### Arrays e Slices — `arrays_e_slices/`
- Array de tamanho fixo: `[5]string`
- Array com inferência de tamanho: `[...]string{ ... }`
- Slice (tamanho dinâmico): `[]string{ ... }`
- `append` para adicionar elementos ao slice
- `reflect.TypeOf` para inspecionar o tipo em runtime

### Maps — `maps/`
- Declaração e inicialização: `map[string]string{ ... }`
- Acesso por chave
- Maps aninhados: `map[string]map[string]string`
- `delete(map, chave)` para remover elemento
- Adição de elementos após criação

### Funções — `funcoes/`
- Função com parâmetros e retorno único
- Múltiplos valores de retorno
- `_` para ignorar retornos desnecessários

### Structs — `structs/`
- Definição de struct com `type`
- Structs aninhadas (composição)
- Inicialização com e sem nome dos campos
- Struct com campos parcialmente preenchidos

### Métodos — `metodos/`
- Método com receiver por valor: `func (p pessoa) metodo()`
- Método com receiver por ponteiro: `func (p *pessoa) metodo()` — modifica o original
- Diferença entre receiver por valor e por ponteiro

### Ponteiros — `ponteiros/`
- `&variavel` — obtém o endereço de memória
- `*ponteiro` — desreferencia (acessa o valor apontado)
- Ponteiro como variável: `var p *int = &variavel`

### Herança por Composição — `"heranca"/`
- Go não tem herança clássica
- Embedding de struct dentro de struct
- Acesso direto aos campos da struct embutida
- Promoção de campos e métodos

### Interfaces — `interfaces/`

**Formas (`interfaces/formas/`):**
- Definição de interface com `type nome interface { ... }`
- Implementação implícita — sem `implements`
- Polimorfismo: diferentes tipos satisfazem a mesma interface
- Exemplo prático: `retangulo` e `circulo` implementando `forma`

**Tipo Genérico (`interfaces/tipo-generico/`):**
- `interface{}` como tipo genérico (pré-Go 1.18)
- Uso em parâmetros e maps com qualquer tipo
- Considerado gambiarra — substituído por Generics na 1.18

### Generics — `genericos/`
- Introduzidos no Go 1.18
- Definição de constraint com interface: `type number interface { int | float64 }`
- Função genérica: `func calculos[T number](n1, n2 T) T`
- Elimina a necessidade de `interface{}` para tipos numéricos

### Funções Avançadas — `funcoes_avancadas/`

| Tópico | Arquivo | Conceito |
|---|---|---|
| Anônimas | `anonimas/` | Funções sem nome, IIFE, funções como valor |
| Closure | `closure/` | Função que captura variáveis do escopo externo |
| Defer | `defer/` | Execução adiada para o fim da função (LIFO) |
| Init | `init/` | Executada antes do `main`, uma por arquivo, para inicialização de pacote |
| Panic/Recover | `panic-recover/` | `panic` interrompe execução; `recover` restaura dentro de `defer` |
| Ponteiros | `ponteiros/` | Passagem por referência em funções |
| Recursivas | `recursivas/` | Funções que chamam a si mesmas — exemplo: Fibonacci |
| Retorno Nomeado | `retorno_nomeado/` | Nomear os valores de retorno na assinatura da função |
| Variádicas | `variaticas/` | `...tipo` — aceita número variável de argumentos, sempre o último parâmetro |

### Pacotes — `pacotes/`
- Criação de pacote local (`package auxiliar`)
- Funções exportadas (letra maiúscula) vs não exportadas (letra minúscula)
- `go.mod` e `go.sum` — gerenciamento de módulos
- Importação de pacote externo (`github.com/badoux/checkmail`)
- Organização de múltiplos arquivos no mesmo pacote

---

## Como executar

Cada pasta é um programa independente. Para rodar um exemplo:

```bash
go run variaveis/variaveis.go
go run loops/loops.go
go run funcoes_avancadas/closure/closure.go
```

Para o exemplo com pacotes e dependências externas:

```bash
cd pacotes
go mod tidy
go run main.go
```
