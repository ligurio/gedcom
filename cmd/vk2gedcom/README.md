### `vk2gedcom.go`

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
