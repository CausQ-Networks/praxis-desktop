# Praxis desktop CI signing secrets

## Windows MSI (`windows.yml`) — DigiCert ONE / KeyLocker via jsign

Set on **CausQ-Networks/praxis-desktop** (or org-wide):

| Secret | What it is |
|--------|------------|
| `SM_API_KEY` | DigiCert ONE / Software Trust Manager API key |
| `SM_CLIENT_CERT_FILE_B64` | Base64 of the DigiCert client auth `.p12` |
| `SM_CLIENT_CERT_PASSWORD` | Password for that `.p12` |
| `UPDATE_SIGNING_KEY` | signify `release.sec` private key (update manifests) |

Optional **variable** (not secret):

| Variable | Default |
|----------|---------|
| `SM_KEY_ALIAS` | DigiCert certificate ID (workflow has a placeholder UUID) |

Encode the p12:

```bash
base64 -i Certificate_pkcs12.p12 | pbcopy   # macOS
```

Set secrets (values from env, never commit):

```bash
gh secret set SM_API_KEY --repo CausQ-Networks/praxis-desktop
gh secret set SM_CLIENT_CERT_FILE_B64 --repo CausQ-Networks/praxis-desktop
gh secret set SM_CLIENT_CERT_PASSWORD --repo CausQ-Networks/praxis-desktop
gh secret set UPDATE_SIGNING_KEY --repo CausQ-Networks/praxis-desktop
```

Then tag: `git tag windows-0.1.0 && git push origin windows-0.1.0`

## Linux (`linux.yml`)

| Secret | What it is |
|--------|------------|
| `UPDATE_SIGNING_KEY` | same signify key as Windows |
| `AUR_SSH_PRIVATE_KEY` | only if publishing to AUR |

## Apple DMG

`praxis-apple` has **no** GitHub Actions release workflow yet (local Xcode / dmgcanvas).
For notarized DMG CI later you typically need:

- `APPLE_ID` / `APPLE_TEAM_ID` / `APPLE_APP_PASSWORD` (or App Store Connect API key)
- `MACOS_CERTIFICATE` + `MACOS_CERTIFICATE_PWD` (Developer ID Application .p12, base64)
- `MATCH_*` or manual provisioning if using Fastlane

## GHCR public packages

Images live at `ghcr.io/causq-networks/{praxis-gerbil,praxis-newt,praxis-node}`.
Until API visibility works, in GitHub UI:

1. https://github.com/orgs/CausQ-Networks/packages
2. Open each package → Package settings → Change visibility → Public
