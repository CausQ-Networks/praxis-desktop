#!/bin/sh
# deb "purge" also removes saved secrets, state and logs.
[ -d /run/systemd/system ] && systemctl daemon-reload >/dev/null 2>&1
if [ "$1" = purge ]; then
    rm -rf /var/lib/praxis-client /var/log/praxis-client /etc/praxis-client
fi
exit 0
