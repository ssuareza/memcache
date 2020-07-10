#!/bin/bash
# Invalidate all cache
echo 'flush_all' | ncat localhost 11211
