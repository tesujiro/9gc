#!/bin/bash
try() {
  expected="$1"
  input="$2"

  ./9gc "$input" > tmp.s
  gcc -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$expected expected, but got $actual"
    exit 1
  fi
}

try 0 0
try 42 42
try 24 "1+23"
try 8 "23-15"
try 17 "23 - 6"
try 31 "5*(3+4)-12/3"

echo OK
