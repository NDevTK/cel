#! /usr/bin/env bash

if [ ! -e index.md ]; then
  echo This script should be run inside the //docs directory.
  exit 1
fi

set -e

for f in *.md; do
  if [ $f != "index.md" ]; then
    sed -i '' -e '/-- INSERT-INDEX --/rindex.md
/-- BEGIN-INDEX --/,/-- END-INDEX --/d' $f
  fi
done

