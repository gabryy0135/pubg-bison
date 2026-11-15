package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.4" )

func q1BF8RzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsOyJJfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCT4JHJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNs9WcoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuXxwrWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fuR9YayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wM5bKutTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mluiPd8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEPyodk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRqLqp0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBKd5qz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKa9VQBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X10orXbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4mYUklOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XE4iRJ5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNHJOwsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WI6zTqTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biiocPscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jHVocHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMqwTSX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7eeGg09hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sk6rqYzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0m8ajtoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1b6zaLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGuof69xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 261gN27CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyuZcf3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0xUf3bUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqi4zMtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIFrQJxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0ROunbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQCDhPt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFfoKw9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kg5wAPxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTrhXmKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJPrZXucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tY5VJxUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxcdfwrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkpePYgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGQVtyeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgovRaDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjuPXkwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdKCsrO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5xLwFUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XK81ODyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAgHPDeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMWtAI0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func detJ8FegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUMgTNPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpOJpsg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMSxLe8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPXZ5rqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEmUvPjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Tu1G75GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WShPVGWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fepbep3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l0kK43csWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBTYdXRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJCr10MJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9HgRHVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZcTx8dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2b7368iEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WasmKrVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NCkpDRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NjDRDEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tr8ua6rKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55zJBw9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0VSwDGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Msx0T7lpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oxndvNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzrlSwxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xh7BpgqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuGpqEC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBj3Loy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy49L4emWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmfjOpDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiUTPv3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qZFonq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhAo0lKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a15NEZUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rJqNgjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29L5VCv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ua9g7EsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohtHFT1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CooszLeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxMr7zfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujRLpO9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3wbomHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXaFQdNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2W9y7XFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qOB7g0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mm3TxyP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKiRCcDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5jKZ02TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4CHN8fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHKEbCuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7HbG5tsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MS9mqqGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owzNyajRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0naPtZAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxzZupvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZifmVCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1Q8Ryy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9QQ5i36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgAv46hOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncm0sbqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5t8De27AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YefZWPlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUXNtwKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHzONNXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XilrkMtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKbdwQblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrFIfRFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3WvHexPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWLH8Z9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdSEQd1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWSzHbWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbvQoB3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3i2GdOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLHSfCe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD2EJfcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDoGmGAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDs7PFN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x21oSYiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhtAcc3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDD2BtAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQr03qFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IrLi16wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4w1CygjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBZhWuGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgaI2CVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGTvfqNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHXQkBasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRh2v9dSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UibhoEy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OCWxiPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjZxNsH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFrd1Xu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GY3PBNSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNctNlXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgGC246pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3QfHSY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqjOQnplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcvsrakoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4T4M8kZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrgDDyhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2hes7OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rVH8Ix5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCxlstWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9nsca6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fvlYRgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7P486pSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JP1I1v7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03UCrhNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L71sViKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wopq51wAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgjkpIOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbvYvGioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CA8yjZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wgRyrR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyZuLypCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYMwEf2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUp4GOJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQNTQEVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmKupphnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWFtE4znWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SedRIQloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LahMmaypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cqrA2unWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iubtBOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCeL1egeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T78vEP9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TScSBd8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJI6FfWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxBWSqhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJTcS5xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoadO2FvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmCDwllkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Mbh9jh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXRXquvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9ZXiZcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6h9Q1xLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhymJm36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJjwpmpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hx3ONh0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfuQPqq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCpfvUBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuRq3ao2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgRMRz6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8IPpDuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pfe7AdV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33N60nWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MEj2AxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozB6AtaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wtabGDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUjTlShDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTp15EehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyL2Ax4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imWZgIViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func coDIetBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKWZh4EuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXVJKx2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLOiwEXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUa3OKHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAIfWdDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3L3VrWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yX3spc4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JB40ATkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXBubxXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yh4f79iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaGD7H15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6tbJ8wMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofz8jHupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1cfgVL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfBjZgBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlzJQ0UoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AE7opRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJDjOKfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K24APtDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qMqbrR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYddZM8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MOkyUO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUfKKkl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNN0ywNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REo2gnNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ecs84OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RY2FqGc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJyvCMzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUtfxb29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gitN6pUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYh5v7WnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJY3B2PTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNVL698kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3trw6Ln8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yK53Xf1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0gSnxhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBuX23fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2fH2RCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCtnxJF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gdqo6FPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ta4uhmFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OnXJp9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gvEpR8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItHunRc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoKPQbsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpUD0VWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H71QnpHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o22K1MT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfH6dbDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNMcvWNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XWQ4COCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYcjun92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func coLK2RaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iM1o0DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kp7DGzqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eexVlsWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCWfHwNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95IreKWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IFg93h7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AcigmY9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ee26QXcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWd7dFdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8c8WzkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iq2dbhD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlXZx597Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHI6V2m2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIJMd87XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csULmQuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykj35MIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8ju7g9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJqQ5GYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OH5D2EQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4lKTjegiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaKaWHEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExdFyI1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4i5W86aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgV9HhemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtPw6b6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjM1EwEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLW98GrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnVD1ExtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrF4nGaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nBtuAD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfpmbMuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4e7wJfKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXbcCkzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oshaz1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPST7dKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFHP5M4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lgGoMxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOsBCvwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6ZLpSfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mr9LCFaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdwDEcyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTmRUCSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1R2Bn5igWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jmARzUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmxdEDcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UA2NtnA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov5mqevRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qlmmTuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyeYXFdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fwe7uI5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62RvVZHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sc03agimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pX5AiCdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViQpipuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYC3TpIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxnQOkVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBJ5nhZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGJ3ViTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmPVs9kjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uiLPQWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAFBxeYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYrUGW6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4AXmrEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX1ktyfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2TfxwqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZR19mdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRIxY34KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VogCv1m1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHijgM5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96SsKKPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5TKgLScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjEf0G4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdfF5mL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1YTP4qNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQLxamfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcZcuu3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsWXly1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cixMo9vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbT5sgX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbsLAXpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCGwzEc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yV1925YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5L2oFEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6SycmJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dknjm9VOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGLvd9mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUUfLCyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLn0KTwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhDSc9k6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSNVCYICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ihy6SGS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maFJuSCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edaYriP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WD3IzSUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pV7dsjgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2cCElwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKAPRIhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlrXqY4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEc9XKFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0nithi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7EYCjmZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6klrhxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i777m9HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBkvblyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2p1GWcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSYROkd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKBohCHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENIzzfNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqgYfbGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUVsczJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIG9CbJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7TIrUqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOny8k7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhILDUyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bGo4VirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLp8YcfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IFIiwMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcmfweB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2m9MBS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPjO4DzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWFnMlXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4QWvgiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GblaEwWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2ik93BWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYIokBHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VE7I6aUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K33EkbmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5P8q01iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8r4aceiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qT3SrKXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2YSJbr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2fYAGPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfwnhorDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SStpWlohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmFWUHkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RehcgNWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tRKpvSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B490O5PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qyy47JapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6eUN2QYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFWcAlqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZdlimypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XUKpVh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TtMCtliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpHL62ppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KB79UwSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXQ1vFM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaIIEnFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1u7ksgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWxaeWYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCgQIKKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Okg4pF67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2a4IwNpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfeouqHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cjekfTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhqdYL7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOGccYS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSvuZ5wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9kopKZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6cBFHJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xK7kj0gCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eh6D1lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOBZhTRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHfWIdEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAn3PDRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVLC59iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJOGmtBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVrSERR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqPxxNvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KpSNwFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3fOD6fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeYdomocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lwwp4YTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMctDl7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIwhF9JlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9W8YmqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI0DMRODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfbM1S5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ix7GXxVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvDHaUEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2mubOe6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBLMmmMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o60a1DMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5UJqO5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beijmrdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1nF6xE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMY6u0EPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HITeZ1yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTPDXLzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l2w3TGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsnv4ot4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuRbYL2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCJlbLoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtHjiukiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvMvDwA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHt5aX9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ox5uKtfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJfxxOlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAfEnjq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebFoRHFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjdTLxuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHw69wV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeW5wFg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y05SjJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gqExgFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhvRpFk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDVTsMvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dNaOB0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1sLxYdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcY4cwW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkeG1AVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UEymx1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZmh1haGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDW1vw1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bi6BD2GPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOMha9QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5kPCRGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtQ79SUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsmTovBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiPlW7aRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQCXITJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72UE7KeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uK1x5ZVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GO5eTktpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

