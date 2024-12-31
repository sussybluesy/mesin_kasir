async function fetchItems() {
    const response = await fetch('http://localhost:8080/items');
    const items = await response.json();

    const itemsDiv = document.getElementById('items');
    items.forEach(item => {
        const div = document.createElement('div');
        div.innerHTML = `${item.name} - Rp${item.price} <input type="number" id="quantity-${item.id}" placeholder="Quantity">`;
        itemsDiv.appendChild(div);
    });
}

async function checkout() {
    const response = await fetch('http://localhost:8080/items');
    const items = await response.json();

    const cart = items.map(item => {
        const quantity = document.getElementById(`quantity-${item.id}`).value;
        return { ...item, quantity: parseInt(quantity) || 0 };
    }).filter(item => item.quantity > 0);

    const checkoutResponse = await fetch('http://localhost:8080/checkout', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(cart),
    });

    const result = await checkoutResponse.json();
    alert(`Total Belanja: Rp${result.total}`);
}

fetchItems();