# 更新库并支持https下载
```
sudo apt-get update

sudo apt-get install \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  lsb-release
```

# 添加 Docker 的官方 GPG 密钥
// 这一步可能会出现访问超时 需要翻墙

```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```

# 设置稳定存储库

```
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

# 安装docker

```
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io
```

# 使用 systemd 来管理容器的 cgroup
```
sudo mkdir -p /etc/docker

// 镜像地址找自己的aliyun加速地址
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "registry-mirrors": ["https://wkgfzwct.mirror.aliyuncs.com"],
  "storage-driver": "overlay2"
}
EOF

```

# 重新启动 Docker 并在开机时启用
```
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker
```

# 建立docker组 并将当前用户添加
```
sudo groupadd docker

sudo usermod -aG docker $USER
重新登陆生效
```