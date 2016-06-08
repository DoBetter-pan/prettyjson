/**
* @file prettyjson.go
* @brief make pretty json
* @author yingx
* @date 2016-05-27
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"bytes"
	"io/ioutil"
	"strings"
)

const (
	PADDINGCOUNT = 4
	PADDINGSTRING = " "
)

type params struct {
	paddingCount        int
	paddingString       string
}

func handleCommandLine() (p *params, files []string) {
	p = &params{}
	flag.IntVar(&p.paddingCount, "paddingCount", PADDINGCOUNT, "padding width")
	flag.StringVar(&p.paddingString, "paddingString", PADDINGSTRING, "padding string")
	flag.Parse()

	files = flag.Args()
	return p, files
}

func makePrettyJson(json []byte, p *params) string {
    l := len(json) + 512
    prettyjson := make([]byte, 0, l)
    buffer := bytes.NewBuffer(prettyjson)

    indent := 0
    s_quoted := 0
    d_quoted := 0
    last_newline := 1

    for _, v := range json {
        if v == '\'' {
            s_quoted = s_quoted ^ 0x1
        }
        if v == '"' {
            d_quoted = d_quoted ^ 0x1
        }
        if s_quoted == 0 && d_quoted == 0 {
            switch v {
            case '{':
                if last_newline == 0 {
                    buffer.WriteByte('\n')
                    buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
                }
                buffer.WriteByte(v)
                buffer.WriteByte('\n')
                last_newline = 1
                indent++
                buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
            case '}':
                indent--
                buffer.WriteByte('\n')
                last_newline = 1
                buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
                buffer.WriteByte(v)
                //buffer.WriteByte('\n')
                //buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
            case '[':
                buffer.WriteByte(v)
                buffer.WriteByte('\n')
                last_newline = 1
                indent++
                buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
            case ']':
                indent--
                buffer.WriteByte('\n')
                last_newline = 1
                buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
                buffer.WriteByte(v)
                //buffer.WriteByte('\n')
                //buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
            case ',':
                buffer.WriteByte(v)
                buffer.WriteByte('\n')
                last_newline = 1
                buffer.WriteString(strings.Repeat(p.paddingString, indent * p.paddingCount))
            //skip space
            case '\n':
            case '\t':
            case '\r':
            case ' ':
            case ':':
                buffer.WriteByte(v)
                buffer.WriteByte(' ')
                last_newline = 0
            default:
                buffer.WriteByte(v)
                last_newline = 0
            }
            continue
        }
        buffer.WriteByte(v)
        last_newline = 0
    }

    return buffer.String()
}

func handleOneFile(file string, p *params) {
    f, err := os.OpenFile(file, os.O_RDONLY , 0660)
    if err != nil {
        fmt.Println("Failed to open file,", err)
        return
    }
    defer f.Close()

    data, err := ioutil.ReadAll(f)
    if err != nil {
        fmt.Println("Failed to read file, ", err)
        return
    }

    prettyjson := makePrettyJson(data, p)
    fmt.Println(prettyjson)
}

func handleFiles(files []string, p *params) {
	for _, file := range files {
        handleOneFile(file, p)
	}
}

func handleInput( p *params) {
    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Println("Failed to read data, ", err)
    }
    prettyjson := makePrettyJson(data, p)
    fmt.Println(prettyjson)
}

func main() {
    p, files := handleCommandLine()
    if len(files) >= 1 {
        handleFiles(files, p)
    } else {
        handleInput(p)
    }
}
