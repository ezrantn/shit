class Shit < Formula
    desc "A simple CLI tool to kill processes by name"
    homepage "https://github.com/ezrantn/shit"
    url "https://github.com/ezrantn/shit/archive/refs/tags/v1.0.0.tar.gz"
    sha256 "d21393e9f6c8bbc26d187cd325be171c8081105c42862f3e5845dc33a62a045a"
    head "https://github.com/ezrantn/shit.git", branch: "main"

    depends_on "go" => :build

    def install
        system "go", "build", "-o", bin/"shit"
    end
end