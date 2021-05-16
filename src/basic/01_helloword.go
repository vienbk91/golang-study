/**
Mỗi file go đều bắt đầu với câu lệnh package name.
Package được dùng để phân chia code cho các mục đích sử dụng khác nhau và tăng tính tái sử dụng,
Ở đây, package name là main.
 */
package main

/**
import “fmt” - package fmt được import vào file và
được sử dụng bên trong hàm chính để in ra màn hình dữ liệu dạng text.
 */
import "fmt"

/**
func main() - phần này là một hàm đặc biệt.
Chương trình được thực thi bắt đầu từ một hàm chính.
Hàm chính của chương trình luôn được đặt trong package main.
Cặp dấu ngoặc nhọn “{ và }” xác định phần bắt đầu và kết thúc của một hàm.
 */
func main() {
	/**
	Hàm Println của package fmt dùng để in dữ liệu ra màn hình.
	 */
	fmt.Println("Helloword")
}

