# Configuração de Ambiente

## Windows 

### Instalar o Cliente Nativo SSH (opcional)

Fonte: [Instalar o OpenSSH usando as Configurações do Windows](https://docs.microsoft.com/pt-br/windows-server/administration/openssh/openssh_install_firstuse)

1. Inicie o PowerShell em modo Administrador: clique com o botão direito no menu iniciar e depois selecione _Windows PowerShell (Admin)_. 
2. Execute o comando:  
```
PS C:\Users\Administrator> Add-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0
```
3. Execute apenas o comando ```ssh``` e veja se a saída abaixo aparece:
```
PS C:\Users\Administrator> ssh
usage: ssh [-46AaCfGgKkMNnqsTtVvXxYy] [-B bind_interface]
           [-b bind_address] [-c cipher_spec] [-D [bind_address:]port]
           [-E log_file] [-e escape_char] [-F configfile] [-I pkcs11]
           [-i identity_file] [-J [user@]host[:port]] [-L address]
           [-l login_name] [-m mac_spec] [-O ctl_cmd] [-o option] [-p port]
           [-Q query_option] [-R address] [-S ctl_path] [-W host:port]
           [-w local_tun[:remote_tun]] destination [command]
```

Nós não precisamos do servidor SSH, apenas do cliente. 

### Instalar o Git

1. Acesso [Git](https://git-scm.com/) e baixe a última versão. 
2. Inicie o instalador, aceite os termos de uso e finalize a instalação sem alterar nenhuma das opções padrão. São várias, deixe como estão e vá avançando.
3. Abra um terminal (_Prompt_ de Comando ou. _Powershell_) e execute (substituindo por seu _e-mail_ e usuário do _GitHub_:
```
git config --global user.mail "joao.marcelo@ufc.br"
git config --global user.name "joaomarceloalencar"
```

### Instalando a Linguagem Go

1. Siga os passos acima para instalar o Git.
2. Vá em [Golang](http://golang.org) e baixe a última versão para o Windows.
3. Siga os passos da instalação sem alterar nada . Somente anote o local da instalação (em geral, C:\Program Files\Go ou algo parecido).
4. Verifique que a instalação ocorreu com sucesso executando o comando ```go version``` no terminal do Windows (_Prompt_ de Comandos ou _PowerShell_).

### Instalar o Visual Studio Code

1. Acesse [Visual Studio Code](https://code.visualstudio.com/) e baixe a última versão. 
2. Inicie o instalador, aceite os termos de uso e finalize a instalação sem alterar nenhuma das opções padrão.
3. Com o _Visual Studio Code_ aberto, vá em _Controle de Código Fonte_. É o botão do meio na barra do lado esquerda. Você também pode usar o atalho 
_Ctrl + Shift + G_.
4. Deixe o _Visual Studo Code_ aberto, vá na tela principal do seu repositório no _GitHub_. Logo acima do código, no canto direito, tem um botão verde, 
chamado _Code_ ou Código. Clique nele e anote a URL que aparece quando você seleciona HTTPS.
6. Retorne ao _Visual Studio Code_, escolha _Clone Repository_ ou _Clonar Repositório_.
7. No formulário que aparece no meio da tela do _Visual Studio Code_, no topo, digite a URL do seu repositório (veja no _GitHub_) e aperte Enter. Escolha uma pasta local para guardar sua cópia do repositório. Pode ser qualquer uma que você achar melhor. Você também escolher _Clone from GitHub_.
8. Irá surgir uma tela para fazer a autenticação no _GitHub_. Escolha a opção _Sign in with your browser_. Irá abrir uma nova janela, ou uma aba, do seu navegador.
9. Coloque seu _login_ e senha do _GitHub_ e clique em _Authorize GitCredentialManager_.
10. Volte ao _Visual Studio Code_ e no canto inferior esquerdo clique para abrir o repositório clonado. Aceite para confiar nos autores.
11. Pronto, no lado direito deve aparece o conteúdo do seu repositório.

### Criando um Projeto Go 

1. Vamos seguir o [Tutorial](https://golang.org/doc/tutorial/getting-started)
2. Abra o _Visual Studio Code_ no repositório da disciplina.
3. Crie um diretório chamado _atividade08_ na pasta de _atividades_.
4. No meu do _Visual Studio Code_, selecione _Terminal->New Terminal_.
5. No Terminal, navegue até a pasta criada.
6. Crie e entre em um diretório chamado _hello_.
```
mkdir hello
cd hello
```
7. Inicialize o módulo do projeto Go. Essa etapa facilita a importação do seu módulo. À princípio não seria necessário, mas vamos criar esse hábito.
```
go mod init github.com/joaomarceloalencar/joaomarcelolip20212/atividades/atividade08/hello
```
8. Na mesma pasta, crie um arquivo chamado _hello.go_. O _Visual Studio Code_ pode oferecer a instalação de extensões do Go. Clique em _Install_ e aguarde. Depois, aceite _Install All_. Se der algum problema na instalação das extensões, não se preocupe, podemos continuar sem elas.
9. Coloque o seguinte conteúdo em _hello.go_:
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
10. No terminal, digite:
```
go run hello.go
```
11. A saída deve imprimir a mensagem. "Hello, Word!"
12. Agora, selecione o controle de versão no _Visual Studio Code_, é o botão do meio na barra do lado esquerda. Você também pode usar o atalho 
_Ctrl + Shift + G_.
13. Logo abaixo de _Changes_, clique no símbolo + para cada arquivo. Eles devem ir para _Staged Changes_.
14. Na caixa de texto logo acima de _Staged Changes_, coloque "Resolução da Atividade 08" e aperte Ctrl+Enter.
15. Clique nos "..." e selecione _Push_. Pronto, atividade submetida. 
