package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// this is a read closer
// io.Reader, io.Writer, io.Closer, etc. are nothing but abstractions.
//  don’t care what you’re reading from (a file, a TCP socket, memory, even a compressed stream).
func getLinesChannel(r io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {

		defer close(out)
		defer r.Close()
		// cline := ""
		// for {
		// 	//since tcp is oneshotdata
		// 	// data := make([]byte, 8)
		// 	// n, err := r.Read(data)

		// 	// if err != nil || errors.Is(err, io.EOF) {
		// 	// 	break
		// 	// }

		// 	// str := string(data[:n])
		// 	// parts := strings.Split(str, "\n")

		// 	// for i := 0; i < len(parts)-1; i++ {
		// 	// 	out <- fmt.Sprintf("read: %s%s\n", cline, parts[i])
		// 	// 	cline = ""
		// 	// }
		// 	// cline += parts[len(parts)-1]

		// 	//we use a buffer

		// }

		sc := bufio.NewScanner(r)

		for sc.Scan() {
			out <- "read\t" + sc.Text()
		}
	}()

	return out
}

const address string = ":42069"

func main() {

	//taking data 8bytes at a time

	// msgFrom, err := os.Open("messages.txt")
	// if err != nil {
	// 	log.Fatal("eror opening the file")
	// }

	// lines := getLinesChannel(f)

	// for line := range lines {
	// 	fmt.Println("%s\n", line)
	// }

	msgListner, err := net.Listen("tcp", address)
	defer msgListner.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := msgListner.Accept()

		fmt.Println("accepting from %s\n", conn.RemoteAddr())

		if err != nil {
			log.Fatal(err)
		}

		for line := range getLinesChannel( conn  ) {
			fmt.Println("read %s\n", line)
		}

	}

}
