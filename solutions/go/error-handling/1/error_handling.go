package erratum

// import (
//     "errors"
// )

func Use(opener ResourceOpener, input string) (err error) {
    var res Resource

    for {
        res, err = opener()
        if err == nil {
            break
        }
        if _, ok := err.(TransientError); ok {
            continue
        }
        return err
    }
    defer res.Close()

    defer func() {
        if r := recover(); r != nil {
            if frobErr, ok := r.(FrobError); ok {
                res.Defrob(frobErr.defrobTag)
            }
            if recoveredErr, ok := r.(error); ok {
                err = recoveredErr
            } else {
                panic(r)
            }
        }
    } ()

    res.Frob(input)

    return nil
}
