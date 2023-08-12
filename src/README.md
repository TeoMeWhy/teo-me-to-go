# Téo Me To Go

> Caso prefira navegar nos arquivos, acesse [github.com/teomewhy/teo-me-to-go](https://github.com/teomewhy/teo-me-to-go) e confira nosso projeto.

Seja bem-vindo(a) à minha trilha de estudos para GoLang, meu desejo é aprender sobre backend estudando esta linguagem. Para alcançar este objetivo, montei este material de estudos junto ao chatGPT, segue o prompt utilizados:

```
Desejo me tornar um desenvolvedor backend.
Gostaria de estudar GoLang e preciso de sua ajuda para elaborar um plano de estudo detalhado.

Para cada etapa deste plano, considere um desafio de programação em formato de projeto.
```

A partir disso, temos as etapas abaixo. Vale lembrar que ao avançarmos os estudos, este plano pode sofrer alterações.

E ah, tudo isso está sendo feio ao vivo na Twitch em meu canal [Téo Me Why](https://www.twitch.tv/teomewhy), sempre às 9AM (UTC-3) de segunda a sexta. Conheça mais sobre mim [aqui](https://github.com/teocalvo).

## Etapa 1: Fundamentos do GoLang (em progresso)

- Aprenda os conceitos básicos da linguagem Go, como sintaxe, tipos de dados, estruturas de controle (if, for, switch), funções e pacotes.
- Recursos recomendados:
    - [Tour pelo Go](https://tour.golang.org/welcome/1): Um tutorial interativo e introdutório que cobre os conceitos básicos do Go.
    - [Effective Go](https://golang.org/doc/effective_go.html): Um guia oficial do Go que descreve as melhores práticas e convenções para escrever código Go.

Utilizaremos o livro [Introdução à Linguagem Go - Crie programas escaláveis e confiáveis de Caleb Doxsey (O'Reilly). Copyrigth 2016 Caleb Doxsey, _ISBN_ 978-1-4919-4195-9](https://amzn.to/3QtuU0R)

**Desafio de Projeto: Criar uma aplicação CLI (Command Line Interface)** Crie uma aplicação de linha de comando simples que realiza uma tarefa específica, como calcular estatísticas básicas em um conjunto de dados ou converter unidades de medida. Isso ajudará a aplicar os conceitos básicos aprendidos e a familiarizar-se com a sintaxe e a estrutura do Go.

## Etapa 2: Entendendo algoritmos

Vamos seguir os estudos de Go Lang a partir od livro [Entendendo algoritmos - Um guia ilustrado para programadores e outros curiosos, Aditya Y. Bhargava](https://amzn.to/3YtSVux).

Nossa motivação é entender e implementar os algoritmos abordados pelo livro, fixando assim conceitos não só da sintaxe de Go, mas aplicando conhecimentos voltados à computação que podem ser utilizado em desafios no mercado de trabalho.

## Etapa 3: Concorrência e Paralelismo (não iniciada)

- Aprenda a trabalhar com goroutines (threads leves) e canais (channels) para escrever código concorrente em Go.
- Explore técnicas de paralelismo para processar tarefas em paralelo e obter um melhor desempenho.
- Recursos recomendados:
    - [Concurrency in Go](https://www.youtube.com/watch?v=f6kdp27TYZs): Uma palestra em vídeo do criador do Go, Rob Pike, sobre concorrência em Go.
    - [Share Memory By Communicating](https://blog.golang.org/codelab-share): Um tutorial do blog oficial do Go que explora o uso de canais para comunicação entre goroutines.

**Desafio de Projeto: Desenvolver uma Aplicação de Web Scraping Paralela** Crie uma aplicação que faça web scraping de várias páginas simultaneamente usando goroutines e canais. Isso demonstrará como aproveitar a concorrência e o paralelismo em Go para melhorar a eficiência de operações intensivas em E/S.

## Etapa 4: Banco de Dados e Persistência (não iniciada)

- Aprenda a usar bancos de dados em Go, como o MySQL, PostgreSQL ou MongoDB.
- Explore técnicas de persistência, como a criação, leitura, atualização e exclusão de dados (CRUD) em um banco de dados.
- Recursos recomendados:
    - [Database/SQL](https://golang.org/pkg/database/sql/): A documentação oficial do pacote `database/sql`, que é uma biblioteca padrão para interagir com bancos de dados SQL em Go.
    - [MongoDB Go Driver Documentation](https://pkg.go.dev/go.mongodb.org/mongo-driver): A documentação oficial do driver Go para MongoDB.

**Desafio de Projeto: Construir um Sistema de Gerenciamento de Tarefas** Crie um sistema de gerenciamento de tarefas que permita criar, atualizar, listar e excluir tarefas em um banco de dados. Use um banco de dados de sua escolha e aplique técnicas adequadas de persistência de dados.

## Etapa 5: Segurança e Autenticação (não iniciada)

- Aprenda a implementar recursos de segurança e autenticação em aplicativos Go.
- Explore a autenticação baseada em token (como JWT) e a proteção de rotas e endpoints.
- Recursos recomendados:
    - [Authentication and Authorization](https://pkg.go.dev/github.com/dgrijalva/jwt-go): A documentação oficial do pacote JWT para autenticação em Go.
    - [Securing Your Go APIs With JWTs](https://www.youtube.com/watch?v=sVq2T4HtVus): Um tutorial em vídeo que mostra como implementar autenticação baseada em JWT em uma API Go.

**Desafio de Projeto: Implementar Autenticação JWT em uma API** Adicione recursos de autenticação baseada em JWT a uma API existente que você tenha desenvolvido. Isso ajudará você a entender como proteger suas APIs e autenticar usuários usando tokens JWT.

Essa é uma sugestão de plano de estudo para se tornar um desenvolvedor backend usando GoLang. Lembre-se de ajustar o plano de acordo com suas necessidades e interesses. Além disso, não hesite em explorar outras bibliotecas, frameworks ou áreas específicas do desenvolvimento backend em Go que você esteja interessado. Boa sorte em sua jornada de aprendizado!