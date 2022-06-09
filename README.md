# loaders

```
Мини-игра грузчики.
Есть заказчик, есть грузчики. Заказчику необходимо переносить тяжелые грузы. 
Заказчик обладает следующими свойствами:
- стартовый капитал (10 000р - 100 000р)
- умение нанимать грузчика
- набор заданий, которые нужно выполнить
Каждый грузчик обладает следующими свойствами: 
- максимально переносимый вес (5кг-30кг)
- "пьянство" (true,false)
- усталость (0-100%)
- зарплата (10 000р - 30 000р)

Суть следующая: генерируется N случайных заданий ("название переносимых предметов", "вес"). 
Есть грузчики зарегистрировавшиеся на работу и получившие случайные свойства. 
Задача заказчика - выбрать нужный набор грузчиков и выполнить задания.

Технически:
Авторизация и регистрация пользователей - в графе они выбирают свойство "заказчик" или "грузчик".

По API:
публичное:
/login - вход по логину и паролю
/register - регистрация логин пароль
/tasks - создание случайного набора заданий

 под авторизацией для грузчика:
/me - показать свои характеристики (вес, деньги, пьянство, усталость)
/tasks - показать список выполненных заданий

 под авторизацией для заказчика:
/me - показать свои характеристики (деньги, зарегистрировавшиеся грузчики)
/tasks - показать список доступных заданий
/start - добавить грузчиков и начать выполнение задания (списываются деньги, рассчитывается выполнено задание или нет)
```
