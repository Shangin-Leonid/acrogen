#--------------------------------------------------------------------------------------------------
# Extensions
#--------------------------------------------------------------------------------------------------

ifeq ($(OS),Windows_NT)
	EXE_EXT := exe
else
	EXE_EXT := out
endif

#--------------------------------------------------------------------------------------------------
# File names
#--------------------------------------------------------------------------------------------------

EXE_FILENAME = ./acrogen.$(EXE_EXT)

TEST_COVERAGE_FILENAME = coverage.out

ALL_GO_PACKAGES = ag algo cont fio ui utils

#--------------------------------------------------------------------------------------------------
# Input command line params
#--------------------------------------------------------------------------------------------------

PACKAGES ?= $(ALL_GO_PACKAGES)

#--------------------------------------------------------------------------------------------------
# Run params
#--------------------------------------------------------------------------------------------------

LOG_FILENAME := ./data/acrogen.log
ifeq ($(OS),Windows_NT)
	LOG_REDIRECT_COMMAND =
else
	LOG_REDIRECT_COMMAND = 2> $(LOG_FILENAME)
endif

#--------------------------------------------------------------------------------------------------
# Golang flags
#--------------------------------------------------------------------------------------------------

BASIC_TEST_COMMAND = go test -race -cover -count=1 -coverprofile=$(TEST_COVERAGE_FILENAME)

#--------------------------------------------------------------------------------------------------
# Rules / goals / targets
#--------------------------------------------------------------------------------------------------

all: run

build:
	go build -trimpath -o $(EXE_FILENAME) *.go

run:
	$(EXE_FILENAME) $(LOG_REDIRECT_COMMAND)

test:
	$(BASIC_TEST_COMMAND) $(addprefix github.com/Shangin-Leonid/acrogen/, $(PACKAGES))

testv:
	$(BASIC_TEST_COMMAND) -v $(addprefix github.com/Shangin-Leonid/acrogen/, $(PACKAGES))

open_coverage_report:
	go tool cover -html=$(TEST_COVERAGE_FILENAME)

debug: build
	dlv debug --

clean:
	rm -f $(EXE_FILENAME)

clean_all:
	rm -f $(EXE_FILENAME)
	rm -f $(OUTP_FILENAME)
	rm -f $(DUMP_FILENAME)
