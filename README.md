# GitLab File Checker

## Descrição
Esta aplicação em Go utiliza a API do GitLab para buscar projetos dentro de um grupo específico e verifica se o arquivo `values.yaml` contém o trecho:

```yaml
livenessProbe:
  enabled: true
```
Se o trecho não for encontrado, a URL do projeto será exibida no terminal.

> Este trecho pode ser modificado de acordo com a necessidade da busca, bem como o arquivo a ser buscado.
---

## Requisitos
- Go 1.20 ou superior
- Token de acesso pessoal (PAT) do GitLab

---

## Configuração
1. Clone o repositório:

```bash
git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>
```

2. Crie um arquivo `.env` na raiz do projeto e insira o token PAT:

```env
GITLAB_TOKEN="seu_token_aqui"
```

---

## Instalação
1. Baixe as dependências:

```bash
go mod tidy
```

2. Compile o projeto:

```bash
go build -o app
```

---

## Execução
Para executar a aplicação, rode o comando:

```bash
./app
```
ou ou pode executar direto com o run.

```bash
go run main.go
```
---

## Bibliotecas utilizadas
- [github.com/joho/godotenv](https://github.com/joho/godotenv) — Para leitura do arquivo `.env`

---

## Licença
Este projeto é disponibilizado sob a licença MIT.
