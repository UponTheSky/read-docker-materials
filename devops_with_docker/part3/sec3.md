# Using a non-root user
- security issue: the application could escape the container due to a bug in Docker or Linux kernel
- we need to add a non-root user to our container and run the processes with that user
    - add user: `RUN useradd -m appuser` (`useradd`: linux command)
    - change the user: `USER appuser`: all the commands after this directive is executed by this new user

- what if we want to give permissions to the new user?
    - be aware that giving permissions should be done while we're still the root user in the container(i.e. before `USER`)
