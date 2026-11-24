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
validator selection as CI. It checks workflows, action pinning policy, configuration,
Markdown, spelling, shell scripts, Dockerfiles and infrastructure security.
Language lint runs through the project-owned tools in `make lint-code` and the shared
language CI, so dependency resolution and tool configuration stay consistent.

## Continuous integration

Pull requests, merge queues and main-branch pushes run the shared language checks
and repository linter. A weekly run detects newly disclosed dependency vulnerabilities.
Shared CI reports consolidate native reports after the producer jobs finish; producer
failures continue to block validation. The semantic PR workflow and repository linter
share [.commitlintrc.json](../../.commitlintrc.json).

The `ci-github-go` references intentionally use `main` during adoption.
The pinning policy permits that repository only; other external actions use reviewed
SHAs with version comments. Replace those temporary branch references with reviewed
release pins after adoption.

Dependabot groups weekly updates for actions, language dependencies, Docker images
and devcontainer features, with a seven-day cooldown and Conventional Commit titles.
Review the lockfile and run `make ci` when updating dependencies or features.

## Image and chart checks

Tests run once with race detection and atomic coverage. Coveralls receives that same
coverage profile. Image builds wait for code and repository checks; container
health/authentication and Helm installation checks validate the built digest before
release promotion. Chart publication and Docker Hub mirroring retain their existing
destinations and credentials.

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
