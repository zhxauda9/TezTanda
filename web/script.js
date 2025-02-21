let cart = [];
let totalAmount = 0;
let productsData = [];

async function loadProducts() {
    try {
        const response = await fetch("http://localhost:8080/products");
        productsData = await response.json();
        displayProducts(productsData);
    } catch (error) {
        console.error("Error loading products:", error);
    }
}

function displayProducts(products) {
    const productList = document.getElementById("products-list");
    productList.innerHTML = "";
    products.forEach(product => {
        productList.innerHTML += `
            <div class="col-md-4 d-flex justify-content-center">
                <div class="card shadow-sm product-card">
                    <img src="${product.image}" class="card-img-top" alt="${product.name}">
                    <div class="card-body text-center">
                        <h5 class="card-title">${product.name}</h5>
                        <p class="text-muted">${product.description}</p>
                        <p class="text-success fw-bold">${product.price} ₸</p>
                        <p class="text-danger">Stock: ${product.stock}</p>
                        <button class="btn btn-success w-100" onclick="addToCart('${product.id}', '${product.name}', '${product.image}', '${product.price}')">Add to Cart</button>
                    </div>
                </div>
            </div>
        `;
    });
}

function addToCart(id, name, image, price) {
    let item = cart.find(p => p.id === id);
    if (item) {
        item.quantity++;
    } else {
        cart.push({ id, name, image, price: Number(price), quantity: 1 });
    }
    totalAmount += Number(price);
    updateCartDisplay();
}

function updateCartDisplay() {
    const cartItemsContainer = document.getElementById("cart-items");
    cartItemsContainer.innerHTML = "";
    document.getElementById("cart-count").innerText = cart.length;
    document.getElementById("total-amount").innerText = `${totalAmount.toLocaleString()} ₸`;

    cart.forEach(item => {
        cartItemsContainer.innerHTML += `
            <div class="cart-item">
                <img src="${item.image}" alt="${item.name}">
                <span>${item.name} x${item.quantity}</span>
                <button onclick="removeFromCart('${item.id}')">❌</button>
            </div>
        `;
    });
}

function removeFromCart(id) {
    let itemIndex = cart.findIndex(p => p.id === id);
    if (itemIndex > -1) {
        totalAmount -= cart[itemIndex].price * cart[itemIndex].quantity;
        cart.splice(itemIndex, 1);
    }
    updateCartDisplay();
}

function clearCart() {
    cart = [];
    totalAmount = 0;
    updateCartDisplay();
}

function toggleCartPopup() {
    document.getElementById("cart-popup").classList.toggle("active");
}

function filterByCategory() {
    let selectedCategory = document.getElementById("categoryFilter").value;
    if (selectedCategory === "all") {
        displayProducts(productsData);
    } else {
        let filtered = productsData.filter(p => p.category === selectedCategory);
        displayProducts(filtered);
    }
}

loadProducts();
