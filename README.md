### Agora que pegamos a base do TDD em go vamos partir para os testes em aplicacoes com:

Servidor HTTP
JSON, roteamento e aninhamento
IO e sorting
Linha de comando e estrutura de pacotes
Tempo
Websockets

### Application requirements

# GET /jogadores/{nome}

# POST /jogadores/{nome}

# Ao receber o post devemos icrementar uma vitoria para aquele jogador

# N pode obter jogador sem telo registrado

# E parece ser dificil ver se a chamada do POST funcionou sem o registro do metodo antes

Curiosidade legal: mocking significa zombar, mas usamos para definir classes com valores pre setados em testes

## Agora vamos criar o seguinte caso de teste:

# Quando o usuario n existir deve retornar um erro 404 de jogador not found com nossa mensagem

# Proximos passos fazer o PostgresArmazenamentoJogador implementando o Armazenamento jogador

# Conceitos

### Requisito:

Criar endpoint GET /liga que retorna todos os Jogadores

# Agora vamos usar arquivos como a base de dados para salvar
