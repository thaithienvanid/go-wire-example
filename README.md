# go-wire-example

## Local Development

```sh
> ./scripts/bin start
```

## Local Deployment

```sh
> docker compose -f docker-compose.yml up -d
```

To open shell:

```sh
> docker compose -f docker-compose.yml run app sh
```

## Lint Code

```sh
> ./scripts/bin code_lint
```

## Format Code

```sh
> ./scripts/bin code_format
```
