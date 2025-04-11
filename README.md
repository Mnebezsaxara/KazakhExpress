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

## Тестирование gRPC сервиса через BloomRPC

### Настройка BloomRPC

1. Скачайте и установите [BloomRPC](https://github.com/bloomrpc/bloomrpc/releases)
2. Запустите BloomRPC
3. Импортируйте proto-файлы:
   - Нажмите на кнопку "+" в левом верхнем углу
   - Выберите файлы `inventory-service/proto/product.proto` и `inventory-service/proto/category.proto`
4. В поле "Server Address" укажите: `localhost:50051`

### Тестирование Category Service

#### 1. CreateCategory

```json
// Метод: inventory.CategoryService/CreateCategory
{
    "name": "electronics",
    "description": "Electronic devices and gadgets"
}

// Ожидаемый ответ:
{
    "id": "generated-id",
    "name": "electronics",
    "description": "Electronic devices and gadgets",
    "product_count": 0,
    "created_at": "timestamp",
    "updated_at": "timestamp"
}
```

#### 2. GetCategory

```json
// Метод: inventory.CategoryService/GetCategory
{
    "id": "your-category-id"
}

// Ожидаемый ответ:
{
    "category": {
        "id": "your-category-id",
        "name": "electronics",
        "description": "Electronic devices and gadgets",
        "product_count": 0,
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
}
```

#### 3. ListCategories

```json
// Метод: inventory.CategoryService/ListCategories
{}

// Ожидаемый ответ:
{
    "categories": [
        {
            "id": "category-id-1",
            "name": "electronics",
            "description": "Electronic devices and gadgets",
            "product_count": 5,
            "created_at": "timestamp",
            "updated_at": "timestamp"
        },
        {
            "id": "category-id-2",
            "name": "clothing",
            "description": "Fashion items",
            "product_count": 3,
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    ],
    "total": 2
}
```

#### 4. DeleteCategory

```json
// Метод: inventory.CategoryService/DeleteCategory
{
    "id": "your-category-id"
}

// Ожидаемый ответ:
{}
```

### Тестирование Product Service

#### 1. CreateProduct

```json
// Метод: inventory.ProductService/CreateProduct
{
    "name": "iPhone 13",
    "description": "Latest iPhone model",
    "price": 499999,
    "image_url": "https://example.com/iphone13.jpg",
    "category": "electronics",
    "stock": 10
}

// Ожидаемый ответ:
{
    "id": "generated-id",
    "name": "iPhone 13",
    "description": "Latest iPhone model",
    "price": 499999,
    "image_url": "https://example.com/iphone13.jpg",
    "category": "electronics",
    "stock": 10,
    "created_at": "timestamp",
    "updated_at": "timestamp"
}
```

#### 2. GetProduct

```json
// Метод: inventory.ProductService/GetProduct
{
    "id": "your-product-id"
}

// Ожидаемый ответ:
{
    "product": {
        "id": "your-product-id",
        "name": "iPhone 13",
        "description": "Latest iPhone model",
        "price": 499999,
        "image_url": "https://example.com/iphone13.jpg",
        "category": "electronics",
        "stock": 10,
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
}
```

#### 3. ListProducts

```json
// Метод: inventory.ProductService/ListProducts
{
    "filter": {
        "category": "electronics",
        "min_price": 100000,
        "max_price": 500000,
        "page": 1,
        "limit": 10
    }
}

// Ожидаемый ответ:
{
    "products": [
        {
            "id": "product-id-1",
            "name": "iPhone 13",
            "description": "Latest iPhone model",
            "price": 499999,
            "image_url": "https://example.com/iphone13.jpg",
            "category": "electronics",
            "stock": 10,
            "created_at": "timestamp",
            "updated_at": "timestamp"
        }
    ],
    "total": 1
}
```

#### 4. UpdateProduct

```json
// Метод: inventory.ProductService/UpdateProduct
{
    "id": "your-product-id",
    "product": {
        "name": "iPhone 13",
        "description": "Latest iPhone model with updates",
        "price": 459999,
        "image_url": "https://example.com/iphone13.jpg",
        "category": "electronics",
        "stock": 15
    }
}

// Ожидаемый ответ:
{
    "product": {
        "id": "your-product-id",
        "name": "iPhone 13",
        "description": "Latest iPhone model with updates",
        "price": 459999,
        "image_url": "https://example.com/iphone13.jpg",
        "category": "electronics",
        "stock": 15,
        "created_at": "timestamp",
        "updated_at": "timestamp"
    }
}
```

#### 5. DeleteProduct

```json
// Метод: inventory.ProductService/DeleteProduct
{
    "id": "your-product-id"
}

// Ожидаемый ответ:
{}
```

### Сценарии тестирования

1. **Создание и управление категориями:**

   - Создайте новую категорию
   - Получите список всех категорий
   - Получите конкретную категорию по ID
   - Попробуйте удалить категорию

2. **Создание и управление продуктами:**

   - Создайте новый продукт в существующей категории
   - Получите список всех продуктов
   - Отфильтруйте продукты по категории
   - Обновите информацию о продукте
   - Удалите продукт

3. **Проверка валидации:**

   - Попробуйте создать продукт с отрицательной ценой
   - Попробуйте создать продукт с отрицательным количеством
   - Попробуйте создать продукт в несуществующей категории

4. **Проверка фильтрации и пагинации:**
   - Создайте несколько продуктов
   - Используйте фильтры по цене
   - Проверьте работу пагинации

### Обработка ошибок

При тестировании вы можете получить следующие ошибки:

1. `NOT_FOUND` - запрашиваемый ресурс не найден
2. `INVALID_ARGUMENT` - неверные параметры запроса
3. `INTERNAL` - внутренняя ошибка сервера
4. `ALREADY_EXISTS` - ресурс уже существует

### Советы по тестированию

1. Сохраняйте ID созданных ресурсов для использования в последующих запросах
2. Проверяйте все поля в ответах
3. Тестируйте граничные случаи (пустые значения, большие числа)
4. Проверяйте обработку ошибок

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

## Изменения во втором задании

### Архитектурные изменения

1. **Реализация gRPC сервиса**

   - Добавлен gRPC сервер для микросервисной архитектуры
   - Реализованы proto-файлы для Product и Category сервисов
   - Настроена двусторонняя связь между сервисами

2. **Улучшение чистой архитектуры**

   - Унифицированы интерфейсы репозиториев и usecase'ов
   - Добавлена поддержка контекста (context.Context) во всех методах
   - Улучшена обработка ошибок с использованием gRPC статусов

3. **Улучшения в работе с MongoDB**
   - Добавлена поддержка ObjectID для корректной работы с MongoDB
   - Улучшена обработка ошибок при работе с базой данных
   - Добавлена поддержка транзакций для атомарных операций

### Новые возможности

1. **Расширенная фильтрация продуктов**

   - Поддержка фильтрации по категориям
   - Фильтрация по диапазону цен
   - Пагинация с поддержкой общего количества элементов

2. **Улучшенная обработка категорий**

   - Добавлен счетчик продуктов в категории
   - Атомарное обновление счетчика при добавлении/удалении продуктов
   - Валидация при удалении категорий с продуктами

3. **Улучшения в обработке ошибок**
   - Добавлены кастомные ошибки для различных сценариев
   - Улучшена обработка ошибок на всех уровнях
   - Добавлено логирование критических ошибок

### Технические улучшения

1. **Улучшения в коде**

   - Удалены дублирующиеся интерфейсы
   - Унифицированы сигнатуры методов
   - Добавлена поддержка контекста для отмены операций
   - Улучшена типизация возвращаемых значений

2. **Безопасность**

   - Добавлена валидация входных данных
   - Улучшена обработка некорректных ID
   - Добавлены проверки на отрицательные значения

3. **Производительность**
   - Оптимизированы запросы к MongoDB
   - Добавлена поддержка индексов
   - Улучшена работа с курсорами

### Примеры использования

#### gRPC вызовы

```go
// Пример создания продукта через gRPC
product := &pb.Product{
    Name:        "Test Product",
    Description: "Test Description",
    Price:       1000,
    Category:    "test",
    Stock:       10,
}

response, err := productClient.CreateProduct(ctx, &pb.CreateProductRequest{
    Product: product,
})
```

#### HTTP API (остался для обратной совместимости)

```http
POST /products
Content-Type: application/json

{
    "name": "Test Product",
    "description": "Test Description",
    "price": 1000,
    "category": "test",
    "stock": 10
}
```

### Мониторинг и логирование

1. **Логирование**

   - Добавлено структурированное логирование
   - Логирование критических операций
   - Трейсинг запросов

2. **Метрики**
   - Счетчики успешных/неуспешных операций
   - Время выполнения операций
   - Количество активных соединений

### Планы на будущее

1. **Дальнейшие улучшения**

   - Добавление кэширования
   - Реализация очередей для асинхронных операций
   - Улучшение мониторинга

2. **Масштабирование**

   - Поддержка шардирования MongoDB
   - Балансировка нагрузки
   - Репликация данных

3. **Безопасность**
   - Добавление аутентификации
   - Шифрование чувствительных данных
   - Улучшение валидации
