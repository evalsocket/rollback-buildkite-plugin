# Rollback Buildkite Plugin

Can be used to perform rollback in a pipeline

## Example

Rollback one pipeline

```yml
steps:
  - command: echo "It's a dummy test"
    label: Deploy to production
    env:
    - BUILDKITE_API_TOKEN: ""
    plugins:
      - evalsocket/rollback-buildkite-plugin#v0.0.1:
```

## Configuration

NA

## Developing

To run the tests:

```shell
docker-compose run --rm tests
```

## Contributing

1. Fork the repo
2. Make the changes
3. Run the tests
4. Commit and push your changes
5. Send a pull request
