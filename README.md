# m2j

This library will convert Markdown into [Jira's text formatting](https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=all).
It is a simple regex-based library built to help migrate Github issues into Jira, and it is therefore
assuming [Github-flavored Markdown](https://github.github.com/gfm/), and the focus is on the styling
frequently seen in Github issues.

# Progress
- [x] Headings (h1-h6)
- [x] Basic text formatting (bold/italics/mono)
- [x] Code blocks
- [ ] Unordered lists
- [ ] Ordered lists
- [ ] Tables

# Credits
* https://github.com/StevenACoffman/j2m for the basic structure that I based this library on.
* https://github.com/kylefarris/J2M for a nice source of Markdown to Jira regexen.
