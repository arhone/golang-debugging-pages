# golang-debugging-pages
Страницы для тестирования и настройки nginx

## Установка из проекта arhone

Создание директории проекта
```
cd /srv
sudo mkdir golang-debugging-pages
sudo chown $USER:$USER golang-debugging-pages
```

Клонирование проекта
```
git clone git@github.com:arhone/golang-debugging-pages.git ./golang-debugging-pages
```

## Установка из своего проекта после форка

Создания deploy ключа для git

Имя файла указать golang-debugging-pages (первый вопрос в ssh-keygen)
```
cd ~/.ssh/
ssh-keygen
```

```
nano config
```
Добавить в файл config
```
Host golang-debugging-pages.github.com
    HostName github.com
    IdentityFile ~/.ssh/golang-debugging-pages
```

Скопировать и добавить ключ в git, в ваш репозиторий
```
cat ~/.ssh/golang-debugging-pages.pub
```

Создание директории проекта
```
cd /srv
sudo mkdir golang-debugging-pages
sudo chown $USER:$USER golang-debugging-pages
```

Клонирование проекта
```
git clone git@golang-debugging-pages.github.com:ВАШ_ПРОЕКТ/golang-debugging-pages.git /srv/golang-debugging-pages
```

## Настройка
Копирование общего примера файла настроек в локальную версию файла
```
cd golang-debugging-pages/
cp .env.example .env
```

Сборка и запуск контейнера

Используется docker compose v2 [Установка](https://github.com/arhone/debian-server-guide/blob/main/docker.md)
```
sudo docker compose -f docker-compose.yml up -d --build --remove-orphans
```

## Управление
Войти в контейнер
```
sudo docker exec -it golang-debugging-pages-01 /bin/sh
```

Остановка/Запуск контейнера
```
sudo docker stop golang-debugging-pages-01
sudo docker start golang-debugging-pages-01
```

Просмотр логов контейнера
```
sudo docker logs --tail 50 --follow --timestamps golang-debugging-pages-01
```

## Deploy
Загрузка и разворачивание проекта на удалённом сервере
```
. deploy.sh username@example.com
```

### Разрешить команду sudo docker без подтверждения паролем
```
sudo visudo
```
Добавить запись
```
username ALL=NOPASSWD: /usr/bin/docker
```
