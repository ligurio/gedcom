# GEDCOM Tools

[![Build
Status](https://travis-ci.org/ligurio/gedcom-tools.svg?branch=master)](https://travis-ci.org/ligurio/gedcom-tools)

Gedcom is a format for storing genealogical information designed by The Church
of Jesus Christ of Latter-Day Saints (http://www.lds.org).

## `gedcom2sql.go`

- https://gramps-project.org/wiki/index.php/Gramps_SQL_Database

This module allows to convert of data in GEDCOM format to SQL.

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

## `gedcom-analysis.go`

Проверка файла GEDCOM на ошибки.

```
gedcom-analysis -verbose -file samples/tsars.ged
```

### `vk2gedcom.go`

Приложение строит генеалогическое дерево на основании данных из социальных
сетей. Пока реализована поддержка только для ВКонтакте.

### `vis.go`

Визуализация GEDCOM

* [Tracing Genealogical Data with TimeNets](http://vis.stanford.edu/papers/timenets)
* [Family Tree Visualization](http://vis.berkeley.edu/courses/cs294-10-sp10/wiki/images/f/f2/Family_Tree_Visualization_-_Final_Paper.pdf)
* Gramps: [GEPS 030: New Visualization Techniques](https://www.gramps-project.org/wiki/index.php/GEPS_030:_New_Visualization_Techniques)
* [Geneaquilts](https://aviz.fr/geneaquilts/)
* [familytreemaker](https://github.com/adrienverge/familytreemaker)
* https://www.ctan.org/pkg/genealogytree
* https://github.com/bencrowder/emperor
* https://github.com/vmiklos/ged2dot
- https://github.com/nicolaskruchten/genealogy

### `pedigree.ipynb`

- https://jupyter.brynmawr.edu/services/public/dblank/jupyter.cs/Genealogy/Gramps%205.0,%20Getting%20Started.ipynb
- https://github.com/brad-do/query-gen-dbs
- https://bencrowder.net/blog/2013/genealogy-notebook-proof-of-concept/
- http://nicolas.kruchten.com/content/2015/08/family-trees/
- https://dadoverflow.com/2018/07/05/roots-magic-and-jupyter-notebook-like-peas-and-carrots/

Copyright (c) 2017-2019, [Sergey Bronnikov](https://bronevichok.ru/)
