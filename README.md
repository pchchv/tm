<div align="center">

# **Simple Turing Machine implement in Golang**

</div>

## What is this Turing Machine

A Turing machine is an abstract "machine" that manipulates symbols on a strip of tape according to a table of rules; to be more exact, it is a mathematical model that defines such a device. Despite the model's simplicity, given any computer algorithm, a Turing machine can be constructed that is capable of simulating that algorithm's logic. ([Wiki](https://en.wikipedia.org/wiki/Turing_machine) Quote)

## Install

```
go get github.com/pchchv/tm
```

## Usage

```go
package main

import (
    "fmt"
    . "github.com/pchchv/tm"
)

func main() {
    nTM := NewTM()
    //Input State and declare if it is final state
    nTM.InputState("0", false)
    nTM.InputState("1", true)
    //Input config
    // InputConfig parameter as follow:
    // - SourceState, 
    // - Input
    // - Modified Value
    // - DestinationState
    // - Tape Head Move Direction
    nTM.InputConfig("0", "1", "1", "1", MoveRight)
    nTM.InputConfig("0", "0", "1", "0", MoveLeft)
    nTM.InputConfig("1", "0", "1", "0", MoveLeft)
    nTM.InputConfig("1", "1", "1", "1", MoveRight)
    //Input tape data
    nTM.InputTape("0", "0", "1", "1", "0", "0", "0")
    //Run TM to the finish (if exist)
    nTM.Run()
    fmt.Println("New Tape:=", nTM.ExportTape())
}

```
