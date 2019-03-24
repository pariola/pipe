#!/bin/sh

# cleanup
echo "Cleaning up!"
rm -rf "$(pwd)/example"


git clone https://blessingpariola@bitbucket.org/blessingpariola/pipe-test.git "$(pwd)/example"
echo "Cloning complete"