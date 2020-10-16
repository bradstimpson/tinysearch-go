FROM alpine:3.12

ENV HUGO_VERSION 0.76.5
ENV HUGO_BINARY hugo_${HUGO_VERSION}_linux-64bit

# Install pygments (for syntax highlighting) and bash
RUN apk update && apk add py-pygments && apk add bash && apk add git && rm -rf /var/cache/apk/*

# Download and Install hugo
RUN mkdir /usr/local/hugo
ADD https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_BINARY}.tar.gz /usr/local/hugo/
RUN tar xzf /usr/local/hugo/${HUGO_BINARY}.tar.gz -C /usr/local/hugo/ \
	&& ln -s /usr/local/hugo/hugo /usr/local/bin/hugo \
	&& rm /usr/local/hugo/${HUGO_BINARY}.tar.gz

WORKDIR /srv/hugo/
RUN hugo new site test

RUN cd test && \
    git init && \
    git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke && \
    echo 'theme = "ananke"' >> config.toml

WORKDIR /srv/hugo/test/

ADD posts/* content/posts/
ADD list.tjson layouts/_default/list.json.json
ADD config.toml .

EXPOSE 8889
CMD [ "hugo", "server", "-D", "--bind", "0.0.0.0", "-p", "8889", "-e", "production", "--renderToDisk"]