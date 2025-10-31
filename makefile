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

# SRC_FILENAME ?= data/src.txt
# DICT_FILENAME ?= data/russian_words.txt
# DUMP_FILENAME ?= acrs_dump.txt
# OUTP_FILENAME ?= acrs.txt

EXE_FILENAME = ./acrogen.$(EXE_EXT)

#--------------------------------------------------------------------------------------------------
# Input coomand line params
#--------------------------------------------------------------------------------------------------

# INPUT_PARAMS = $(SRC_FILENAME) $(DICT_FILENAME) $(DUMP_FILENAME) $(OUTP_FILENAME)

#--------------------------------------------------------------------------------------------------
# Rules / goals / targets
#--------------------------------------------------------------------------------------------------

all: build run

build:
	go build -o $(EXE_FILENAME) *.go

run:
	$(EXE_FILENAME)
#	$(EXE_FILENAME) $(INPUT_PARAMS)

debug:
	dlv debug -- $(INPUT_PARAMS)

clean:
	rm -f $(EXE_FILENAME)

clean_all:
	rm -f $(EXE_FILENAME)
	rm -f $(OUTP_FILENAME)
	rm -f $(DUMP_FILENAME)
