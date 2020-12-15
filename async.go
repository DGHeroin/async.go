package async
// MIT License
// Blake@github.com/DGHeroin/async.go

type F func(err error, args ...interface{})
type FF func(cb F, args ...interface{})
type FFs []FF

func Waterfall(fs FFs, result F) {
    var (
        next func(err error, args ...interface{})
    )
    next = func(err error, args ...interface{}) {
        if len(fs)==0{
            result(err, args)
            return
        }
        if err != nil {
            result(err, args...)
            return
        }
        err = nil
        v := fs[0]
        fs = fs[1:]
        v(next, args...)
    }
    next(nil)
}