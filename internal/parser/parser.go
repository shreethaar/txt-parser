package parser

import (
    "io/ioutil"
    "os"
)

func ReadFile(filePath string) (string,error) {
    data,err:=ioutil.ReadFile(filePath)
    if err!=nil{
        return "",err
    }
    return string(data),nil
}

func UpdateFile(filePath,content string) error {
    return ioutil.WriteFile(filePath, []byte(content),os.ModePerm)
}

