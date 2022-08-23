# gouache

Relaxation project. A micro-service system.

## Components

- `client`: Vue 3 frontend written with TypeScript and built on Vite with Quasar. Integration tests with Cypress, unit tests with Vitest.
- `resource`: Golang REST API built atop packages I've built i.e. corset (CORS spec-compliant middleware) and turnpike (trie-based HTTP multiplexer). Unit tests with Go testing standard library.
- `session`: Java REST API built on Spring Boot w/ Redis for session caching. Testing with JUnit 5 and Mockito.
- `db`: AWS DynamoDB
