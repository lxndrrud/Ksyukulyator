const AddProductPage = () => {
    let [categories, setCategories] = React.useState(null)
    let [title, setTitle] = React.useState('')
    let [cost, setCost] = React.useState('')
    let [isLoading, setIsLoading] = React.useState(true)
    let [categoryId, setCategoryId] = React.useState(0)
    let [error, setError] = React.useState(null)
    //const navigate = ReactRouterDOM.useNavigate()
    React.useEffect(async () => {
        try {
            let result = await loadCategories()
            setCategories(result)
            setIsLoading(false)
        } catch (e) {
            setError(e)
        }
        
    }, [])
    async function saveProduct(event) {
        try {
            let result = await addProduct(title, cost, categoryId)
            if (result) window.location.href = "/frontend/pages/main/main.html"
        } catch (e) {
            setError(e)
        }
    }
    /*
<!--button @click="await navigateTo('/frontend/pages/main/main.html')">Отменить</button-->
{'<ReactRouterDOM.Link to="/main" >Отменить</ReactRouterDOM.Link>'}
    */

    return (
    <div>
        {
            isLoading
            ? 
                <span>Загрузка...</span>
            : 
                <div className="column">
                    <h1 className="ksyu">Ксюкулятор • Добавление продукта</h1>
                    <nav>
                        <a href="/frontend/pages/main/main.html" className="link" >Калькулятор</a>
                        <a href="/frontend/pages/categories/categories.html" className="link">Категории</a>
                    </nav>
                    {
                        error 
                        ?
                            <span>{error}</span>
                        : null
                    }
                    
                    <div id="content">
                        <div className="input_group">
                            <label for="#productTitle" >Название продукта</label>
                            <input type="text" id="productTitle" className="input_text" 
                                onChange={(e) => setTitle(e.target.value) } />
                        </div>
                        
                        <div className="input_group">
                            <label for="#productCost">Стоимость продукта, г углеводов / 100 г продукта</label>
                            <input type="text" id="productCost" className="input_text"
                                onChange={(e) => setCost(parseFloat(e.target.value)) } />
                        </div>
                        
                        <select id="productCategory" onChange={(e) => setCategoryId(parseInt(e.target.value)) } >
                            <option categoryid="0">
                                Без категории
                            </option>
                            {categories && categories.map(category => (
                                <option value={category.category_id}>
                                    {category.category_title}
                                </option>
                            ))}
                        </select>
                        <button onClick={saveProduct} className="button">Сохранить</button>
                        
                        <a href="/frontend/pages/main/main.html" className="button">Отменить</a>
                    </div>
                    
                </div>
        }
        
    </div>)
    
}

ReactDOM.render(<AddProductPage />, document.getElementById('root'))