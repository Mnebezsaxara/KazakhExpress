# Документация проекта KazakhExpress

## Обзор проекта

KazakhExpress - это маркетплейс нового поколения, который позволяет пользователям покупать товары со всего мира с быстрой доставкой по Казахстану. Проект состоит из следующих компонентов:

1. **Frontend** - веб-интерфейс на HTML, CSS и JavaScript
2. **Backend** - серверная часть на Go с использованием Gin и MongoDB
3. **База данных** - MongoDB для хранения информации о продуктах и категориях

## Архитектура проекта

Проект следует принципам чистой архитектуры:

- **Domain** - содержит бизнес-модели (Product, Category)
- **Repository** - отвечает за взаимодействие с базой данных
- **Usecase** - содержит бизнес-логику
- **Handler** - обрабатывает HTTP-запросы

## API Endpoints

### Продукты (Products)

#### 1. Создание продукта

- **Метод**: POST
- **URL**: `/products` http://localhost:8081/products
- **Описание**: Создает новый продукт в базе данных
- **Тело запроса**:
  ```json
  {
    "name": "Redmi Note 12",
    "description": "Latest Redmi smartphone with great features",
    "price": 199999,
    "image_url": "https://example.com/redmi-note-12.jpg",
    "category": "electronics",
    "stock": 10
  }
  ```
- **Ответ**: Созданный продукт с ID
  ```json
  {
    "id": "60d21b4667d0d8992e610c85",
    "name": "Redmi Note 12",
    "description": "Latest Redmi smartphone with great features",
    "price": 199999,
    "image_url": "https://example.com/redmi-note-12.jpg",
    "category": "electronics",
    "stock": 10
  }
  ```

#### 2. Получение продукта по ID

- **Метод**: GET
- **URL**: `/products/:id`
- **Описание**: Возвращает продукт по его ID
- **Параметры URL**: `id` - ID продукта
- **Ответ**: Продукт с указанным ID
  ```json
  {
    "id": "60d21b4667d0d8992e610c85",
    "name": "Redmi Note 12",
    "description": "Latest Redmi smartphone with great features",
    "price": 199999,
    "image_url": "https://example.com/redmi-note-12.jpg",
    "category": "electronics",
    "stock": 10
  }
  ```

#### 3. Обновление продукта

- **Метод**: PATCH
- **URL**: `/products/:id`
- **Описание**: Обновляет существующий продукт
- **Параметры URL**: `id` - ID продукта
- **Тело запроса**:
  ```json
  {
    "price": 189999,
    "stock": 15
  }
  ```
- **Ответ**: Сообщение об успешном обновлении
  ```json
  {
    "message": "Обновлено успешно"
  }
  ```

#### 4. Удаление продукта

- **Метод**: DELETE
- **URL**: `/products/:id`
- **Описание**: Удаляет продукт по его ID
- **Параметры URL**: `id` - ID продукта
- **Ответ**: Сообщение об успешном удалении
  ```json
  {
    "message": "Удалено успешно"
  }
  ```

#### 5. Получение списка продуктов

- **Метод**: GET
- **URL**: `/products`
- **Описание**: Возвращает список продуктов с возможностью фильтрации и пагинации
- **Параметры запроса**:
  - `category` - фильтрация по категории
  - `minPrice` - минимальная цена
  - `maxPrice` - максимальная цена
  - `page` - номер страницы (по умолчанию 1)
  - `limit` - количество продуктов на странице (по умолчанию 10)
- **Ответ**: Массив продуктов
  ```json
  [
    {
      "id": "60d21b4667d0d8992e610c85",
      "name": "Redmi Note 12",
      "description": "Latest Redmi smartphone with great features",
      "price": 199999,
      "image_url": "https://example.com/redmi-note-12.jpg",
      "category": "electronics",
      "stock": 10
    },
    {
      "id": "60d21b4667d0d8992e610c86",
      "name": "iPhone 13",
      "description": "Latest iPhone with great features",
      "price": 499999,
      "image_url": "https://example.com/iphone-13.jpg",
      "category": "electronics",
      "stock": 5
    }
  ]
  ```

### Категории (Categories)

#### 1. Создание категории

- **Метод**: POST
- **URL**: `/categories` http://localhost:8081/categories
- **Описание**: Создает новую категорию
- **Тело запроса**:
  ```json
  {
    "name": "electronics"
  }
  ```
- **Ответ**: Созданная категория с ID
  ```json
  {
    "id": "60d21b4667d0d8992e610c87",
    "name": "electronics"
  }
  ```

#### 2. Получение категории по ID

- **Метод**: GET
- **URL**: `/categories/:id`
- **Описание**: Возвращает категорию по ее ID
- **Параметры URL**: `id` - ID категории
- **Ответ**: Категория с указанным ID
  ```json
  {
    "id": "60d21b4667d0d8992e610c87",
    "name": "electronics"
  }
  ```

#### 3. Получение списка категорий

- **Метод**: GET
- **URL**: `/categories`
- **Описание**: Возвращает список всех категорий
- **Ответ**: Массив категорий
  ```json
  [
    {
      "id": "60d21b4667d0d8992e610c87",
      "name": "electronics"
    },
    {
      "id": "60d21b4667d0d8992e610c88",
      "name": "clothing"
    }
  ]
  ```

#### 4. Удаление категории

- **Метод**: DELETE
- **URL**: `/categories/:id`
- **Описание**: Удаляет категорию по ее ID
- **Параметры URL**: `id` - ID категории
- **Ответ**: Сообщение об успешном удалении
  ```json
  {
    "message": "Категория удалена"
  }
  ```

## Инструкция по тестированию API в Postman

### Настройка Postman

1. Скачайте и установите [Postman](https://www.postman.com/downloads/)
2. Создайте новую коллекцию "KazakhExpress API"
3. Убедитесь, что сервер запущен на `http://localhost:8081`

### Тестирование эндпоинтов продуктов

#### 1. Создание продукта (POST /products)

1. Создайте новый запрос в Postman
2. Выберите метод POST
3. Введите URL: `http://localhost:8081/products`
4. Перейдите на вкладку "Body"
5. Выберите "raw" и "JSON"
6. Введите тело запроса:
   ```json
   {
     "name": "Redmi Note 12",
     "description": "Latest Redmi smartphone with great features",
     "price": 199999,
     "image_url": "https://example.com/redmi-note-12.jpg",
     "category": "electronics",
     "stock": 10
   }
   ```
7. Нажмите "Send"
8. Проверьте ответ - должен вернуться созданный продукт с ID
9. Сохраните ID продукта для последующих запросов

#### 2. Получение продукта по ID (GET /products/:id)

1. Создайте новый запрос в Postman
2. Выберите метод GET
3. Введите URL: `http://localhost:8081/products/{id}` (замените {id} на ID, полученный в предыдущем шаге)
4. Нажмите "Send"
5. Проверьте ответ - должен вернуться продукт с указанным ID

#### 3. Обновление продукта (PATCH /products/:id)

1. Создайте новый запрос в Postman
2. Выберите метод PATCH
3. Введите URL: `http://localhost:8081/products/{id}` (замените {id} на ID продукта)
4. Перейдите на вкладку "Body"
5. Выберите "raw" и "JSON"
6. Введите тело запроса:
   ```json
   {
     "price": 189999,
     "stock": 15
   }
   ```
7. Нажмите "Send"
8. Проверьте ответ - должно вернуться сообщение об успешном обновлении

#### 4. Удаление продукта (DELETE /products/:id)

1. Создайте новый запрос в Postman
2. Выберите метод DELETE
3. Введите URL: `http://localhost:8081/products/{id}` (замените {id} на ID продукта)
4. Нажмите "Send"
5. Проверьте ответ - должно вернуться сообщение об успешном удалении

#### 5. Получение списка продуктов (GET /products)

1. Создайте новый запрос в Postman
2. Выберите метод GET
3. Введите URL: `http://localhost:8081/products`
4. Нажмите "Send"
5. Проверьте ответ - должен вернуться массив продуктов

#### 6. Фильтрация продуктов

1. Создайте новый запрос в Postman
2. Выберите метод GET
3. Введите URL: `http://localhost:8081/products?category=electronics&minPrice=100000&maxPrice=200000&page=1&limit=10`
4. Нажмите "Send"
5. Проверьте ответ - должен вернуться массив продуктов, соответствующих фильтрам

### Тестирование эндпоинтов категорий

#### 1. Создание категории (POST /categories)

1. Создайте новый запрос в Postman
2. Выберите метод POST
3. Введите URL: `http://localhost:8081/categories`
4. Перейдите на вкладку "Body"
5. Выберите "raw" и "JSON"
6. Введите тело запроса:
   ```json
   {
     "name": "electronics"
   }
   ```
7. Нажмите "Send"
8. Проверьте ответ - должна вернуться созданная категория с ID
9. Сохраните ID категории для последующих запросов

#### 2. Получение категории по ID (GET /categories/:id)

1. Создайте новый запрос в Postman
2. Выберите метод GET
3. Введите URL: `http://localhost:8081/categories/{id}` (замените {id} на ID, полученный в предыдущем шаге)
4. Нажмите "Send"
5. Проверьте ответ - должна вернуться категория с указанным ID

#### 3. Получение списка категорий (GET /categories)

1. Создайте новый запрос в Postman
2. Выберите метод GET
3. Введите URL: `http://localhost:8081/categories`
4. Нажмите "Send"
5. Проверьте ответ - должен вернуться массив категорий

#### 4. Удаление категории (DELETE /categories/:id)

1. Создайте новый запрос в Postman
2. Выберите метод DELETE
3. Введите URL: `http://localhost:8081/categories/{id}` (замените {id} на ID категории)
4. Нажмите "Send"
5. Проверьте ответ - должно вернуться сообщение об успешном удалении

## Структура проекта

```
inventory-service/
├── cmd/
│   └── main.go                 # Точка входа в приложение
├── config/
│   └── mongo.go                # Конфигурация MongoDB
├── internal/
│   ├── domain/
│   │   ├── product.go          # Модель продукта
│   │   └── category.go         # Модель категории
│   ├── handler/
│   │   ├── product_handler.go  # Обработчик запросов для продуктов
│   │   └── category_handler.go # Обработчик запросов для категорий
│   ├── repository/
│   │   ├── product_mongo.go    # Репозиторий для работы с продуктами в MongoDB
│   │   └── category_mongo.go   # Репозиторий для работы с категориями в MongoDB
│   └── usecase/
│       ├── product_usecase.go  # Бизнес-логика для продуктов
│       └── category_usecase.go # Бизнес-логика для категорий
└── public/
    ├── css/
    │   └── style.css           # Стили для фронтенда
    ├── js/
    │   └── main.js             # JavaScript для фронтенда
    └── index.html              # Главная страница
```

## Запуск проекта

1. Убедитесь, что MongoDB запущена на `localhost:27017`
2. Запустите сервер:
   ```
   cd inventory-service
   go run cmd/main.go
   ```
3. Откройте браузер и перейдите по адресу `http://localhost:8081`

## Технологии

- **Backend**: Go, Gin, MongoDB
- **Frontend**: HTML, CSS, JavaScript
- **База данных**: MongoDB

## Дополнительная информация

- Сервер запускается на порту 8081
- MongoDB должна быть запущена на порту 27017
- База данных называется "kazakhexpress"
- Коллекции: "products" и "categories"
