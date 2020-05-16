#!/bin/bash

rm Mp3Listaren.exe

dotnet publish -c Release -r win10-x64 --self-contained true -p:PublishSingleFile=true -p:PublishTrimmed=true --output .
rm -rf bin
rm -rf obj
rm Mp3Listaren.pdb

 
