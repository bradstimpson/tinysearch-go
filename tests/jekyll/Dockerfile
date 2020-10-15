FROM ruby:2.7-alpine

RUN apk add --no-cache build-base gcc bash cmake git

# install both bundler 1.x and 2.x
RUN gem install bundler -v "~>1.0" && gem install bundler jekyll

WORKDIR /srv/jekyll/

RUN bundle init

RUN jekyll new test

WORKDIR /srv/jekyll/test/

RUN gem install public_suffix -v '4.0.6' --source 'https://rubygems.org/'

RUN bundle install && jekyll build

ADD ./test/_posts/* /srv/jekyll/test/_posts/

EXPOSE 8890

CMD [ "bundle", "exec", "jekyll", "serve", "--watch", "--incremental", "--force_polling", "-H", "0.0.0.0", "-P", "8890" ]