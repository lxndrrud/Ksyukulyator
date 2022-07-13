const CategoriesPage = () => {
    let [categories, setCategories] = React.useState([])
    let [error, setError] = React.useState(null)
    let [isLoading, setIsLoading] = React.useState(true)
    async function delCategory(idCategory) {
        try {
            let result = await deleteCategory(parseInt(e.target.getAttribute('categoryid')))
            if (result) setCategories(await loadCategories())
        } catch (e) {
            setError(e)
        }
    }
    React.useEffect(async () => {
        try {
            setCategories(await loadCategories())
            setIsLoading(false)
        } catch (e) {
            setError(e)
        }
        
    }, [])
    return (
        <div>
            {
                isLoading
                ? <div>Загрузка...</div>
                :
                <div className="column"> 
                    <h1 className="ksyu">Ксюкулятор • Категории</h1>
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
                    <div id="content"  >
                        <div className="overflow_container">
                            <table id="productCategoryTable">
                                <thead>
                                    <tr>
                                        <td>Название категории</td>
                                        <td>Контроль</td>
                                    </tr>
                                </thead>
                                <tbody>
                                    {categories && categories.map(category => (
                                        <tr>
                                            <td>{category.category_title}</td>
                                            <td className="column">
                                                {/*<button>Изменить</button>*/}
                                                <button 
                                                    onClick={async () => { await delCategory(category.category_id)}} 
                                                    className="button">
                                                    Удалить
                                                </button>
                                            </td>
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        </div>
                        <a href="/frontend/pages/addCategory/addCategory.html" className="button">
                            Добавить категорию
                        </a>
                    </div>
                    
                </div>
            }
        </div>
        
        
    )
}

ReactDOM.render(<CategoriesPage />, document.getElementById("root"))