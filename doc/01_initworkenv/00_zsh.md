# zsh
linux有很多shell，一般默认的是bash。zsh相比于bash有更强大的功能。默认的zsh配置麻烦，于是有个oh-my-zsh项目，可以帮助使用者更便捷使用zsh的各种插件。

## oh-my-zsh
repo: https://github.com/ohmyzsh/ohmyzsh
管理zsh配置文件

```
// 使用方式
git clone https://github.com/robbyrussell/oh-my-zsh.git ~/.oh-my-zsh

cp ~/.oh-my-zsh/templates/zshrc.zsh-template ~/.zshrc

chsh -s /bin/zsh
```

### oh-my-zsh 主题插件 PowerLevel10k
```
git clone https://github.com/romkatv/powerlevel10k.git $ZSH_CUSTOM/themes/powerlevel10k

ZSH_THEME="powerlevel10k/powerlevel10k"
```

p10k configure 进行配置


# 参考
1. https://zhuanlan.zhihu.com/p/58073103