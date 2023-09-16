package product


type MyMapProd map[string]int

func NewProduct() *MyMapProd {
    product := make(MyMapProd)
    product["хлеб"] = 50
    product["молоко"] = 100
    product["масло"] = 200
    product["колбаса"] = 500
    product["соль"] = 20
    product["огурцы"] = 200
    product["сыр"] = 600
    product["ветчина"] = 700
    product["буженина"] = 900
    product["помидоры"] = 250
    product["рыба"] = 300
    product["хамон"] = 1500
    return &product
}

func GetDelecacies (product *MyMapProd) MyMapProd {
    dalecacies := make(MyMapProd)
    for name, price := range *product {
        if price > 500 {
            dalecacies[name] = price
        }
    }
    return dalecacies
}
