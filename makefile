PACKAGE=github.com/robsignorelli/isodates

TESTING_FLAGS=
ifeq ($(VERBOSE),true)
	TESTING_FLAGS=-v
endif

#
# Runs through our suite of all unit tests
#
test:
	go test $(TESTING_FLAGS) -timeout 5s $(PACKAGE)/...

benchmark:
	go test -bench=.
