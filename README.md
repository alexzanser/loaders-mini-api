# loaders

Запуск:  
make build

Зарегистрировать заказчика и залогиниться:  
    make customer_register  
    make customer_login

Зарегистрировать три грузчика и залогиниться ими:  
    make loaders_register  
    make loaders_login

Создать случайны набор заказов для каждого заказчика:  
    make generate_tasks

Все остальные запросы пока что отправляются в ручную, нужно подставлять  
токены, полученные при авторизации:  

 - получить информацию о себе:  
    `curl -H "Authorization: Bearer <user_token>" -X GET http://localhost:8080/me`

 - получить информацию о заказах:  
    `curl -H "Authorization: Bearer <user_token>" -X GET http://localhost:8080/tasks`

 - начать игру заказчиком, loaders - список id выбранных грузчиков:  
    `curl -d "loaders=1,2,3" -H "Authorization: Bearer <customer_token>" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/start`

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
