set operating_systems linux darwin windows
set architectures amd64 arm64
for operating_system in $operating_systems
    set file_extension (test $operating_system = windows && echo ".exe" || echo "")
    for architecture in $architectures
        set output_file_path "build/mdn-$operating_system-$architecture$file_extension"
        echo "Building $output_file_path."
        GOOS=$operating_system GOARCH=$ARCHITECTURE go build -o $output_file_path
    end
end
