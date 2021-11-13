***Pet project***

---

# Cortex

Монорепозиторий сервисов, разворачиваемых в докер-контейнерах, взаимодействующих меж собой.

## Services

Краткое описание внутренних сервисов

### Authentication

```bash
internal/auth
```

Код сервиса аутентификации. Из особенностей - работа с БД согласно парадигме [CQRS](https://en.wikipedia.org/wiki/CQRS)

### Frontend

```bash
web
```

Код SSR-сервера фронтенда, который уже работает на "внешний мир"

## Deployment

Для запуска необходимо/достаточно выполнять комманду докера на сборку контейнерров

### Production environment

```bash
  docker-compose up
```

### Develop environment

```bash
  docker-compose up --build
```

или

```bash
  docker-compose up --no-cache
```
