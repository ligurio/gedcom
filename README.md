# GEDCOM Tools

[![Build
Status](https://travis-ci.org/ligurio/gedcom-tools.svg?branch=master)](https://travis-ci.org/ligurio/gedcom-tools)

Gedcom is a format for storing genealogical information designed by The Church
of Jesus Christ of Latter-Day Saints (http://www.lds.org).

### Импорт файла GEDCOM в SQLite

This module allows to converting data in GEDCOM format to SQLite DB.

```sql
$ gedcom -file sergeyb.gedcom
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

### Проверка файла GEDCOM на ошибки.

```
$ gedcom -file samples/bronte.ged

Person (I0003) Maria /Brontë/
EI112: person has no family pointers
Person (I0004) Elizabeth /Brontë/
EI112: person has no family pointers
Person (I0006) Patrick Branwell /Brontë/
EI112: person has no family pointers
Person (I0007) Emily Jane /Brontë/
EI112: person has no family pointers
Person (I0008) Anne /Brontë/
EI112: person has no family pointers
Person (I0014) Elizabeth /Branwell/
EI112: person has no family pointers
Family (F001)
EP102: child is born before mother
Family (F002)
EF107: family missing pointer to child
Family (F002)
EF100: family has no members
[Brunty] Brontë
Family (F003)
EP106: child doesn't inherit father's surname
```

### Родословная из социальной сети

Приложение строит генеалогическое дерево на основании данных из социальных
сетей. Пока реализована поддержка только для ВКонтакте.

Facebook:

- https://developers.facebook.com/docs/graph-api/reference
- [Заполнить](https://www.facebook.com/help/1557948767777120): Профиль -> Информация -> Семейные отношения

VK API:

- Методы: [account.getProfileInfo](https://vk.com/dev/account.getProfileInfo), [users.get](https://vk.com/dev/users.get)
- Частотные ограничения: "К методам API ВКонтакте можно обращаться не чаще 3 раз в секунду."
- Ошибки: https://vk.com/dev/errors
- [Приложение](https://vk.com/editapp?id=6130272&section=options)
- Аккаунты для отладки:
  - [id181554010](https://vk.com/id181554010)
  - [id144273654](https://vk.com/id144273654)
  - [id233293686](https://vk.com/id233293686)
  - [denis.ivanov1988](https://vk.com/denis.ivanov1988)
  - [am.ivanov](https://vk.com/am.ivanov)
  - [id140022470](https://vk.com/id140022470)
  - [id1107203](https://vk.com/id1107203)

## Полезные инструменты

- [gedcom](https://github.com/elliotchance/gedcom/) ([example](https://gedcom.app/royals/places.html))
- [emperor](https://github.com/bencrowder/emperor)
- [genealogytree](https://www.ctan.org/pkg/genealogytree)
- [Using Prolog to Analyze your Family Tree](https://gramps-project.org/blog/2015/05/using-prolog-to-analyze-your-family-tree/)

Copyright (c) 2017-2021, [Sergey Bronnikov](https://bronevichok.ru/)
