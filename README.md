# cbyt
cbyt : compress, encrypt, and publish data as a single archive file.

A CLI that packs files into encrypted, chunked payloads, uploads to IPFS/pinning services, and produces a small manifest that is the only secret needed to reconstruct the data.

## Quickstart (dev)
Install Go 1.20+ and make.

```bash
git clone https://github.com/cbyt-cli/cbyt.git
cd cbyt
make build
```

