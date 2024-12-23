class Shit < Formula
    desc "A simple CLI tool to kill processes by name"
    homepage "https://github.com/ezrantn/shit"
    url "https://github.com/ezrantn/shit/archive/refs/tags/v1.0.0.tar.gz"
    sha256 "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
    head "https://github.com/ezrantn/shit.git", branch: "main"

    depends_on "go" => :build

    def install
        system "go", "build", "-o", bin/"shit"
    end
end