GO ?= go
COVERAGEDIR = coverage
ifdef CIRCLE_ARTIFACTS
	COVERAGEDIR=$(CIRCLE_ARTIFACTS)/coverage
endif

.PHONY: test

TEST_LIST = $(foreach pkg, $(ALL_PACKAGES), $(pkg)_test)
COVER_LIST = $(foreach pkg, $(ALL_PACKAGES), $(pkg)_cover)

deps:
	dep ensure

gen-mocks:
	mockery -dir=./logger/ -all
	mockery -dir=./metrics/ -all
	mockery -dir=./models/ -all
	mockery -dir=./middleware/ -all
	mockery -dir=./utils/ -all

test:
	if [ ! -d $(COVERAGEDIR) ]; then mkdir $(COVERAGEDIR); fi
	AWS_ACCESS_KEY_ID=1 AWS_SECRET_ACCESS_KEY=1 go test -v ./$* -race -cover -covermode=atomic -coverprofile=$(COVERAGEDIR)/$(subst /,_,$*).coverprofile
	$(GO) test -ldflags -s -v ./logger -race -cover -coverprofile=$(COVERAGEDIR)/logger.coverprofile
	$(GO) test -ldflags -s -v ./metrics -race -cover -coverprofile=$(COVERAGEDIR)/metrics.coverprofile
	$(GO) test -ldflags -s -v ./middleware -race -cover -coverprofile=$(COVERAGEDIR)/middleware.coverprofile
	$(GO) test -ldflags -s -v ./models -race -cover -coverprofile=$(COVERAGEDIR)/models.coverprofile
	$(GO) test -ldflags -s -v ./utils -race -cover -coverprofile=$(COVERAGEDIR)/utils.coverprofile

cover:
	go tool cover -html=$(COVERAGEDIR)/logger.coverprofile -o $(COVERAGEDIR)/logger.html
	go tool cover -html=$(COVERAGEDIR)/middleware.coverprofile -o $(COVERAGEDIR)/middleware.html
	go tool cover -html=$(COVERAGEDIR)/metrics.coverprofile -o $(COVERAGEDIR)/metrics.html
	go tool cover -html=$(COVERAGEDIR)/models.coverprofile -o $(COVERAGEDIR)/models.html
	go tool cover -html=$(COVERAGEDIR)/utils.coverprofile -o $(COVERAGEDIR)/utils.html

docs:
	@godoc -http=:6060 2>/dev/null &
	@printf "To view geo-api docs, point your browser to:\n"
	@printf "\n\thttp://127.0.0.1:6060/pkg/github.com/hashknife/common/$(pkg)\n\n"
	@sleep 1
	@open "http://127.0.0.1:6060/pkg/github.com/hashknife/common/$(pkg)"

tc: test cover

coveralls:
	gover $(COVERAGEDIR) $(COVERAGEDIR)/coveralls.coverprofile
	goveralls -coverprofile=$(COVERAGEDIR)/coveralls.coverprofile -service=circle-ci -repotoken=$(COVERALLS_TOKEN); echo "Coveralls finished"

bench:
	go test -bench ./...

clean:
	$(GO) clean
	rm -rf coverage/
