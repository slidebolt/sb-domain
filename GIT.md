# Git Workflow for sb-domain

This repository contains the core Slidebolt domain models and shared logic. It is a fundamental library with zero internal Slidebolt dependencies.

## Dependencies
- **Internal:** None. This is a base repository.
- **External:** Standard Go library.

## Build Process
- **Type:** Pure Go Library (Shared Module).
- **Consumption:** Imported as a module dependency in other Go projects via `go.mod`.
- **Artifacts:** No standalone binary or executable is produced.
- **Validation:** 
  - Validated through unit tests: `go test -v ./...`
  - Validated by its consumers during their respective build/test cycles.

## Pre-requisites & Publishing
As a base repository, `sb-domain` must be updated and published early in the change cycle.

**Before publishing:**
1. Determine current tag: `git tag | sort -V | tail -n 1`
2. Ensure all local tests pass: `go test -v ./...`
3. If changing domain logic, check impact on major consumers like `sb-api`, `sb-manager`, and `plugin-*` repositories.

**Publishing Order:**
1. Update `sb-domain`.
2. Determine next semantic version (e.g., `v1.0.4`).
3. Commit and push the changes.
4. Tag the repository: `git tag v1.0.4`.
5. Push the tag: `git push origin v1.0.4`.
6. Update dependent repositories using `go get github.com/slidebolt/sb-domain@v1.0.4`.

## Update Workflow & Verification
1. **Modify:** Update domain models or shared logic.
2. **Verify Local:**
   - Run `go mod tidy`.
   - Run `go test ./...`.
3. **Commit:** Ensure the commit message clearly describes the domain change.
4. **Tag & Push:** (Follow the Publishing Order above).
