const allProducts = {
  featured: [
    { name: "Ноутбук Acer", price: 259990, image: "images/aser.jpg" },
    { name: "Redmi Note 12", price: 99990, image: "images/redmi.webp" },
    { name: "Наушники JBL", price: 19990, image: "images/jbl.jpg" },
    { name: "Mi Band 7", price: 28990, image: "images/miband.webp" },
    { name: "Умная колонка", price: 44990, image: "images/speaker.jpg" },
    { name: "Клавиатура Logitech", price: 30990, image: "images/keyboard.webp" },
    { name: "Мышка Razer", price: 19990, image: "images/mouse.jpg" },
    { name: "Монитор Samsung", price: 129990, image: "images/monitor.png" },
    { name: "Внешний SSD", price: 37990, image: "images/ssd.jpg" },
    { name: "Смарт-лампа", price: 14990, image: "images/lamp.jpg" },
    { name: "Геймпад Xbox", price: 39990, image: "images/gamepad.webp" },
    { name: "Чехол iPhone", price: 5990, image: "images/case.jpg" },
  ],
  bestseller: [
    { name: "Хитовый товар 1", price: 19990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 2", price: 24990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 3", price: 28990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 4", price: 33990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 5", price: 37990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 6", price: 41990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 7", price: 44990, image: "images/hit2.jpeg" },
    { name: "Хитовый товар 8", price: 48990, image: "images/hit2.jpeg" },
  ],
  recommended: [
    { name: "Рекомендуемый 1", price: 12990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 2", price: 15990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 3", price: 17990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 4", price: 19990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 5", price: 21990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 6", price: 23990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 7", price: 26990, image: "images/for_you.jpg" },
    { name: "Рекомендуемый 8", price: 28990, image: "images/for_you.jpg" },
  ],
};

const loaded = { featured: 0, bestseller: 0, recommended: 0 };

function renderProducts(category, count = 4) {
  const list = document.getElementById(`${category}-list`);
  const products = allProducts[category];
  const start = loaded[category];
  const end = Math.min(start + count, products.length); // <= это важно

  products.slice(start, end).forEach(product => {
    const card = document.createElement("div");
    card.className = "card";
    card.innerHTML = `
      <img src="${product.image}" alt="${product.name}">
      <h3>${product.name}</h3>
      <p>₸${product.price}</p>
    `;
    list.appendChild(card);
  });

  loaded[category] = end;

  if (loaded[category] >= products.length) {
    document.querySelector(`button[onclick="loadMore('${category}')"]`).style.display = "none";
  }
}


function loadMore(category) {
  renderProducts(category);
}

['featured', 'bestseller', 'recommended'].forEach(c => renderProducts(c, 4));
