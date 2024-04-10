# Desafio de 1bi de linhas python

## Descrição

Este script em Python foi desenvolvido para processar dados de um arquivo contendo medições. Ele utiliza multiprocessamento para dividir o arquivo em pedaços e processá-los em paralelo, melhorando assim o desempenho.

## Uso

Para executar o script, use o seguinte comando:

```
python3 calculate.py
```

## Como Funciona

1. **Divisão em Pedados do Arquivo**:
   - A função `get_file_chunks` divide o arquivo de entrada em pedaços menores, considerando o número de núcleos de CPU disponíveis. Ela calcula o tamanho do pedaço dividindo o tamanho do arquivo pelo número de núcleos da CPU.
   - Os pedaços são identificados procurando por caracteres de nova linha dentro do arquivo, garantindo que cada pedaço termine em uma linha completa.

2. **Processamento dos Pedaços**:
   - Cada pedaço do arquivo é processado em paralelo usando multiprocessamento. A função `_process_file_chunk` lê o pedaço, calcula o mínimo, máximo, soma e contagem das medições para cada localização.
   - Os resultados de cada pedaço são armazenados em um dicionário.

3. **Combinação dos Resultados**:
   - Após processar todos os pedaços, o script combina os resultados de cada pedaço. Ele agrega as medições para cada localização, considerando mínimo, máximo, soma e contagem.
   - Por fim, imprime os resultados agregados para cada localização no formato especificado.

## Otimização de Desempenho

- **Multiprocessamento**: Ao utilizar o módulo `multiprocessing`, o script aproveita múltiplos núcleos de CPU, melhorando significativamente a velocidade de processamento.
- **Divisão em Pedaços**: Dividir o arquivo em pedaços menores permite o processamento em paralelo, reduzindo o tempo total de processamento.
- **Manipulação Eficiente de Dados**: O script lida eficientemente com os dados processando cada pedaço separadamente e combinando os resultados posteriormente, reduzindo o consumo de memória.

## Tempo de Execução

O tempo de execução do script foi aproximadamente 8 minutos para o processamento completo dos dados. Esse tempo pode variar dependendo do tamanho do arquivo e da capacidade do hardware utilizado.