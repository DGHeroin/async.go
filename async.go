package async
// MIT License
// Blake@github.com/DGHeroin/async.go

type NextCallback func(err error, args ...interface{})
type Callback func(cb NextCallback, args ...interface{})

func Waterfall(fs []Callback, result NextCallback) {
    var (
        next func(err error, args ...interface{})
        v    Callback
    )
    next = func(err error, args ...interface{}) {
        if len(fs) == 0 {
            if result != nil {
                result(err, args)
            }
            return
        }
        if err != nil {
            if result != nil {
                result(err, args...)
            }
            return
        }
        err = nil
        v, fs = fs[0], fs[1:]
        v(next, args...)
    }
    next(nil)
}

