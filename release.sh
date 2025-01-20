#!/bin/bash

# Function to bump version parts
bump_version_parts() {
    local version_parts=("$@")
    local bump_type=$1

    case $bump_type in
        major)
            version_parts[0]=$((version_parts[0] + 1))
            version_parts[1]=0
            version_parts[2]=0
            ;;
        minor)
            version_parts[1]=$((version_parts[1] + 1))
            version_parts[2]=0
            ;;
        patch|*)
            version_parts[2]=$((version_parts[2] + 1))
            ;;
    esac

    echo "${version_parts[0]}.${version_parts[1]}.${version_parts[2]}"
}

# Function to bump version in a file
bump_version_in_file() {
    local file=$1
    local version_regex=$2
    local bump_type=$3
    local version_line=$(grep -E "$version_regex" "$file")
    local current_version=$(echo "$version_line" | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
    
    if [ -z "$current_version" ]; then
        echo "No version found in $file"
        return
    fi

    IFS='.' read -r -a version_parts <<< "$current_version"
    local new_version=$(bump_version_parts "${version_parts[@]}" "$bump_type")

    # Replace the old version with the new version
    sed -i.bak -E "s/$current_version/$new_version/" "$file"
    echo "Updated $file from $current_version to $new_version"
}

# Ask the user for the type of version bump
echo "Enter the type of version bump (major, minor, patch):"
read -r bump_type
bump_type=${bump_type:-patch}  # Default to patch if no input

# Clean up backup files created by sed
find . -name "*.bak" -type f -delete

# Ask the user which SDKs to bump
echo "Select SDKs to bump (python, go, both):"
read -r sdk_choice
sdk_choice=${sdk_choice:-both}  # Default to both if no input

# Bump version for selected SDKs
if [[ "$sdk_choice" == "python" || "$sdk_choice" == "both" ]]; then
    python_setup_file="python/pyproject.toml"
    python_version_regex="version = \"[0-9]+\.[0-9]+\.[0-9]+\""
    bump_version_in_file "$python_setup_file" "$python_version_regex" "$bump_type"
    git add python/pyproject.toml
    git commit -m "Bump version to $new_version"

    cd python || exit
    # clean up dist folder
    rm -rf dist
    # build the package
    python -m build

    # Upload the package to PyPI
    twine upload dist/*

    cd - || exit

    # Clean up backup files created by twine
    find . -name "*.bak" -type f -delete

    git push
fi

if [[ "$sdk_choice" == "go" || "$sdk_choice" == "both" ]]; then
    # Fetch the latest tags
    git fetch --tags

    # Get the latest tag
    latest_tag=$(git describe --tags --abbrev=0)
    echo "Latest tag: $latest_tag"

    # Extract the version numbers
    IFS='.' read -r -a version_parts <<< "${latest_tag#v}"

    # Bump the version based on the bump type
    new_tag="v$(bump_version_parts "${version_parts[@]}" "$bump_type")"
    echo "Creating new tag: $new_tag"

    # Create and push the new tag
    git tag "$new_tag"

    # Push commits and new tag
    git push origin HEAD
    git push origin "$new_tag"
fi

echo "Version bumping completed."