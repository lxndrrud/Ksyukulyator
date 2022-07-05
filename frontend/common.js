function setError(text) {
    error = text
} 

async function parseProducts() {
    let products = await loadProducts()
    let table = $('#productsTable').find("tbody")
    table.empty()

    for (let product of products) {
        table.append(`
        <tr productId=${product.product_id}>
            <td>${product.product_title}</td>
            <td>${product.product_cost}</td>
            <td>${product.category_title}</td>
            <td class="column">
                <button productId=${product.product_id}>Изменить</button>
                <button productId=${product.product_id} @click="await delProduct(${product.product_id})">
                    Удалить
                </button>
            </td>
        </tr>`)
        
    }
}

async function parseCategories(notForDelete=false) {
    let categories = await loadCategories()
    let selector = $("#productCategory")
    selector.empty() 
    if (notForDelete) {
        selector.append(`
        <option categoryid="0">
            Без категории
        </option>
    `)
    }
    for (let category of categories) {
        selector.append(`
        <option categoryid=${category.category_id}>
            ${category.category_title}
        </option>
        `)
    }
} 