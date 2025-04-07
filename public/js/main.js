// Функция для получения товаров с сервера
const getProducts = async (
  category = "",
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  try {
    // Формируем URL с параметрами
    let url = `http://localhost:8081/products?page=${page}&limit=${limit}`;

    // Добавляем параметры только если они не пустые
    if (category) url += `&category=${encodeURIComponent(category)}`;
    if (minPrice) url += `&minPrice=${encodeURIComponent(minPrice)}`;
    if (maxPrice) url += `&maxPrice=${encodeURIComponent(maxPrice)}`;

    console.log("Запрос к API:", url);

    const response = await fetch(url);

    if (!response.ok) {
      console.error(`HTTP error! status: ${response.status}`);
      return [];
    }

    // Проверяем тип контента
    const contentType = response.headers.get("content-type");
    console.log("Content-Type:", contentType);

    // Получаем текст ответа для отладки
    const responseText = await response.text();
    console.log("Raw response:", responseText);

    // Если ответ пустой или null, возвращаем пустой массив
    if (!responseText || responseText === "null") {
      console.warn("API вернул пустой ответ или null");
      return [];
    }

    // Пытаемся распарсить JSON
    let data;
    try {
      data = JSON.parse(responseText);
    } catch (e) {
      console.error("Error parsing JSON:", e);
      return [];
    }

    console.log("Parsed data:", data);

    // Проверяем, что данные не null и являются массивом
    if (data === null) {
      console.warn("API вернул null, возвращаем пустой массив");
      return [];
    }

    if (!Array.isArray(data)) {
      console.warn("API вернул не массив:", data);
      // Если данные не массив, но содержат массив products, используем его
      if (data && data.products && Array.isArray(data.products)) {
        return data.products;
      }
      return [];
    }

    return data;
  } catch (error) {
    console.error("Ошибка при получении данных:", error);
    return []; // Возвращаем пустой массив в случае ошибки
  }
};

// Функция для получения категорий с сервера
const getCategories = async () => {
  try {
    console.log("Запрос категорий к API");
    const response = await fetch("http://localhost:8081/categories");

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log("Полученные категории:", data);

    // Проверяем, что данные не null и являются массивом
    if (data === null) {
      console.warn("API вернул null для категорий, возвращаем пустой массив");
      return [];
    }

    if (!Array.isArray(data)) {
      console.warn("API вернул не массив для категорий:", data);
      return [];
    }

    return data;
  } catch (error) {
    console.error("Ошибка при получении категорий:", error);
    return [];
  }
};

// Функция для отображения продуктов
const renderProducts = async (
  sectionId,
  category = "",
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  const list = document.getElementById(`${sectionId}-list`);

  if (!list) {
    console.error(`Element with id ${sectionId}-list not found!`);
    return;
  }

  // Показываем индикатор загрузки
  list.innerHTML = '<div class="loading">Загрузка товаров...</div>';

  try {
    const products = await getProducts(
      category,
      minPrice,
      maxPrice,
      page,
      limit
    );
    console.log(
      `Получено ${products.length} продуктов для раздела ${sectionId}`
    );

    if (!products || products.length === 0) {
      list.innerHTML = "<p class='no-products'>Товары не найдены</p>";
      return;
    }

    // Очищаем текущий список перед добавлением новых товаров
    list.innerHTML = "";

    products.forEach((product) => {
      // Проверяем наличие необходимых полей
      const productId = product.id || product._id || "";
      const productName = product.name || "Без названия";
      const productPrice = product.price || 0;
      const productStock = product.stock || 0;
      const productImage = product.image_url || "images/placeholder.jpg";

      const card = document.createElement("div");
      card.className = "card";
      card.innerHTML = `
        <div class="card-image">
          <img src="${productImage}" alt="${productName}" onerror="this.src='images/placeholder.jpg'">
        </div>
        <div class="card-content">
          <h3>${productName}</h3>
          <p class="price">₸${productPrice.toFixed(2)}</p>
          <p class="stock">В наличии: ${productStock}</p>
          <button class="add-to-cart" data-id="${productId}">В корзину</button>
        </div>
      `;
      list.appendChild(card);
    });

    // Добавляем обработчики для кнопок "В корзину"
    const addToCartButtons = list.querySelectorAll(".add-to-cart");
    addToCartButtons.forEach((button) => {
      button.addEventListener("click", (e) => {
        const productId = e.target.getAttribute("data-id");
        console.log(`Добавление в корзину: ${productId}`);
        // Здесь можно реализовать логику добавления в корзину
        alert(`Товар ${productId} добавлен в корзину`);
      });
    });
  } catch (error) {
    console.error(
      `Ошибка при отображении продуктов для раздела ${sectionId}:`,
      error
    );
    list.innerHTML = "<p class='error'>Ошибка при загрузке товаров</p>";
  }
};

// Загружаем продукты для разных разделов
const loadProducts = async (
  sectionId,
  category = "",
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  console.log(`Загрузка продуктов для раздела ${sectionId}`, {
    category,
    minPrice,
    maxPrice,
    page,
    limit,
  });

  // Если это первая загрузка, используем разные фильтры для разных разделов
  if (page === 1 && limit === 4 && !category && !minPrice && !maxPrice) {
    // Для разных разделов используем разные фильтры
    switch (sectionId) {
      case "featured":
        // Для featured используем все продукты
        break;
      case "bestseller":
        // Для bestseller можно использовать сортировку по популярности
        // В данном случае просто используем все продукты
        break;
      case "recommended":
        // Для recommended можно использовать персонализированные рекомендации
        // В данном случае просто используем все продукты
        break;
    }
  }

  await renderProducts(sectionId, category, minPrice, maxPrice, page, limit);
};

// Функция для загрузки дополнительных товаров
const loadMore = (sectionId) => {
  const category = document.getElementById("category-filter").value;
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  console.log(`Загрузка всех товаров для раздела ${sectionId}`, {
    category,
    minPrice,
    maxPrice,
  });

  // Получаем текущий список продуктов
  const list = document.getElementById(`${sectionId}-list`);
  const currentProducts = list.querySelectorAll(".card").length;

  // Если уже загружено 12 или более продуктов, показываем сообщение
  if (currentProducts >= 12) {
    alert("Все доступные товары уже загружены");
    return;
  }

  // Проверяем, является ли выбранное значение продуктом или категорией
  if (category && category.includes(" ")) {
    // Если это продукт, используем специальную функцию для фильтрации по продукту
    loadProductsByProductName(sectionId, category, minPrice, maxPrice, 1, 12);
  } else {
    // Если это категория, используем стандартную функцию
    loadProducts(sectionId, category, minPrice, maxPrice, 1, 12);
  }

  // Обновляем текст кнопки, если все товары загружены
  const loadMoreButton = document.querySelector(
    `#${sectionId}-list`
  ).nextElementSibling;
  if (loadMoreButton && loadMoreButton.classList.contains("load-more")) {
    loadMoreButton.textContent = "Все товары загружены";
    loadMoreButton.disabled = true;
  }
};

// Фильтрация по категории
const filterByCategory = async (category) => {
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  // Обновляем все разделы с выбранной категорией
  ["featured", "bestseller", "recommended"].forEach((sectionId) => {
    loadProducts(sectionId, category, minPrice, maxPrice);
  });
};

// Применение фильтров
const applyFilters = async () => {
  const category = document.getElementById("category-filter").value;
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  console.log("Применение фильтров:", { category, minPrice, maxPrice });

  // Проверяем, что minPrice не больше maxPrice
  if (minPrice && maxPrice && parseFloat(minPrice) > parseFloat(maxPrice)) {
    alert("Минимальная цена не может быть больше максимальной");
    return;
  }

  // Проверяем, что категория существует в базе данных
  if (category && category !== "") {
    console.log(`Фильтрация по категории: ${category}`);

    // Проверяем, является ли выбранное значение продуктом или категорией
    // Если это продукт (содержит пробел), то ищем его категорию
    if (category.includes(" ")) {
      console.log(`Выбрано название продукта: ${category}`);

      // Получаем все продукты, чтобы найти категорию для выбранного продукта
      try {
        const allProducts = await getProducts("", "", "", 1, 100);
        const selectedProduct = allProducts.find((p) => p.name === category);

        if (selectedProduct) {
          console.log(
            `Найдена категория для продукта ${category}: ${selectedProduct.category}`
          );

          // Создаем специальный параметр для фильтрации по имени продукта
          const productName = category;

          // Обновляем все разделы с примененными фильтрами
          ["featured", "bestseller", "recommended"].forEach((sectionId) => {
            console.log(
              `Загрузка продуктов для раздела ${sectionId} с фильтром по продукту: ${productName}`
            );
            loadProductsByProductName(
              sectionId,
              productName,
              minPrice,
              maxPrice
            );
          });

          return; // Выходим из функции, так как используем специальную функцию для фильтрации по продукту
        } else {
          console.warn(`Продукт ${category} не найден в базе данных`);
          alert(`Продукт ${category} не найден в базе данных`);
          return;
        }
      } catch (error) {
        console.error("Ошибка при поиске категории продукта:", error);
        return;
      }
    } else {
      // Если это категория, нормализуем её (приводим к нижнему регистру)
      const normalizedCategory = category.toLowerCase();
      if (normalizedCategory !== category) {
        console.log(
          `Нормализация категории: ${category} -> ${normalizedCategory}`
        );
        document.getElementById("category-filter").value = normalizedCategory;
      }
    }
  }

  // Обновляем все разделы с примененными фильтрами
  ["featured", "bestseller", "recommended"].forEach((sectionId) => {
    console.log(`Загрузка продуктов для раздела ${sectionId} с фильтрами`);
    loadProducts(
      sectionId,
      document.getElementById("category-filter").value,
      minPrice,
      maxPrice
    );
  });
};

// Функция для загрузки продуктов по имени продукта
const loadProductsByProductName = async (
  sectionId,
  productName,
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  console.log(
    `Загрузка продуктов для раздела ${sectionId} по имени продукта: ${productName}`,
    {
      productName,
      minPrice,
      maxPrice,
      page,
      limit,
    }
  );

  // Получаем все продукты
  const allProducts = await getProducts("", "", "", 1, 100);

  // Фильтруем продукты по имени
  const filteredProducts = allProducts.filter(
    (product) => product.name === productName
  );

  // Применяем фильтры по цене, если они заданы
  let priceFilteredProducts = filteredProducts;
  if (minPrice) {
    priceFilteredProducts = priceFilteredProducts.filter(
      (product) => product.price >= parseFloat(minPrice)
    );
  }
  if (maxPrice) {
    priceFilteredProducts = priceFilteredProducts.filter(
      (product) => product.price <= parseFloat(maxPrice)
    );
  }

  // Применяем пагинацию
  const startIndex = (page - 1) * limit;
  const endIndex = startIndex + limit;
  const paginatedProducts = priceFilteredProducts.slice(startIndex, endIndex);

  console.log(
    `Отфильтровано ${paginatedProducts.length} продуктов для раздела ${sectionId}`
  );

  // Отображаем отфильтрованные продукты
  renderFilteredProducts(sectionId, paginatedProducts);
};

// Функция для отображения отфильтрованных продуктов
const renderFilteredProducts = (sectionId, products) => {
  const list = document.getElementById(`${sectionId}-list`);

  if (!list) {
    console.error(`Element with id ${sectionId}-list not found!`);
    return;
  }

  // Показываем индикатор загрузки
  list.innerHTML = '<div class="loading">Загрузка товаров...</div>';

  try {
    console.log(
      `Отображение ${products.length} продуктов для раздела ${sectionId}`
    );

    if (!products || products.length === 0) {
      list.innerHTML = "<p class='no-products'>Товары не найдены</p>";
      return;
    }

    // Очищаем текущий список перед добавлением новых товаров
    list.innerHTML = "";

    products.forEach((product) => {
      // Проверяем наличие необходимых полей
      const productId = product.id || product._id || "";
      const productName = product.name || "Без названия";
      const productPrice = product.price || 0;
      const productStock = product.stock || 0;
      const productImage = product.image_url || "images/placeholder.jpg";

      const card = document.createElement("div");
      card.className = "card";
      card.innerHTML = `
        <div class="card-image">
          <img src="${productImage}" alt="${productName}" onerror="this.src='images/placeholder.jpg'">
        </div>
        <div class="card-content">
          <h3>${productName}</h3>
          <p class="price">₸${productPrice.toFixed(2)}</p>
          <p class="stock">В наличии: ${productStock}</p>
          <button class="add-to-cart" data-id="${productId}">В корзину</button>
        </div>
      `;
      list.appendChild(card);
    });

    // Добавляем обработчики для кнопок "В корзину"
    const addToCartButtons = list.querySelectorAll(".add-to-cart");
    addToCartButtons.forEach((button) => {
      button.addEventListener("click", (e) => {
        const productId = e.target.getAttribute("data-id");
        console.log(`Добавление в корзину: ${productId}`);
        // Здесь можно реализовать логику добавления в корзину
        alert(`Товар ${productId} добавлен в корзину`);
      });
    });
  } catch (error) {
    console.error(
      `Ошибка при отображении продуктов для раздела ${sectionId}:`,
      error
    );
    list.innerHTML = "<p class='error'>Ошибка при загрузке товаров</p>";
  }
};

// Инициализация категорий в выпадающем списке
const initCategories = async () => {
  try {
    console.log("Запрос категорий к API");
    const response = await fetch("http://localhost:8081/categories");

    if (!response.ok) {
      console.error(`HTTP error! status: ${response.status}`);
      return;
    }

    // Проверяем тип контента
    const contentType = response.headers.get("content-type");
    console.log("Content-Type для категорий:", contentType);

    // Получаем текст ответа для отладки
    const responseText = await response.text();
    console.log("Raw response для категорий:", responseText);

    // Если ответ пустой или null, выходим
    if (!responseText || responseText === "null") {
      console.warn("API вернул пустой ответ или null для категорий");
      return;
    }

    // Пытаемся распарсить JSON
    let data;
    try {
      data = JSON.parse(responseText);
    } catch (e) {
      console.error("Error parsing JSON для категорий:", e);
      return;
    }

    console.log("Полученные категории:", data);

    // Проверяем, что данные не null и являются массивом
    if (data === null) {
      console.warn("API вернул null для категорий, возвращаем пустой массив");
      return;
    }

    if (!Array.isArray(data)) {
      console.warn("API вернул не массив для категорий:", data);
      return;
    }

    // Если массив категорий пуст, добавляем стандартные категории
    if (data.length === 0) {
      console.log("Добавление стандартных категорий");
      data = [
        { id: "1", name: "electronics" },
        { id: "2", name: "clothing" },
      ];
    }

    const categoryFilter = document.getElementById("category-filter");

    // Очищаем текущие опции, кроме "Все товары"
    while (categoryFilter.options.length > 1) {
      categoryFilter.remove(1);
    }

    // Добавляем категории из базы данных
    data.forEach((category) => {
      const option = document.createElement("option");
      option.value = category.name;
      option.textContent = category.name;
      categoryFilter.appendChild(option);
    });

    // Добавляем разделитель
    const separator = document.createElement("option");
    separator.disabled = true;
    separator.textContent = "──────────";
    categoryFilter.appendChild(separator);

    // Получаем все продукты, чтобы добавить их имена в выпадающий список
    try {
      const allProducts = await getProducts("", "", "", 1, 100);
      console.log(
        "Получены все продукты для добавления в фильтр:",
        allProducts
      );

      // Добавляем продукты в выпадающий список
      allProducts.forEach((product) => {
        const option = document.createElement("option");
        option.value = product.name;
        option.textContent = product.name;
        categoryFilter.appendChild(option);
      });
    } catch (error) {
      console.error("Ошибка при получении продуктов для фильтра:", error);
    }

    console.log("Категории и продукты успешно загружены");
  } catch (error) {
    console.error("Ошибка при инициализации категорий:", error);
  }
};

// Инициализация при загрузке страницы
document.addEventListener("DOMContentLoaded", async () => {
  console.log("Страница загружена, инициализация...");

  try {
    // Загружаем категории
    console.log("Загрузка категорий...");
    await initCategories();

    // Загружаем продукты для разных разделов
    console.log("Загрузка продуктов...");
    ["featured", "bestseller", "recommended"].forEach((sectionId) => {
      console.log(`Загрузка продуктов для раздела ${sectionId}`);
      loadProducts(sectionId);
    });

    // Добавляем обработчик для кнопки фильтрации
    const filterButton = document.getElementById("filter-button");
    if (filterButton) {
      console.log("Добавление обработчика для кнопки фильтрации");
      filterButton.addEventListener("click", () => {
        console.log("Применение фильтров");
        applyFilters();
      });
    } else {
      console.warn("Кнопка фильтрации не найдена");
    }

    // Добавляем обработчик для поиска
    const searchButton = document.querySelector(".search-button");
    const searchInput = document.querySelector(".search-input");

    if (searchButton && searchInput) {
      console.log("Добавление обработчика для поиска");
      searchButton.addEventListener("click", () => {
        const searchTerm = searchInput.value.trim();
        if (searchTerm) {
          console.log("Поиск:", searchTerm);
          // Здесь можно реализовать поиск по названию товара
        }
      });
    } else {
      console.warn("Элементы поиска не найдены");
    }

    console.log("Инициализация завершена");
  } catch (error) {
    console.error("Ошибка при инициализации:", error);
  }
});
