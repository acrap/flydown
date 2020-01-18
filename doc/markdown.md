# Markdown format

Markdown is a human-readable text format that also can be rendered to popular formats like: pdf, epub, html and etc.

So, these features make markdown a very handful for using in documentation, wiki pages and knowledge bases. Also, it's great with CVS, because it's a text format.

`flydown` support all popular markdown features, like images, tables and etc.

## Subchapters in summary

Use tabulation to create subchapters. For example:

```markdown
* [Chapter](chapter.md)
    * [Subchapter](subchapter.md)
```

* [Chapter](chapter.md)
    * [Subchapter](subchapter.md)

And even more nesting:

```markdown
* [Chapter](chapter.md)
    * [Subchapter](subchapter.md)
        * [Subchapter](subchapter.md)
```

* [Chapter](chapter.md)
    * [Subchapter](subchapter.md)
        * [Subchapter](subchapter.md)

> It shows differently in summary, because of a special stylesheet.

## Headers

Just use `#` symbol as in the following example:

```markdown
# H1
## H2
### H3
#### H4
```

# H1
## H2
### H3
#### H4


## Links

```markdown
[google](http://google.com)
```

## Image

```markdown
![flydown logo](https://i.ibb.co/sq28KyP/flydown-logo-small.png)
```

![flydown logo](https://i.ibb.co/sq28KyP/flydown-logo-small.png)

## More

Markdown cheatsheet can be found by [the link](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet).

## Reference

Also, you can use flydown documentation as a reference because it's written on Markdown itself.