const MainPage = () => {
    //let [idProduct, setIdProduct] = React.useState(0)
    let [amount, setAmount] = React.useState(0.0)
    let [amountCost, setAmountCost] = React.useState(0.0)
    let [sum, setSum] = React.useState(0.0)
    let [isLoading, setIsLoading] = React.useState(true) 
    let [error, setError] = React.useState(null)
    let [products, setProducts] = React.useState([])
    let [productsAmounts, setProductsAmounts] = React.useState([])
    let [isAmountDisabled, setIsAmountDisabled ] = React.useState(true)

    React.useEffect(async () => {
        await parseProducts()
        await parseProductsAmounts()
    }, [])

    async function parseProducts() {
        try {
            setIsLoading(true)
            let productsLocal = await loadProducts()
            setProducts(productsLocal)
            setIsLoading(false)
        } catch(e) {
            setError(e)
        }
    }
    async function parseProductsAmounts() {
        try {
            let productsAmountsLocal = await getAmounts()
            setProductsAmounts(productsAmountsLocal.products)
            setSum(productsAmountsLocal.sum)
        } catch(e) {
            setError(e)
        }
    }
    async function calculate(idProduct) {
        try {
            let table = await addAmountToCalc(
                parseInt(idProduct), 
                parseFloat(amount), parseFloat(amountCost))
            setProductsAmounts(table.products)
            setSum(table.sum)
            //await parseProductsAmounts()
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
    async function removeFromCalc(idProduct) {
        try {   
            let table = await deleteAmountFromCalc(parseInt(idProduct))
            setSum(table.sum)
            setProductsAmounts(table.products)
            //await parseProductsAmounts()
        } catch (e) {
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
                    <div className="column">
                        {
                            error
                            ? <span>{error}</span>
                            : null
                        }
                        <div id="content"> 
                            <div className="column">
                                <h3>Продукты</h3>
                                <div className="overflow_container">
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
                                                (<tr>
                                                        <td>{product.product_title}</td>
                                                        <td>{product.product_cost}</td>
                                                        <td>{product.category_title}</td>
                                                        <td className="column">
                                                            <button 
                                                                className="button" 
                                                                onClick={async (e) => { await calculate(product.product_id)}}>Посчитать</button>
                                                            <button
                                                                className="button" 
                                                                onClick={async (e) => { await delProduct(product.product_id)}}>Удалить</button>
                                                        </td>
                                                </tr>)
                                                
                                            )}
                                        </tbody>
                                        
                                    </table>
                                </div>
                            </div>
                            <div className="column">
                                <h3>Контроль продуктов</h3>
                                <div id="controlButtons">
                                    <button onClick={async (e) => {await parseProducts()}} className="button">
                                            Обновить список
                                    </button>
                                    <a href="/pages/addProduct/addProduct.html/" className="button">
                                        Добавить продукт
                                    </a>
                                    <select id="productCategory" onChange={async () => { await filterProductsByCategory() }}>
                                        
                                    </select>
                                </div>
                            </div>
                            <div className="column">
                                <h3>Продукты на расчёт</h3>
                                <div className="overflow_container">
                                    <table id="amountsTable">
                                        <thead>
                                            <tr>
                                                <td>Название продукта</td>
                                                <td>Количество продукта, г</td>
                                                <td>Стоимость количества продукта, х.е.</td>
                                                <td>Контроль</td>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {productsAmounts && productsAmounts.map(amount => (
                                                <tr key={amount.product_amount_id}>
                                                    <td>{amount.product_amount_product.product_title}</td>
                                                    <td>{amount.product_amount_amount}</td>
                                                    <td>{amount.product_amount_cost}</td>
                                                    <td>
                                                        <button 
                                                            onClick={async () => { 
                                                                await removeFromCalc(amount.product_amount_id)
                                                            }} 
                                                            className="button">
                                                                Убрать из расчёта
                                                        </button>
                                                    </td>
                                                </tr>
                                            ))
                                            }
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                            <div className="column">
                                <h3>Контроль расчёта продуктов</h3>
                                <div id="calculateControllers">
                                    <h4 style={{marginTop: "2px", marginBottom: "10px;"}} >Выбрать тип расчёта</h4>
                                    <label for="#amount">Количество продукта, г</label>
                                    <input type="number" id="amount" 
                                        className="input_text" disabled={isAmountDisabled} 
                                        defaultValue={0} onChange={(e) => { setAmount(e.target.value) }}
                                        value={amount} />
                                    <button onClick={(e) => { setIsAmountDisabled(false); setAmountCost(0) }} 
                                    className="button">
                                        Расчитать стоимость в хлебных единицах
                                    </button>
                                    <label for="#amountCost">Количество хлебных единиц, х.е.</label>
                                    <input type="number" id="amountCost"  
                                        className="input_text" disabled={!isAmountDisabled} 
                                        defaultValue={0} onChange={(e) => { setAmountCost(e.target.value) }}
                                        value={amountCost} />
                                    <button onClick={(e) => { setIsAmountDisabled(true); setAmount(0.0) }} className="button">
                                        Рассчитать количество продукта
                                    </button>

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