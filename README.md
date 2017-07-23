# gedcom2sql

This module allows to convert of data in GEDCOM format to SQL. Gedcom is a
format for storing genealogical information designed by The Church of Jesus
Christ of Latter-Day Saints (http://www.lds.org).

```sql
.print "Total number of persons:"
SELECT COUNT(*) FROM person_st;
.print "Total number of childs:"
SELECT COUNT(*) FROM famchild;
.print "Total number of families:"
SELECT COUNT(*) FROM family;
.print "Total number of male persons:"
SELECT COUNT(*) FROM person_st WHERE sex='M';
.print "Total number of female persons:"
SELECT COUNT(*) FROM person_st WHERE sex='F';
.print "Total number of death dates:"
SELECT COUNT(*) FROM person_st WHERE deat_date IS NOT "";
.print "Total number of burial dates:"
SELECT COUNT(*) FROM person_st WHERE buri_date IS NOT "";
.print "Total number of birth dates:"
SELECT COUNT(*) FROM person_st WHERE birt_date IS NOT "";
```

Copyright (c) 2017, Sergey Bronnikov sergeyb@bronevichok.ru
