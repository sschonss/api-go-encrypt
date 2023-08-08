# API Encrypt

## Instalação

Clone o repositório e execute o seguinte comando:

```bash
$ go mod tidy
```

## Execução

```bash
$ go run main.go
```

## API

### Testando com o Postman:

O Postman é uma ferramenta popular para testar APIs e fazer solicitações HTTP. Siga estas etapas para testar sua API com o Postman:

- Abra o Postman.

- Crie uma nova solicitação:
- Selecione o método POST.
- Insira a URL completa da sua API, por exemplo: http://localhost:8080/encrypt.

- Na seção "Body", selecione "raw" e escolha "JSON (application/json)" como tipo de dado.

- Cole o seguinte exemplo de payload JSON no corpo da solicitação:

```json
{
  "password": "mysecretpassword",
  "encryption_method": "sha256"
}
```

- Clique em "Send" para enviar a solicitação.

### Testando com o cURL:

O cURL é uma ferramenta de linha de comando que permite fazer solicitações HTTP. Para testar sua API com o cURL, siga estas etapas:

- Abra o terminal.

- Execute o seguinte comando:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"password": "mysecretpassword", "encryption_method": "sha256"}' http://localhost:8080/encrypt
```
Isso enviará uma solicitação POST com o payload JSON especificado para a URL da sua API.

- ocê verá a resposta no terminal, que conterá o hash da senha criptografada.

### Observação:

Lembre-se de substituir `mysecretpassword` pelo valor da senha que você deseja criptografar e `sha256` pelo método de criptografia que deseja usar, como por exemplo o `bcrypt`.

Com essas instruções, você poderá testar sua API tanto com o Postman quanto com o cURL. Certifique-se de que o servidor da sua aplicação esteja em execução durante os testes.
