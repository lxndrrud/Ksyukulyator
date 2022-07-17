const AddProductPage = () => {
    let [categories, setCategories] = React.useState(null)
    let [title, setTitle] = React.useState("")
    let [cost, setCost] = React.useState("")
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
            if (!title) {
                setError("Ксю, ты не ввела название продукта!")
                return
            }
            if (!cost || parseFloat(cost) === 0 ) {
                setError("Ксю, ты не ввела стоимость продукта!")
                return
            }
            let result = await addProduct(title, parseFloat(cost), parseInt(categoryId))
            if (result) window.location.href = "/pages/main/main.html"
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
                                onChange={(e) => setCost(e.target.value) } />
                        </div>
                        
                        <select id="productCategory" onChange={(e) => setCategoryId(e.target.value) } >
                            <option value="0">
                                Без категории
                            </option>
                            {categories && categories.map(category => (
                                <option value={category.category_id}>
                                    {category.category_title}
                                </option>
                            ))}
                        </select>
                        <button onClick={async() => { await saveProduct() }} className="button">Сохранить</button>
                        
                        <a href="/pages/main/main.html" className="button">Отменить</a>
                    </div>
                    
                </div>
        }
        
    </div>)
    
}

ReactDOM.render(<AddProductPage />, document.getElementById('root'))