// тема: Разбираемся с интерфейсами в Go. 
// Модульное тестирование и заглушки
package main()
import "fmt"
// Создаём свой интерфейс ShopModel. Он прекрасно подходит для
// интерфейса с описанием нескольких методов, и он должен описывать
// входные параметры-типы, а также типы возвращаемых значений.
type ShopModel interface{
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}
// Тип ShopDB удовлетворяет новому интерфейсу ShopModel, потому что
// у него есть два необходимых метода -- CountCustomers() и CountSales().
type ShopDB struct{
	*sql.DB
}
func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}
func(sdb *ShopDB) CountSales(since time.Time) (int, error){
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}
func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	shopDB := &ShopDB{db}
	sr := calculateSalesRate(shopDB)
	fmt.Printf(sr)
}
func calculateSalesRate(sm ShopModel) string {
	since := time.Now().Sub(24 * time.Hour)
	sales, err := sm.CountSales(since)
	if err != nil {
		return "", err
	}
	customers, err := sm.CountCustomers(since)
	if err != nil {
		return "", err
	}
	rate := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", rate), nil
}