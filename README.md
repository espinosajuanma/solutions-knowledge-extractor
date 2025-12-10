# Solutions Knowledge Extractor

## Install

You can just grab the latest binary [release](https://github.com/espinosajuanma/solutions-knowledge-extractor/releases).

This command can be installed as a standalone program.

Standalone

```bash
go install github.com/espinosajuanma/solutions-knowledge-extractor/cmd/solutions-knowledge-extractor@latest
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C solutions-knowledge-extractor solutions-knowledge-extractor
```

If you don't have bash or tab completion check use the shortcut
commands instead.

