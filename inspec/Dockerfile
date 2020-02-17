FROM ruby:2.5

ENV AWS_REGION="us-east-1"
ENV AWS_ACCESS_KEY_ID=""
ENV AWS_SECRET_ACCESS_KEY=""
ENV CHEF_LICENSE="accept-silent"

ADD . /inspec

RUN wget https://raw.githubusercontent.com/inspec/inspec-aws/master/Gemfile

RUN bundle install

WORKDIR /inspec

CMD /bin/bash -c "inspec exec . -t aws://"