FROM kindest/node:v1.31.2@sha256:18fbefc20a7113353c7b75b5c869d7145a6abd6269154825872dc59c1329912e AS base

ARG TARGETOS
ARG TARGETARCH

RUN apt-get update -yq && \
    apt-get install -yq --no-install-recommends wget apparmor apparmor-utils jq openssh-server sudo

# remove kind's kubelet unit
RUN rm -f /etc/systemd/system/kubelet.service && \
    rm -rf /etc/systemd/system/kubelet.service.d/

# copy containerd hosts configurations for local registry mirrors
COPY hosts.toml /etc/containerd/certs.d/garden.local.gardener.cloud:5001/hosts.toml

# add our userdata executor unit
COPY run-userdata.sh /run-userdata.sh
COPY run-userdata.service /etc/systemd/system
RUN systemctl enable run-userdata.service

# TODO(marc1404): Remove once kindest/node uses runc >= v1.2.4
# workaround issue with runc v1.2.3 provided by kindest/node:v1.32.0 by installing runc v1.2.4 manually (https://github.com/opencontainers/runc/pull/4555)
COPY patch-runc.sh /patch-runc.sh
RUN chmod +x /patch-runc.sh \
    && /patch-runc.sh \
    && rm -f /patch-runc.sh

ENTRYPOINT ["/usr/local/bin/entrypoint", "/sbin/init"]
