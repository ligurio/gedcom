# gedcom2sql

[![Build Status](https://travis-ci.org/ligurio/gedcom2sql.svg?branch=master)](https://travis-ci.org/ligurio/gedcom2sql)

This module allows to convert of data in GEDCOM format to SQL. Gedcom is a
format for storing genealogical information designed by The Church of Jesus
Christ of Latter-Day Saints (http://www.lds.org).

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

Copyright (c) 2017-2018, Sergey Bronnikov sergeyb@bronevichok.ru
