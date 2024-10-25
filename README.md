# BlackWebFilter

该项目是将开源情报网站提供的[maravento/blackweb](https://github.com/maravento/blackweb.git)中的网站域名黑名单作为域名黑名单库，
实现了匹配输入的url中域名是否在黑名单中的功能。

使用[Sensitive-lexicon](https://github.com/konsheng/Sensitive-lexicon.git)库提供的敏感词库，作为url敏感词检测的库，当然也可以使用自己定义的敏感词库。

## 使用方法

clone本项目编译二进制可执行文件，并且clone最新的webblack数据源文件和最新的敏感词库：

修改配置文件，`.env`。启动程序，或者可以自行添加守护执行，可参考service文件。

### 恶意域名检测

该项目将域名黑名单使用布隆过滤器加载到内存中，提供恶意域名检测。


```bash
git clone --depth=1 https://github.com/maravento/blackweb.git
tar zxvf blackweb.tar.gz
md5sum  blackweb.txt
```

### 敏感词检测

该项目将敏感词文件用ac自动机算法生成了trie树，提供敏感词检测。

```
git clone https://github.com/konsheng/Sensitive-lexicon.git
```

## API

访问`http://localhost:yourport/swagger/index.html`
