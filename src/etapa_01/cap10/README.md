# Capítulo 10

## Exemplos

### Concorrência dentro de Main

```go
{{#include exemplo01/main.go}}
```

### Chamando de forma concorrente a mesma função várias vezes
```go
{{#include exemplo02/main.go}}
```

### Demonstração com time.Sleep()
```go
{{#include exemplo03/main.go}}
```

### Troca de informações entre Goroutines com canais
```go
{{#include exemplo04/main.go}}
```

### Mais de uma Goroutine enviando dados para o mesmo canal
```go
{{#include exemplo05/main.go}}
```

### Direção dos canais no momento de definição da função
```go
{{#include exemplo06/main.go}}
```

### Select para receber informações de canais que tem dados disponíveis
```go
{{#include exemplo07/main.go}}
```

### Timeout no Select quando não há dados disponíveis em nenhum canal
```go
{{#include exemplo08/main.go}}
```

### Valor default em Select
```go
{{#include exemplo09/main.go}}
```

### Buffer de canal
```go
{{#include exemplo10/main.go}}
```

---

## Exercícios

### Exercício 01

```go
{{#include ex01/main.go}}
```

### Exercício 02

```go
{{#include ex02/main.go}}
```

### Exercício 03

```go
{{#include ex03/main.go}}
```