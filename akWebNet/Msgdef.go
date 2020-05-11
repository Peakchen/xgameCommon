package akWebNet

//message ID 
const (
	MID_login = 1  //上线 (获取自己的数据)
	MID_logout = 2 //离线
	MID_move = 3   //移动
	MID_Bump = 4   //撞击
	MID_HeartBeat = 5 //心跳
	MID_StarBorn = 6  //星星出生
	MID_GM  = 7 	  //gm
	MID_Online4Other = 8 //上线获取别人数据显示或者广播自己数据给别人显示
	MID_Register = 9 	 //注册
	MID_SyncPos = 10     //同步位置
	MID_MonsterInfo = 11   //同步moster信息
	MID_LogicFrameSync = 12  //逻辑帧同步其他客户端moster信息
)