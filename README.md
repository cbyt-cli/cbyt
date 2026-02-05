# cbyt

`cbyt` — compress, encrypt, and publish byte chunks (IPFS-ready) as a single manifest.

Short: a CLI that packs files into encrypted, chunked payloads, uploads to IPFS/pinning services, and produces a small manifest that is the only secret needed to reconstruct the data.

## Quickstart (dev)
Install Go 1.20+ and make.

```bash
git clone https://github.com/YOUR_GITHUB_ORG/cbyt.git
cd cbyt
make build

