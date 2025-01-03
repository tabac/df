.PHONY: build
build:
	@cd df-rs/df-rs-ffi && cargo build --release
	@cp df-rs/df-rs-ffi/target/release/libdf_rs_ffi.dylib df-go/lib/
	@cp df-rs/df-rs-ffi/df.h df-go/lib/
	@cd df-go &&  go build .

.PHONY: clean
clean:
	@cd df-rs/ && cargo clean
	@cd df-rs/df-rs-ffi && cargo clean
	@rm -f df-go/lib/* df-go/df
