#!/bin/bash
set -x

mkdir -p mmdb && cd mmdb || exit
wget "https://cdn.jsdelivr.net/npm/geolite2-city@1.0.0/GeoLite2-City.mmdb.gz"
wget "https://cdn.jsdelivr.net/npm/geolite2-country@1.0.2/GeoLite2-Country.mmdb.gz"
gunzip GeoLite2-City.mmdb.gz
gunzip GeoLite2-Country.mmdb.gz
rm -f GeoLite2-City.mmdb.gz
rm -f GeoLite2-Country.mmdb.gz
echo "Download geo mmdb files successfully"
