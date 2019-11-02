# GEDCOM Tools

[![Build
Status](https://travis-ci.org/ligurio/gedcom-tools.svg?branch=master)](https://travis-ci.org/ligurio/gedcom-tools)

Gedcom is a format for storing genealogical information designed by The Church
of Jesus Christ of Latter-Day Saints (http://www.lds.org).

## `gedcom2sql.go`

This module allows to converting data in GEDCOM format to SQLite DB.

```sql
$ gedcom2sql -file sergeyb.gedcom
$ sqlite3 sergeyb.sqlite
SQLite version 3.20.1 2017-08-24 16:21:36
Enter ".help" for usage hints.

sqlite> -- "Total number of persons:"
sqlite> SELECT COUNT(*) FROM person_st;
sqlite> -- "Total number of childs:"
sqlite> SELECT COUNT(*) FROM famchild;
sqlite> -- "Total number of families:"
sqlite> SELECT COUNT(*) FROM family;
sqlite> -- "Total number of male persons:"
sqlite> SELECT COUNT(*) FROM person_st WHERE sex='M';
sqlite> -- "Total number of female persons:"
sqlite> SELECT COUNT(*) FROM person_st WHERE sex='F';
sqlite> -- "Total number of death dates:"
sqlite> SELECT COUNT(*) FROM person_st WHERE deat_date IS NOT "";
sqlite> -- "Total number of burial dates:"
sqlite> SELECT COUNT(*) FROM person_st WHERE buri_date IS NOT "";
sqlite> -- "Total number of birth dates:"
sqlite> SELECT COUNT(*) FROM person_st WHERE birt_date IS NOT "";
```

- [Gramps_SQL_Database](https://gramps-project.org/wiki/index.php/Gramps_SQL_Database)

## `gedcom2errors.go`

Проверка файла GEDCOM на ошибки.

```
gedcom2errors -verbose -file samples/tsars.ged
```

## `vk2gedcom.go`

Приложение строит генеалогическое дерево на основании данных из социальных
сетей. Пока реализована поддержка только для ВКонтакте.

## `gedcom2timenet.go`

Визуализация GEDCOM.

* [Javascript script to embed pedigree charts in web pages](https://github.com/bencrowder/emperor)
* [Pedigree and genealogical tree diagrams with LaTeX](https://www.ctan.org/pkg/genealogytree)

## `gedcom2report.go`

Отчёты для GEDCOM.

* [gedcom2html](https://godoc.org/github.com/elliotchance/gedcom/gedcom2html) ([example](https://gedcom.app/royals/places.html))
* [gedcom2json](https://godoc.org/github.com/elliotchance/gedcom/gedcom2json)
* [gedcom2text](https://godoc.org/github.com/elliotchance/gedcom/gedcom2text)
* [gedcomdiff](https://godoc.org/github.com/elliotchance/gedcom/gedcomdiff)

## `pedigree.ipynb`

- https://jupyter.brynmawr.edu/services/public/dblank/jupyter.cs/Genealogy/Gramps%205.0,%20Getting%20Started.ipynb
- https://github.com/brad-do/query-gen-dbs
- https://bencrowder.net/blog/2013/genealogy-notebook-proof-of-concept/
- http://nicolas.kruchten.com/content/2015/08/family-trees/
- https://dadoverflow.com/2018/07/05/roots-magic-and-jupyter-notebook-like-peas-and-carrots/
- [gedcomq](https://godoc.org/github.com/elliotchance/gedcom/gedcomq)


Copyright (c) 2017-2019, [Sergey Bronnikov](https://bronevichok.ru/)
