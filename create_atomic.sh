#!/bin/bash

# Cria o arquivo CSV e escreve o cabeçalho
echo "N;Tempo de Execução" > atomic.csv

for ((N=2; N<=2048; N*=2)); do
    echo "Running with N=$N"
    
    # Executa o programa Go e mede o tempo de execução
    duration=$( (time go run atomic.go -N $N) 2>&1 | grep real | awk '{print $2}' )

    # Adiciona a linha ao arquivo CSV
    echo "$N;$duration" >> atomic.csv
done
