import pandas as pd
import matplotlib.pyplot as plt

# Ler os arquivos csv
df1 = pd.read_csv('bakery.csv', delimiter=';', decimal=',')
df2 = pd.read_csv('atomic.csv', delimiter=';', decimal=',')

# Converter a coluna 'Tempo de Execução' para segundos
df1['Tempo de Execução'] = df1['Tempo de Execução'].apply(lambda x: int(x.split('m')[0])*60 + float(x.split('m')[1].replace('s','').replace(',', '.')))
df2['Tempo de Execução'] = df2['Tempo de Execução'].apply(lambda x: int(x.split('m')[0])*60 + float(x.split('m')[1].replace('s','').replace(',', '.')))

plt.figure(figsize=(10,6))

# Plotar os dados
plt.plot(df1['N'], df1['Tempo de Execução'], marker='o', color='b', label='bakery.csv')
plt.plot(df2['N'], df2['Tempo de Execução'], marker='o', color='r', label='atomic.csv')

plt.title('Gráfico do Tempo de Execução')
plt.xlabel('N')
plt.ylabel('Tempo de Execução (s)')
plt.grid(True)

# Adicionando anotações para cada ponto
for i in range(len(df1['N'])):
    plt.annotate(df1['N'][i], (df1['N'][i], df1['Tempo de Execução'][i]))
for i in range(len(df2['N'])):
    plt.annotate(df2['N'][i], (df2['N'][i], df2['Tempo de Execução'][i]))

plt.legend()
plt.show()
