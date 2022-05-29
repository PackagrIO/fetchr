

# References

https://www.jfrog.com/confluence/display/JFROG/Xray+REST+API#XrayRESTAPI-ComponentIdentifiers

| Package Type | Identifier | Example | 
| --- | --- | --- |
| Maven | gav://group:artifact:version | gav://ant:ant:1.6.5 |
| Docker | docker://Namespace/name:tag | docker://jfrog/artifactory-oss:latest |
| RPM | rpm://os-version:package:epoch-version:version | rpm://7:rpm-python:7:4.11.3-43.el7 |
| Debian | deb://vendor:dist:package:version | deb://ubuntu:trustee:acl:2.2.49-2 |
| NuGet | nuget://module:version | nuget://log4net:9.0.1 |
| Generic file | generic://sha256:<Checksum>/name | generic://sha256:244fd47e07d1004f0aed9c156aa09083c82bf8944eceb67c946ff7430510a77b/foo.jar |
| NPM | npm://package:version | npm://mocha:2.4.5 |
| Python | pypi://package:version | pypi://raven:5.13.0 |
| Composer | composer://package:version | composer://nunomaduro/collision:1.1 |
| Golang | go://package:version | go://github.com/ethereum/go-ethereum:1.8.2 |
| Alpine | alpine://branch:package:version | alpine://3.7:htop:2.0.2-r0 | 
| Conan | conan://vendor:name:version | conan://openssl:openssl:1.1.1g | 

https://github.com/package-url/purl-spec

```
pkg:bitbucket/birkenfeld/pygments-main@244fd47e07d1014f0aed9c

pkg:deb/debian/curl@7.50.3-1?arch=i386&distro=jessie

pkg:docker/cassandra@sha256:244fd47e07d1004f0aed9c
pkg:docker/customer/dockerimage@sha256:244fd47e07d1004f0aed9c?repository_url=gcr.io

pkg:gem/jruby-launcher@1.1.2?platform=java
pkg:gem/ruby-advisory-db-check@0.12.4

pkg:github/package-url/purl-spec@244fd47e07d1004f0aed9c

pkg:golang/google.golang.org/genproto#googleapis/api/annotations

pkg:maven/org.apache.xmlgraphics/batik-anim@1.9.1?packaging=sources
pkg:maven/org.apache.xmlgraphics/batik-anim@1.9.1?repository_url=repo.spring.io/release

pkg:npm/%40angular/animation@12.3.1
pkg:npm/foobar@12.3.1

pkg:nuget/EnterpriseLibrary.Common@6.0.1304

pkg:pypi/django@1.11.1

pkg:rpm/fedora/curl@7.50.3-1.fc25?arch=i386&distro=fedora-25
pkg:rpm/opensuse/curl@7.56.1-1.1.?arch=i386&distro=opensuse-tumbleweed

```


- https://github.com/replit/upm
- CycloneDX - Uniform Resource Name (URN) 
- "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.5.0?type=module",
- https://github.com/package-url/purl-spec
- 