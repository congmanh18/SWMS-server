#!/bin/sh

if [ "$#" -lt 2 ]; then
  echo "Usage: $0 host:port -- command args..."
  exit 1
fi

hostport="$1"
shift

host=$(echo $hostport | cut -d: -f1)
port=$(echo $hostport | cut -d: -f2)

while ! nc -z $host $port; do
  echo "Waiting for $host:$port..."
  sleep 1
done

exec "$@"
