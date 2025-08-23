#!/usr/bin/env fish

echo "Creating fresh virtualenv..."
python3 -m venv build-venv
source build-venv/bin/activate.fish

echo "Installing build tools to virtualenv..."
python3 -m pip install --quiet --upgrade pip build twine

echo "Building and uploading packages..."
python -m build > /dev/null
python -m twine upload dist/*

echo "Deleting build artifacts..."
trash build-venv/ dist/ sabot/sabot.egg-info/

echo "Done."
