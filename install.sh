#!/bin/sh
echo -n "Copying flydown binary..."
cp -f ./flydown /usr/bin
echo "Done"

echo -n "Copying static files..."
cp -af ./static /usr/share/flydown/
cp -af ./templates /usr/share/flydown/
echo "Done"

echo "Flydown installation has been finished"