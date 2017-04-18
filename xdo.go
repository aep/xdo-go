package xdo;

/*
#include <stdlib.h>
#include <xdo.h>
#cgo LDFLAGS: -lxdo
*/
import "C"

import (
    "unsafe"
    "runtime"
)

type Xdo struct {
    x *C.xdo_t
}

func xdo_finalizer(x *Xdo) {
    C.xdo_free (x.x);
}

func NewXdo() *Xdo {
    x := C.xdo_new (nil);
    r := &Xdo{x};
    runtime.SetFinalizer(r, xdo_finalizer);
    return r;
}

type Window struct {
    w C.Window
    x *Xdo
}

func (x *Xdo) GetActiveWindow() Window {
    window := C.Window(0);
    C.xdo_get_active_window(x.x, &window)
    return Window{window, x};
}

func (w *Window) GetName() string {
    var name_ret        *C.uchar;
    var name_len_ret    C.int;
    var whatever        C.int;
    C.xdo_get_window_name(w.x.x, w.w, &name_ret, &name_len_ret, &whatever);
    str := C.GoBytes(unsafe.Pointer(name_ret), name_len_ret);
    C.free(unsafe.Pointer(name_ret));
    return string(str);
}


func (w *Window)  SendKeysequence(seq string, delay uint) error{
    seq_ := C.CString(seq);
    C.xdo_send_keysequence_window (w.x.x, w.w, seq_, C.useconds_t(delay));
    C.free(unsafe.Pointer(seq_));
    return nil
}

