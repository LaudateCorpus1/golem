SHELL:=/bin/bash
default: download-all
LANG=en
download-all:
	mkdir -p data
	$(MAKE) den
	$(MAKE) dsv
	$(MAKE) dfr
	$(MAKE) des
	$(MAKE) dde
	rm data/*.zip
	$(MAKE) build-bin

build-bin:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o dicts/data.go -nocompress data/

en:
	mkdir -p data
	$(MAKE) den
	rm data/*.zip
	$(MAKE) build-bin

den: LANG=en
den: download

dsv: LANG=sv
dsv: download

dfr: LANG=fr
dfr: download

des: LANG=es
des: download

dde: LANG=de
dde: download

download:
	curl http://www.lexiconista.com/Datasets/lemmatization-$(LANG).zip > data/$(LANG).zip
	unzip data/$(LANG).zip -d data
	mv data/lemmatization-$(LANG).txt data/$(LANG).orig
	go run normalizer/normalizer.go --in data/$(LANG).orig --out data/$(LANG)
	gzip data/$(LANG)
