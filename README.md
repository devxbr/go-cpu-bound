# go-cpu-bound

**go-cpu-bound** é uma ferramenta simples escrita em Go para simular carga de CPU controlada, útil para testes de estresse e benchmarks.

## Funcionalidades

- Consome uma porcentagem especificada da CPU (0-100%).
- Suporta múltiplos núcleos de CPU.
- Permite configurar a duração do teste de estresse.

## Como Usar

1. Clone o repositório:
   ```bash
   git clone https://github.com/devxbr/go-cpu-bound.git
   cd go-cpu-bound

2. Compile o projeto
   ```bash
   go build -o go-cpu-bound

3. Execute o projeto
   ```bash
   ./go-cpu-bound -percent 50 -duration 60

