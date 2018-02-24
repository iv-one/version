require 'rbconfig'
class Version < Formula
  desc ""
  homepage "https://github.com/ivan-dyachenko/version"
  version "1.0.0"

  if Hardware::CPU.is_64_bit?
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/ivan-dyachenko/version/releases/download/v1.0.0/version_1.0.0_darwin_amd64.zip"
      sha256 "a923c601d820fe1a4b6dfa2d94e83f288dc3c90ae081b108108b673537ae5306"
    when /linux/
      url "https://github.com/ivan-dyachenko/version/releases/download/v1.0.0/version_1.0.0_linux_amd64.tar.gz"
      sha256 "cbd2e23b51fd1e3605d10692f4013dad5184257786666b239c98045c8c8c708f"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  else
    case RbConfig::CONFIG['host_os']
    when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
      :windows
    when /darwin|mac os/
      url "https://github.com/ivan-dyachenko/version/releases/download/v1.0.0/version_1.0.0_darwin_386.zip"
      sha256 "b2de0656b96448cb3122084ad009d557503ddde099aea62557da1b4c10fc4d43"
    when /linux/
      url "https://github.com/ivan-dyachenko/version/releases/download/v1.0.0/version_1.0.0_linux_386.tar.gz"
      sha256 "d8db8922b8cef1818ea4f15e817ca602bb7141236a0360a8db7e3cb9fb4f8ba7"
    when /solaris|bsd/
      :unix
    else
      :unknown
    end
  end

  def install
    bin.install "version"
  end

  test do
    system "version"
  end

end
