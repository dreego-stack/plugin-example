# Change files

Every pull request adds exactly one Markdown file in this directory. Use a
short, unique name such as `fix-greeting-default.md`:

```markdown
---
version: patch
---

- Bug: keep greeting default consistent
```

Use `version: none` only when the release tag must not change. After merge,
the release workflow combines all pending files into one changelog entry and
removes them.