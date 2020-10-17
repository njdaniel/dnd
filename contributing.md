# Contributing to dnd
These are mostly guidelines to contritubing to this project.
Use your best judgement and feel free to propose changes to this document in a 
pull request.

#### Table of Contents


[Where to ask Questions](#i-have-a-question)

[What to know before I get started](#what-to-know-before-i-get-started)

[How can I contribute?](#how-can-i-contribute)
 * [Reporting bugs](#reporting-bugs)
 * [Suggesting Enhancements](#suggesting-enhancements)
 * [Your First Code Contribution](#your-first-code-contribution)
 * [Pull Requests](#pull-requests)
 
[Styleguides](#styleguides)


## Where to ask Questions

Join slack: 
https://join.slack.com/t/rtp-gophers/shared_invite/enQtNTE3NjIyMTgyODgyLTRhOTcxODBlNjc3NGYxNTI4Mzg5Mzg4OTY5ZjVmOGQ4ZDIyZjIxMzUxZDJlNjc0MzNmZjM2MmI3YmRlMzFjOTk
Ask question in the #go-projects channels.

## What to know before I get started

### Dependency management

Project uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages. This requires a working Go environment with version 1.13 or greater installed.

All dependencies are vendored in the `vendor/` directory.

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

Tidy up the `go.mod` and `go.sum` files and copy the new/updated dependency to the `vendor/` directory:


```bash
# The GO111MODULE variable can be omitted when the code isn't located in GOPATH.
GO111MODULE=on go mod tidy

GO111MODULE=on go mod vendor
```

You have to commit the changes to `go.mod`, `go.sum` and the `vendor/` directory before submitting the pull request.
 
 
## How can I contribute?

### Steps to Contribute
Should you wish to work on an issue, please claim it first by commenting on the GitHub issue that you want to work on it. 
This is to prevent duplicated efforts from contributors on the same issue.

Fork project and create new branch.
Create and reference a 'Draft Pull Request' in the Issue and mark with WIP until ready to be merged.  

For quick compile and testing
```
#running the project
go run main.go

#run tests
make test
```

### Reporting Bugs
Follow this guideline to help contributors to 
understand your report, reproduce the behavior, and find related reports.

Before creating a bug report check under Issues to see if bug is already 
created.

When creating a bug report please include as many details as possible.
Fill out the [required template](https://github.com/dnd/.github/blob/master/.github/ISSUE_TEMPLATE/bug_report.md).

### Suggesting Enhancements
This section guides you through submitting an enhancement suggestion for Esim, including completely new features and minor improvements to existing functionality. Following these guidelines helps maintainers and the community understand your suggestion :pencil: and find related suggestions :mag_right:.

#### Before Submitting An Enhancement Suggestion

* **Check the [debugging guide](https://flight-manual.dnd.io/hacking-dnd/sections/debugging/)** for tips — you might discover that the enhancement is already available. 
* **Check if there's already which provides that enhancement.
* **Determine [which repository the enhancement should be suggested in]
* **Perform a [cursory search] to see if the enhancement has already been suggested. If it has, add a comment to the existing issue instead of opening a new one.

#### How Do I Submit A (Good) Enhancement Suggestion?

Enhancement suggestions are tracked as [GitHub issues](https://guides.github.com/features/issues/). After you've determined [which repository](#dnd-and-packages) your enhancement suggestion is related to, create an issue on that repository and provide the following information:

* **Use a clear and descriptive title** for the issue to identify the suggestion.
* **Provide a step-by-step description of the suggested enhancement** in as many details as possible.
* **Provide specific examples to demonstrate the steps**. Include copy/pasteable snippets which you use in those examples, as [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines).
* **Describe the current behavior** and **explain which behavior you expected to see instead** and why.
* **Include screenshots and animated GIFs** which help you demonstrate the steps or point out the part of dnd which the suggestion is related to. You can use [this tool](https://www.cockos.com/licecap/) to record GIFs on macOS and Windows, and [this tool](https://github.com/colinkeenan/silentcast) or [this tool](https://github.com/GNOME/byzanz) on Linux.
* **Explain why this enhancement would be useful** to most dnd users and isn't something that can or should be implemented as a [community package](#dnd-and-packages).
* **List some other text editors or applications where this enhancement exists.**
* **Specify which version of dnd you're using.** You can get the exact version by running `dnd -v` in your terminal, or by starting dnd and running the `Application: About` command from the [Command Palette](https://github.com/dnd/command-palette).
* **Specify the name and version of the OS you're using.**


### Pull Request Checklist

* Branch from the master branch and, if needed, rebase to the current master branch before submitting your pull request. If it doesn't merge cleanly with master you may be asked to rebase your changes.

* Commits should be as small as possible, while ensuring that each commit is correct independently (i.e., each commit should compile and pass tests).

* If your patch is not getting reviewed or you need a specific person to review it, you can @-reply a reviewer asking for a review in the pull request or a comment, or you can ask for a review on slack channel.

* Add tests relevant to the fixed bug or new feature.


