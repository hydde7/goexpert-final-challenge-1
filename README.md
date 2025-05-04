# Rate Limiter

Este projeto implementa um sistema de **rate limiting** flexível e configurável em Go, utilizando o framework [Gin](https://github.com/gin-gonic/gin). Seu propósito é controlar o número de requisições que um cliente pode realizar dentro de um intervalo de tempo, bloqueando temporariamente acessos excessivos.

---

## Funcionalidade

A principal funcionalidade é limitar requisições com base em um identificador (IP ou token) e, ao exceder esse limite dentro de uma janela de tempo, bloquear temporariamente o cliente.

---

## Modos de Bloqueio

O sistema suporta dois modos configuráveis de bloqueio:

- **Por IP** (`IPBlockMode = true`): o endereço IP do cliente é usado como chave para controle do limite.
- **Por Token** (`IPBlockMode = false`): a chave usada é o valor do cabeçalho `API_KEY` enviado na requisição.

---

## Configuração

As configurações são feitas por variáveis de ambiente, com valores padrão seguros. Alguns dos parâmetros configuráveis incluem:

- Porta do servidor (`APP_PORT`)
- Ativação do Redis (`USE_REDIS`)
- Limites por IP ou token (`IP_LIMIT`, `TOKEN_LIMIT`)
- Duração da janela de rate limit (`RATE_LIMIT_WINDOW`)
- Duração do bloqueio após exceder o limite (`BLOCK_DURATION_IP`, `BLOCK_DURATION_TOKEN`)
- Endereço e autenticação do Redis (`REDIS_ADDR`, `REDIS_PASSWORD`, `REDIS_DB`)

---

## Como funciona

1. O middleware identifica o cliente por IP ou token.
2. Verifica se o número de requisições dentro da janela permitida foi excedido.
3. Se não excedeu, permite a requisição normalmente.
4. Se excedeu, bloqueia novas requisições daquele cliente por um tempo determinado.

A contagem pode ser armazenada em memória ou no Redis, dependendo da configuração. Com Redis ativado, a aplicação é capaz de operar em ambientes distribuídos com consistência.

---

## Testes

Existe um teste automatizado para cada mecanismo de armazenamento (`InMemoryStore`, `RedisStore`). Esses testes garantem que o limite é respeitado, que o bloqueio ocorre após exceder o número permitido de requisições e que o desbloqueio acontece corretamente após o tempo definido.

---

## 🐳 Executando com Docker

O projeto está preparado para rodar com Docker e Docker Compose. A aplicação utiliza uma imagem com suporte a hot reload (`cosmtrek/air`) e pode ser iniciada com:

```bash
docker-compose up --build
```

O ambiente inicia dois serviços:
- **Redis** (para armazenamento persistente de contagens e bloqueios)
- **App** (servidor Go com suporte a rate limiting)

A aplicação ficará disponível em: `http://localhost:8080/ping`

---

## Respostas esperadas

- Se a requisição for permitida, a resposta será:
  ```json
  { "message": "pong" }
  ```

- Se o limite for excedido, a resposta será:
  ```json
  { "error": "you have reached the maximum number of requests or actions allowed within a certain time frame" }
  ```

- Se a requisição exigir token e o usuário não possuir, a resposta será:
  ```json
  { "error": "API_KEY header is required" }
  ```

---
