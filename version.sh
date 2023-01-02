#!/bin/bash

major() {
if IFS=. read -r major rest <version.txt || [ -n "$major" ]; then
  echo "$((major + 1)).0.0" >"version.txt"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

minor() {
if IFS=. read -r major minor patch <version.txt || [ -n "$major" ]; then
  echo "$major.$((minor + 1)).0" >"version.txt"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

patch() {
if IFS=. read -r major minor patch <version.txt || [ -n "$major" ]; then
  echo "$major.$minor.$((patch + 1))" >"version.txt"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

print_major() {
if IFS=. read -r major minor patch <version.txt || [ -n "$major" ]; then
  echo "$major"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

print_major_minor() {
if IFS=. read -r major minor patch <version.txt || [ -n "$major" ]; then
  echo "$major.$minor"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

print_major_minor_patch() {
if IFS=. read -r major minor patch <version.txt || [ -n "$major" ]; then
  echo "$major.$minor.$patch"
else
  echo "ERROR: Unable to read version number from version.txt" >&2
  exit 1
fi
}

case "$1" in
  major)
    major $2
    ;;
  minor)
    minor $2 
    ;;
  patch)
    patch $2
    ;;
  print_major)
    print_major $2
    ;; 
  print_major_minor)
    print_major_minor $2
    ;;   
  print_major_minor_patch)
    print_major_minor_patch $2
    ;;  
  *)
    echo "Usage: bash version.sh {major|minor|patch|print_major_minor_patch}"
    exit 1
esac
exit 0