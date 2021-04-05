#源镜像
FROM golang:latest
#设置工作目录
WORKDIR $GOPATH/src/datum_textbook
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/datum_textbook
#go构建可执行文件
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go build .
#暴露端口
EXPOSE 8080


#最终运行docker的命令
ENTRYPOINT  ["./datum_textbook"]
#docker build -t datum_textbook .
#打包完成上传到阿里云镜像仓库
#docker tag [ImageId]
#docker push registry.cn-beijing.aliyuncs.com/a233/datum

#上传完进行下载
#docker pull registry.cn-beijing.aliyuncs.com/a233/datum
#docker run --name datum_common01 --restart=always -p 41376:41376 -d registry.cn-beijing.aliyuncs.com/a233/datum
