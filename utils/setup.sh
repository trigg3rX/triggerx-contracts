#!/bin/bash

source .env

yarn install

git submodule update --init --recursive