#!/usr/bin/env sh
#
# Modified from https://www.cyberciti.biz/faq/how-to-set-up-wireguard-firewall-rules-in-linux/

IPT="/sbin/iptables"
if [ -f /usr/sbin/iptables ]; then
  IPT="/usr/sbin/iptables";
fi
#IPT6="/usr/sbin/ip6tables"

IN_FACE="eth4"                       # NIC connected to the internet
WG_FACE="wg0"                        # WG NIC
SUB_NET="10.1.2.0/24"                # WG IPv4 sub/net aka CIDR
WG_PORT="51820"                      # WG udp port
SUB_NET_6="fda4:a774:9757::1:0/112"  # WG IPv6 sub/net

## IPv4 rules ##
# Route traffic to the internet for wireguard clients
$IPT -t nat -I POSTROUTING -s $SUB_NET -o $IN_FACE -j MASQUERADE
# Accept packets targeted to wg0
$IPT -I INPUT -i $WG_FACE -j ACCEPT
# Allow for packets routed through the wireguard server
$IPT -I FORWARD -i $IN_FACE -o $WG_FACE -j ACCEPT
$IPT -I FORWARD -i $WG_FACE -o $IN_FACE -j ACCEPT
# UDP for new connections
$IPT -I INPUT -i $IN_FACE -p udp --dport $WG_PORT -j ACCEPT

# IPv6 rules (uncomment) #
## $IPT6 -t nat -I POSTROUTING -s $SUB_NET_6 -o $IN_FACE -j MASQUERADE
## $IPT6 -I INPUT -i $WG_FACE -j ACCEPT
## $IPT6 -I FORWARD -i $IN_FACE -o $WG_FACE -j ACCEPT
## $IPT6 -I FORWARD -i $WG_FACE -o $IN_FACE -j ACCEPT
