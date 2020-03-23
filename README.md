# Digital Account


### Instação e execução

```sh
$ cd digital-account
$ docker-compose up -d
```

Após a inicilização o serviço [transactions](https://github.com/dalmarcogd/digital-account/tree/master/transactions), [accounts](https://github.com/dalmarcogd/digital-account/tree/master/accounts), [transactions-persist](https://github.com/dalmarcogd/digital-account/tree/master/transactions-persist) e [accounts-persist](https://github.com/dalmarcogd/digital-account/tree/master/accounts-persist) estão executando. Assim como o [RabbitMQ](https://www.rabbitmq.com/), [Redis](https://redis.io/) e [PostgreSQL](https://www.postgresql.org/).

Os serviços [accounts](https://github.com/dalmarcogd/digital-account/tree/master/accounts) e [transactions](https://github.com/dalmarcogd/digital-account/tree/master/transactions) são APIs que disponibilizam métodos HTTP para criação de contas digitais dos usuários e criação de transações para as contas dos usuários. O [RabbitMQ](https://www.rabbitmq.com/) é utilizado como gerenciador da fila que persiste estes respectivos cadastros, o [Redis](https://redis.io/) é utilizado como camada de cache até que o regsitro seja persistido no banco de dados [PostgreSQL](https://www.postgresql.org/).

### Bibliotecas

Este projeto utiliza os seguintes bibliotecas.

| Lib | LINK |
| ------ | ------ |
| Echo | https://github.com/labstack/echo |
| Gorm | https://github.com/jinzhu/gorm |

[Download Collections Postman](https://github.com/dalmarcogd/digital-account/blob/master/Digital%20Account.postman_collection.json)
