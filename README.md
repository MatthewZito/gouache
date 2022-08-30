# gouache

Relaxation project. A micro-service system.

## Components

- `auth`: Golang REST API built atop packages I've built i.e. corset (CORS spec-compliant middleware) and turnpike (trie-based HTTP multiplexer). Unit tests with Go testing standard library.
- `cache` Redis session cache.
- `client`: Vue 3 frontend written with TypeScript and built on Vite with Quasar. Integration tests with Cypress, unit tests with Vitest.
- `db`: AWS DynamoDB
- `resource`: Java REST API built on Spring Boot. Testing with JUnit 5 and Mockito.
- `reporting`: Python analytics API built on Flask. Testing with Python unittest.
