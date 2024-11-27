# To Do List - REST API

🌍 *[English](README.md) ∙ [Português](README_pt.md)*

Neste repositório você encontrará um exemplo de uma API REST feita em Go com o tema To Do List. Esta aplicação foi
construída para melhorar minhas habilidades em Go e APIs REST, e agora eu gostaria de compartilhar esses conhecimentos
adquiridos com vocês. Este repositório ainda é monitorado de perto pelo criador, então sinta-se à vontade para contribuir
com o repositório relatando problemas ou sugerindo novos recursos. Se este projeto for útil para você, considere dar uma
estrela ao repositório.

## Conceitos importantes aplicados nesse repositório

* API REST com autenticação simples usando token JWT
* CRUD usando PostgreSQL e Go
* Arquitetura hexagonal
* Dockerização

## Como Executar

### Pré-requisitos

Você terá duas maneiras de aproveitar todos os recursos do repositório. A primeira maneira é a mais simples, você só
precisará baixar este repositório e ter o [Docker](https://www.docker.com/get-started/) instalado na sua máquina. Após
fazer isso, vá até à pasta do projeto e execute o comando abaixo, deste modo toda a API estará em execução sem precisar
de nenhuma configuração adicional:

```shell
$ docker-compose up
```

No entanto, se você for um programador e quiser fazer alguma modificação no projeto, aconselho-o a utilizar outro
container que também está disponível no projeto e que terá apenas a base de dados. Então, para que a API funcione em sua
máquina, primeiramente será necessário ter o [Go](https://go.dev/dl/) instalado, além deste repositório e do Docker já
citados acima. Com todos os requisitos anteriores atendidos e dentro da pasta do projeto, para executar a API basta
executar os seguintes comandos:

```shell
$ cd res/docker
$ docker compose up --detach
$ cd ../..
$ go run main.go
```

### Documentação

Você poderá testar a execução de todas as rotas da API usando apenas seu navegador. Se estiver usando os valores padrões
do projeto para `HOST` e `PORT`, basta digitar o seguinte endereço na barra de pesquisa do seu navegador
(Note que se você alterar esses valores, precisará atualizar esse endereço de pesquisa colocando os valores escolhidos
para o host e a porta):

```
localhost:8000/api/documentation/index.html
```

Observe que para utilizar todas as rotas da API você precisará estar autenticado, o que é um processo simples de fazer.
Para isso, procure as rotas **Sign Up** e **Sign In**. Nessas rotas, você encontrará todas as informações necessárias
para se registrar e fazer login na API. Após realizar uma das requisições anteriores, você receberá no corpo da resposta
o seu `access_token`. O valor desse token de acesso deve ser copiado e colado no campo que aparece quando você clica no
botão **Authorize** acima de todas as rotas no lado esquerdo, como você pode ver na imagem abaixo.

![todo-rest-api](https://user-images.githubusercontent.com/89457923/169172172-1c112bf0-14d0-43c2-89d9-ba52c8391ac2.png)

### Configurações

Você pode modificar facilmente algumas configurações da API, como o endereço do seu servidor ou o banco de dados
PostgreSQL que ela acessará, por exemplo. Para isso, basta modificar os valores presentes no [arquivo de configuração de
ambiente](.env) e tudo ficará como você definiu.
