EOSIO Elliptic Curve Cryptography Wrapper
=========================================

This is a simple wrapper for `btcec`, that handles the specificities
of the format for keys in YTA.

It was crafted in reference to `eosjs-ecc`, `eosjs-keygen` and the C++
codebase of YTA.IO Software.

This handles the `YTA` prefix on public keys, manages the version and
checksums in public and private keys.
