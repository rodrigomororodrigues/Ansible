# Serasa Challenge

Repositório para armazenar o Laboratório Desafio de criação Página em GO

# Dependências
Para a criação do laboratório é necessário ter pré instalado os seguintes softwares:

Git
VirtualBox
Vagrant

Para o máquinas com Windows aconselhamos, se possível, que as instalações sejam feitas pelo gerenciador de pacotes Cygwin.

Para as máquinas com MAC OS aconselhamos, se possível, que as instalações sejam feitas pelo gerenciador de pacotes brew.

# Laboratório

Para criar o laboratório é necessário fazer o git clone desse repositório e, dentro da pasta baixada realizar a execução do vagrant up, conforme abaixo:

git clone https://github.com/rodrigomororodrigues/Ansible

cd Ansible/

vagrant up

User:rodrigomoro 
pass:/
RUN apk --no-cache add gcc g++ make git

- ./certs/:/etc/nginx/ssl
- ./nginx/nginx.conf:/etc/nginx/conf.d

 ssl_certificate /etc/nginx/conf.d/localhost.crt;
  ssl_certificate_key /etc/nginx/conf.d/localhost.key;

docker-compose -f  nginx-compose.yaml down  --volumes
docker-compose  -f go-compose.yaml down