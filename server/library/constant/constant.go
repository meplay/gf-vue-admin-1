package constant

const (
	ConfigEnv      = "GVA_CONFIG"
	ConfigFile     = "config/viper.yaml"
	SerializerGob  = "gob"
	SerializerJson = "json"

	BasePath = "./template"
	AutoPath = "generate/"

	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

	ExcelDir      = "./public/excel"
	FinishDir     = "./fileDir/"
	BreakpointDir = "./breakpointDir/"

	End              string = "end"        // 结束节点
	Flow             string = "flow"       // 连线
	Start            string = "start"      // 开始节点
	Process          string = "process"    // 基础节点
	UserTask         string = "userTask"   // 审批节点
	MailTask         string = "mailTask"   //邮件节点
	TimerStart       string = "timerStart" // 定时节点
	ScriptTask       string = "scriptTask" //脚本节点
	ReceiveTask      string = "receiveTask"
	MessageStart     string = "messageStart"     // 消息节点
	ParallelGateway  string = "parallelGateway"  // 并行网关
	ExclusiveGateway string = "exclusiveGateway" // 排他网关
	InclusiveGateway string = "inclusiveGateway" // 包容网关
)
