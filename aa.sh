#!/bin/bash

# Cores ANSI
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' 


display_flag() {
    echo -e "${GREEN}"
    cat << "EOF"
__      __       .__                                  __                               .___      .__
/  \    /  \ ____ |  |   ____  ____   _____   ____   _/  |_  ____     _____ _____     __| _/______|__|_____   ____   ___________
\   \/\/   // __ \|  | _/ ___\/  _ \ /     \_/ __ \  \   __\/  _ \   /     \__  \   / __ |\_  __ \  \____ \ /  _ \ /  _ \_  __ \
 \         /\  ___/|  |_\  \__(  <_> )  Y Y  \  ___/  |  | (  <_> ) |  Y Y  \/ __ \_/ /_/ | |  | \/  |  |_> >  <_> |  <_> )  | \/
  \__/\  /  \___  >____/\___  >____/|__|_|  /\___  >  |__|  \____/  |__|_|  (____  /\____ | |__|  |__|   __/ \____/ \____/|__|
       \/       \/          \/            \/     \/                       \/     \/      \/          |__|
EOF
    echo -e "${NC}"
    echo -e "${YELLOW}Bem-vindo ao Madripoor!${NC}"
    echo -e "${YELLOW}Script de instalação de ferramentas!${NC}"
}

success_message() {
    echo -e "${GREEN} =======$1 Instalado com sucesso=======${NC}"
}

error_message() {
    echo -e "${RED}=======Erro ao instalar $1=======${NC}"
    exit 1
}


display_flag

# Atualiza a lista de pacotes
apt update

# Função para instalar pacote e exibir mensagem de sucesso ou erro
install_package() {
    package_name="$1"
    echo -e "${YELLOW}======= Instalando $package_name =======${NC}"
    apt install -y "$package_name"
    if [ $? -eq 0 ]; then
        success_message "$package_name"
    else
        error_message "$package_name"
    fi
}


install_vscode(){
    echo -e "${YELLOW}======= Instalando o vscode =======${NC}"
    apt --fix-broken install -y
    wget "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64" -O vscode.deb
    apt-get -f install -y
    dpkg -i vscode.deb
}

install_awscli(){
    echo -e "${YELLOW}======= Instalando o AWS CLI =======${NC}" 
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    ./aws/install
    if [ $? -eq 0 ]; then
        success_message "AWS CLI"
    else
        error_message "AWS CLI"
    fi
}

install_kubectl(){
    echo -e "${YELLOW}======= Instalando o kubectl =======${NC}"
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
    if [ $? -eq 0 ]; then
        success_message "kubectl"
    else
        error_message "kubectl"
    fi
}

install_ansible(){
    echo -e "${YELLOW}======= Instalando o Ansible =======${NC}"
    python3 -m pip install --user ansible
}

# Instalação dos pacotes default
install_package "vim"
install_package "firefox"
install_package "curl"
install_package "unzip"
install_package "wget"
install_package "git"
install_package "python3"
install_package "python3-pip"
install_package "lsb-release"

install_ansible
install_vscode
install_awscli
install_kubectl


# Instala o terraform
echo -e "${YELLOW}======= Instalando Terraform =======${NC}"
apt-get clean all
apt --fix-broken install -y
apt-get update && apt-get install -y gnupg software-properties-common
wget -O- https://apt.releases.hashicorp.com/gpg | \
gpg --dearmor | \
tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
gpg --no-default-keyring \
--keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg \
--fingerprint
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
tee /etc/apt/sources.list.d/hashicorp.list
apt update
install_package "terraform" 

# Fim da instalação
echo -e "${GREEN}Instalação de softwares concluída.${NC}"