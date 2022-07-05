package window_controllers

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"strconv"

	"github.com/jchv/go-webview-selector"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/webviewKsyukulyator/dto"
	"github.com/lxndrrud/webviewKsyukulyator/storage"
)

func NewWindowController(db *sqlx.DB) *WindowController {
	return &WindowController{
		productStorage:  storage.NewProductModel(db),
		categoryStorage: storage.NewCategoryModel(db),
	}
}

type WindowController struct {
	Window          webview.WebView
	productStorage  *storage.ProductModel
	categoryStorage *storage.CategoryModel
	server          *httptest.Server
}

func (c *WindowController) LoadBindings() {
	c.Window.Bind("loadProducts", c.LoadProducts)
	c.Window.Bind("addProduct", c.AddProduct)
	c.Window.Bind("deleteProduct", c.DeleteProduct)
	c.Window.Bind("loadCategories", c.LoadCategories)
	c.Window.Bind("addCategory", c.AddCategory)
	c.Window.Bind("deleteCategory", c.DeleteCategory)
}

func (c *WindowController) SetupWindow(srv *httptest.Server) {
	c.server = srv
	url := fmt.Sprintf("%s/frontend/pages/main/main.html", srv.URL)
	c.Window = webview.New(false)
	defer c.Window.Destroy()

	c.Window.SetTitle("Ксюкулятор")
	c.Window.SetSize(800, 600, webview.HintNone)
	c.Window.Navigate(url)
	c.LoadBindings()
	c.Window.Run()
}

func (c *WindowController) SetError(message string) {
	c.Window.Eval(fmt.Sprintf("setError(%s)", message))
}

func (c *WindowController) LoadProducts() []dto.ProductFull {
	products, err := c.productStorage.GetAll()
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось загрузить продукты: " + err.Error())
		return []dto.ProductFull{}
	}
	fmt.Println(products)
	return products
}

func (c *WindowController) AddProduct(title, cost string, categoryId int64) bool {
	fmt.Println(title, cost, categoryId)
	parsedCost, _ := strconv.ParseFloat(cost, 32)
	idCategory := sql.NullInt64{}
	if categoryId != 0 {
		idCategory.Valid = true
		idCategory.Int64 = categoryId
	}
	err := c.productStorage.AddProduct(dto.Product{
		Title:      title,
		Cost:       float32(parsedCost),
		IdCategory: idCategory,
	})
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось добавить продукт: " + err.Error())
		return false
	}
	return true
}

func (c *WindowController) DeleteProduct(id int64) bool {
	err := c.productStorage.DeleteProduct(id)
	if err != nil {
		fmt.Println(err)
		c.SetError("Ошибка во время удаления продукта: " + err.Error())
		return false
	}
	return true
}

func (c *WindowController) LoadCategories() []dto.Category {
	categories, err := c.categoryStorage.GetAll()
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось загрузить категории: " + err.Error())
		return []dto.Category{}
	}
	fmt.Println(categories)
	return categories

}

func (c *WindowController) AddCategory(title string) bool {
	err := c.categoryStorage.AddCategory(dto.Category{
		Title: title,
	})
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось добавить категорию: " + err.Error())
		return false
	}
	return true
}

func (c *WindowController) DeleteCategory(idCategory int64) bool {
	fmt.Println("Удаление категории")
	err := c.categoryStorage.DeleteCategory(idCategory)
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось удалить категорию: " + err.Error())
		return false
	}
	return true
}