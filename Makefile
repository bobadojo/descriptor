all:	descriptor

descriptor:
	git submodule init
	git submodule update
	protoc `find apis/bobadojo -name "*.proto"` \
		--proto_path='apis' \
		--include_imports \
		--include_source_info \
		--descriptor_set_out=descriptor.pb
