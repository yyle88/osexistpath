COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/neatjson/blob/18a93d8accdeb633f25d4200fdd003362bfcc6cb/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
