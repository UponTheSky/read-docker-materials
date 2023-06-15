# Official Images and trust

## Look into official images
- docker official images
    - stored in the github repo("library of images")
    - introduced into the library by PR
    - verifying image: described in README.md
    - images are maintained by either the docker-library organization or separate organizations

- investigate official images: ubuntu
    - `docker pull <image> && docker image history <image>`
    - `FROM scratch`: means just empty
    - `ADD`: copies new files, dirs, or remote file urls from `src`

- so can we trust them?
    - see `ADD` command
    - see the checksum hash value(one of PR commits?)
    - security issues on the official image page

**Official images are nothing special**
~                                                          
