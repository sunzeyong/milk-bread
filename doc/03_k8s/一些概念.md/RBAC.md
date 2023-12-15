# 权限控制
1. Role: 一组权限的集合
2. RoleBinding: 绑定role和用户
3. Subject: 用户，可以是ServiceAccount，k8s内置用户，用于在pod内部获取pod信息。 SA会放在pod中，表示有哪些权限，若不明确会用系统默认sa，可能会有一些安全问题
   
Role和Subject都是有Namespace限制的，如果要在集群内获取权限，使用Cluster和ClusterRoleBinding。