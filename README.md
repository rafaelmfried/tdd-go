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

# Interface do Reader:

type Reader interface {
Read(p []byte) (n int, err error)
}

Vemos a interface ReadSeeker que junta a interface Read e a Seeker que te ajuda a ler e percorrer arquivos

type ReadSeeker interface {
Reader
Seeker
}

type Seeker interface {
Seek(offset int64, whence int) (int64, error)
}

# Temos agora uma nova funcionalidade que sera a contrucao do seguinte fluxo:

> Entrada do numero de jogadores
> A entrada sera usada para calcular o intervalo entre as dobras de bets ->
> {
> n = (playerQtd \* 1) + 5
> }
> n sera representado em minutos e quando o tempo chegar a bet dobra comecando em 100 fichas
> O sistema deve avisar a cada iteracao

# html/template sera usada para formar-mos templates html

criaremos a rota /jogo que deve retornar um template html por meio de websocket
