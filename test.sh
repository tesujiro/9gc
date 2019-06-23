#!/bin/bash
try() {
  expected="$1"
  input="$2"

  ./9gc "$input" > tmp.s
  gcc -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    #echo "$input => $actual"
    :
  else
    echo "$expected expected, but got $actual"
    echo "input: $input"
    exit 1
  fi
}

try 0 "0;"
try 42 "42;"
try 24 "1+23;"
try 8 "23-15;"
try 17 "23 - 6;"
try 31 "5*(3+4)-12/3;"
try 47 "5+6*7;"
try 15 "5*(9-6);"
try 4 "(3+5)/2;"
try 15 "+5+10;"
try 5 "-5+10;"
try 1 "-100==-100;"
try 0 "-100==100;"
try 1 "123!=-100;"
try 0 "123!=123;"
try 1 "-100==-100;"
try 1 "5<10;"
try 0 "5<5;"
try 1 "5<=7;"
try 1 "5<=5;"
try 0 "5<=3;"
try 1 "1>0;"
try 0 "0>0;"
try 1 "0>=0;"
try 0 "-100>=0;"
try 0 "i=0;"
try 42 "x=42;"
try 3 "a=1;b=2;c=3;"
try 42 "return 42;"
try 44 "return (+5-3)*22;"
try 42 "i=42;return i;"
try 5  "j=5;return j;"
try 47 "i=42;j=5;k=i+j;return k;"
try 63 "a=1;b=20;c=3;return (a+b)*c;"

echo OK
