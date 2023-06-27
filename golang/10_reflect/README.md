# reflect
## 概念
- 运行时，修改。

## 常用场景：
- ORM
- 文件解析
- 结构体反射

## reflect一般用法 (./reflect1.go && ./reflect2.go)
- TypeOf：获取类型； 返回值：*reflect.rtype类型；
  - reflect.rtype.Name()
  - reflect.rtype.Kind()
  - 基础类型中Name() = Kind()
- ValueOf：获取变量值；

## struct reflect (./struct_reflect.go)
- Field
  - Field(int): 根据id，获取字段;
  - FieldByName(string): 根据string，获取字段
- Method
  - Method(int)：
  - MethodByName(string):