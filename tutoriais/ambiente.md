# Configuração de Ambiente

## Windows 

### Instalar o Cliente Nativo SSH

Fonte: [Instalar o OpenSSH usando as Configurações do Windows](https://docs.microsoft.com/pt-br/windows-server/administration/openssh/openssh_install_firstuse)

1. Inicie o PowerShell em modo Administrador: clique com o botão direito no menu iniciar e depois selecione _Windows PowerShell (Admin)_. 
2. Execute o comando:  
```
PS C:\Users\Administrator> Add-WindowsCapability -Online -Name OpenSSH.Client
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

### Instalar o Visual Studio Code

1. Acesse [Visual Studio Code](https://code.visualstudio.com/) e baixe a última versão. 
2. Inicie o instalador, aceite os termos de uso e finalize a instalação sem alterar nenhuma das opções padrão.
3. Com o _Visual Studio Code_ aberto, vá em _Controle de Código Fonte_. É o botão do meio na barra do lado esquerda. Você também pode usar o atalho 
_Ctrl + Shift + G_.
4. Deixe o _Visual Studo Code_ aberto, vá na tela principal do seu repositório no _GitHub_. Logo acima do código, no canto direito, tem um botão verde, 
chamado _Code_ ou Código. Clique nele e anote a URL que aparece quando você seleciona HTTPS.
6. Retorne ao _Visual Studio Code_, escolha _Clone Repository_ ou _Clonar Repositório_.
7. No formulário que aparece no meio da tela do _Visual Studio Code_, no topo, digite a URL do seu repositório e aperte Enter. Escolha uma pasta local para guardar
sua cópia do repositório. Pode ser qualquer uma que você achar melhor.
8. Irá surgir uma tela para fazer a autenticação no _GitHub_. Escolha a opção _Sign in with your browser_. Irá abrir uma nova janela, ou uma aba, do seu navegador.
9. Coloque seu _login_ e senha do _GitHub_ e clique em _Authorize GitCredentialManager_.
10. Volte ao _Visual Studio Code_ e no canto inferior esquerdo clique para abrir o repositório clonado. Aceite para confiar nos autores.
11. Pronto, no lado direito deve aparece o conteúdo do seu repositório.

