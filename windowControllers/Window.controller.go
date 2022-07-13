package window_controllers

import (
	"database/sql"
	"fmt"
	"net/http/httptest"

	"github.com/jchv/go-webview-selector"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/webviewKsyukulyator/controllers"
	"github.com/lxndrrud/webviewKsyukulyator/dto"
	"github.com/lxndrrud/webviewKsyukulyator/storage"
)

func NewWindowController(db *sqlx.DB) *WindowController {
	return &WindowController{
		productStorage:  storage.NewProductModel(db),
		categoryStorage: storage.NewCategoryModel(db),
		calcController:  controllers.NewCalculationController(),
	}
}

type WindowController struct {
	Window          webview.WebView
	productStorage  *storage.ProductModel
	categoryStorage *storage.CategoryModel
	calcController  *controllers.CalculationController
	server          *httptest.Server
}

func (c *WindowController) LoadBindings() {
	c.Window.Bind("loadProducts", c.LoadProducts)
	c.Window.Bind("addProduct", c.AddProduct)
	c.Window.Bind("deleteProduct", c.DeleteProduct)
	c.Window.Bind("loadCategories", c.LoadCategories)
	c.Window.Bind("addCategory", c.AddCategory)
	c.Window.Bind("deleteCategory", c.DeleteCategory)
	c.Window.Bind("addAmountToCalc", c.AddProductToCalculation)
	c.Window.Bind("deleteAmountFromCalc", c.DeleteProductFromCalculation)
	c.Window.Bind("getAmounts", c.GetProductsAmounts)
	c.Window.Bind("filterByCategory", c.productStorage.GetByCategoryId)
}

func (c *WindowController) SetupWindow(srv *httptest.Server) {
	c.server = srv
	url := fmt.Sprintf("%s/frontend/pages/main/main.html", srv.URL)
	c.Window = webview.New(false)
	defer c.Window.Destroy()

	c.Window.SetTitle("Ксюкулятор")
	c.Window.SetSize(1200, 800, webview.HintNone)
	c.LoadBindings()
	c.Window.Navigate(url)
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
	return products
}

func (c *WindowController) AddProduct(title string, cost float32, categoryId int64) bool {
	idCategory := sql.NullInt64{}
	if categoryId != 0 {
		idCategory.Valid = true
		idCategory.Int64 = categoryId
	}
	err := c.productStorage.AddProduct(dto.Product{
		Title:      title,
		Cost:       cost,
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
	err := c.categoryStorage.DeleteCategory(idCategory)
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось удалить категорию: " + err.Error())
		return false
	}
	return true
}

func (c *WindowController) AddProductToCalculation(idProduct int64, amount float32, amountCost float32) dto.ProductAmountTable {
	product, err := c.productStorage.GetById(idProduct)
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось расчитать стоимость продукта: " + err.Error())
		return dto.ProductAmountTable{}
	}
	return c.calcController.AddProductAmount(&dto.ProductAmount{
		Product:    product.Product,
		Amount:     amount,
		AmountCost: amountCost,
	})
}

func (c *WindowController) DeleteProductFromCalculation(idProduct int64) dto.ProductAmountTable {
	return c.calcController.DeleteProductAmount(idProduct)
}

func (c *WindowController) GetProductsAmounts() dto.ProductAmountTable {
	return c.calcController.GetProducts()
}

func (c *WindowController) FilterByCategory(idCategory int64) []dto.ProductFull {
	fmt.Println("Фильтр по категориям!")
	products, err := c.productStorage.GetByCategoryId(idCategory)
	if err != nil {
		fmt.Println(err)
		c.SetError("Не удалось расчитать стоимость продукта: " + err.Error())
		return []dto.ProductFull{}
	}
	return products
}
