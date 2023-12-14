# 介绍
Helm 是 Kubernetes 的包管理器

# 安装方式
```
curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
sudo apt-get install apt-transport-https --yes
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

# 基本概念
## Chart
Helm 所管理的包

## Release
Release 就是 chart 在 K8S 上部署后的实例

## Repository
存储 chart 的仓库

## Config
前面提到了 chart 是应用程序所必须的资源，当然我们实际部署的时候，可能就需要有些自定义的配置了。Config 便是用于完成此项功能的，在部署时候，会将 config 与 chart 进行合并，共同构成我们将部署的应用。