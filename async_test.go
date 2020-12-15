package async

import (
    "errors"
    "testing"
    "time"
)

func TestWaterfall(t *testing.T) {
    Waterfall(FFs{
        func(cb F, args ...interface{}) {
            t.Log("run 1")
            time.AfterFunc(time.Second, func() {
                cb(nil, "1", 2, 3.0)
            })
        },
        func(cb F, args ...interface{}) {
            t.Log("run 2")
            for _, v := range args {
                t.Log("arg:", v)
            }
            cb(nil, 1, 2, 3)
        },
    }, func(err error, args ...interface{}) {
        t.Log("result:", err)
        for _, v := range args {
            t.Log("arg:", v)
        }
    })
    time.Sleep(time.Second * 2)
}

func TestWaterfall2(t *testing.T) {
    Waterfall(FFs{
        func(cb F, args ...interface{}) {
            t.Log("run 1")
            time.AfterFunc(time.Second, func() {
                cb(errors.New("stop running"), "stopped", 2, 3.0)
            })
        },
        func(cb F, args ...interface{}) {
            t.Log("run 2")
            cb(nil, 1, 2, 3)
        },
    }, func(err error, args ...interface{}) {
        t.Log("result:", err, args)
        for _, v := range args {
            t.Log("arg:", v)
        }
    })
    time.Sleep(time.Second * 2)
}