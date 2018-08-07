VERSION = $(shell gobump show -r)

.PHONY: build
build:
	gox \
		-osarch="darwin/386" \
		-osarch="darwin/amd64" \
		-osarch="linux/386" \
		-osarch="linux/amd64" \
		-osarch="windows/386" \
		-osarch="windows/amd64"
	zip medicine_v$(VERSION)_darwin_386.zip medicine_darwin_386
	zip medicine_v$(VERSION)_darwin_amd64.zip medicine_darwin_amd64
	tar -czvf medicine_v$(VERSION)_linux_386.tar.gz medicine_linux_386
	tar -czvf medicine_v$(VERSION)_linux_amd64.tar.gz medicine_linux_amd64
	zip medicine_v$(VERSION)_windows_386.zip medicine_windows_386.exe
	zip medicine_v$(VERSION)_windows_amd64.zip medicine_windows_amd64.exe
	rm -f \
		medicine_darwin_386 \
		medicine_darwin_amd64 \
		medicine_linux_386 \
		medicine_linux_amd64 \
		medicine_windows_386.exe \
		medicine_windows_amd64.exe
	mkdir out
	mv medicine_* out
	ghr v$(VERSION) out
	rm -rf out
