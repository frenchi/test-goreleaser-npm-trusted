# Testing goreleaser's npm Trusted Publishing flow

Testing for [Goreleaser-pro Issue #33](https://github.com/goreleaser/goreleaser-pro/discussions/33) against [@frenchi/test-goreleaser-npm-trusted](https://www.npmjs.com/package/@frenchi/test-goreleaser-npm-trusted) and [@frenchi/test-goreleaser-npm-trusted-control](https://www.npmjs.com/package/@frenchi/test-goreleaser-npm-trusted)

## Preconditions

- valid `GORELEASER_KEY` set in all cases
- For base case (`v0.0.1`) no token or OIDC configuration exists, expect failure.
- For standard publishing flow (`v0.0.2` & `v0.0.3`) `NPM_TOKEN` exists with package publishing permissions, no NPM Trusted Publishers config.
- For OIDC scenario (`v0.0.4`): remove `NODE_AUTH_TOKEN` or `NPM_TOKEN` from the github environment secrets. NPM Trusted Publishers configured for the repo.

from: https://www.npmjs.com/package/@frenchi/test-goreleaser-npm-trusted/access

![Publishing access](publishing_access.png)

## Testing matrix

Run a tag-triggered release for each scenario and capture logs.

| Tag    | Scenario                                                           | NPM_TOKEN | Publishing access (package setting)                                         | Setup (summary)                                                              |
| ------ | ------------------------------------------------------------------ | --------- | --------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| v0.0.1 | Any; no token (standard flow)                                      | No        | Any                                                                         | Do not set `NPM_TOKEN`.                                                      |
| v0.0.2 | Tokens allowed; automation token present                           | Yes       | Require two-factor authentication or an automation or granular access token | Package allows tokens; export an automation token (e.g., `NPM_TOKEN`).       |
| v0.0.3 | Tokens disallowed; automation token present; No Trusted Publishing | Yes       | Require two-factor authentication and disallow tokens (recommended)         | Package: Require 2FA and disallow tokens; keep `NPM_TOKEN` set.              |
| v0.0.4 | Tokens disallowed; no token (OIDC)                                 | No        | Require two-factor authentication and disallow tokens (recommended)         | Remove token envs; rely on OIDC; ensure preconditions and Trusted Publisher. |

Control runs: for each tag above, also cut a matching `-control` tag (e.g., `v0.0.1-control`) to trigger a direct npm publish via `control.yml`. For GoReleaser runs, cut matching `-goreleaser` tags trigger the testing workflow: release.yml

## Results

| Tag    | Expected outcome                                                                         | Control run                                                                                             | Goreleaser outcome                                                                                      | Test passed? |
| ------ | ---------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------ |
| v0.0.1 | Publish fails with ENEEDAUTH (no auth provided).                                         | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17902305355/job/50897305627) | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17902821877/job/50898605291) | ✅           |
| v0.0.2 | Publish succeeds via token (baseline).                                                   | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903146406)                 | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903373304/job/50900148300) | ✅           |
| v0.0.3 | 403 error: 2FA required but an automation token was specified, and no Trusted Publishing | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903562298)                 | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903562329/job/50900839382) | ✅           |
| v0.0.4 | Publish succeeds; provenance shown for public repo + public package.                     | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903842964/job/50901612810) | [logs](https://github.com/frenchi/test-goreleaser-npm-trusted/actions/runs/17903685543/job/50901169457) | ✅           |

Successfully published: https://www.npmjs.com/package/@frenchi/test-goreleaser-npm-trusted/v/0.0.4-goreleaser

## References

- npm Trusted Publishing (OIDC): https://docs.npmjs.com/trusted-publishers
- GoReleaser npm pipe: https://goreleaser.com/customization/npm/
- GoReleaser on GitHub Actions: https://goreleaser.com/ci/actions/
- Discussion: https://github.com/goreleaser/goreleaser-pro/discussions/33
