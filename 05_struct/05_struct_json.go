package main

import (
    "fmt"
    "encoding/json"
)

type student struct {
    name string
    id int
}

type Student struct {
    Name string
    Id int
}
func Newstundet(name string, id int) student {
    return student {
        name: name,
        id: id,
    }
}
type class struct {
    title string
    students []student
}

type ClassT struct {
    Title string `json:"title"`
    Students []Student `json:"student_list"`
}

func NewStudent(name string, id int) Student {
    return Student {
        Name: name,
        Id: id,
    }
}
func main() {
    class1 := class {
        title: "3-1",
        students: make([]student, 0, 10),
    }
    class2 := ClassT {
        Title: "3-4",
        Students: make([]Student, 0, 10),
    }
    for i := 0; i < 10; i ++ {
        newstudent := Newstundet(fmt.Sprintf("student%02d", i), i)
        class1.students = append(class1.students, newstudent)

        newStudent := NewStudent(fmt.Sprintf("student%02d", i), i)
        class2.Students = append(class2.Students, newStudent)
    }
    fmt.Println(class1)



    // marshal, the diffence of class1_json and class2_json
    class1_json, err := json.Marshal(class1)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%T, %s\n", class1_json, class1_json)

    class2_json, err2 := json.Marshal(class2)
    if err2 != nil {
        fmt.Println(err2)
        return
    }
    fmt.Printf("%T, %s\n", class2_json, class2_json)
}
