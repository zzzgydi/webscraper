HOMEDIR := $(shell pwd)
PROJECTNAME := $(shell basename $(HOMEDIR))
OUTDIR  := $(HOMEDIR)/output

GOBUILD := go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"

build:
	mkdir -p $(OUTDIR)
	$(GOBUILD) -o $(OUTDIR)/$(PROJECTNAME) ./cmd/main.go


dev: build
	mkdir -p $(OUTDIR)/log

	@if [ ! -f $(HOMEDIR)/config/dev.yaml ]; then \
		echo "Copying temp.yaml to dev.yaml..."; \
		cp $(HOMEDIR)/config/temp.yaml $(HOMEDIR)/config/dev.yaml; \
	fi

	cp -r $(HOMEDIR)/assets  $(OUTDIR)/
	cp -r $(HOMEDIR)/config  $(OUTDIR)/
	cp -r $(HOMEDIR)/static  $(OUTDIR)/

	CONFIG_ENV=dev $(OUTDIR)/$(PROJECTNAME)
