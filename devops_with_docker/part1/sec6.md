# Utilizing tools from the Registry
How do we create a Dockerfile for a website project?

Follow the order specified below:

1. `FROM` layer: pull the very basic image for your project

```Dockerfile
FROM python:3.9-slim
```

2. Setting `WORKDIR`: you need to set this in order to work in the container!!

```Dockerfile
WORKDIR /usr/src/app
```

3. Install the dependencies: separating depedency installation from the source code makes the subsequent image building way faster
  - this is because the source code always changes, whereas dependencies change less often
  - so we can treat the collection of the dependencies as a separate layer to be cached

```Dockerfile
# install another tools if necessary
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
```

4. Copy the source code and set additioanl configs

```Dockerfile
COPY . .

ENV FOO=bar
EXPOSE 3000

RUN db_migration

CMD ["python", "./main.py"]
```

