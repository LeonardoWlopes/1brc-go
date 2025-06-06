# 1brc-go

## One Billion Row Challenge (1BRC)

O One Billion Row Challenge é um desafio de processamento de dados que consiste em processar um arquivo de texto contendo 1 bilhão de linhas de medições de temperatura.

### Regras do Desafio

1. O arquivo de entrada contém 1 bilhão de linhas no formato:
   ```
   <nome da estação>;<temperatura>
   ```
   Exemplo: `Hamburg;12.0`

2. A temperatura é um número decimal com uma casa decimal, variando de -99.9 a 99.9

3. O nome da estação é uma string UTF-8 com comprimento variável

4. O objetivo é calcular as estatísticas (min, max, média) para cada estação

5. A saída deve ser formatada como:
   ```
   {nome da estação}={min}/{média}/{max}
   ```
   Exemplo: `Hamburg=12.0/15.3/18.0`

6. As estações devem ser ordenadas alfabeticamente na saída

7. O programa deve ser executado em um único thread

8. O tempo de execução é medido do início ao fim do programa

9. O arquivo de entrada é fornecido como argumento de linha de comando

10. A saída deve ser escrita em stdout

### Objetivo

O objetivo principal é processar o arquivo o mais rápido possível, mantendo a precisão dos cálculos e seguindo todas as regras especificadas.

### Checklist de Implementação

#### Estrutura Básica
- [ ] Criar estrutura do projeto Go
- [ ] Implementar leitura do arquivo de entrada
- [ ] Implementar parser para cada linha
- [ ] Criar estrutura de dados para armazenar estatísticas

#### Processamento de Dados
- [ ] Implementar extração do nome da estação
- [ ] Implementar conversão da temperatura
- [ ] Implementar cálculo de min/max/média
- [ ] Implementar ordenação alfabética das estações

#### Otimizações
- [ ] Otimizar alocação de memória
- [ ] Otimizar parsing de strings
- [ ] Otimizar cálculos matemáticos
- [ ] Otimizar estrutura de dados para busca

#### Validações
- [ ] Validar formato do arquivo de entrada
- [ ] Validar range de temperaturas
- [ ] Validar formato UTF-8 dos nomes
- [ ] Validar precisão dos cálculos

#### Testes
- [ ] Criar testes unitários
- [ ] Criar testes de performance
- [ ] Testar com arquivo pequeno
- [ ] Testar com arquivo grande
- [ ] Validar tempo de execução

#### Documentação
- [ ] Documentar decisões de design
- [ ] Documentar otimizações realizadas
- [ ] Documentar resultados de performance
- [ ] Atualizar README com instruções de uso

### Implementação

Esta implementação utiliza Go para resolver o desafio, focando em otimizações de performance e uso eficiente de memória.
