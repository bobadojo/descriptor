#!/bin/sh

#grpcurl -plaintext -protoset descriptor.pb localhost:8080 bobadojo.stores.v1.Stores.ListStores

grpcurl -plaintext -protoset descriptor.pb localhost:443 bobadojo.stores.v1.Stores.ListStores
