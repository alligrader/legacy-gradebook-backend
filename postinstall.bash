#!/bin/bash

export GO15VENDOREXPERIMENT=1
export SHAMAN_ENV='DEVELOPMENT'

sudo chown -R vagrant /opt/gopath
go get github.com/Masterminds/glide
glide up
