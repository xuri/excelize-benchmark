#!/bin/bash
# SheetJS
cd SheetJS
cleanup_nodejs()
{
    if [ -d "node_modules" ]; then rm -rf node_modules; fi
    if [ -f "package-lock.json" ]; then rm -f package-lock.json; fi
    rm -rf *.xlsx
}
cleanup_nodejs
npm install --registry='https://registry.npmjs.org'
NODE_OPTIONS=--max_old_space_size=10240 node index.js
cleanup_nodejs
cd ..

# unidoc
cd unidoc
go build
./unidoc
rm -f unidoc
rm -rf *.xlsx
cd ..

# tealeg/xlsx
cd xlsx
go build
./xlsx
rm -f xlsx
rm -rf *.xlsx
cd ..
