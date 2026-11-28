# Development

## Environment

Open this repository in Visual Studio Code and select "Reopen in Container". The devcontainer
installs Go tooling, Docker and the GitHub CLI, then runs `make setup`.
The committed feature lockfile makes the feature resolution reproducible.
Runtime and tool versions live in [devcontainer.json](../../.devcontainer/devcontainer.json).

For a host setup, install Docker and the same language tools, then run `make setup`.
Use the supported Go toolchain and golangci-lint declared in the devcontainer. `go.mod` declares the minimum language version.

## Commands

| Command         | Purpose                                           |
| --------------- | ------------------------------------------------- |
| `make setup`    | Install or verify dependencies                    |
| `make lint`     | Check application code and repository files       |
| `make lint-fix` | Apply supported formatting fixes                  |
| `make test`     | Run tests and collect coverage                    |
| `make audit`    | Check production dependencies for vulnerabilities |
| `make build`    | Build locally without publishing                  |
| `make ci`       | Run the complete local validation sequence        |

`make lint-repository` uses the same pinned Hoverkraft Super-Linter image and
tool configuration as CI, without a repository validator allowlist. It checks
workflows, action pinning policy, configuration,
Markdown, spelling, shell scripts, Dockerfiles and infrastructure security.
Language lint runs through the project-owned tools in `make lint-code` and the shared
language CI, so dependency resolution and tool configuration stay consistent.
Super-Linter checks Go recursively through `GO_MODULES`; its standalone-file mode
cannot combine files from multiple packages and is disabled. `GOTOOLCHAIN=local`
keeps its bundled compiler compatible with its golangci-lint binary. The dedicated
Go CI still uses the project toolchain.

## Continuous integration

Pull requests, merge queues and main-branch pushes run the shared language checks
and repository linter. A weekly run detects newly disclosed dependency vulnerabilities.
Shared CI reports consolidate native reports after the producer jobs finish; producer
failures continue to block validation. The semantic PR workflow and repository linter
share [.commitlintrc.json](../../.commitlintrc.json).

External actions and reusable workflows use reviewed release SHAs with matching
version comments. Dependabot maintains these references. The pending shared workflow
capabilities use immutable candidate commits with linked TODOs; replace those pins
with the first containing release after the supporting PRs merge.

Dependabot is the primary updater for Go modules, GitHub Actions, Docker images
and devcontainer features, with weekly groups and a seven-day cooldown.
Renovate is limited to versions outside those supported manifests: Go and
golangci-lint feature options, and the govulncheck version embedded in the Makefile.
It also maintains the explicit Go `toolchain` directive and the Helm CLI in
`.tool-versions`. It cannot update module dependencies or raise the minimum Go version.
The scheduled [Renovate workflow](../../.github/workflows/__renovate.yml) uses
`CI_BOT_APP_CLIENT_ID` and `CI_BOT_APP_PRIVATE_KEY`; its
[configuration](../../.github/renovate.json) explicitly limits enabled managers.
Review the lockfile and run `make ci` when updating dependencies or features.

## Image and chart checks

Tests run on Go 1.26 and 1.27 with race detection and detailed JUnit results for
both runtimes. Lint, audit, build and atomic coverage run once on Go 1.27. The
shared Go action retains the native coverage profile and validates source paths.

One terminal `reports` job waits for all code, container and Helm checks. It
publishes annotations and a consolidated summary with expandable details.
Coverage uses GitHub Code Quality when available, otherwise the pull request
comment, and is also submitted to Coveralls. Native Go coverage is shown as
statement coverage; a normalized LCOV projection supplies Coveralls without
changing the retained profile.

Helm validation and publication
use the dedicated container workflows. The Harbor robot username and password
remain secrets passed only to chart publication. Image builds wait for code and repository checks; container
health/authentication and Helm installation checks validate the built digest before
release promotion. The shared `docker-promote-images.yml` publishes the validated
digest to GHCR and the explicitly configured Docker Hub destination in one locked,
journaled operation. It preserves and verifies the copied signatures and provenance.
Chart publication waits for both registries to succeed. Registry destinations and
credentials remain unchanged.

## Container and chart policy

The image runs as UID/GID 10001 and has an HTTP health check. The chart disables
service account token mounting, drops capabilities, uses a read-only root filesystem,
sets resource budgets and denies outbound traffic. Adjust the budgets for your workload.

Raw YAML formatters exclude Helm Go templates; the chart checks and security scanners
validate their rendered manifests. Scoped scanner annotations document Helm namespace
and image inputs, plus the existing environment-based authentication interface.
CI installs the built image by digest. Deploy into a dedicated namespace and pass a
strong authentication value explicitly.

Chart validation uses the synthetic token in
[ci/default-values.yaml](../../helm/chart/ci/default-values.yaml). Production installs
must provide `securityHeader.value`; the chart has no default authentication token.

The existing GitHub CodeQL default setup continues to scan Go. The shared
linter disables its additional CodeQL job because GitHub rejects advanced-setup
results while default setup is enabled. Actionlint, Zizmor and Checkov still check
workflow correctness and security.
