#!/bin/sh
# Runs before removal and, for deb and rpm, before an upgrade replaces the
# package. Only a real removal stops the service:
#   deb: remove | upgrade | ...    rpm: 0 on removal, 1 on upgrade
#   Arch pre_remove: <old version> (removal only)
[ -d /run/systemd/system ] || exit 0
case "$1" in
    upgrade|failed-upgrade|1) exit 0 ;;
esac
systemctl disable --now praxis-client.service >/dev/null 2>&1 || true
# Tells postinstall.sh that a later install is a reinstall, not an upgrade.
mkdir -p -m 0700 /var/lib/praxis-client && touch /var/lib/praxis-client/.removed
exit 0
