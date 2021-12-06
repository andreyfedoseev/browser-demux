APP_IMAGE_DIR = ./build/AppImage
ARCH = x64

.PHONY: build
build:
	cd ./browser-demux && go build -o ../build/browser-demux

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
	appimagetool-x86_64.AppImage $(APP_IMAGE_DIR) ./build/BrowserDemux-$(ARCH).AppImage
