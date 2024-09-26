#!/bin/bash

cd contracts && forge build

cd ../trx-contracts && tronbox compile