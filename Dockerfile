FROM docker.repo-ci.sfera.inno.local/gdcr-docker/docker-base-images/python:3.10.4

ENV TZ=Europe/Moscow

ARG req=requirements.txt
ARG NEXUS_CI_USERNAME
ARG NEXUS_CI_PASSWORD
ARG PULL_CI_REGISTRY

WORKDIR app

COPY . .

RUN pip3 install --upgrade pip \
    --index-url "https://${NEXUS_CI_USERNAME}:${NEXUS_CI_PASSWORD}@sfera.inno.local/app/repo-ci-misc/api/repository/gdcr-pypi/simple"\
    --trusted-host "sfera.inno.local"

# Install requirements
RUN pip3 install -r $req \
    --index-url "https://${NEXUS_CI_USERNAME}:${NEXUS_CI_PASSWORD}@sfera.inno.local/app/repo-ci-misc/api/repository/gdcr-pypi/simple"\
    --trusted-host "sfera.inno.local"

ENTRYPOINT ["python3", "main.py"]