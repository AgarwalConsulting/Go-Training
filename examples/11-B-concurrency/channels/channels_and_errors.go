ch := make(chan Value)
errCh := make(chan error)

select {
case val := <- ch:
  // Do something with val
case <-errCh:
  // Handle th error
}

---

type ValueWithErr struct{
    val Value
    err error
}

dataCh := make(chan ValueWithErr)
