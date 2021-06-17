package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	//config := netgsm.Config{SmsCompany: "NETGSM", SmsMsgHeader: "", SmsUserCode: "", SmsPassword: "", ApiUrl: "https://api.netgsm.com.tr/sms/send/xml"}
	//api := &netgsm.API{config}
	//request := netgsm.SendSmsRequest{}
	//request.MainBody.Body.Msg = "test"
	//request.MainBody.Body.No = "905555555555"
	//send, _ := api.Sms(request)
	//if send {
	//	fmt.Println("mesaj iletildi")
	//} else {
	//	fmt.Println("hata oluştu")
	//}

	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("MessageBoxW")
	var MB_YESNOCANCEL = 0x00000003

	ret, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
		uintptr(MB_YESNOCANCEL))
	fmt.Printf("Return: %d\n", ret)
}
