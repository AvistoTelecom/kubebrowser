#!/usr/bin/env bash

set -o noclobber
set -o errexit
set -o pipefail

SCRIPT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/.."

# Example usage: ./release.sh
version=""
IMAGE_REGISTRY="brawdunoir"
HELM_REPO="oci://rgy.k8s.devops-svc-ag.com/avisto/helm"
APP_NAME="kubebrowser"
CHART_PACKAGE=""

LIGHT_GREEN="\033[1;32m"
RED="\033[1;31m"
RESET="\033[0m"

usage() {
  echo "Usage: $0"
  echo "This script automatically parses the version from chart/Chart.yaml."
  exit 1
}

print_green() {
  echo -e "${LIGHT_GREEN}$1${RESET}"
}

print_red() {
  echo -e "${RED}$1${RESET}"
}

run_command() {
  local command=$1

  print_green "Do - $command"
  if ! eval $command; then
    print_red "Error - $command"
    exit 1
  fi
  print_green "Done - $command"
}

check_dependencies() {
  for cmd in skaffold helm yq; do
    if ! command -v $cmd &>/dev/null; then
      print_red "Error: $cmd is not installed or not in PATH."
      exit 1
    fi
  done
}

check_existing_version_in_registry() {
  local image="$IMAGE_REGISTRY/$APP_NAME:$version"
  if ! docker inspect "$image" &>/dev/null; then
    print_red "Error: Version $version already exists in the Docker registry ($image)."
    exit 1
  fi
}

validate_version_in_footer() {
  local footer_file="$SCRIPT_ROOT/ui/src/components/AppFooter.vue"
  if ! grep -q "$version" "$footer_file"; then
    print_red "Error: Version $version not changed in $footer_file."
    exit 1
  fi
}

validate_version_in_values() {
  local values_file="$SCRIPT_ROOT/chart/values.yaml"
  if ! grep -q "tag: \"$version\"" "$values_file"; then
    print_red "Error: Version $version not found in $values_file."
    exit 1
  fi
}

validate_variables() {
  if [[ -z "$version" ]]; then
    print_red "Error: version is not set. Ensure Chart.yaml has a valid version."
    exit 1
  fi
  if [[ -z "$IMAGE_REGISTRY" || -z "$HELM_REPO" || -z "$APP_NAME" ]]; then
    print_red "Error: One or more required variables are not set."
    exit 1
  fi
  CHART_PACKAGE="${APP_NAME}-${version}.tgz"
}

cleanup() {
  print_green "Cleaning up..."
  if [[ -n "$CHART_PACKAGE" ]]; then
    touch "$CHART_PACKAGE"
    rm "$CHART_PACKAGE"
  fi
}

trap cleanup EXIT

# Handle --help
if [[ $# -gt 0 && $1 == "--help" ]]; then
  usage
fi

version=$(yq -r '.version' "$SCRIPT_ROOT/chart/Chart.yaml")

check_dependencies
validate_variables
validate_version_in_footer
validate_version_in_values
check_existing_version_in_registry

run_command "skaffold build -t $version -d $IMAGE_REGISTRY --cache-artifacts=false"
run_command "helm package chart"
run_command "helm push $CHART_PACKAGE $HELM_REPO"

print_green "Release process completed successfully."
