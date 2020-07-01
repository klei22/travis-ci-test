#!/bin/bash

sudo apt install -y ruby ruby-dev

gem install travis --user-install
gem install bundler --user-install

echo "make sure to add the right directory for gem to your PATH"
