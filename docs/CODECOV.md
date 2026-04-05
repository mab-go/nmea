# Codecov Setup

CI uploads coverage from GitHub Actions using
[`codecov/codecov-action@v5`](https://github.com/codecov/codecov-action) with
**OIDC** (`use_oidc: true`), so no upload token is required in the workflow once
Codecov trusts GitHub for this repository.

1. Install the [Codecov GitHub app](https://github.com/apps/codecov) for
   **`mab-go`** and grant access to **`nmea`** (maintainers: also confirm the
   repo is enabled on [codecov.io](https://codecov.io) and the default branch is
   **`main`**).
2. If uploads fail or your Codecov setup requires a repository token, add a
   **`CODECOV_TOKEN`** repository secret and follow the action README to pass
   `token: ${{ secrets.CODECOV_TOKEN }}` (and drop OIDC if you switch to
   token-based upload).

Coverage status checks are **informational** (see
[`codecov.yml`](../codecov.yml)); a failed upload step still fails CI.
