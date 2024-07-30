-- ************************* --
-- *** Informacoes Uteis *** --
-- ************************* --

-- 1 - Instalação Golang

-- Site golang - Download aplicacoes
Link: https://golang.org
ou
Link: https://go.dev/

-- Download aplicação Golang
Link: https://go.dev/dl/

-- Configurar variável de ambiente Windows
Caso não tenha sido adicionada durante a instalacao, acrescentar a variavel ao path da conta do usuario:

%USERPROFILE%\go\bin
ou
C:\Program Files\Go\bin

-- Caminho workspace GO
GOPATH=C:\Users\Will Oliveira\go

-- *********************** -- 
-- *** Comandos Golang *** --
-- *********************** -- 

-- 2 - Comando Golang

-- versao GO
comando: go version

-- Variaveis de ambiente
comando: go env

-- Executando um arquivo GO
comando go run <nome do arquivo>
exemplo: go run primeiro.go

-- Criando pacote modulo
comando: go mod init <nome do modulo>
exemplo: go mod init module

-- Buildando o projeto
comando: go build
(irá criar um arquivo binário com o código compilado do projeto)

-- Comando de instalação
comando: go install
(irá gerar o arqivo main para a raiz do projeto e não para seua pacotes)