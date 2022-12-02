# ansible
Created: 2022-12-02 06:07

## Setup
 - From docs [1], you need python3 and pip, after that, run:
```sh
python3 -m pip install --user ansible
```
 - Command completion:
```sh
python3 -m pip install --user argcomplete
```
 - Alternatively, if that not work, you can clone project and build it, or use ```
export PATH=$PATH:/path/to/ansible``` in the bash or install with root:
```sh
sudo pip3 install ansible
sudo pip3 install ansible --upgrade
```
 - Test Ansible using a ping to localhost:
```sh
echo 'localhost ansible_connection=local' >> hosts
ansible -i hosts all -m ping 
```

## References
1. [Installation](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
2. [FullCycle](https://plataforma.fullcycle.com.br/courses/184/168/131/conteudos?capitulo=131&conteudo=7331)