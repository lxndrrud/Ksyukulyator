const AddCategoryPage = () => {
    let [error, setError] = React.useState(null)
    let [categoryTitle, setCategoryTitle] = React.useState("")

    async function saveCategory(event) {
        try {
            if (!categoryTitle) {
                setError("Ксю, ты не ввела название категории!") 
                return
            }
            let result = await addCategory(categoryTitle)
            if (result) window.location.href = "/pages/categories/categories.html"
        } catch(e) {
            setError(e)
        }
    }

    return (
        <div className="column">
            {
                error 
                ?
                    <span>{error}</span>
                : null
            }
            <div id="content">
                <div className="input_group">
                    <label for="#categoryTitle">Название категории</label> 
                    <input type="text" className="input_text" id="categoryTitle" 
                        onChange={(e) => setCategoryTitle(e.target.value) }/>
                </div>
                
                <button onClick={saveCategory} className="button">Сохранить</button>
                <a href="/pages/categories/categories.html" className="button">Отменить</a>
            </div>
        </div>
    )
} 

ReactDOM.render(<AddCategoryPage />, document.getElementById('root'))