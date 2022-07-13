function setError(text) {
    $("#error").text(text)
} 

async function parseProducts(products = null) {
    if (products === null) {
        products = await loadProducts()
    }
    let table = $('#productsTable').find("tbody")
    table.empty()

    for (let product of products) {
        table.append(`
        <tr>
            <td>${product.product_title}</td>
            <td>${product.product_cost}</td>
            <td>${product.category_title}</td>
            <td class="column">
                <button class="button" @click="await calculate(${product.product_id})">Посчитать</button>
                <button class="button" productId=${product.product_id} @click="await delProduct(${product.product_id})">
                    Удалить
                </button>
            </td>
        </tr>
        <hr>
        `)
        
    }
}

async function parseCategories(forDelete=false) {
    let categories = await loadCategories()
    let selector = $("#productCategory")
    selector.empty() 
    if (!forDelete) {
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

async function parseProductsAmounts() {
    let productsAmounts = await getAmounts()
    let table = $("#amountsTable").find("tbody")
    table.empty()
    $("#sum").text(productsAmounts.sum)
    for (let amount of productsAmounts.products) {
        table.append(`
        <tr>
            <td>${amount.product_title}</td>
            <td>${amount.product_amount}</td>
            <td>${amount.product_amount_cost}</td>
            <td>
                <button @click="await removeFromCalc(${amount.product_id})" class="button">Убрать из расчёта</button>
            </td>
        </tr>
        `)
    }
}



