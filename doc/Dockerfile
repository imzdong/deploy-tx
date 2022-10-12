FROM nginx
RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
# docker build [选项] <上下文路径/URL/->
# docker build -t nginx:v3 . 构建dockerfile

# shell 格式：RUN <命令>
# exec 格式：RUN ["可执行文件", "参数1", "参数2"]，这更像是函数调用中的格式

# scratch 空白镜像

# 错误做法 start
# FROM debian:stretch
# 
# RUN apt-get update
# RUN apt-get install -y gcc libc6-dev make wget
# RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz"
# RUN mkdir -p /usr/src/redis
# RUN tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1
# RUN make -C /usr/src/redis
# RUN make -C /usr/src/redis install
# 错误做法 end

# Dockerfile 中每一个指令都会建立一层 曾经是最大不得超过 42 层，现在是不得超过 127 层。

# 正确做法 start
# FROM debian:stretch
# 
# RUN set -x; buildDeps='gcc libc6-dev make wget' \
#     && apt-get update \
#     && apt-get install -y $buildDeps \
#     && wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
#     && mkdir -p /usr/src/redis \
#     && tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1 \
#     && make -C /usr/src/redis \
#     && make -C /usr/src/redis install \
#     && rm -rf /var/lib/apt/lists/* \
#     && rm redis.tar.gz \
#     && rm -r /usr/src/redis \
#     && apt-get purge -y --auto-remove $buildDeps
	
# 正确做法 end	