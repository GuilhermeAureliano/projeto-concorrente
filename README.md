# Projeto de Programação Concorrente: Bakery vs Atomic Bakery

Este projeto compara o desempenho entre duas implementações de uma padaria concorrente: Bakery e Atomic Bakery. Abaixo estão as instruções para executar o código e gerar os gráficos correspondentes.

## Como executar o código?

### 1° Passo: Clonar o repositório

```bash
git clone https://github.com/GuilhermeAureliano/projeto-concorrente.git
```

### 2° Passo: Entrar no repositório do projeto

```bash
cd projeto-concorrente/
```

### 3° Passo: Rodar o script shell que executa o bakery.go e cria o seu respectivo csv com os resultados usados no gráfico

```bash
chmod +x create_bakery.sh
```
```
./create_bakery.sh
```

### 4° Passo: Rodar o script shell que executa o atomic.go e cria o seu respectivo csv com os resultados usados no gráfico

```bash
chmod +x create_atomic.sh
```
```
./create_atomic.sh
```

## Criação dos gráficos

### 1° Passo: Baixar bibliotecas necessárias para manipulação e visualização dos dados
```
pip install pandas matplotlib numpy
```

### 2° Passo: Criar o gráfico de linhas

```bash
python3 linhas.py
```

### 3° Passo: Criar o gráfico de barras

```bash
python3 barras.py
```

Certifique-se de ter o Python3 instalado no seu sistema antes de executar os comandos acima.

## Observação
Os scripts em shell realizam um loop que executam nossos programas em Go várias vezes com diferentes valores de `threads (N)`, variando de 2 a 2048, na base de 2.

Dependendo da máquina que está sendo executada, esse código shell pode demorar muito, então se for o seu caso, recomendamos alterar os arquivos shell para diminuir a quantidade de threads para um valor menor.

### Diminuir as threads do create_bakery.sh
```
#!/bin/bash

# Cria o arquivo CSV e escreve o cabeçalho
echo "N;Tempo de Execução" > bakery.csv

for ((N=2; N<=256; N*=2)); do
    echo "Running with N=$N"
    
    # Executa o programa Go e mede o tempo de execução
    duration=$( (time go run bakery.go -N $N) 2>&1 | grep real | awk '{print $2}' )

    # Adiciona a linha ao arquivo CSV
    echo "$N;$duration" >> bakery.csv
done
```

### Diminuir as threads do create_atomic.sh
```
#!/bin/bash

# Cria o arquivo CSV e escreve o cabeçalho
echo "N;Tempo de Execução" > atomic.csv

for ((N=2; N<=256; N*=2)); do
    echo "Running with N=$N"
    
    # Executa o programa Go e mede o tempo de execução
    duration=$( (time go run atomic.go -N $N) 2>&1 | grep real | awk '{print $2}' )

    # Adiciona a linha ao arquivo CSV
    echo "$N;$duration" >> atomic.csv
done
```

Agora pode tentar executar novamente o código. Caso persista a demora, diminua para um valor menor ainda.
