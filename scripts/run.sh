#!/bin/sh

cd /home/ec2-user/

export PATH=$PATH:/usr/local/go/bin
cd /home/ec2-user/book-freelancer/
scripts/build.sh
build/gql-server &

cd /home/ec2-user/book-freelancer-frontend/
npm install
npm start