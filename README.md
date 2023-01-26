# golang-goroutine
## Membuat Test Goroutine dengan perintah go (function)

**Goroutine** merupakan salah satu bagian paling penting dalam concurrent programming di Go. Salah satu yang membuat goroutine sangat istimewa adalah eksekusi-nya dijalankan di multi core processor. Kita bisa tentukan berapa banyak core yang aktif, makin banyak akan makin cepat.

Buat Project Golang Baru dan 1 file berikut :
> goroutine_test.go
```go
package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display Number ", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
```

## Penjelasan Function 

**Func DisplayNumber** adalah function menampilkan text Display Number dengan menyisipkan variable number yang akan dipanggil oleh goroutine.

**Func TestDisplayNumber** adalah function main atau function utama yang akan menampung testing goroutine dengan metode looping number yang dimana terdapat perintah
go Displaynumber(i) yaitu perintah proses concurrency (multithreads) yang akan menjalankan peroses kinerja atau perintah mana yang akan terselesaikan secara acak (yang pertama selesai).

### Berikut Hasil dari testing goroutine :
![alt text](https://github.com/febysuasaf/golang-goroutine/blob/main/hasil_run.png)

#### Dari hasil tersebut kita lihat number yang sudah kita looping tidak berurutan dikarenakan proses tersebut dilakukan secara bersamaan namun dengan mesin yang berbeda (cpu/engine)
