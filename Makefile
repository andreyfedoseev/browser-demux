BUILD_DIR = ./build
DIST_DIR = ./dist
APP_IMAGE_DIR = $(BUILD_DIR)/AppImage
OS = linux
ARCH = amd64

.PHONY: setup
setup:
	cd ./browser-demux && go get

.PHONY: build
build: setup
	cd ./browser-demux && env GOOS=$(OS) GOARCH=$(ARCH) go build -o ../build/browser-demux .

.PHONY: test
test: setup
	cd ./browser-demux && go test -v ./...

.PHONY: all
all: build

.PHONY: appimage
appimage: build
	rm -rf $(APP_IMAGE_DIR)
	mkdir -p $(APP_IMAGE_DIR)
	cp ./build/browser-demux $(APP_IMAGE_DIR)
	cp ./packaging/browser-demux.svg $(APP_IMAGE_DIR)
	cp ./packaging/browser-demux.desktop $(APP_IMAGE_DIR)
	ln -s -r $(APP_IMAGE_DIR)/icon.png $(APP_IMAGE_DIR)/.DirIcon
	ln -s -r $(APP_IMAGE_DIR)/browser-demux $(APP_IMAGE_DIR)/AppRun
	mkdir -p $(DIST_DIR)
	appimagetool-x86_64.AppImage $(APP_IMAGE_DIR) $(DIST_DIR)/BrowserDemux-$(ARCH).AppImage

.PHONY: clean
clean:
	rm -rf BUILD_DIR
	rm -rf DIST_DIR