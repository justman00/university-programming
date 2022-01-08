#!/bin/bash

nasm -f macho64 lab1-32.asm
ld ./lab1-32.o -lSystem -L/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/lib
./a.out
