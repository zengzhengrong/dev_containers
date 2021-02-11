### 开启5个节点的
1个 master
3个 work节点
1个 测试节点

### 将 master 节点 的ssh key 加入到其他节点 以用来免密登陆

```
ssh-keygen -t rsa -b 4096 -N '' -f ~/.ssh/id_rsa
```

### 将key 导入到各个节点
```
ssh-copy-id node1-container
ssh-copy-id node2-container
ssh-copy-id node3-container
ssh-copy-id node-test-container
```

### 各个节点安装docker cli
```
ansible-playbook -i .devcontainer/hosts .devcontainer/prepare-docker-cli-nodes.yml --user root
```