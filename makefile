# TODO improve goals, improve build instructions

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

ALL_GO_PACKAGES = ag algo cont fio ui utils

# SRC_FILENAME ?= data/src.txt
# DICT_FILENAME ?= data/russian_words.txt
# DUMP_FILENAME ?= acrs_dump.txt
# OUTP_FILENAME ?= acrs.txt

#--------------------------------------------------------------------------------------------------
# Input command line params
#--------------------------------------------------------------------------------------------------

PACKAGES ?= $(ALL_GO_PACKAGES)

# INPUT_PARAMS = $(SRC_FILENAME) $(DICT_FILENAME) $(DUMP_FILENAME) $(OUTP_FILENAME)

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

BASIC_TEST_COMMAND = go test -race -cover -count=1

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

test_coverage_in_details:
	go test -coverprofile=coverage.out $(addprefix github.com/Shangin-Leonid/acrogen/, $(PACKAGES))
	go tool cover -html=coverage.out

debug: build
	dlv debug --

clean:
	rm -f $(EXE_FILENAME)

clean_all:
	rm -f $(EXE_FILENAME)
	rm -f $(OUTP_FILENAME)
	rm -f $(DUMP_FILENAME)
