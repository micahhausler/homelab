#!/usr/bin/env bash
set -e
set -x

METADATA_DIR=${METADATA_DIR:-/run/metadata}

# network creates directories and files under $METADATA_DIR/networks. Given a
# machine with a single external interface 'eno1', the following files and
# directories would be created
#
# /run/metadata/networks
# /run/metadata/networks/default -> /run/metadata/networks/iface/eno1
# /run/metadata/networks/iface
# /run/metadata/networks/iface/eno1
# /run/metadata/networks/iface/eno1/ipv6
# /run/metadata/networks/iface/eno1/ipv4
# /run/metadata/networks/iface/eno1/delegated_ipv6
network(){
    local -r default_iface=$(ip route | grep default | awk '{print $5}')
    local -r links=$(ip link | grep -v link | awk '{print $2}' | sed -e 's/://g')
    for link in ${links}; do
        if [ ${link} ==  docker0 ] || [ ${link} == lo ] ; then
            continue
        fi
        link_dir=$METADATA_DIR/networks/iface/$link
        mkdir -p $link_dir
        ipv4=$(ip -4 addr show $link | grep 'inet ' | awk '{print $2}' | cut -f1 -d/)
        echo $ipv4 | tee $link_dir/ipv4
        chmod 444 $link_dir/ipv4
        ipv6=$(ip -6 addr show $link | grep 'inet'  | awk '{print $2}' | grep '/128' | cut -f1 -d/)
        echo $ipv6 | tee $link_dir/ipv6
        chmod 444 $link_dir/ipv6
        delegated_ipv6=$(ip -6 addr show $link scope global mngtmpaddr | grep inet6 | awk '{print $2}')
        echo $delegated_ipv6| tee $link_dir/delegated_ipv6
        chmod 444 $link_dir/delegated_ipv6
        if [ ${link} == $default_iface ]; then
            ln -s $link_dir ${METADATA_DIR}/networks/default
        fi
    done
}

# hostname writes the machine hostname to
# $METADATA_DIR/hostname
write_hostname(){
    echo $(hostname) | tee ${METADATA_DIR}/hostname
    chmod 444 ${METADATA_DIR}/hostname
}


mkdir -p ${METADATA_DIR}
write_hostname
network
