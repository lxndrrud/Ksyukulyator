const MainPage = () => {
    let [idProduct, setIdProduct] = React.useState(0)
    let [amount, setAmount] = React.useState(0.0)
    let [amountCost, setAmountCost] = React.useState(0.0)
    let [sum, setSum] = React.useState(0.0)
    let [isLoading, setIsLoading] = React.useState(true) 
    let [error, setError] = React.useState(null)
    let [products, setProducts] = React.useState([])
    let [productsAmounts, setProductsAmounts] = React.useState([])

    React.useEffect(async () => {
        await parseProducts()
    }, [])

    async function parseProducts() {
        try {
            let productsLocal = await loadProducts()
            setProducts(productsLocal)
        } catch(e) {
            setError(e)
        }
    }
    async function calculate(idProduct) {
        try {
            let summ = await addAmountToCalc(
                parseInt(idProduct), 
                parseFloat($("#amount").val()), parseFloat( $("#amountCost").val()))
            setSum(summ)
            await parseProductsAmounts()
        } catch(e) {
            setError(e)
        }
    }
    async function delProduct(idProduct) {
        try {
            let result = await deleteProduct(idProduct)
            if (result) await parseProducts()
        }
        catch (e) {
            setError(e)
        }
    }
    async function filterProductsByCategory(idCategory) {
        let products = await filterByCategory(parseInt(idCategory))
        await parseProducts(products)
    }
    return (
        <div>
            {
                isLoading
                ?
                    <span>Загрузка...</span>
                :
                    <div>
                        <h1 class="ksyu">Ксюкулятор</h1>
                        <nav>
                            <a href="/frontend/pages/main/main.html" class="link" >Калькулятор</a>
                            <a href="/frontend/pages/categories/categories.html" class="link">Категории</a>
                        </nav>
                        {
                            error
                            ? <span>{error}</span>
                            : null
                        }
                        <div className="content"> 
                            <div className="column">
                                <h3>Продукты</h3>
                                <div class="overflow_container">
                                    <table id="productsTable">
                                        <thead>
                                            <tr>
                                                <td>Название продукта</td>
                                                <td>Количество углеводов, г / 100 г продукта</td>
                                                <td>Название категории</td>
                                                <td>Контроль</td>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {products && products.map(product => 
                                                (<div key={product.product_id}>
                                                    <tr>
                                                        <td>${product.product_title}</td>
                                                        <td>${product.product_cost}</td>
                                                        <td>${product.category_title}</td>
                                                        <td class="column">
                                                            <button productId={product.product_id} 
                                                                onClick={async () => { await calculate(product.product_id)}}>Посчитать</button>
                                                            <button productId={product.product_id} 
                                                                onClick={async () => { await delProduct(product.product_id)}}>Удалить</button>
                                                        </td>
                                                    </tr>
                                                    <hr />
                                                </div>)
                                                
                                            )}
                                        </tbody>
                                        
                                    </table>
                                </div>
                            </div>
                            <div className="column">
                                <h3>Контроль продуктов</h3>
                                <div id="controlButtons">
                                    <button onClick={parseProducts} className="button">Обновить список</button>
                                    <a href="/frontend/pages/addProduct/addProduct.html/" className="button">
                                        Добавить продукт
                                    </a>
                                    <select id="productCategory" onChange={async () => { await filterProductsByCategory() }}>
                                        
                                    </select>
                                </div>
                            </div>
                            <span className="ksyu">
                                ИТОГО: <span id="sum">{sum}</span> хлебных единиц
                            </span>
                        </div>
                    </div>
            }
            
        </div>
    )
}

ReactDOM.render( <MainPage />, document.getElementById('root'))