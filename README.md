## FTerceiraIdade 
&nbsp;
### Description

Estudo de backend em GO, que tem como proposta gerenciar os dados referentes aos usuários do sistema (professores e estudantes), aos cursos e turmas a eles relacionados. Por fim, gerir questões individualmente definidas que poderão ser utilizadas em um processo de avaliações personalizáveis. Uma API opcional em suporte ao frontend FTerceiraIdade - https://github.com/jobson-almeida/fterceiraidade-frontend.


&nbsp;

### Linguagens e Ferramentas


<p>  
    <a
    href="https://developer.mozilla.org/en-US/docs/Web/JavaScript"
    target="_blank"
    rel="noreferrer"
  >
    <img    
      src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg"
      alt="javascript"
      title="javascript"
      width="40"
      height="40"
    />
  </a>
  &nbsp;
  <a href="https://www.docker.com/" target="_blank" rel="noreferrer">
    <img src="https://cdn.worldvectorlogo.com/logos/docker-4.svg" alt="docker" title="docker" width="40" height="40" />
  </a>
  &nbsp;
  <a href="https://www.postgresql.org/" target="_blank" rel="noreferrer">
    <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-plain.svg" alt="postgres" title="postgres" width="40" height="40" />
  </a>
</p>
 
&nbsp;

### Instalar dependências

```dosini
go mod tidy
```
&nbsp;

### Criar e executar o container do banco de dados

```dosini
docker compose up
```
&nbsp;

### Opcional: criar ou configurar o arquivo .env.local do frontend (FTerceiraIdade) com a nova url base

```dosini
API_URL_BASE=http://localhost:8888
```
&nbsp;

### Executar os endpoints da aplicação
Utilize o arquivo test.http (Rest Client) presente na raiz do projeto ou outra plataforma de sua preferência. Endpoints de exemplos estão contidos no arquivo.


&nbsp;
&nbsp;

---

developed by Jobson Almeida