# Praxis client for Linux

This archive is for distributions without a Praxis package (Gentoo, Void,
Alpine, Slackware, ...). On Debian, Ubuntu, Fedora, openSUSE or Arch, install
the deb, rpm or Arch package instead: those update themselves from the app.

## Requirements

- x86_64 or ARM64, glibc 2.35 or newer (musl systems need a glibc
  compatibility layer)
- GTK 3 and WebKitGTK 4.1 (Gentoo: `net-libs/webkit-gtk:4.1`; Void:
  `libwebkit2gtk41`)
- systemd or OpenRC to run the background service, or start
  `praxis-client --daemon` as root yourself

## Install

    tar xzf praxis-client-<arch>-<version>.tar.gz
    cd praxis-client-<arch>-<version>
    sudo ./install.sh

This installs to `/usr/local` (set `PREFIX` to change it) and starts the
background service. Open Praxis from your applications menu. On GNOME, log
out and back in, then run `gnome-extensions enable praxis@causq.com` to
show Praxis in the top bar.

Run `./install.sh --help` for the options, including `DESTDIR` for building a
distribution package from this archive.

## Update

The app tells you when a new version is out but can't install it from this
archive. Download the new archive and run its `install.sh` again: it replaces
the files and restarts the service, keeping your sign-ins.

## Uninstall

    sudo /usr/local/lib/praxis-client/uninstall.sh

Your sign-ins and logs stay in `/var/lib/praxis-client` and
`/var/log/praxis-client`; delete those too to remove them.

## What's inside

| File | Installed to |
| --- | --- |
| `bin/praxis-client` | `$PREFIX/bin/` |
| `share/applications/praxis-client.desktop` | `$PREFIX/share/applications/` |
| `share/icons/hicolor/256x256/apps/praxis-client.png` | `$PREFIX/share/icons/hicolor/256x256/apps/` |
| `share/gnome-shell/extensions/praxis@causq.com/` | `$PREFIX/share/gnome-shell/extensions/` |
| `share/polkit-1/actions/com.causq.praxis.client.policy` | `/usr/share/polkit-1/actions/` |
| `etc/xdg/autostart/praxis-client-autostart.desktop` | `/etc/xdg/autostart/` |
| `lib/systemd/system/praxis-client.service` | `/etc/systemd/system/` (systemd) |
| `etc/init.d/praxis-client` | `/etc/init.d/` (OpenRC) |
