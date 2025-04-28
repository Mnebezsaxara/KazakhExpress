// Функция для получения товаров с сервера
const getProducts = async (
  category = "",
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  try {
    let url = `http://localhost:8082/products?page=${page}&limit=${limit}`;
    if (category) url += `&category=${encodeURIComponent(category)}`;
    if (minPrice) url += `&minPrice=${encodeURIComponent(minPrice)}`;
    if (maxPrice) url += `&maxPrice=${encodeURIComponent(maxPrice)}`;

    console.log("Запрос к API:", url);
    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    return data.products || [];
  } catch (error) {
    console.error("Ошибка при получении данных:", error);
    return [];
  }
};

// Функция для получения категорий с сервера
const getCategories = async () => {
  try {
    const url = "http://localhost:8082/categories";
    console.log("Запрос категорий к API:", url);

    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    return data.categories || [];
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

    if (!products.length) {
      list.innerHTML = "<p class='no-products'>Товары не найдены</p>";
      return;
    }

    list.innerHTML = "";

    products.forEach((product) => {
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

    const addToCartButtons = list.querySelectorAll(".add-to-cart");
    addToCartButtons.forEach((button) => {
      button.addEventListener("click", (e) => {
        const productId = e.target.getAttribute("data-id");
        console.log(`Добавление в корзину: ${productId}`);
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

// Функция для загрузки товаров
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
  await renderProducts(sectionId, category, minPrice, maxPrice, page, limit);
};

// Функция для загрузки дополнительных товаров
window.loadMore = function (sectionId) {
  const category = document.getElementById("category-filter").value;
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  const list = document.getElementById(`${sectionId}-list`);
  const currentProducts = list.querySelectorAll(".card").length;

  if (currentProducts >= 12) {
    alert("Все доступные товары уже загружены");
    return;
  }

  if (category && category.includes(" ")) {
    loadProductsByProductName(sectionId, category, minPrice, maxPrice, 1, 12);
  } else {
    loadProducts(sectionId, category, minPrice, maxPrice, 1, 12);
  }
};

// Фильтрация по категории
const filterByCategory = async (category) => {
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  ["featured", "bestseller", "recommended"].forEach((sectionId) => {
    loadProducts(sectionId, category, minPrice, maxPrice);
  });
};

// Применение фильтров
const applyFilters = async () => {
  const category = document.getElementById("category-filter").value;
  const minPrice = document.getElementById("min-price").value;
  const maxPrice = document.getElementById("max-price").value;

  if (minPrice && maxPrice && parseFloat(minPrice) > parseFloat(maxPrice)) {
    alert("Минимальная цена не может быть больше максимальной");
    return;
  }

  if (category && category.includes(" ")) {
    const allProducts = await getProducts("", "", "", 1, 100);
    const selectedProduct = allProducts.find((p) => p.name === category);
    if (selectedProduct) {
      ["featured", "bestseller", "recommended"].forEach((sectionId) => {
        loadProductsByProductName(sectionId, category, minPrice, maxPrice);
      });
      return;
    } else {
      alert(`Продукт ${category} не найден`);
      return;
    }
  }

  ["featured", "bestseller", "recommended"].forEach((sectionId) => {
    loadProducts(sectionId, category, minPrice, maxPrice);
  });
};

// Загрузка товаров по названию
const loadProductsByProductName = async (
  sectionId,
  productName,
  minPrice = "",
  maxPrice = "",
  page = 1,
  limit = 4
) => {
  try {
    const products = await getProducts("", minPrice, maxPrice, page, limit);
    const filteredProducts = products.filter((product) =>
      product.name.toLowerCase().includes(productName.toLowerCase())
    );
    renderFilteredProducts(sectionId, filteredProducts);
  } catch (error) {
    console.error(`Ошибка при загрузке продуктов по названию: ${error}`);
    const list = document.getElementById(`${sectionId}-list`);
    if (list) {
      list.innerHTML = "<p class='error'>Ошибка при загрузке товаров</p>";
    }
  }
};

// Отображение отфильтрованных товаров
const renderFilteredProducts = (sectionId, products) => {
  const list = document.getElementById(`${sectionId}-list`);
  if (!list) {
    console.error(`Element with id ${sectionId}-list not found!`);
    return;
  }

  if (!products.length) {
    list.innerHTML = "<p class='no-products'>Товары не найдены</p>";
    return;
  }

  list.innerHTML = "";

  products.forEach((product) => {
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

  const addToCartButtons = list.querySelectorAll(".add-to-cart");
  addToCartButtons.forEach((button) => {
    button.addEventListener("click", (e) => {
      const productId = e.target.getAttribute("data-id");
      console.log(`Добавление в корзину: ${productId}`);
      alert(`Товар ${productId} добавлен в корзину`);
    });
  });
};

// Инициализация категорий
const initCategories = async () => {
  try {
    const categories = await getCategories();
    const categoryFilter = document.getElementById("category-filter");

    while (categoryFilter.options.length > 1) {
      categoryFilter.remove(1);
    }

    categories.forEach((category) => {
      const option = document.createElement("option");
      option.value = category.name;
      option.textContent = category.name;
      categoryFilter.appendChild(option);
    });

    const separator = document.createElement("option");
    separator.disabled = true;
    separator.textContent = "──────────";
    categoryFilter.appendChild(separator);

    const allProducts = await getProducts("", "", "", 1, 100);
    allProducts.forEach((product) => {
      const option = document.createElement("option");
      option.value = product.name;
      option.textContent = product.name;
      categoryFilter.appendChild(option);
    });
  } catch (error) {
    console.error("Ошибка при инициализации категорий:", error);
  }
};

// Инициализация при загрузке страницы
document.addEventListener("DOMContentLoaded", async () => {
  console.log("Страница загружена, инициализация...");

  await initCategories();

  ["featured", "bestseller", "recommended"].forEach((sectionId) => {
    loadProducts(sectionId);
  });

  const filterButton = document.getElementById("filter-button");
  if (filterButton) {
    filterButton.addEventListener("click", applyFilters);
  }
});
