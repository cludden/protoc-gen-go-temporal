#!/bin/sh

# modified from https://github.com/release-lab/install

set -e

# eg. release-lab/whatchanged
target="cludden/protoc-gen-go-temporal"
owner="cludden"
repo="protoc-gen-go-temporal"
exe_name="protoc-gen-go_temporal"
githubUrl=""
githubApiUrl=""
version=""

get_arch() {
    a=$(uname -m)
    case ${a} in
        "x86_64" | "amd64" )
            echo "amd64"
        ;;
        "aarch64" | "arm64" | "arm")
            echo "arm64"
        ;;
        *)
            echo ${NIL}
        ;;
    esac
}

get_os(){
    # darwin: Darwin
    echo $(uname -s | awk '{print tolower($0)}')
}

# parse flag
for i in "$@"; do
    case $i in
        -v=*|--version=*)
            version="${i#*=}"
            shift # past argument=value
        ;;
        *)
            # unknown option
        ;;
    esac
done

if [ -z "$githubUrl" ]; then
    githubUrl="https://github.com"
fi
if [ -z "$githubApiUrl" ]; then
    githubApiUrl="https://api.github.com"
fi

downloadFolder="${TMPDIR:-/tmp}"
mkdir -p ${downloadFolder} # make sure download folder exists
os=$(get_os)
arch=$(get_arch)
executable_folder="/usr/local/bin" # Eventually, the executable file will be placed here

# fetch latest version if version is empty
if [ -z "$version" ]; then
    version=$(curl --silent https://api.github.com/repos/cludden/protoc-gen-go-temporal/releases/latest|jq -r .tag_name)
fi
file_name="${exe_name}_${version#"v"}_${os}_${arch}.tar.gz" # the file name should be download
downloaded_file="${downloadFolder}/${file_name}" # the file path should be download
asset_uri="${githubUrl}/${owner}/${repo}/releases/download/${version}/${file_name}"

echo "[1/3] Download ${asset_uri} to ${downloadFolder}"
rm -f ${downloaded_file}
curl --fail --location --output "${downloaded_file}" "${asset_uri}"

echo "[2/3] Install ${exe_name} to ${executable_folder}"
tar -xz -f ${downloaded_file} -C ${executable_folder}
exe=${executable_folder}/${exe_name}
chmod +x ${exe}

echo "${exe_name} was installed successfully to ${exe}"
if command -v $exe_name --version >/dev/null; then
    echo "Run '$exe_name -version' to get started"
else
    echo "Manually add the directory to your \$HOME/.bash_profile (or similar)"
    echo "  export PATH=${executable_folder}:\$PATH"
    echo "Run '$exe_name -version' to get started"
fi

exit 0
