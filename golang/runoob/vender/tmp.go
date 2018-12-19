package vender
import "fmt"

func ShowTemp() {
   var n [10]int /* n 是一个长度为 10 的数组 */

   /* 为数组 n 初始化元素 */         
   for i := 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j := 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func TestStruct(){
	var Book1 Books        /* Declare Book1 of type Book */
   var Book2 Books        /* Declare Book2 of type Book */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook(Book1)

   /* 打印 Book2 信息 */
   printBook(Book2)
}
func printBook(book Books) {
	fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}

func TestSlice(){
	// var numbers []int
	numbers := []int{}
	printSlice(numbers)
	numbers = append(numbers,0, 1)
	printSlice(numbers)
	// numbers = append(numbers,1)
	// printSlice(numbers)
	numbers = append(numbers, 2,3,4)
    printSlice(numbers)
    numbers = append(numbers,6)
    printSlice(numbers)
    var numbers1 = make([]int, len(numbers), (cap(numbers))*2)
    copy(numbers1,numbers)
    printSlice(numbers1)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}