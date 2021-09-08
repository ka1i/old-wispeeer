PREFIX  := wispeeer
SOURCE  := cmd/${PREFIX}/main.go
BINARY  := bin/${PREFIX}

all: build

.PHONY: build
build:          ## build with native env.
	@./scripts/build.sh ${SOURCE} ${BINARY}


.PHONY: clean
clean:          ## Clean build cache.
	@rm -rf bin
	@echo "clean [ ok ]"