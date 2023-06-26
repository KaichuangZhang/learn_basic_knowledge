# reflect
运行时，修改。

常用场景：
- 1
- 结构体反射

reflect包。


### reflect
- TypeOf：获取类型； 返回值：*reflect.rtype类型；
  - reflect.rtype.Name()
  - reflect.rtype.Kind()
  - 基础类型中Name() = Kind()
- ValueOf：获取变量值；
- 