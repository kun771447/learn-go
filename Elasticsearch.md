
## docker 安装 elasticsearch

https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html

https://www.elastic.co/guide/en/elasticsearch/reference/7.17/docker.html
```
docker run -d -p 9200:9200 -p 9300:9300 --name elasticsearch docker.elastic.co/elasticsearch/elasticsearch:7.10.2
```

这个命令会在后台运行 Elasticsearch 容器，并将容器的 9200 和 9300 端口映射到主机的相应端口。--name elasticsearch 为容器指定了一个名字（在这里是 "elasticsearch"），以方便后续管理。



```shell
# 新建 es 的 config 配置文件夹
mkdir -p /elasticsearch/config
# 新建 es 的 data 目录
mkdir -p /elasticsearch/data

# 设置所有用户的读写权限
chmod 777 -R /Users/didi/Desktop/elastic-search

# 写入配置到配置文件中 elasticsearch.yml 中
echo 'http.host: "0.0.0.0"' >> /elasticsearch/config/elasticsearch.yml
```


```shell
# 安装 es
sudo docker run --name elasticsearch -p 9200:9200 -p 9300:9300 \
    -e "discovery.type=single-node" \
    -e "ES_JAVA_OPTS=-Xms128m -Xmx256m" \
    -v /Users/didi/Desktop/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
    -v /Users/didi/Desktop/elasticsearch/data:/usr/share/elasticsearch/data \
    -v /Users/didi/Desktop/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
    -d elasticsearch:7.17.17
```

```shell
 -v 本机路径:Docker路径
```

## docker 安装 kibana


```shell
docker run -d --name kibana -e ELASTICSEARCH_HOSTS="http://192.168.51.14:9200" -p 5601:5601 kibana:7.17.17
```


```shell
# 获取本机 ip
ipconfig getifaddr en0
```

```shell
docker run -d --name kibana -e ELASTICSEARCH_HOSTS="http://本机ip:9200" -p 5601:5601 kibana:7.17.17
```