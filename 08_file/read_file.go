package main

import (
    "io/ioutil"
    "os/exec"
    "fmt"
    "strings"
)

func main() {
    // read file
    b, err := ioutil.ReadFile("./nginx.conf")
    if err == nil {
        fmt.Println(b)
    } else {
        fmt.Println(err)
    }
    run_cmd := fmt.Sprintf("sed -i 's#access_log.*#access_log /var/log/nginx/access.log main;#g' nginx.conf")
    fmt.Println(run_cmd)
    cmd_array := strings.Split(run_cmd, " ")
    shell_cmd := cmd_array[0]
    shell_opt := cmd_array[1]
    //shell_sub_cmd :=  cmd_array[2] + " " + cmd_array[3] + " " + cmd_array[4]
    shell_sub_cmd := "s#access_log.*#access_log off;#g"
    shell_file := cmd_array[5]
    fmt.Println(shell_cmd)
    fmt.Println(shell_opt)
    fmt.Println(shell_sub_cmd)
    fmt.Println(shell_file)
    //new_err := exec.Command(cmd_array[0], cmd_array[1], cmd_array[2] + " " + cmd_array[3] + " " + cmd_array[4], cmd_array[5]).Run()
    new_err := exec.Command(shell_cmd, shell_opt, shell_sub_cmd, shell_file).Run()
    if new_err != nil {
        fmt.Println(new_err)
    }
    // output the file content
}
