
�
proto/MSG_Login.proto	MSG_Login"C
CS_UserBind_Req
Account (	RAccount
Passwd (	RPasswd"�
CS_UserRegister_Req
Account (	RAccount
Passwd (	RPasswd"
DeviceSerial (	RDeviceSerial

DeviceName (	R
DeviceName"=
SC_UserRegister_Rsp&
Ret (2.MSG_Login.ErrorCodeRRet"�
CS_Login_Req
Account (	RAccount
Passwd (	RPasswd"
DeviceSerial (	RDeviceSerial

DeviceName (	R
DeviceName"6
SC_Login_Rsp&
Ret (2.MSG_Login.ErrorCodeRRet*{
SUBMSG	
Begin 
CS_UserBind
SC_UserBind
CS_UserRegister
SC_UserRegister
CS_Login
SC_Login*L
	ErrorCode
Invalid 
Success
Fail
UserNotExistOrPasswdErrbproto3