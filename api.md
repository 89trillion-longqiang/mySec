
请求参数一个字符串，返回表达式的计算结果
接口地址  /C1/calculator
请求方式  GET
请求示例  http://127.0.0.1:9000/C1/calculator?exp=3+3*4/2
参数说明   exp  类型 string  输入需要一个字符串包含正整数、加(+)、减(-)、乘(*) 、除(/)的算数表达式(括号除外)
返回参数说明 condition 表示状态，pass表示输入的字符串合法  error表示有误 
           result 表示结果 如果condition为pass则显示表达式的计算结果，如果condition为error则显示出现错误的原因
成功示例  {"condition":"pass","result":9}
错误示列  {"condition":"error","result":"exp error"}
