# FTerCeiraIdade Backend

API REST em Go para apoiar o frontend FTerceiraIdade no gerenciamento de estudantes, professores, cursos, turmas, questões e avaliações. O projeto usa PostgreSQL para persistência e oferece um endpoint provisório para demonstrar gráficos no frontend.

## Funcionalidades

- Criar, consultar, atualizar e remover estudantes, professores, cursos, turmas, questões e avaliações.
- Consultar dados de demonstração para gráficos pelo endpoint `/chart_data/`.

### Limitação do endpoint de gráficos

O endpoint `GET /chart_data/` é um recurso provisório. Suas regras de negócio ainda não foram definidas e os valores são fixos, usados somente para demonstrar os gráficos no frontend. Para esse objetivo atual, o recurso cumpre seu propósito; os dados não devem ser considerados resultados de regras de negócio consolidadas.

## Tecnologias

- Go 1.27.4
- PostgreSQL 17
- Chi e GORM
- Docker Compose

## Pré-requisitos

- Go 1.27.4
- Docker com Docker Compose

## Configuração e execução

Na raiz do projeto, copie o arquivo de exemplo e preencha as configurações locais:

```sh
cp .env.example .env
```

Configure `POSTGRES_USER`, `POSTGRES_PASSWORD` e `POSTGRES_DB`. Atualize `CONNECTION` com os mesmos valores de usuário, senha e banco; `DIALECT` deve permanecer como `postgres`. O arquivo `.env` contém configurações locais e não deve ser versionado.

Baixe as dependências:

```sh
go mod download
```

Inicie o PostgreSQL:

```sh
docker compose up -d database
```

Quando o banco estiver pronto, inicie a API em outro terminal:

```sh
go run ./cmd
```

A API ficará disponível em `http://localhost:5000`. Para parar a API, use `Ctrl+C`; para parar o banco sem remover os dados persistidos, execute:

```sh
docker compose down
```

## API

Os recursos `/courses`, `/students`, `/teachers`, `/questions`, `/classrooms` e `/assessments` oferecem operações de consulta e criação na coleção, além de consulta, atualização e remoção por identificador.

Para consultar os dados de demonstração dos gráficos:

```sh
curl http://localhost:5000/chart_data/
```

## Frontend

O frontend relacionado está disponível em [jobson-almeida/fterceiraidade-frontend](https://github.com/jobson-almeida/fterceiraidade-frontend). Para apontá-lo para esta API local, configure no arquivo `.env.local` do frontend:

```env
API_URL_BASE=http://localhost:5000
```

## Créditos

Desenvolvido por Jobson Almeida.
