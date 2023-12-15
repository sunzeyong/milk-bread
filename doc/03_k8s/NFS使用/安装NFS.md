# 服务端
## 安装服务端
sudo apt update
sudo apt install nfs-kernel-server
sudo systemctl status nfs-server

## 创建共享文件夹
sudo mkdir -p /mnt/nfs
sudo chown nobody:nogroup /mnt/nfs
sudo chmod -R 777 /mnt/nfs

## 允许访问的客户端，并使配置生效
sudo vim /etc/exports 添加下面信息
/mnt/nfs 10.211.55.0/24(rw,sync,no_root_squash,no_subtree_check)

sudo exportfs -a
sudo exportfs -s

# 客户端
## 安装客户端
sudo apt install nfs-common

## 将nfs端共享目录挂载到本地的/home/ubuntu/node-1/mnt/nfs
sudo mount -t nfs 10.211.55.5:/mnt/nfs /home/ubuntu/node-1/mnt/nfs